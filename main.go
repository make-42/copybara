package main

import (
	"copybara/config"
	"copybara/listener"
	"copybara/urlclean"
	"fmt"
)

var Banner = `Copybara v0.0.1
₍ᐢ•(ܫ)•ᐢ₎`

func main() {
	fmt.Println(Banner)
	urlclean.Init()
	config.Init()
	listener.ListenerThread()
}
