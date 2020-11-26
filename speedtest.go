package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net/http"
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
	insecure = flag.Bool("insecure", false, "Skip server certificate verification.")

	httpGet func(url string) (resp *http.Response, err error)
)

func init() {
	flag.Var(&serverID, "server", "Select servers to speedtest. List the server ID separated by comma.")
	flag.Parse()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: *insecure},
	}
	httpGet = (&http.Client{Transport: tr}).Get
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
