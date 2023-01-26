package cmd

import (
	"droidsh/utils"
	"fmt"
	"github.com/chzyer/readline"
	"net"
	"os"
)

func Location(conn net.Conn, rl *readline.Instance, args []string) {
	appId := os.Getenv("APPID")
	if appId == "" {
		fmt.Println("APPID env var not set.")
		return
	}

	location, err := utils.FetchLocation(appId, args[0])

	if err != nil {
		fmt.Println("ERROR: failed to fetch location data:", err)
		return
	}

	fmt.Println("Name:", location.Name)
	fmt.Println("Latitude:", location.Latitude)
	fmt.Println("Longitude:", location.Longitude)
	fmt.Println("Country:", location.Country)
	fmt.Println("State:", location.State)
	fmt.Println("URL:", fmt.Sprintf("https://www.google.com/maps/@%f,%f,17z", location.Latitude, location.Longitude))

}
