package NekoCAS

type user struct {
	Name    string
	Email   string
	Token   string
	Message string
}

type casResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Token string `json:"token"`
	} `json:"data"`
	Message string `json:"message"`
}
