package models

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"testing"
)

func TestSetGet(t *testing.T) {
	_, err := GetUserTokenRedis(1234)
	//T.Log(token)
	//t.Log(err.Error())
	if err == redis.ErrNil {
		t.Log("Error OK")
	}

	err = SetUserTokenRedis(1234, "UGUESS")
	if err != nil {
		t.Error(err.Error())
	}

	token, err := GetUserTokenRedis(1234)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("--------------------" + token)
	t.Log("--------------------" + token)

}
