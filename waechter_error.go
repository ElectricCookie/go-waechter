package waechter

//AuthError describes all errors that can happen in this package
type AuthError struct {
	ErrorCode   string `json:"errorCode"`
	Description string `json:"description"`
	IsInternal  bool   `json:"isInternal"`
	Err         error
}

func internalError(err error) *AuthError {
	return &AuthError{
		ErrorCode:   "internalError",
		Description: "Unknown internal error",
		IsInternal:  true,
		Err:         err,
	}
}

func emailNotVerifiedError() *AuthError {
	return &AuthError{
		ErrorCode:   "emailNotVerified",
		Description: "The email address is not verified",
		IsInternal:  false,
	}
}

func cryptError(err error) *AuthError {
	return &AuthError{
		ErrorCode:   "cryptError",
		Description: "Encryption failed",
		IsInternal:  true,
		Err:         err,
	}
}

func userNotFoundError(err error) *AuthError {
	return &AuthError{
		ErrorCode:   "userNotFound",
		Description: "User was not found",
		IsInternal:  true,
		Err:         err,
	}
}

func randomError(err error) *AuthError {
	return &AuthError{
		ErrorCode:   "randomError",
		Description: "Something went wrong while generating random data for tokens",
		IsInternal:  true,
		Err:         err,
	}
}

func invalidPasswordError() *AuthError {
	return &AuthError{
		ErrorCode:   "wrongPassword",
		Description: "The passsword provided was not correct",
		IsInternal:  false,
	}
}

//InvalidParametersError is thrown if the parameters passed to a request are invalid
func InvalidParametersError(err error) *AuthError {
	return &AuthError{
		ErrorCode:   "invalidParameters",
		Description: "The parameters passed were not valid",
		IsInternal:  false,
		Err:         err,
	}
}

//NotLoggedInError obviously.
func NotLoggedInError() *AuthError {
	return &AuthError{
		ErrorCode:   "notLoggedIn",
		Description: "There was no refresh token present.",
		IsInternal:  false,
	}
}

func dbWriteError(err error) *AuthError {
	return &AuthError{
		ErrorCode:   "dbWriteErr",
		Description: "Could not write data to the db",
		IsInternal:  true,
		Err:         err,
	}
}

func emailSendError(err error) *AuthError {
	return &AuthError{
		ErrorCode:   "emailSendErr",
		Description: "Could not send the email",
		IsInternal:  true,
		Err:         err,
	}
}
