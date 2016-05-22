package models

import (
	"errors"
)

var (
	ErrNoRows = errors.New("No row found")
	ErrArgs   = errors.New("args error may be empty")
	ErrExist  = errors.New("Have already exist")
)

func GetToken(tokenid int) (string, error) {
	token, err := GetUserTokenRedis(tokenid)
	if err == nil {
		return token
	}
	token, err = GetTokenById(tokenid)
	if err != nil {
		return "", err
	}
	SetUserTokenRedis(tokenid, token)
	return token
}
