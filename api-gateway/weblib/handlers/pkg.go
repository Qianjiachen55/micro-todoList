package handlers

import (
	"api-gateway/pkg/logging"
	"errors"
)

// error

func PanicIfUserError(err error)  {
	if err != nil{
		err = errors.New("userService--"+err.Error())
		logging.Info(err)
		panic(err)
	}
}