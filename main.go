package main

import (
	"github.com/sethvargo/go-githubactions"
)

func main() {
	val := githubactions.GetInput("name")
	if val == "" {
		githubactions.Fatalf("missing 'val'")
	}
}
