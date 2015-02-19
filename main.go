package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"time"
)

var (
	git = "/usr/bin/git"
)

func main() {
	projectName := os.Args[1]
	message := os.Args[2]
	createGitRepo(projectName)
	timestamps := getTimestamps(message)
	formatCommitDates(timestamps)
}

func formatCommitDates(timestamps []int64) {
	// Loop over each timestamp to create 10 commits at it
	for i, t := range timestamps {
		// Loop over the corresponding 10 commits we need to edit
		for j := i * 20; j < (i*20)+20; j++ {
			createCommit()
			date := time.Unix(t, 0)
			formattedDate := date.Format("Mon Jan 2 15:04:05 2006 -0700")
			c := exec.Command(git, "commit", "-m", "\"poof\"", "--amend", "--date", "\""+formattedDate+"\"")
			c.Env = []string{"GIT_COMMITTER_DATE=\"" + formattedDate + "\""} //os.Environ()
			err := c.Run()
			if err != nil {
				panic(err)
			}
		}
	}
}

func getTimestamps(message string) []int64 {
	holder := []int64{}
	c := exec.Command("node", "./../../graffiti/index.js", message)
	data, err := c.Output()
	if err != nil {
		panic(err)
	}
	json.Unmarshal(data, &holder)
	return holder
}

func createCommit() {
	f, err := os.OpenFile("README.md", os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString("0")
	err = exec.Command(git, "add", ".").Run()
	if err != nil {
		panic(err)
	}
	err = exec.Command(git, "commit", "-m", "\"Poof\"").Run()
	if err != nil {
		panic(err)
	}
}

func createGitRepo(name string) {

	// Create new project
	os.Mkdir("./projects/"+name, 0777)
	os.Chdir("./projects/" + name)
	f, err := os.Create("README.md")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.WriteString("1")
	if err != nil {
		panic(err)
	}
	err = exec.Command(git, "init").Run()
	if err != nil {
		panic(err)
	}
}
