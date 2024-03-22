package main

import (
	"bufio"
)

// TODO: Handle multiple camera inputs.
// TODO: Concurrently process the incoming video streams if necessary.

func main() {
	ffmpegConfig := FfmpegConfig{
		// TODO: This should be based on which room the client is entering into.
		// TODO: This must be one of the valid input devices configured upon startup.
		input_device: "0",

		// TODO: These might have to vary depending on the connection quality.
		bit_rate:   "1M",
		frame_rate: "30",
		resolution: "640x480",
	}

	stdout := *ffmpeg(ffmpegConfig)
	defer stdout.Close()

	// TODO: Process the bytes out of stdout to be streamed though WebRTC later.
	reader := bufio.NewReader(stdout)

	// TODO: Implement an HTTP server to handle the WebRTC handshake.
	// It should also expect to receive the connection quality to adjust for the right quality.
}
