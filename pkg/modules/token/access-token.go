package token

type AccessToken struct {
	Token          string
	ExpirationTime int // in seconds
}
