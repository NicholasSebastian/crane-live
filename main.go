package main

import "fmt"
import "os/exec"

// TODO: Handle multiple camera inputs.
// TODO: Concurrently process the incoming video streams if necessary.

var DEVICE_TYPE = "avfoundation" // NOTE: This should be "dshow" on windows, "v4l2" on linux.
var PRESET = "ultrafast" // Fastest processing time, but probably worst compression.
var FORMAT = "h265"      // NOTE: If this doesn't work, either install the lib it needs or change to h264.

func main() {
	input_device := "0" // TODO: This should be based on which room the client is entering into.
	// TODO: This must be one of the valid input devices configured upon startup.

	bit_rate := "1M" // TODO: These might have to vary depending on the connection quality.
	frame_rate := "30"
	resolution := "640x480"

	response := exec.Command("ffmpeg",
		"-f", DEVICE_TYPE,
		"-i", input_device,
		"-preset", PRESET,
		"-f", FORMAT,
		"-b", bit_rate,
		"-r", frame_rate,
		"-s", resolution,
	)

	stdout, error := response.Output()
	if error != nil {
		fmt.Printf("ERROR: %s\n", error) // TODO: Fix this from erroring status 251.
		return
	}

	fmt.Println(stdout)

	// TODO: Implement an HTTP server to handle the WebRTC handshake.
	// It should also expect to receive the connection quality to adjust for the right quality.
}
