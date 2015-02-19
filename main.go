package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	projectName := os.Args[1]
	createGitRepo(projectName)
	numCommits := 2
	createCommits(numCommits)
}

func createCommits(numCommits int) {

	// defer f.Close()
	for i := 0; i < numCommits; i++ {
		f, err := os.OpenFile("README.md", os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		f.Write([]byte{'1'})
		buff := make([]byte, 50)
		_, err = f.Read(buff)
		f.Close()
		if err != nil {
			panic(err)
		}
		fmt.Println("README contents: ", buff)
		err = exec.Command("/usr/bin/git", "add", ".").Run()
		if err != nil {
			panic(err)
		}
		err = exec.Command("/usr/bin/git", "commit", "-m", "\"Poof\"").Run()
		if err != nil {
			panic(err)
		}
	}
}

func createGitRepo(name string) {

	// Create new project
	os.Mkdir("./projects/"+name, 0666)
	os.Chdir("./projects/" + name)
	f, err := os.Create("README.md")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.Write([]byte{'1'})
	if err != nil {
		panic(err)
	}
	err = exec.Command("/usr/bin/git", "init").Run()
	if err != nil {
		panic(err)
	}
}
