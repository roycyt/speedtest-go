package main

import (
	"flag"
	"log"
	"os"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

var (
	showList = flag.Bool("list", false, "Show available speedtest.net servers.")
	serverID ServerIDList
)

func init() {
	flag.Var(&serverID, "server", "Select servers to speedtest. List the server ID separated by comma.")
	flag.Parse()
}

func main() {
	user := fetchUserInfo()
	user.Show()

	list := fetchServerList(user)
	if *showList {
		list.Show()
		return
	}

	targets := list.FindServer(serverID)
	targets.StartTest()
	targets.ShowResult()
}
