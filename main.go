package main

import (
	"github.com/Averianov/authservice"
)

func main() {
	err := authservice.run()
	if err != nil {
		panic(err)
	}
}
