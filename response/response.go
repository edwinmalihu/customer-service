package response

type SuccessAddCustomer struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email" binding:"required,email`
}
