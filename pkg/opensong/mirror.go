package opensong

import (
	"bitbucket.org/cezary_butler/opensong-remote-display/pkg/display/fyne"
	"bitbucket.org/cezary_butler/opensong-remote-display/pkg/xml"
	"fmt"
	"golang.org/x/net/websocket"
	"image/jpeg"
	"io"
	"log"
	"net/http"
)

type OpenSong struct {
	display *fyne.Display
	width   int
	height  int
	quality int
	host    string
}

const subscribe = "/ws/subscribe/presentation"

func (s *OpenSong) InitMirroring() {
	conn, err := websocket.Dial(fmt.Sprintf("ws://%s/ws", s.host), "", "http://localhost")

	if err != nil {
		panic(err)
	}
	defer closeRes(conn)

	go func() {
		response := ""
		for {
			log.Print("waiting for OK")
			err = websocket.Message.Receive(conn, &response)
			if err != nil {
				log.Print("error while waiting for ok", err)
				continue
			}
			if response != "" {
				log.Printf("received %s, waiting for subscribed messages", response)
				go func() {
					for { //todo check context
						log.Print("going to receive")

						mgs := Response{}
						//mgs := ""
						err = xml.XML.Receive(conn, &mgs)
						if err != nil {
							log.Printf("received error %+e", err)
							continue
						}
						log.Printf("received %d (%s:%s)\n", mgs.Presentation.Slide.ItemNumber,
							mgs.Presentation.Slide.Name, mgs.Presentation.Slide.Title)
						uriStr := fmt.Sprintf("http://%s/presentation/slide/%d/preview/height:%d/widht:%d/quality:%d",
							s.host, mgs.Presentation.Slide.ItemNumber, s.height, s.width, s.quality)

						func() {
							img, err := http.Get(uriStr)
							if err != nil {
								log.Printf("received error while getting image %+e", err)
								return
							}
							//todo func
							defer closeRes(img.Body)

							all, err := jpeg.Decode(img.Body)
							if err != nil {
								log.Printf("received error while reading image %+e", err)
								return
							}

							s.display.DisplayImage(all)
						}()
						//log.Printf("received %s", mgs)
					}
				}()
				return
			}
		}
	}()

	go func() {

		log.Print("will send subscription", subscribe)

		err = websocket.Message.Send(conn, subscribe)
		if err != nil {
			log.Print("error while subscribing")
			panic(err)
		}

		log.Print("will show and run")
	}()
}

func NewOpenSong(display *fyne.Display, host string, width, height, quality int) *OpenSong {
	return &OpenSong{
		display: display,
		host:    host,
		width:   width, height: height, quality: quality,
	}
}

func closeRes(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Printf("error occurred while closing resource: %v\n", err)
	}
}
