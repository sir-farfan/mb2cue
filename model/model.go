package model

import (
	"fmt"
)

type Track struct {
	Title    string `json:"title"`
	Length   int64  `json:"length"`
	Position int    `json:"position"`
}

func (t Track) FormatCue(offset, gap int64) int64 {
	fmt.Printf("  TRACK %02d AUDIO\n", t.Position)
	fmt.Printf("    TITLE \"%s\"\n", t.Title)
	fmt.Printf("    INDEX 01 %s\n", FormatIndex(offset))
	if gap > 0 && offset > 0 {
		fmt.Printf("    INDEX 02 %s\n", FormatIndex(offset+gap))
	}
	return offset + t.Length + gap
}

type Media struct {
	Position int     `json:"position"`
	FormatID string  `json:"format-id"`
	Tracks   []Track `json:"tracks"`
}

func (m Media) FormatCue(gap int64) {
	offset := int64(0)
	fmt.Printf("FILE \"%s\" WAVE\n", m.FormatID)
	for _, t := range m.Tracks {
		offset = t.FormatCue(offset, gap)
	}
}

type Release struct {
	ID    string  `json:"id"`
	Title string  `json:"title"`
	Media []Media `json:"Media"`
}

func (r Release) FormatCue(gap int64) {
	fmt.Println("REM release id " + r.ID)
	fmt.Printf("TITLE \"%s\"\n", r.Title)
	for _, m := range r.Media {
		m.FormatCue(gap)
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
