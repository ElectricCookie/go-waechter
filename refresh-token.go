package waechter

// RefreshToken is a token that can be used to gain access to an access_token
type RefreshToken struct {
	Token   string `json:"tokenId" bson:"_id"`
	UserID  string `json:"userId" bson:"userId"`
	Expires int64  `json:"expires" bson:"expires"`
}
