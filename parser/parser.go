package parser

import (
	"fmt"
	"github.com/200sc/klangsynthese/audio"
	"github.com/200sc/klangsynthese/synth"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	notes = map[string]synth.Pitch{
		"C0":  synth.C0,
		"C0s": synth.C0s,
		"D0":  synth.D0,
		"D0s": synth.D0s,
		"E0":  synth.E0,
		"F0":  synth.F0,
		"F0s": synth.F0s,
		"G0":  synth.G0,
		"G0s": synth.G0s,
		"A0":  synth.A0,
		"A0s": synth.A0s,
		"B0":  synth.B0,
		"C1":  synth.C1,
		"C1s": synth.C1s,
		"D1":  synth.D1,
		"D1s": synth.D1s,
		"E1":  synth.E1,
		"F1":  synth.F1,
		"F1s": synth.F1s,
		"G1":  synth.G1,
		"G1s": synth.G1s,
		"A1":  synth.A1,
		"A1s": synth.A1s,
		"B1":  synth.B1,
		"C2":  synth.C2,
		"C2s": synth.C2s,
		"D2":  synth.D2,
		"D2s": synth.D2s,
		"E2":  synth.E2,
		"F2":  synth.F2,
		"F2s": synth.F2s,
		"G2":  synth.G2,
		"G2s": synth.G2s,
		"A2":  synth.A2,
		"A2s": synth.A2s,
		"B2":  synth.B2,
		"C3":  synth.C3,
		"C3s": synth.C3s,
		"D3":  synth.D3,
		"D3s": synth.D3s,
		"E3":  synth.E3,
		"F3":  synth.F3,
		"F3s": synth.F3s,
		"G3":  synth.G3,
		"G3s": synth.G3s,
		"A3":  synth.A3,
		"A3s": synth.A3s,
		"B3":  synth.B3,
		"C4":  synth.C4,
		"C4s": synth.C4s,
		"D4":  synth.D4,
		"D4s": synth.D4s,
		"E4":  synth.E4,
		"F4":  synth.F4,
		"F4s": synth.F4s,
		"G4":  synth.G4,
		"G4s": synth.G4s,
		"A4":  synth.A4,
		"A4s": synth.A4s,
		"B4":  synth.B4,
		"C5":  synth.C5,
		"C5s": synth.C5s,
		"D5":  synth.D5,
		"D5s": synth.D5s,
		"E5":  synth.E5,
		"F5":  synth.F5,
		"F5s": synth.F5s,
		"G5":  synth.G5,
		"G5s": synth.G5s,
		"A5":  synth.A5,
		"A5s": synth.A5s,
		"B5":  synth.B5,
		"C6":  synth.C6,
		"C6s": synth.C6s,
		"D6":  synth.D6,
		"D6s": synth.D6s,
		"E6":  synth.E6,
		"F6":  synth.F6,
		"F6s": synth.F6s,
		"G6":  synth.G6,
		"G6s": synth.G6s,
		"A6":  synth.A6,
		"A6s": synth.A6s,
		"B6":  synth.B6,
		"C7":  synth.C7,
		"C7s": synth.C7s,
		"D7":  synth.D7,
		"D7s": synth.D7s,
		"E7":  synth.E7,
		"F7":  synth.F7,
		"F7s": synth.F7s,
		"G7":  synth.G7,
		"G7s": synth.G7s,
		"A7":  synth.A7,
		"A7s": synth.A7s,
		"B7":  synth.B7,
		"C8":  synth.C8,
		"C8s": synth.C8s,
		"D8":  synth.D8,
		"D8s": synth.D8s,
		"E8":  synth.E8,
		"F8":  synth.F8,
		"F8s": synth.F8s,
		"G8":  synth.G8,
		"G8s": synth.G8s,
		"A8":  synth.A8,
		"A8s": synth.A8s,
		"B8":  synth.B8,
	}
)

func ParseMusic(music string) []audio.Audio {
	fmt.Println("Start parsing")
	notes := strings.Split(music, "-")
	a := make([]audio.Audio, len(notes))
	for i, s := range notes {
		a[i] = parseNote(s)
	}

	fmt.Println("Parsing success")
	return a
}

func parseNote(note string) audio.Audio {
	r, _ := regexp.Compile("^([A-G][0-8][b|s]?)\\[([0-9]+)\\]$")
	matches := r.FindStringSubmatch(note)
	if len(matches) < 3 {
		fmt.Printf("Invalid note %s\n", note)
	}

	fmt.Printf("Note %s duration %s ms\n", matches[1], matches[2])

	pitch := synth.AtPitch(notes[matches[1]])
	durationMs, err := strconv.ParseInt(matches[2], 10, 0)
	if err != nil {
		panic(err)
	}

	duration := synth.Duration(time.Duration(durationMs) * time.Millisecond)
	a, err := synth.Int16.Saw(pitch, duration)
	if err != nil {
		panic(err)
	}

	return a
}
