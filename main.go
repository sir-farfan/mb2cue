package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/sir-farfan/mb2cue/model"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("Must pass the release URL")
		os.Exit(1)
	}

	release := os.Args[1]

	album, err := getRelease(release)
	if err != nil {
		log.Println(err, album)
	}

	album.FormatCue()
}

func getRelease(r string) (release model.Release, err error) {
	r = strings.Replace(r, "musicbrainz.org/release", "musicbrainz.org/ws/2/release", 1)

	req, err := http.NewRequest(http.MethodGet, r, nil)
	if err != nil {
		return
	}

	q := req.URL.Query()
	q.Add("inc", "aliases+artist-credits+labels+discids+recordings")
	q.Add("fmt", "json")

	req.URL.RawQuery = q.Encode()

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &release)

	return
}
