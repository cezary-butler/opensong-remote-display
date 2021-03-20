package main

import (
	"bitbucket.org/cezary_butler/opensong-remote-display/pkg/display/fyne"
	"bitbucket.org/cezary_butler/opensong-remote-display/pkg/opensong"
	"flag"
	"log"
)

var (
	width   = flag.Int("width", 1280, "width of the screen in pixels, max=4096")
	height  = flag.Int("height", 720, "height of the screen in pixels, max=4096")
	host    = flag.String("host", "localhost:8082", "hostname and port where the OpenSong API can be found")
	quality = flag.Int("quality", 96, "image quality passed to OpenSongAPI, has to be in the range 0-100")
)

func main() {
	flag.Parse()

	validate()
	log.Printf("Running with width: %d, height %d, host: %s, quality: %d", *width, *height, *host, *quality)

	display := fyne.InitDisplay(*width, *height)

	os := opensong.NewOpenSong(display, *host, *width, *height, *quality)

	os.InitMirroring()

	display.ShowAndRun()

}

func validate() {
	if *width < 1 || *width > 4096 {
		log.Fatalf("got width %d but it has to be between 1 and 4096", *width)
	}

	if *height < 1 || *height > 4096 {
		log.Fatalf("got height %d but it has to be between 1 and 4096", *height)
	}

	if *quality < 0 || *quality > 100 {
		log.Fatalf("got quality %d but it has to be between 0 and 100", *quality)
	}
}
