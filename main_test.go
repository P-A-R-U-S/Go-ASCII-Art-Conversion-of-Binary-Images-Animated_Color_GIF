package main

import (
	"os"
	"testing"
)

func Test_getFrames(t *testing.T) {
	testDatas := []struct {
		imageName      string
		numberOfFrames int
	}{
		{"1.gif", 25},
		{"2.gif", 4},
		{"3.gif", 8},
	}

	for _, td := range testDatas {

		// read file
		reader, err := os.Open("./GIF-Images/" + td.imageName)
		if err != nil {
			t.Errorf("failed due to error %s", err)
		}
		pf, err := getFrames(reader)

		// check the number of frames
		if len(pf) != td.numberOfFrames {
			t.Fail()
		}
	}
}
