package main

import "time"

// User is the general user structure
type User struct {
	ID                string
	Username          string
	FirstName         string
	Email             string
	LastName          string
	PasswordHash      string
	Language          string
	EmailVerfied      bool
	VerificationToken string
	Salt              string
	Registered        time.Time
	LastLogin         time.Time
}
