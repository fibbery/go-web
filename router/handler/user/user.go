package user

type CreateRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type CreateResponse struct {
	UserName string `json:"username"`
}
