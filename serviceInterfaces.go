package main

import (
)

//http request msg struct
type WelcomeMsgReq struct {
	Name string `json:"name"`
}

//response msg struct
type WelcomeMsgRes struct {
	Greeting string `json:"greeting"`
}
