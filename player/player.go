package player

import (
	"fmt"
	"github.com/200sc/klangsynthese/audio"
	"time"
)

type Note struct {
	audio.Audio
	start int64
}

func NewNote(a audio.Audio, start int64) Note {
	return Note{
		a,
		start,
	}
}

func PlaySequence(notes []Note) {
	timeElapsed := int64(0)
	for i, n := range notes {
		n.Play()
		fmt.Println(n.PlayLength())
		time.Sleep(n.PlayLength())
		timeElapsed += n.PlayLength().Milliseconds()
		if i < len(notes)-1 {
			nextNote := notes[i+1]
			time.Sleep(time.Duration(nextNote.start-timeElapsed) * time.Millisecond)
		}
	}
}
