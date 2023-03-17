package main

import (
	"fmt"
	"go-tabs-player/parser"
	"os"
	"time"
)

func main() {
	b, err := os.ReadFile("notes.txt")
	if err != nil {
		panic(err)
	}

	str := string(b)
	a := parser.ParseMusic(str)
	for _, s := range a {
		s.Play()
		fmt.Println(s.PlayLength())
		time.Sleep(s.PlayLength())
	}
}
