package user

// Login model is binded on the /login function, must contain username and password in json payloads
type Login struct {
	Username string `json:"username"`
	Password string `json:"password`
}
