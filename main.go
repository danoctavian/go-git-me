package main

import (
  "fmt"
  "bufio"
  "io/ioutil"
  "github.com/google/go-github/github"
  "encoding/json"
  "log"
  "os"
  "os/user"
  "golang.org/x/oauth2"
  ui "github.com/gizak/termui"
)

func main() {
	runUI()
	config := loadConfiguration()
	fmt.Println(config)

  fmt.Printf("Hello, world.\n")

	ts := oauth2.StaticTokenSource(
    &oauth2.Token{AccessToken: config.githubAPIToken},
  )

  tc := oauth2.NewClient(oauth2.NoContext, ts)

  client := github.NewClient(tc)

  opt := &github.ListOptions{Page: 1, PerPage: 200}
	users, _, _ := client.Activity.ListStargazers("danoctavian", "bluntly", opt)

	fmt.Printf("%v", len(users))

	computeRepoStats(Repository{stargazers: users})
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
func computeRepoStats(repo Repository) {
	fmt.Println("computing...")
	fmt.Println(repo.stargazers)


}

func runUI() {
	 err := ui.Init()
	  if err != nil {
	      panic(err)
	  }
	  defer ui.Close()

	  p := ui.NewPar(":PRESS q TO QUIT DEMO")
	  p.Height = 3
	  p.Width = 50
	  p.TextFgColor = ui.ColorWhite
	  p.Border.Label = "Text Box"
	  p.Border.FgColor = ui.ColorCyan



	bc := ui.NewBarChart()
	data := []int{3, 2, 5, 3, 9, 5, 3, 2, 5, 8, 3, 2, 4, 5, 3, 2, 5, 7, 5, 3, 2, 6, 7, 4, 6, 3, 6, 7, 8, 3, 6, 4, 5, 3, 2, 4, 6, 4, 8, 5, 9, 4, 3, 6, 5, 3, 6}
	bclabels := []string{"S0", "S1", "S2", "S3", "S4", "S5"}
	bc.Border.Label = "Bar Chart"
	bc.Data = data
	bc.Width = 26
	bc.Height = 10
	bc.DataLabels = bclabels
	bc.TextColor = ui.ColorGreen
	bc.BarColor = ui.ColorRed
	bc.NumColor = ui.ColorYellow

	  ui.Render(p, bc)

	  reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
}

