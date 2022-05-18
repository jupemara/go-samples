package main

import (
	"context"
	"io"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2/google"
	admin "google.golang.org/api/admin/reports/v1"
)

func main() {
	ctx := context.TODO()
	// credential.json: service account json key
	b, err := ioutil.ReadFile("credential.json")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	// previous
	j, err := google.JWTConfigFromJSON(b, admin.AdminReportsAuditReadonlyScope)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	// needs to replace with your workspace domain email
	j.Subject = "workspace-user@example.com"
	// workspace audit log API docs: https://developers.google.com/admin-sdk/reports/v1/appendix/activity/data-studio
	res, err := j.Client(ctx).Get("https://admin.googleapis.com/admin/reports/v1/activity/users/all/applications/data_studio?eventName=EDIT&maxResults=10")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	log.Printf("reports API response: %s", body)
}
