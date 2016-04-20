// Package wsdl currently provides two parse function for pre-defined elements
package wsdl

import (
	"encoding/xml"
	"io/ioutil"
	"log"
)

// ParseWSDL return a parsed definition or empty
func ParseWSDL(data []byte) Definitions {
	d := Definitions{}
	err := xml.Unmarshal(data, &d)
	if err != nil {
		log.Fatal("ParseWSDL:", err)
		return Definitions{}
	}
	return d
}

// ParseWSDLFile return a parsed definition or empty by given xml name
func ParseWSDLFile(filename string) Definitions {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("ParseWSDLFile:", err)
		return Definitions{}
	}
	return ParseWSDL(data)
}
