package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)


func parseEpisodeName(origin string) string {

	s := strings.Split(origin, "/")

	str := s[len(s)-1]

	pattern := regexp.MustCompile("[0-9]+")

	out := pattern.FindAllString(str, -1)

	return fmt.Sprintf("s%se%s.mp4", out[0], out[1])
	
} // parseEpisodeName


func era() {

	for {

		log.Printf("jobs: %d", len(jobs))

		if len(jobs) == 0 {
			wg.Done()
			break
		}

		select {
		case j := <-jobs:
			download(j.Url, j.Origin)
		}

	}

} // era


func download(origin string, url string) {

	if url == "" || origin == "" {
		return
	}

	if !strings.Contains(url, "ova") && !strings.Contains(url, "movie") {

		name := parseEpisodeName(url)

		res, err := http.Get(origin)

		if err != nil {
			log.Println(err)
		} else {

			defer res.Body.Close()

			f, err := os.Create(name)

			defer f.Close()

			if err != nil {
				log.Println(err)
			} else {
				io.Copy(f, res.Body)
			}

		}		

	}

} // downloader
