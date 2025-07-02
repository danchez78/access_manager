package token_generator

type Token struct {
	UserID         string
	String         string
	IssuedTime     int64
	ExpirationTime int64
}
