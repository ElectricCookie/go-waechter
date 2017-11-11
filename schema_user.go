package waechter

import "time"

// User is the general user structure
type User struct {
	ID       string `bson:"_id"`
	Username string `bson:"username"`
	Email    string `bson:"email"`

	FirstName string `bson:"firstName"`
	LastName  string `bson:"lastName"`

	PasswordHash string `bson:"passwordHash"`
	Salt         string `bson:"salt"`

	ForgotPasswordToken       string `bson:"forgotPasswordToken"`
	ForgotPasswordRequestTime int64  `bson:"forgotPasswordTokenRequestTime"`

	Language string `bson:"language"`

	EmailVerfied      bool   `bson:"emailVerified"`
	VerificationToken string `bson:"verificationToken"`

	Registered time.Time `bson:"registeredAt"`
	LastLogin  time.Time `bson:"lastLoginAt"`
}
