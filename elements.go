package wsdl

import (
	"encoding/xml"
)

// Elements' comment please refer to http://www.w3schools.com/xml/xml_wsdl.asp

// Definitions is top-level tag
type Definitions struct {
	XMLName   xml.Name   `xml:"definitions"`
	Name      string     `xml:"name,attr"`
	Types     Types      `xml:"types"`
	Messages  []Message  `xml:"message"`
	PortTypes []PortType `xml:"portType"`
	Bindings  []Binding  `xml:"binding"`
	Service   Service    `xml:"service"`
}

// Types defines the (XML Schema) data types used by the web service
type Types struct {
	Schema interface{} `xml:",any"`
}

// Message defines the data elements for each operation
type Message struct {
	Name  string `xml:"name,attr"`
	Parts []Part `xml:"part"`
}

// Part define input and output
type Part struct {
	Name string `xml:"name,attr"`
	Type string `xml:"type,attr"`
}

// PortType describes the operations that can be performed and the messages involved.
type PortType struct {
	Name       string      `xml:"name,attr"`
	Operations []Operation `xml:"operation"`
}

// Operation defines operation type
type Operation struct {
	Name   string `xml:"name,attr"`
	Input  Input  `xml:"input"`
	Output Output `xml:"output"`
	Fault  Fault  `xml:"fault"`
}

// Input defines input interface
type Input struct {
	Name    string `xml:"name,attr"`
	Message string `xml:"message,attr"`
}

// Output defines output interface
type Output struct {
	Name    string `xml:"name,attr"`
	Message string `xml:"message,attr"`
}

// Fault defines output interface
type Fault struct {
	Name    string `xml:"name,attr"`
	Message string `xml:"message,attr"`
}

// Binding defines the protocol and data format for each port type
type Binding struct {
	Name       string      `xml:"name,attr"`
	Type       string      `xml:"type,attr"`
	Operations []Operation `xml:"operation"`
}

// Service defines expose service
type Service struct {
	Name  string `xml:"name,attr"`
	Ports []Port `xml:"port"`
}

// Port defines parameters
type Port struct {
	Name    string `xml:"name,attr"`
	Binding string `xml:"binding,attr"`
}

// utility elemets

// Import other WSDL documents or XML Schemas
type Import struct {
	NameSpace string `xml:"namespace,attr"`
	location  string `xml:"location,attr"`
}

// Documentation provides text doc
type Documentation struct {
	Doc interface{} `xml:",any"`
}
