package tests

import (
	"fmt"
	"log"

	"github.com/panprogramadorgh/goquickjwt/pkg/jwt"
)

func Foo() {
	p := jwt.Payload{
		"username": "user1",
	}
	token, err := p.SignWithHS256("helloworld")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("token signed with hs256 alg: %v\n", token)

	originalP, err := jwt.VerifyWithHS256("helloworld", token)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("original payload: %v\n", originalP)
}
