package middleware

import (
	"fmt"
	"log"
	"net/http"
	"util"
)

type AuthenticateTokenHandler func(http.ResponseWriter, *http.Request)

type MakesureToken struct {
	handler AuthenticateTokenHandler
}

func (authenticate *MakesureToken) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("USR_ID")
	authorization := r.Header.Get("Authorization")
	emailfromtoken := util.DecodeToken(authorization[7:])

	if err != nil {
		log.Println(err)
		fmt.Fprint(w, "Please login correctly")
		return
	}

	if authorization == "" || r.FormValue("email") != emailfromtoken {
		log.Println(err)
		fmt.Fprint(w, "Failed to Authorize")
		return
	}

	authenticate.handler(w, r)
}

func NewEntranceToken(wrapHandler AuthenticateTokenHandler) *MakesureToken {
	return &MakesureToken{wrapHandler}
}
