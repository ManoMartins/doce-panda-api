package dtos

type InputAuthenticationUserDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type OutputAuthenticationUserDto struct {
	Token string `json:"token"`
	User  struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"user"`
}
