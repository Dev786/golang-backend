package user

type User struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Type      string `json:"type"`
}
