package models


type User struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
type UserPKey struct {
	UserId string `json:"user_id"`
	Login  string `json:"login"`
}

type GetListUserRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListUserResponse struct {
	Count int     `json:"count"`
	Users []*User `json:"Users"`
}