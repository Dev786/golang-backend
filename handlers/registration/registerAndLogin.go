package registerAndLogin

import (
	"net/http"
	"time"

	"../../models/user"

	"../../helpers/user"

	"../../helpers/registration"

	"../../helpers/session"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitializeDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/social_media")
	registerAndLoginHelpers.FailOnError(err)
	return db
}

func Register(responseWriter http.ResponseWriter, request *http.Request) {
	var userToRegister user.User
	err := userHelpers.ParseJSONToUser(&userToRegister, request.Body)
	registerAndLoginHelpers.FailOnError(err)

	db := InitializeDB()
	defer db.Close()

	registerAndLoginHelpers.SetUserPasswordAsHash(&userToRegister)
	isValidUser := registerAndLoginHelpers.ValidateUser(&responseWriter, db, userToRegister)

	if isValidUser {
		registerAndLoginHelpers.InsertUser(db, userToRegister)
		responseWriter.Write([]byte("User is registered"))
	}
}

func Login(responseWriter http.ResponseWriter, request *http.Request) {
	//email ID & password
	var user user.User
	db := InitializeDB()
	defer db.Close()
	err := userHelpers.ParseJSONToUser(&user, request.Body)
	if err != nil {
		registerAndLoginHelpers.FailOnError(err)
	}
	if registerAndLoginHelpers.ValidateLogin(db, &responseWriter, user) {
		sessionId, err := sessionHelpers.CreateSession(user.Username)
		if err != nil {
			registerAndLoginHelpers.FailOnError(err)
		} else {
			cookie := http.Cookie{
				Name:    "session_token",
				Value:   sessionId,
				Expires: time.Now().Add(120 * time.Second),
			}
			http.SetCookie(responseWriter, &cookie)
			responseWriter.Write([]byte("User Login Successfully "))
		}
	}
}
