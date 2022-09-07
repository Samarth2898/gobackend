package token

import "time"

type Maker interface {
	// CreateToken creates a token for a specific username and duration.
	CreateToken(username string, duration time.Duration) (string, *Payload, error)
	// VerfifyToken checks if token is valid or not
	VerifyToken(token string) (*Payload, error)
}
