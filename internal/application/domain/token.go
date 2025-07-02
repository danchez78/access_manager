package domain

type AccessToken string

type RefreshToken struct {
	UserID         UserID
	String         string
	IssuedTime     int64
	ExpirationTime int64
}

func NewRefreshToken(userID UserID, token string, issuedTime, ExpirationTime int64) *RefreshToken {
	return &RefreshToken{
		UserID:         userID,
		String:         token,
		IssuedTime:     issuedTime,
		ExpirationTime: ExpirationTime,
	}
}
