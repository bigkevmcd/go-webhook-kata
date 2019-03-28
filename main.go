package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", hookHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hookHandler(w http.ResponseWriter, r *http.Request) {
	githubEvent := r.Header.Get("X-GitHub-Event")
	if githubEvent != "deployment_status" {
		fmt.Fprintf(w, "ignored")
		return
	}

	decoder := json.NewDecoder(r.Body)
	var event DeploymentStatusEvent
	if err := decoder.Decode(&event); err != nil {
		log.Printf("error decoding JSON: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	message := fmt.Sprintf("There has been a new deployment of %s to %s ref '%s' with state %s",
		event.Repository.Fullname,
		event.Deployment.Environment,
		event.Deployment.Ref,
		event.DeploymentStatus.State)
	title := fmt.Sprintf("Deployment of %s", event.Repository.Fullname)

	sendNotificationToPushover(title, message, os.Getenv("NOTIFICATION_USER"), os.Getenv("PUSHOVER_TOKEN"))
	fmt.Fprintf(w, message)
}

func sendNotificationToPushover(title string, message string, user string, token string) (string, error) {
	notification := &notification{
		Token:   token,
		User:    user,
		Message: message,
		Title:   title,
	}
	b, err := json.Marshal(notification)
	if err != nil {
		return "", err
	}
	res, err := http.Post("https://api.pushover.net/1/messages.json", "application/json", bytes.NewBuffer(b))
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	output, err := ioutil.ReadAll(res.Body)

	return string(output), nil
}

type notification struct {
	Token   string `json:"token"`
	User    string `json:"user"`
	Message string `json:"message"`
	Title   string `json:"title"`
}

type DeploymentStatusEvent struct {
	Repository struct {
		Fullname string `json:"full_name"`
	} `json:"repository"`

	DeploymentStatus struct {
		State       string `json:"state"`
		Description string `json:"description"`
	} `json:"deployment_status"`

	Deployment struct {
		Ref         string `json:"ref"`
		Environment string `json:"environment"`
	} `json:"deployment"`
}
