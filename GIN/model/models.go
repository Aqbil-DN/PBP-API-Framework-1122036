package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type GetAllUsersResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []User `json:"data"`
}
