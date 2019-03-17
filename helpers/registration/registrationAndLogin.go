package registerAndLoginHelpers

import (
	"database/sql"
	"net/http"

	"../../models/user"

	"../../helpers/user"

	_ "github.com/go-sql-driver/mysql"

	"../../helpers/password"
)

func FailOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func EmailExists(db *sql.DB, userToRegister user.User) bool {
	userCheckQuery := "Select username from users where email = ?"
	rows := db.QueryRow(userCheckQuery, userToRegister.Email)
	var email string
	err := rows.Scan(&email)
	// fmt.Println(username)
	if err == sql.ErrNoRows {
		return false
	}
	return true
}

func UserExists(db *sql.DB, userToRegister user.User) bool {
	userCheckQuery := "Select username from users where username = ?"
	rows := db.QueryRow(userCheckQuery, userToRegister.Username)
	var username string
	err := rows.Scan(&username)
	// fmt.Println(username)
	if err == sql.ErrNoRows {
		return false
	}
	return true
}

func InsertUser(db *sql.DB, userToRegister user.User) bool {
	stmt, err := db.Prepare("INSERT INTO users(username,firstname,lastname,email,password,type) VALUES(?,?,?,?,?,?)")
	FailOnError(err)

	result, err := stmt.Exec(userToRegister.Username, userToRegister.Firstname, userToRegister.Lastname, userToRegister.Email, userToRegister.Password, userToRegister.Type)
	FailOnError(err)

	rowsEffected, err := result.RowsAffected()
	FailOnError(err)
	if rowsEffected > 0 {
		return true
	}

	return false
}

func SetUserPasswordAsHash(user *user.User) {
	hashedPassword := password.GenerateHashedPassword([]byte(user.Password))
	user.Password = hashedPassword
}

func ValidateUser(responseWriter *http.ResponseWriter, db *sql.DB, userToRegister user.User) bool {
	// fmt.Println(isValidUsername(db, userToRegister))
	if UserExists(db, userToRegister) {
		(*responseWriter).Write([]byte("username already exists"))
	} else if EmailExists(db, userToRegister) {
		(*responseWriter).Write([]byte("Email already exists"))
	} else {
		if !userHelpers.ValidUser(&userToRegister) {
			(*responseWriter).Write([]byte("Invalid User Data"))
		} else {
			return true
		}
	}
	return false
}

func IsValidPassword(db *sql.DB, user user.User) bool {
	userCheckQuery := "Select password from users where username = ?"
	rows := db.QueryRow(userCheckQuery, user.Username)
	var passwordInDB string
	err := rows.Scan(&passwordInDB)
	if !password.VerifyPassword(passwordInDB, user.Password) || err == sql.ErrNoRows {
		return false
	}
	return true
}
