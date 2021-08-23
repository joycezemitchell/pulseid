package client

type Client struct {
	Id         int64  `json:"id"`
	Token      string `json:"token"`
	Status     string `json:"status"`
	Created_at string `json:"created_at,omitempty"`
	Updated_at string `json:"updated_at,omitempty"`
}

type ValidateToken struct {
	Id         int64  `json:"id"`
	Status     string `json:"status"`
	Created_at string `json:"created_at,omitempty"`
	Updated_at string `json:"updated_at,omitempty"`
}
