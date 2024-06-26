package middleware

import (
	"errors"
	"net/http"

	"github.com/jantoniogonzalez/go-web-server/api"
	"github.com/jantoniogonzalez/go-web-server/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Invalid username or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" {
			// Throw Error
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		var database *tools.DatabaseInterface // We use the pointer to interface because we won't change the value
		database, err = tools.NewDatabase()
		// Check if db is created
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)
		// Check if login details match in db
		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		// Go next middleware
		next.ServeHTTP(w, r)
	})
}
