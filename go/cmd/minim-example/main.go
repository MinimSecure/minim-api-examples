package main

import (
	"log"
	"os"

	minim "github.com/MinimSecure/minim-api-examples/go"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalln("expected two args, application ID and secret")
	}
	appID := os.Args[1]
	secret := os.Args[2]
	cl := minim.New(appID, secret)

	isps, err := cl.GetIDs("/api/v1/isps")
	if err != nil {
		log.Fatalln("failed to get isps:", err)
	}
	if len(isps) == 0 {
		log.Fatalln("no isps are available for your user")
	}
	unums, err := cl.MultiGet("/api/v1/isps/"+isps[0]+"/unums", nil)
	if err != nil {
		log.Fatalln("failed to get unums:", err)
	}
	for _, u := range unums {
		log.Println(u["lan_mac_address"], "->", u)
	}
}
