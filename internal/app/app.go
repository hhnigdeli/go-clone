package app

import (
	"fmt"
	"log"
	"os"

	"com.bosch/rhp-pm/internal/pkg/getGitCredentials"
	"com.bosch/rhp-pm/internal/pkg/parser"
)

func exitOnError(err error, msg string) {
	if err != nil {
		log.Println(msg)
		os.Exit(1)
	}
}

func Run() {
	fmt.Println("Running the app")
	repoURL := "https://sourcecode.socialcoding.bosch.com/scm/rhp/rhp-devops.git"
	username, password, err := getGitCredentials.Get(repoURL)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Username: %s\n", username)
	fmt.Printf("Password: %s\n", password)

	//cloner.Clone("c:/Dev/test2", repoURL, username, password)
	parser.Parse("C:/Dev/Go/rhp-pm/package.json")
}
