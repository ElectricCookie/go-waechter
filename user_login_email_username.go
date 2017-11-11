package waechter

// UserLoginEmailOrUsernameParams is the required information for logging in
type UserLoginEmailOrUsernameParams struct {
	UsernameOrEmail string `json:"usernameOrEmail" binding:"required"`
	Password        string `json:"password" binding:"required"`
	RememberMe      bool   `json:"rememberMe"`
}

// UserLoginWithUsernameOrEmail logs a user in using email or username and password. Returns a new refresh token.
func (waechter *Waechter) UserLoginWithUsernameOrEmail(parameters UserLoginEmailOrUsernameParams) (string, *AuthError) {
	var user *User
	var err error
	// Retrieve the desired user
	if user, err = waechter.getDBAdapter().GetUserByUsernameOrEmail(parameters.UsernameOrEmail); err != nil {
		return "", userNotFoundError(err)
	}

	return waechter.performLogin(parameters.Password, user, parameters.RememberMe)
}
