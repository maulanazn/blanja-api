package middleware

import (
	"fmt"
	"log"
	"net/http"
)

type AuthenticateTokenHandler func(writer http.ResponseWriter, req *http.Request)

type MakeSureToken struct {
	handler AuthenticateTokenHandler
}

func (authenticate *MakeSureToken) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	authorization := r.Header.Get("Authorization")

	if authorization == "" {
		if _, err := fmt.Fprint(w, "Failed to Authorize"); err != nil {
			log.Println(err.Error())
		}
		return
	}

	authenticate.handler(w, r)
}

func NewEntranceToken(wrapHandler AuthenticateTokenHandler) *MakeSureToken {
	return &MakeSureToken{wrapHandler}
}
