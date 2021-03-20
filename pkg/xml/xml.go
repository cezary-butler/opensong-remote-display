package xml

import (
	"encoding/xml"
	"golang.org/x/net/websocket"
)

func xmlMarshal(v interface{}) (msg []byte, payloadType byte, err error) {
	msg, err = xml.Marshal(v)
	return msg, websocket.TextFrame, err
}

func xmlUnmarshal(msg []byte, payloadType byte, v interface{}) (err error) {
	return xml.Unmarshal(msg, v)
}

var XML = websocket.Codec{xmlMarshal, xmlUnmarshal}
