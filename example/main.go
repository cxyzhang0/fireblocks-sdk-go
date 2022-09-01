package main

import (
	SDK "fireblocks-sdk"
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	var secretFile string
	flag.StringVar(&secretFile, "s", "", "path to key file")
	//secretFilePtr := flag.String("s", "../../../certs/fireblocks_secret.key", "path to secret key file")
	var apiKey string
	flag.StringVar(&apiKey, "k", "", "API Key")
	//apiKeyPtr := flag.String("k", "98fe77cb-b459-6508-a945-b01cdda83576", "API Key")
	var baseUrl string
	flag.StringVar(&baseUrl, "url", "https://api.fireblocks.io", "API Key")
	var timeout int
	flag.IntVar(&timeout, "timeout", 5, "timeout in millisecond")
	flag.Parse()
	log.Infoln("secret file: ", secretFile)
	//log.Infoln("api key: ", apiKey)
	log.Infoln("base url: ", baseUrl)
	log.Infoln("timeout (ms): ", timeout)
	privK, err := os.ReadFile(secretFile)
	check(err)
	//log.Infoln("secret key: ", string(privK))

	sdk := SDK.NewInstance(privK, apiKey, baseUrl, time.Duration(timeout)*time.Millisecond)

	// all users
	users, err := sdk.GetUsers()
	check(err)
	fmt.Println("***** Users:")
	for i, user := range users {
		fmt.Println(i, " - ", user.LastName, ", ", user.FirstName)
	}
	/*
		// supported assets - too many
		supportedAssets, err := sdk.GetSupportedAssets()
		check(err)
		fmt.Println("***** Supported Assets:")
		for i, asset := range supportedAssets {
			fmt.Println(i, " - ", asset.AssetType, ":", asset.Name)
		}

	*/

}
