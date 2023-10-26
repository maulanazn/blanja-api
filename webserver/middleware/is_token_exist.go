package middleware

import (
	"fmt"
	"log"
	"net/http"
)

type AuthenticateTokenHandler func(http.ResponseWriter, *http.Request)

type MakesureToken struct {
	handler AuthenticateTokenHandler
}

func (authenticate *MakesureToken) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("USR_ID")
	_, tokenerr := r.Cookie("TKN_ID")
	if err != nil {
		log.Println(err)
		fmt.Fprint(w, "Need Cookie")
		return
	}

	if tokenerr != nil {
		log.Println(err)
		fmt.Fprint(w, "Need Cookie")
		return
	}

	authenticate.handler(w, r)
}

func NewEntranceToken(wrapHandler AuthenticateTokenHandler) *MakesureToken {
	return &MakesureToken{wrapHandler}
}
