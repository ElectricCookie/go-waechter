package waechter

//DBAdapter is an interface used to connect to a database
type DBAdapter interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByUsername(username string) (*User, error)
	GetUserByID(id string) (*User, error)

	GetUserByUsernameOrEmail(input string) (*User, error)

	CreateUser(*User) error

	VerifyEmail(userID string) error

	InsertRefreshToken(token *RefreshToken) error

	SetForgotPasswordToken(userID string, token string) error

	SetVerificationToken(userID string, token string) error

	FindRefreshToken(userID string, tokenID string) (*RefreshToken, error)

	SetPassword(userID string, hash string) error
}
