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
		panic(err)
	}
	if registerAndLoginHelpers.ValidateLogin(db, &responseWriter, user) {
		session, err := sessionHelpers.CreateSession(user.Username)
		if err != nil {
			panic(err)
		} else {
			http.SetCookie(responseWriter, &http.Cookie{
				Name:    "session_token",
				Value:   session,
				Expires: time.Now().Add(120 * time.Second),
			})
			responseWriter.Write([]byte("User Successfully Logged In"))
		}
	}
}
