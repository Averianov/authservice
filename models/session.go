package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// Session is a interface who implements session struct
type Session interface {
	CompareWithExisting(string) error
	Save() error
}

//Session struct contains account{GUID and Token value}
type session struct {
	Account account `bson:"account"`
}

// NewSession is a function who returned new session
func NewSession(account account) *session {
	return &session{Account: account}
}

// CompareWithExisting is method who find hashed token in DB, compare incoming token with him and delete old session from DB if true
func (session *session) CompareWithExisting(incomingToken string) (err error) {
	db, err := ConnectToDB()
	if err != nil {
		return
	}
	defer db.Close()

	filter := bson.D{primitive.E{Key: "account.guid", Value: session.Account.GUID}}
	err = db.collection.FindOne(db.ctx, filter).Decode(session)
	if err != nil {
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(session.Account.Token), []byte(incomingToken))
	return
}

// Save is method who saved or replased session in DB
func (session *session) Save() (err error) {
	var hashedToken []byte
	db, err := ConnectToDB()
	if err != nil {
		return
	}
	defer db.Close()

	hashedToken, err = bcrypt.GenerateFromPassword([]byte(session.Account.Token), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	session.Account.Token = string(hashedToken)

	filter := bson.D{primitive.E{Key: "account.guid", Value: session.Account.GUID}}
	err = db.collection.FindOne(db.ctx, filter).Err()
	if err == nil {
		update := bson.D{
			{"$set", bson.D{
				{"account.token", session.Account.Token},
			}},
		}
		_, err = db.collection.UpdateOne(db.ctx, filter, update)
	} else {
		_, err = db.collection.InsertOne(db.ctx, session)
	}
	return
}
