package user

type Request struct {
	Username   string `json:"username,required"`
	Password   string `json:"password"`
	ActualName string `json:"actual_name"`
	Mobile     string `json:"mobile"`
	Email      string `json:"email"`
}
