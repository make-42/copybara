package main

import (
	"copybara/config"
	"copybara/ipc"
	"copybara/listener"
	"copybara/urlclean"
	"flag"
	"fmt"
	"os"
)

var Banner = `Copybara v0.0.1
₍ᐢ•(ܫ)•ᐢ₎`

func main() {
	fmt.Println(Banner)

	togglePtr := flag.Bool("toggle", false, "toggle running copybara instance through IPC")
	flag.Parse()
	if *togglePtr {
		ipc.SendToggleCommand()
		os.Exit(0)
	}

	ipc.Init()
	urlclean.Init()
	config.Init()
	listener.ListenerThread()
}
