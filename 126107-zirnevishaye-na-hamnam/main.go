package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
)

const (
	srtRegEx = `^(.*)(s|S)(\d{1,2})(e|E)(\d{1,2}).*\.srt$`
	mkvRegEx = `^(.*)(s|S)(\d{1,2})(e|E)(\d{1,2}).*\.mkv$`
)

func Renamify(path string) {
	srtRegex := regexp.MustCompile(srtRegEx)
	mkvRegEx := regexp.MustCompile(mkvRegEx)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	subtitle_addresses := make(map[string]string)
	for _, file := range files {
		if srtRegex.MatchString(file.Name()) {
			matches := srtRegex.FindStringSubmatch(file.Name())
			season, _ := strconv.Atoi(matches[3])
			episode, _ := strconv.Atoi(matches[5])
			subtitle_addresses[fmt.Sprintf("%d-%d", season, episode)] = file.Name()
		}
	}

	for _, file := range files {
		if mkvRegEx.MatchString(file.Name()) {
			matches := mkvRegEx.FindStringSubmatch(file.Name())
			season, _ := strconv.Atoi(matches[3])
			episode, _ := strconv.Atoi(matches[5])
			subtitle_address := subtitle_addresses[fmt.Sprintf("%d-%d", season, episode)]
			new_name := file.Name()[:len(file.Name())-3] + "srt"
			os.Rename(path+subtitle_address, path+new_name)
		}
	}

}
