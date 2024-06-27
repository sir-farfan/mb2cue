package model

import (
	"fmt"
)

type Track struct {
	Title    string `json:"title"`
	Length   int64  `json:"length"`
	Position int    `json:"position"`
}

func (t Track) FormatCue(offset int64) int64 {

	fmt.Printf("  TRACK %02d AUDIO\n", t.Position)
	fmt.Printf("    TITLE \"%s\"\n", t.Title)
	fmt.Printf("    INDEX 01 %s\n", FormatIndex(offset))
	return offset + t.Length
}

type Media struct {
	Position int     `json:"position"`
	FormatID string  `json:"format-id"`
	Tracks   []Track `json:"tracks"`
}

func (m Media) FormatCue() {
	offset := int64(0)
	fmt.Printf("FILE \"%s\" WAVE\n", m.FormatID)
	for _, t := range m.Tracks {
		offset = t.FormatCue(offset)
	}
}

type Release struct {
	ID    string  `json:"id"`
	Title string  `json:"title"`
	Media []Media `json:"Media"`
}

func (r Release) FormatCue() {
	fmt.Println("REM release id " + r.ID)
	fmt.Printf("TITLE \"%s\"\n", r.Title)
	for _, m := range r.Media {
		m.FormatCue()
	}
}

func FormatIndex(length int64) string {
	// index := time.UnixMilli(length)
	// return index.Format("04:05.999999")

	length /= 10
	ff := length % 100
	if ff >= 70 {
		ff = 69 // whatever this is, shn likes it this way
	}

	length /= 100

	ss := length % 60
	mm := length / 60

	// MM:SS:FF
	return fmt.Sprintf("%02d:%02d:%02d", mm, ss, ff)
}
