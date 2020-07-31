package main

import (
	"mmo/game/login"
	"os"
	"strconv"
)

func main()  {
	path,_ := os.Getwd()
	os.Args = []string{path, "-ComponentId", strconv.Itoa(int(222)), "-ListenAddr", "127.0.0.1:14001", ""}

	login := login.NewLoginService()
	login.Run()
}