package main

import (
	"io"
	"log"
	"os/exec"
	"runtime"
)

type FfmpegConfig struct {
	input_device string
	bit_rate     string
	frame_rate   string
	resolution   string
}

var PRESET = "ultrafast" // Fastest processing time, but probably worst compression.
var FORMAT = "h265"      // Just use 'h264' if this doesn't work.

var DEVICE_TYPES = map[string]string{
	"windows": "dshow",
	"darwin":  "avfoundation",
	"linux":   "v412",
}

func ffmpeg(config FfmpegConfig) *io.ReadCloser {
	device_type, os_ok := DEVICE_TYPES[runtime.GOOS]
	if !os_ok {
		log.Fatalln("ERROR: Unknown device type.")
	}

	response := exec.Command("ffmpeg",
		"-f", device_type,
		"-i", config.input_device,
		"-preset", PRESET,
		"-c:v", FORMAT,
		"-b", config.bit_rate,
		"-r", config.frame_rate,
		"-s", config.resolution,
		"pipe:1",
	)

	stdout, error := response.StdoutPipe()
	if error != nil {
		log.Fatalln(error)
	}
	return &stdout
}
