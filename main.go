package main

import "fmt"
import "github.com/google/go-github/github"

func main() {
  fmt.Printf("Hello, world.\n")
  client := github.NewClient(nil)

  opt := &github.RepositoryListByOrgOptions{Type: "public"}
	repos, _, err := client.Repositories.ListByOrg("github", opt)
}

