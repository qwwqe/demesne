// Stub for post-processing SVGs output by PlantUML.
// Not entirely clear on
package main

import (
	"encoding/xml"
)

type XMLNode struct {
	XMLName xml.Name
	Attrs []xml.Attr `xml:",any,attr"`
	Content []byte `xml:",innerxml"`
	XMLNodes []XMLNode `xml:",any"`
}

func main() {

}
