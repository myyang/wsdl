package wsdl

import (
	"encoding/xml"
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func diff_def(d1, d2 Definitions) (bool, error) {

	if d1.Name != d2.Name {
		return false, errors.New("Definitions name are not equal")
	} else if !reflect.DeepEqual(d1.Types, d2.Types) {
		fmt.Printf("Types:\n===\n%#v\n===\n%#v\n===\n", d1.Types, d2.Types)
		return false, errors.New("Types are not equal")
	} else if !reflect.DeepEqual(d1.Messages, d2.Messages) {
		fmt.Printf("Messages:\n===\n%#v\n===\n%#v\n===\n", d1.Messages, d2.Messages)
		return false, errors.New("Messages are not equal")
	} else if !reflect.DeepEqual(d1.PortTypes, d2.PortTypes) {
		fmt.Printf("PortTypes:\n===\n%#v\n===\n%#v\n===\n", d1.PortTypes, d2.PortTypes)
		return false, errors.New("PortTypes are not equal")
	} else if !reflect.DeepEqual(d1.Bindings, d2.Bindings) {
		fmt.Printf("Bindings:\n===\n%#v\n===\n%#v\n===\n", d1.Bindings, d2.Bindings)
		return false, errors.New("Bindings are not equal")
	} else if !reflect.DeepEqual(d1.Service, d2.Service) {
		fmt.Printf("Service:\n===\n%#v\n===\n%#v\n===\n", d1.Service, d2.Service)
		return false, errors.New("Service are not equal")
	}
	return true, nil
}

func TestParseWSDL(t *testing.T) {
	data := `
	<definitions name="testing">
	</definitions>
	`
	supposed := Definitions{
		XMLName:   xml.Name{Space: "", Local: "definitions"},
		Name:      "testing",
		Types:     Types{},
		Messages:  []Message(nil),
		PortTypes: []PortType(nil),
		Bindings:  []Binding(nil),
		Service:   Service{},
	}
	parsed := ParseWSDL([]byte(data))
	if !reflect.DeepEqual(supposed, parsed) {
		diff_def(supposed, parsed)
		t.Fatalf("Supposed: %v, got: %v\n", supposed, parsed)
	}

	data = `
	<definitions name="HelloService">

	<message name="SayHelloRequest">
		<part name="firstName" type="xsd:string"/>
	</message>

	<message name="SayHelloResponse">
		<part name="greeting" type="xsd:string"/>
	</message>

	<portType name="Hello_PortType">
		<operation name="sayHello">
			<input message="tns:SayHelloRequest"/>
			<output message="tns:SayHelloResponse"/>
		</operation>
	</portType>

	<binding name="Hello_Binding" type="tns:Hello_PortType">
		<operation name="sayHello">
			<soap:operation soapAction="sayHello"/>
			<input></input>
			<output></output>
		</operation>
	</binding>

	<service name="Hello_Service">
		<documentation>WSDL File for HelloService</documentation>
		<port binding="tns:Hello_Binding" name="Hello_Port">
		</port>
		</service>
	</definitions>

	`
	supposed = Definitions{
		XMLName: xml.Name{Space: "", Local: "definitions"},
		Name:    "HelloService",
		Types:   Types{},
		Messages: []Message{
			Message{Name: "SayHelloRequest", Parts: []Part{Part{Name: "firstName", Type: "xsd:string"}}},
			Message{Name: "SayHelloResponse", Parts: []Part{Part{Name: "greeting", Type: "xsd:string"}}},
		},
		PortTypes: []PortType{
			PortType{
				Name: "Hello_PortType",
				Operations: []Operation{
					Operation{
						Name:   "sayHello",
						Input:  Input{Message: "tns:SayHelloRequest"},
						Output: Output{Message: "tns:SayHelloResponse"},
					},
				},
			},
		},
		Bindings: []Binding{
			Binding{
				Name: "Hello_Binding",
				Type: "tns:Hello_PortType",
				Operations: []Operation{
					Operation{
						Name:   "sayHello",
						Input:  Input{},
						Output: Output{},
					},
				},
			},
		},
		Service: Service{
			Name: "Hello_Service",
			Ports: []Port{
				Port{
					Name:    "Hello_Port",
					Binding: "tns:Hello_Binding",
				},
			},
		},
	}
	parsed = ParseWSDL([]byte(data))
	if !reflect.DeepEqual(supposed, parsed) {
		diff_def(supposed, parsed)
		t.Fatalf("Supposed: %v, got: %v", supposed, parsed)
	}
}
