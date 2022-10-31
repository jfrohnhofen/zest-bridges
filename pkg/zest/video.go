package zest

import "fmt"

type Video struct {
	dataFile string
}

func NewVideo(dataFile string) Video {
	return Video{dataFile}
}

func (video Video) Play(from, to uint32) {
	cue.setVideo(fmt.Sprintf("play %s %d to %d @ 25 fps\n", video.dataFile, from, to))
}
