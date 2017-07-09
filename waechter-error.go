package waechter

//AuthError describes all errors that can happen in this package
type AuthError struct {
	ErrorCode   string `json:"errorCode"`
	Description string `json:"description"`
	IsInternal  bool   `json:"isInternal"`
	Err         error
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

func invalidParameters(err error) *AuthError {
	return &AuthError{
		ErrorCode:   "invalidParameters",
		Description: "The parameters passed were not valid",
		IsInternal:  false,
		Err:         err,
	}
}

func dbWriteErr(err error) *AuthError {
	return &AuthError{
		ErrorCode:   "dbWriteErr",
		Description: "Could not write data to the db",
		IsInternal:  true,
		Err:         err,
	}
}

func emailSendErr(err error) *AuthError {
	return &AuthError{
		ErrorCode:   "emailSendErr",
		Description: "Could not send the email",
		IsInternal:  true,
		Err:         err,
	}
}
