package main

import (
	"bufio"
	"go-synth-player/parser"
	"go-synth-player/player"
	"os"
)

func main() {
	lines := readLines("notes.txt")
	notes := parser.ParseMusic(lines)
	player.PlaySequence(notes)
}

func readLines(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
