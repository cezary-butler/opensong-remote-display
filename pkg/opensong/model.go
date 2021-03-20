package opensong

type Response struct {
	Presentation Presentation `xml:"presentation"`
}

type Presentation struct {
	Running int   `xml:"running,attr"`
	Slide   Slide `xml:"slide"`
}

type Slide struct {
	ItemNumber int    `xml:"itemnumber,attr"`
	Name       string `xml:"name"`
	Title      string `xml:"title"`
}
