package waechter

//UserLoginEmailParams is the required information for loggin in using the email address
type UserLoginEmailParams struct {
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RememberMe bool   `json:"rememberMe" binding:"required"`
}

//UserLoginEmail logs in using emailemail
func (waechter *Waechter) UserLoginEmail(parameters UserLoginEmailParams) (string, *AuthError) {

	var user *User
	var err error

	// Retrieve the desired user
	if user, err = waechter.DbAdapter.GetUserByEmail(parameters.Email); err != nil {
		return "", userNotFoundError(err)
	}
	return waechter.performLogin(parameters.Password, user, parameters.RememberMe)

}
