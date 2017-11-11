package waechter

//UserLoginUsernameParams is the required infromation for logging in using the username
type UserLoginUsernameParams struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RememberMe bool   `json:"rememberMe" binding:"required"`
}

//UserLoginUsername logs in using email
func (waechter *Waechter) UserLoginUsername(parameters UserLoginUsernameParams) (string, *AuthError) {

	var user *User
	var err error
	// Retrieve the desired user
	if user, err = waechter.DbAdapter.GetUserByUsername(parameters.Username); err != nil {
		return "", userNotFoundError(err)
	}

	return waechter.performLogin(parameters.Password, user, parameters.RememberMe)
}
