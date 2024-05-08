package main

import (
	"context"
	"log"

	apigee "google.golang.org/api/apigee/v1"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()

	service, err := apigee.NewService(ctx, option.WithCredentialsFile("./config.json"))
	if err != nil {
		log.Fatalf("Could not create service: %v", err)
	}

	// export orgnName from config.json     "project_id": "zvdy-helm",

	orgName := "organizations/@org"

	// Create a list request for the APIs in the organization
	req := service.Organizations.Apis.List(orgName)

	// Execute the request
	resp, err := req.Do()
	if err != nil {
		log.Fatalf("Failed to list APIs: %v", err)
	}

	// Print the names of the API proxies
	for _, api := range resp.Proxies {
		log.Println(api.Name)
	}
}
