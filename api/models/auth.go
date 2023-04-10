package models

type Register struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Login       string `json:"login"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	UserType    string `json:"user_type"`
}

type Login struct {
	Login    string `json:"login"`
	Register string `json:"register"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type PasswordResponse struct {
	AccessToken string `json:"access_token"`
}
