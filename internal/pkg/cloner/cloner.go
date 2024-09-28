package cloner

import (
	"log"

	git "github.com/go-git/go-git/v5"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

func Clone(dir string, scmUrl string, username string, password string) {

	// Clones the repository into the given dir, just as a normal git clone does
	_, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL: scmUrl, // Update with your Bitbucket repo URL
		Auth: &http.BasicAuth{
			Username: username, // Bitbucket username
			Password: password, // Your personal access token
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Repository cloned successfully!")

}
