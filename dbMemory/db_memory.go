package dbMemory

import (
	"errors"

	waechter "github.com/ElectricCookie/go-waechter"
)

//MemoryAdapter is a  test adapter that stores all data in RAM
type MemoryAdapter struct {
	users  []*waechter.User
	tokens []*waechter.RefreshToken
}

//Reset clears all data
func (adapter *MemoryAdapter) Reset() {

	adapter.users = []*waechter.User{}
	adapter.tokens = []*waechter.RefreshToken{}

}

// GetUserByEmail get user by email
func (adapter *MemoryAdapter) GetUserByEmail(email string) (*waechter.User, error) {

	for _, user := range adapter.users {

		if user.Email == email {
			return user, nil
		}

	}

	return nil, errors.New("notFound")

}

// GetUserByUsername get user by username
func (adapter *MemoryAdapter) GetUserByUsername(username string) (*waechter.User, error) {

	for _, user := range adapter.users {

		if user.Username == username {
			return user, nil
		}

	}

	return nil, errors.New("notFound")
}

// GetUserByID get user by ID
func (adapter *MemoryAdapter) GetUserByID(id string) (*waechter.User, error) {

	for _, user := range adapter.users {

		if user.ID == id {
			return user, nil
		}

	}

	return nil, errors.New("notFound")
}

// GetUserByUsernameOrEmail get user by username or email
func (adapter *MemoryAdapter) GetUserByUsernameOrEmail(input string) (*waechter.User, error) {
	for _, user := range adapter.users {
		if user.Username == input || user.Email == input {
			return user, nil
		}
	}

	return nil, errors.New("notFound")
}

// CreateUser insert new user
func (adapter *MemoryAdapter) CreateUser(user *waechter.User) error {

	adapter.users = append(adapter.users, user)

	return nil

}

// VerifyEmail verifies the email address of a user
func (adapter *MemoryAdapter) VerifyEmail(userID string) error {

	user, err := adapter.GetUserByID(userID)

	if err != nil {
		return err
	}

	user.EmailVerfied = true

	return nil

}

// InsertRefreshToken insert a token
func (adapter *MemoryAdapter) InsertRefreshToken(token *waechter.RefreshToken) error {

	adapter.tokens = append(adapter.tokens, token)

	return nil

}

// FindRefreshToken retrieve by userID and tokenID
func (adapter *MemoryAdapter) FindRefreshToken(userID string, tokenID string) (*waechter.RefreshToken, error) {

	for _, token := range adapter.tokens {

		if token.Token == tokenID && token.UserID == userID {
			return token, nil
		}

	}

	return nil, errors.New("notFound")

}

// SetForgotPasswordToken writes a "forgotPasswordToken" to the db
func (adapter *MemoryAdapter) SetForgotPasswordToken(userID string, token string) error {

	user, err := adapter.GetUserByID(userID)

	if err != nil {
		return err
	}

	user.ForgotPasswordToken = token

	return nil

}

//SetVerificationToken writes a "verificationToken" to the db
func (adapter *MemoryAdapter) SetVerificationToken(userID string, token string) error {
	user, err := adapter.GetUserByID(userID)

	if err != nil {
		return err
	}

	user.VerificationToken = token

	return nil
}

//SetPassword sets the password
func (adapter *MemoryAdapter) SetPassword(userID string, hash string) error {
	user, err := adapter.GetUserByID(userID)

	if err != nil {
		return err
	}

	user.PasswordHash = hash

	return nil
}
