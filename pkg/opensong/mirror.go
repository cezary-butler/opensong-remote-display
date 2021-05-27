package opensong

import (
	"bitbucket.org/cezary_butler/opensong-remote-display/pkg/display/fyne"
	"bitbucket.org/cezary_butler/opensong-remote-display/pkg/xml"
	"fmt"
	"golang.org/x/net/websocket"
	"image"
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

//TODO find a way to safely reconnect after connection error (detect eof)
//TODO handle sever not in presentation mode
//TODO add context support
//FIXME on first receive no image is being displayed
//FIXME images are different for what is on the screen
//TODO find elegant way to handle subscription in the golang

func (s *OpenSong) InitMirroring() {
	conn, err := websocket.Dial(fmt.Sprintf("ws://%s/ws", s.host), "", "http://localhost")

	if err != nil {
		panic(err)
	}
	//defer closeRes(conn) //TODO find a way to close the connection at just the right moment (context might be useful)

	go s.subscribeEvents(conn)
	go s.listenForConfirmation(conn)

}

func (s *OpenSong) listenForConfirmation(conn *websocket.Conn) {
	response := ""

	log.Print("waiting for OK")
	for {
		err := websocket.Message.Receive(conn, &response)

		if err != nil {
			log.Print("error while waiting for ok", err)
			continue
		}
		//TODO get currently presented image
		if response != "" {
			log.Print("Received ", response, "listening for future events")
			s.listenForEvents(conn)
		}
	}
}

func (s *OpenSong) listenForEvents(conn *websocket.Conn) {
	for {
		log.Print("going to receive")

		mgs := Response{}
		//mgs := ""
		err := xml.XML.Receive(conn, &mgs)
		if err != nil {
			log.Printf("received error %+e", err)
			continue
		}
		log.Printf("received %d (%s:%s)\n", mgs.Presentation.Slide.ItemNumber,
			mgs.Presentation.Slide.Name, mgs.Presentation.Slide.Title)
		uriStr := fmt.Sprintf("http://%s/presentation/slide/%d/preview/height:%d/widht:%d/quality:%d",
			s.host, mgs.Presentation.Slide.ItemNumber, s.height, s.width, s.quality)

		image := s.fetchImage(uriStr)

		s.display.DisplayImage(image)
	}
}

func (s *OpenSong) fetchImage(uriStr string) image.Image {
	img, err := http.Get(uriStr)
	if err != nil {
		log.Printf("received error while getting image %+e", err)
		return nil
	}
	//todo func
	defer closeRes(img.Body)

	all, err := jpeg.Decode(img.Body)
	if err != nil {
		log.Printf("received error while reading image %+e", err)
		return nil
	}
	return all

}

func (s *OpenSong) subscribeEvents(conn *websocket.Conn) error {

	log.Print("will send subscription", subscribe)

	err := websocket.Message.Send(conn, subscribe)
	if err != nil {
		log.Print("error while subscribing", err)
		return err
	}

	log.Print("will show and run")
	return nil
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
