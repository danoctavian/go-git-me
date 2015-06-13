package main

import (
  "fmt"
  "io/ioutil"
  "github.com/google/go-github/github"
  "encoding/json"
  "log"
  "os/user"
)

func main() {
	config := loadConfiguration()
	fmt.Println(config)

	return

  fmt.Printf("Hello, world.\n")
  client := github.NewClient(nil)

  // opt := &github.ListOptions{Type: "public"}
	users, _, _ := client.Activity.ListStargazers("danoctavian", "bluntly", nil)
	fmt.Printf("%v", len(users))
}

type Configuration struct {
	githubAPIToken string
}

type Repository struct {
	stargazers []github.User
}

type ConfigJSON map[string]string

func loadConfiguration() Configuration {
  usr, err := user.Current()
  if err != nil {
      log.Fatal( err )
  }

  var configJSON ConfigJSON
  file, err := ioutil.ReadFile(usr.HomeDir + "/.go-git-me/githubAPIAuth.json")
  if err != nil {
      log.Fatal(err)
  }
  err = json.Unmarshal(file, &configJSON)
  if err != nil {
      log.Fatal(err)
  }
  
  return Configuration{githubAPIToken: configJSON["APIToken"]}
}

// computes
// * popularity rank 
//
// * stargazer stats:
// 		-- distribution by location
//    -- distribution by expertise
// *  
func computeRepoStats() {
}

