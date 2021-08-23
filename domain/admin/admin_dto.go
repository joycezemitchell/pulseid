package admin

type Admin struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type RequestToken struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}
