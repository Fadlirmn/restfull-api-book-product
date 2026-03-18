package models

import "time"

type RefreshToken struct{
	TokenId  int
	UserID string
	Token string
	ExpiresAT time.Time
}