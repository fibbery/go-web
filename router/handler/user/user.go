package user

type CreateRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type CreateResponse struct {
	UserName string `json:"username"`
}

type ListRequest struct {
	Offset int `form:"offset" json:"offset"`
	Limit  int `form:"limit" json:"limit"`
}
