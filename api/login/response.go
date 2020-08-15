package login

type Response struct {
	Token      string `json:"token"`
	Id         uint   `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	ActualName string `json:"actual_name"`
	Mobile     string `json:"mobile"`
	Email      string `json:"email"`
}
