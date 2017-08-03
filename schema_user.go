package waechter

import "time"

// User is the general user structure
type User struct {
	ID                  string    `bson:"_id"`
	Username            string    `bson:"username"`
	FirstName           string    `bson:"firstName"`
	Email               string    `bson:"email"`
	LastName            string    `bson:"lastName"`
	PasswordHash        string    `bson:"passwordHash"`
	ForgotPasswordToken string    `bson:"forgotPasswordToken"`
	Language            string    `bson:"language"`
	EmailVerfied        bool      `bson:"emailVerified"`
	VerificationToken   string    `bson:"verificationToken"`
	Salt                string    `bson:"salt"`
	Registered          time.Time `bson:"registeredAt"`
	LastLogin           time.Time `bson:"lastLoginAt"`
}
