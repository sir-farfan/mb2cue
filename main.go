package main

import (
	"encoding/json"
	"errors"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/sir-farfan/mb2cue/model"
)

var uuid_regex = regexp.MustCompile("[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}")

func main() {
	var release string
	gap := int64(0)

	flag.StringVar(&release, "r", "", "Release id to format the cue file")
	flag.Int64Var(&gap, "g", 0, "Add a gap, in milliseconds, between songs")
	flag.Parse()

	if len(release) == 0 {
		log.Println("Must indicate the release")
		flag.Usage()
		os.Exit(1)
	}

	id, err := getId(release)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println(id)

	album, err := getRelease(id)
	if err != nil {
		log.Println(err, album)
	}

	album.FormatCue(gap)
}

func getId(release string) (s string, err error) {
	// whether they enter the whole URL or only the ID, I only care about the release uuid
	id := uuid_regex.Find([]byte(release))
	if len(id) == 0 {
		return "", errors.New("no release ID found")
	}

	return string(id), nil
}

func getRelease(r string) (release model.Release, err error) {
	r = "https://musicbrainz.org/ws/2/release/" + r

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
