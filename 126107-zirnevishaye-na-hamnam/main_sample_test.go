package main

import (
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func copy(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	return
}

func listDirectory(folder string) (dir []string) {
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		dir = append(dir, f.Name())
	}
	return
}

func TestSample(t *testing.T) {
	folder := "./serialFolder/"
	os.RemoveAll(folder)
	err := os.Mkdir(folder, 0777)
	if err != nil {
		panic(err)
	}

	err = copy("video.mkv", folder+"Serial1.S01E01.mkv")
	if err != nil {
		panic(err)
	}
	err = copy("sub.srt", folder+"S01E01-sub.srt")
	if err != nil {
		panic(err)
	}
	Renamify(folder)
	assert.ElementsMatch(t,
		[]string{"Serial1.S01E01.mkv", "Serial1.S01E01.srt"},
		listDirectory(folder),
	)
}
