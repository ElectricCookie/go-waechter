package waechter

//DBAdapter is an interface used to connect to a database
type DBAdapter interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByUsername(username string) (*User, error)
	GetUserById(id string) (*User, error)

	GetUserByUsernameOrEmail(input string) (*User, error)

	CreateUser(*User) error

	VerifyEmail(*User) error

	InsertRefreshToken(token *RefreshToken) error

	FindRefreshToken(id string, token string) (*RefreshToken, error)
}
