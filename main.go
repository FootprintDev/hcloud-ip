package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/hetznercloud/hcloud-go/hcloud"
)

func main() {

	apiKey := flag.String("key", "HCloud API Key", "-key xaxaxaxax")
	floatIP := flag.String("ip", "Name of Floating IP", "-ip voip")
	flag.Parse()

	// Check for API Key
	if *apiKey == "" {
		log.Fatalf("No API Key specified!")
	}
	// Check for FloatingIP
	if *floatIP == "" {
		log.Fatalf("No Floating IP specified!")
	}

	// Get System Hostname
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	// Initialize HCloud Client
	client := hcloud.NewClient(hcloud.WithToken(*apiKey))

	// Try Get Server by Name
	server, _, err := client.Server.GetByName(context.Background(), name)

	if err != nil {
		// Request Error
		log.Fatalf("error retrieving server: %s\n", err)
	}
	if server != nil {
		// Server Found
		response := fmt.Sprintf("Server called %v was found\n", server.Name)
		fmt.Println(response)

		ip, _, err := client.FloatingIP.Get(context.Background(), *floatIP)

		if err != nil {
			// Request Error
			log.Fatalf("error retrieving floating ip: %s\n", err)
		}
		// IP Found
		if ip != nil {
			// Assign IP
			_, res, err := client.FloatingIP.Assign(context.Background(), ip, server)
			if err != nil {
				// Request Error
				log.Fatalf("error assigning floating ip: %s\n", res.Body)
			}
		} else {
			// IP Not Found
			response := fmt.Sprintf("ip with name %v was not found!", *floatIP)
			fmt.Println(response)
		}

	} else {
		// Server Not Found
		response := fmt.Sprintf("server with name %v was not found!", name)
		fmt.Println(response)
	}
}
