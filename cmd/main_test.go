package main

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/websocket"
	"testing"
)

func Test_messageCanBeUnmarshaled(t *testing.T) {

	var result Response
	msg := "<?xml version=\"1.0\" encoding=\"UTF-8\"?><response resource=\"presentation\" action=\"status\"><presentation running=\"1\"><screen mode=\"N\">normal</screen><slide itemnumber=\"12\"><name>And Can It Be</name><title>And Can It Be</title></slide></presentation></response>"
	xmlUnmarshal([]byte(msg), websocket.TextFrame, &result)

	assert.Equal(t, 12, result.Presentation.Slide.ItemNumber)

}
