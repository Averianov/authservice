package controllers

import (
	"authservice/mocks"
	"authservice/models"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
)

const DOMAIN = "localhost"
const SECRETKEY = "So$0meP3r[hektK&y"

var trueRefreshToken string
var notExistRefreshToken string

func TestAuthController(t *testing.T) {
	models.Secret = SECRETKEY
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testAuthenticate(t, ctrl)
	testRefresh(t, ctrl)

}

func testAuthenticate(t *testing.T, ctrl *gomock.Controller) {
	type testValue struct {
		GUID             string
		expectedBody     string
		expectingUseMosk bool
		expectingErrMosk error
	}

	var testValueArr []testValue

	testValueArr = append(testValueArr, testValue{
		"6F9619FF-8B86-D011-B42D-00CF4FC964FF",
		`"message":"Tokens has been created","status":true}`,
		true,
		nil,
	})

	testValueArr = append(testValueArr, testValue{
		"6F9619FF-8B86-D011-B42D-00CF4FC964FF",
		`"message":"Tokens has been created","status":true}`,
		true,
		nil,
	})

	testValueArr = append(testValueArr, testValue{
		"as8df-a9sd87f9aa9sd8-fasd8f9asd87f98asd",
		`{"message":"guid: invalid format","status":false}`,
		false,
		nil,
	})

	testValueArr = append(testValueArr, testValue{
		"6F9619FF-8B86-D011-B42D-00CF4FC964FF",
		`{"message":"SomeError","status":false}`,
		true,
		fmt.Errorf("SomeError"),
	})

	authTest := func(t *testing.T, GUID string, useSaveMock bool, errSaveMosk error) (resp *http.Response, body []byte) {
		makeBody := fmt.Sprintf("{\"guid\": \"%s\"}", GUID)

		requestBody := strings.NewReader(makeBody)
		r := httptest.NewRequest("POST", "http://"+DOMAIN+"/auth/login", requestBody)
		w := httptest.NewRecorder()

		account := models.NewAccount()
		session := mocks.NewMockSession(ctrl)
		controller := NewAuthController(&account, session, DOMAIN)
		if useSaveMock {
			session.EXPECT().Save().Return(errSaveMosk)
		}
		controller.Authenticate(w, r)

		resp = w.Result()
		body, _ = ioutil.ReadAll(resp.Body)
		return
	}

	for i, tVal := range testValueArr {
		resp, body := authTest(t, tVal.GUID, tVal.expectingUseMosk, tVal.expectingErrMosk)
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Invalid status code %d. Expected: %d. Iteration = %d\n", resp.StatusCode, http.StatusOK, i)
		}

		contentType := resp.Header.Get("Content-Type")
		expectedContentType := "application/json"
		if contentType != expectedContentType {
			t.Fatalf("Invalid content-type %s. Expected: %s. Iteration = %d\n", contentType, expectedContentType, i)
		}

		if !strings.Contains(string(body), tVal.expectedBody) {
			t.Fatalf("Invalid body %s. Expected: %s. Iteration = %d\n", string(body), tVal.expectedBody, i)
		}

		if tVal.expectingUseMosk == true && tVal.expectingErrMosk == nil {
			cookies := resp.Cookies()
			for _, cookie := range cookies {
				if cookie.Name == "refresh_token" {
					if trueRefreshToken != "" {
						notExistRefreshToken = trueRefreshToken
					}
					trueRefreshToken = cookie.Value
					break
				}
			}
		}
	}
}

func testRefresh(t *testing.T, ctrl *gomock.Controller) {
	type testValue struct {
		token                   string
		expectedBody            string
		expectingUseCompareMosk bool
		expectingErrCompareMosk error
		expectingUseSaveMosk    bool
	}

	var testValueArr []testValue

	testValueArr = append(testValueArr, testValue{
		trueRefreshToken,
		`"message":"Tokens has been created","status":true}`,
		true,
		nil,
		true,
	})

	testValueArr = append(testValueArr, testValue{
		"eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDc5ODI3NDcsImp0aSI6IjZmOTYxOWZmLThiODYtZDAxMS1iNDJkLTAwY2Y0ZmM5NjRmZiIsInN1YiI6InJlZnJlc2hfdG9rZW4ifQ.XZ6VvWDXTYV6wCHRj6QqBNuRkjp0JrMY7jgGr3e_zQR_ZoYCkro3AcLBxPyXq8dH",
		`"message":"token is expired by`,
		false,
		nil,
		false,
	})

	testValueArr = append(testValueArr, testValue{
		notExistRefreshToken,
		`{"message":"Not Exist","status":false}`,
		true,
		fmt.Errorf("Not Exist"),
		false,
	})

	testValueArr = append(testValueArr, testValue{
		"eyJleHAiOjE2MDc5ODA2MjAsImp0aSI6IjZmOTYxOWZmLThiODYtZDAxMS1iNDJkLTAwY2Y0ZmM5NjRmZiIsInN1YiI6InJlZnJlc2hfdG9rZW4ifQ",
		`{"message":"token contains an invalid number of segments","status":false}`,
		false,
		nil,
		false,
	})

	authTest := func(t *testing.T, token string, useCompareMock bool, errCompareMock error, useSaveMock bool) (resp *http.Response, body []byte) {

		r := httptest.NewRequest("POST", "http://"+DOMAIN+"/auth/refresh", nil)
		r.Header.Add("Authorization", "Bearer "+token)
		w := httptest.NewRecorder()

		account := models.NewAccount()
		session := mocks.NewMockSession(ctrl)
		controller := NewAuthController(&account, session, DOMAIN)
		if useCompareMock {
			session.EXPECT().CompareWithExisting(token).Return(errCompareMock)
		}
		if useSaveMock {
			session.EXPECT().Save().Return(nil)
		}
		controller.Refresh(w, r)

		resp = w.Result()
		body, _ = ioutil.ReadAll(resp.Body)
		return
	}

	for i, tVal := range testValueArr {
		resp, body := authTest(t, tVal.token, tVal.expectingUseCompareMosk, tVal.expectingErrCompareMosk, tVal.expectingUseSaveMosk)
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Invalid status code %d. Expected: %d. Iteration = %d\n", resp.StatusCode, http.StatusOK, i)
		}

		contentType := resp.Header.Get("Content-Type")
		expectedContentType := "application/json"
		if contentType != expectedContentType {
			t.Fatalf("Invalid content-type %s. Expected: %s. Iteration = %d\n", contentType, expectedContentType, i)
		}

		if !strings.Contains(string(body), tVal.expectedBody) {
			t.Fatalf("Invalid body %s. Expected: %s. Iteration = %d\n", string(body), tVal.expectedBody, i)
		}
	}
}

//*/
