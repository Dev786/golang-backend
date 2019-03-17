package userHelpers

import (
	"encoding/json"
	"io"
	"regexp"

	"../../models/user"
)

func ParseToJSON(u *user.User, reader io.Reader) ([]byte, error) {
	var jsonData []byte
	jsonData, err := json.Marshal(u)
	if err != nil {
		return jsonData, err
	}
	return jsonData, nil
}

func ParseJSONToUser(u *user.User, jsonData io.Reader) error {
	err := json.NewDecoder(jsonData).Decode(u)
	if err != nil {
		return err
	}
	return nil
}

func validateEmail(email *string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(*email)
}

func ValidUser(u *user.User) bool {
	if len(u.Username) == 0 || len(u.Password) == 0 || len(u.Firstname) == 0 || len(u.Lastname) == 0 || len(u.Email) == 0 || len(u.Type) == 0 {
		return false
	} else {
		return validateEmail(&u.Email)
	}
}
