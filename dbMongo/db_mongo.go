package dbMongo

import (
	"log"
	"strings"

	"github.com/ElectricCookie/go-waechter"
	mgo "github.com/go-mgo/mgo"
	"github.com/go-mgo/mgo/bson"
)

//MongoAdapter is a ready adapter to connect go-waechter to mongodb
type MongoAdapter struct {
	Db *mgo.Database
}

//NewMongoAdapter creates a new mongodb adapter
func NewMongoAdapter(address string, db string) *MongoAdapter {

	a := MongoAdapter{}

	session, err := mgo.Dial(address)

	if err != nil {
		log.Fatalf("Could not connect to mongodb")
	}

	a.Db = session.DB(db)

	return &a
}

// GetUserByEmail get user by email
func (adapter *MongoAdapter) GetUserByEmail(email string) (*waechter.User, error) {

	user := waechter.User{}

	err := adapter.Db.C("users").Find(bson.M{
		"email": strings.ToLower(email),
	}).One(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil

}

// GetUserByUsername get user by username
func (adapter *MongoAdapter) GetUserByUsername(username string) (*waechter.User, error) {

	user := waechter.User{}

	err := adapter.Db.C("users").Find(bson.M{
		"username": username,
	}).One(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUserByID get user by ID
func (adapter *MongoAdapter) GetUserByID(id string) (*waechter.User, error) {
	user := waechter.User{}

	err := adapter.Db.C("users").FindId(id).One(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUserByUsernameOrEmail get user by username or email
func (adapter *MongoAdapter) GetUserByUsernameOrEmail(input string) (*waechter.User, error) {
	user := waechter.User{}

	err := adapter.Db.C("users").Find(bson.M{"$or": []bson.M{bson.M{"username": input}, bson.M{"email": strings.ToLower(input)}}}).One(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// CreateUser insert new user in DB
func (adapter *MongoAdapter) CreateUser(user *waechter.User) error {
	return adapter.Db.C("users").Insert(user)
}

// VerifyEmail verifies the email address of a user
func (adapter *MongoAdapter) VerifyEmail(userID string) error {

	return adapter.Db.C("users").Update(bson.M{
		"_id": userID,
	}, bson.M{
		"$set": bson.M{
			"emailVerified": true,
		},
	})

}

// InsertRefreshToken insert a token
func (adapter *MongoAdapter) InsertRefreshToken(token *waechter.RefreshToken) error {
	return adapter.Db.C("refreshTokens").Insert(token)
}

// FindRefreshToken retrieve by userID and tokenID
func (adapter *MongoAdapter) FindRefreshToken(userID string, tokenID string) (*waechter.RefreshToken, error) {

	token := waechter.RefreshToken{}

	err := adapter.Db.C("refreshTokens").Find(bson.M{
		"userId": userID,
		"_id":    tokenID,
	}).One(&token)

	if err != nil {
		return nil, err
	}

	return &token, nil

}

// SetForgotPasswordToken writes a "forgotPasswordToken" to the db
func (adapter *MongoAdapter) SetForgotPasswordToken(userID string, token string) error {

	return adapter.Db.C("users").Update(bson.M{
		"_id": userID,
	}, bson.M{
		"$set": bson.M{
			"forgotPasswordToken": token,
		},
	})

}

//SetVerificationToken writes a "verificationToken" to the db
func (adapter *MongoAdapter) SetVerificationToken(userID string, token string) error {
	return adapter.Db.C("users").Update(bson.M{
		"_id": userID,
	}, bson.M{
		"$set": bson.M{
			"verificationToken": token,
		},
	})
}

//SetPassword sets the password
func (adapter *MongoAdapter) SetPassword(userID string, hash string) error {
	return adapter.Db.C("users").Update(bson.M{
		"_id": userID,
	}, bson.M{
		"$set": bson.M{
			"passwordHash": hash,
		},
	})
}
