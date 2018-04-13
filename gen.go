package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"text/template"
	"time"
)

//go:generate go run gen.go

type Wsdl struct {
	XMLName   struct{}        `xml:"http://schemas.xmlsoap.org/wsdl/ definitions"`
	Types     WsdlTypes       `xml:"http://schemas.xmlsoap.org/wsdl/ types"`
	Services  []*WsdlService  `xml:"http://schemas.xmlsoap.org/wsdl/ service"`
	PortTypes []*WsdlPortType `xml:"http://schemas.xmlsoap.org/wsdl/ portType"`
	Messages  []*WsdlMessage  `xml:"http://schemas.xmlsoap.org/wsdl/ message"`
}

type WsdlTypes struct {
	XsSchema XsSchema `xml:"http://www.w3.org/2001/XMLSchema schema"`
}

type XsSchema struct {
	ComplexTypes []*XsComplexType `xml:"http://www.w3.org/2001/XMLSchema complexType"`
	Elements     []*XsElement     `xml:"http://www.w3.org/2001/XMLSchema element"`
}

type XsComplexType struct {
	XsComplexContent XsComplexContent `xml:"http://www.w3.org/2001/XMLSchema complexContent"`
	XsAttributes     []*XsAttribute   `xml:"http://www.w3.org/2001/XMLSchema attribute"`
	XsSequence       XsSequence       `xml:"http://www.w3.org/2001/XMLSchema sequence"`
	Name             string           `xml:"name,attr"`
}

type XsComplexContent struct {
	XsExtension XsExtension `xml:"http://www.w3.org/2001/XMLSchema extension"`
}

type XsExtension struct {
	XsAttributes []*XsAttribute `xml:"http://www.w3.org/2001/XMLSchema attribute"`
	Base         string         `xml:"base,attr"`
}

type XsAttribute struct {
	Name string `xml:"name,attr"`
	Type string `xml:"type,attr"`
	Use  string `xml:"use,attr"`
}

type XsElement struct {
	Name        string        `xml:"name,attr"`
	ComplexType XsComplexType `xml:"http://www.w3.org/2001/XMLSchema complexType"`
}

type XsSequence struct {
	XsSequenceElements []*XsSequenceElement `xml:"http://www.w3.org/2001/XMLSchema element"`
}

type XsSequenceElement struct {
	Name         string       `xml:"name,attr"`
	Type         string       `xml:"type,attr"`
	XsSimpleType XsSimpleType `xml:"http://www.w3.org/2001/XMLSchema simpleType"`
}

type XsSimpleType struct {
	XsRestriction XsRestriction `xml:"http://www.w3.org/2001/XMLSchema restriction"`
}

type XsRestriction struct {
	Base string `xml:"base,attr"`
}

type WsdlService struct {
	Name string   `xml:"name,attr"`
	Port WsdlPort `xml:"http://schemas.xmlsoap.org/wsdl/ port"`
}

type WsdlPort struct {
	SoapAddress SoapAddress `xml:"http://schemas.xmlsoap.org/wsdl/soap/ address"`
}

type SoapAddress struct {
	Location string `xml:"location,attr"`
}

type WsdlPortType struct {
	Name       string          `xml:"name,attr"`
	Operations []WsdlOperation `xml:"http://schemas.xmlsoap.org/wsdl/ operation"`
}

type WsdlOperation struct {
	Name   string     `xml:"name,attr"`
	Input  WsdlInput  `xml:"http://schemas.xmlsoap.org/wsdl/ input"`
	Output WsdlOutput `xml:"http://schemas.xmlsoap.org/wsdl/ output"`
}

type WsdlInput struct {
	Message string `xml:"message,attr"`
}

type WsdlOutput struct {
	Message string `xml:"message,attr"`
}

type WsdlMessage struct {
	Name  string     `xml:"name,attr"`
	Parts []WsdlPart `xml:"http://schemas.xmlsoap.org/wsdl/ part"`
}

type WsdlPart struct {
	Name    string `xml:"name,attr"`
	Element string `xml:"element,attr"`
}

func main() {
	wsdl, err := ioutil.ReadFile("api/betdaq-api.wsdl")
	die(err)

	f, err := os.Create("api/generated_model.go")
	die(err)
	defer f.Close()

	var parsed Wsdl
	err = xml.Unmarshal(wsdl, &parsed)
	die(err)

	var serviceMap = make(map[string]*WsdlService)
	for _, service := range parsed.Services {
		fmt.Println("Service", service.Name)
		serviceMap[service.Name] = service
	}

	var messageMap = make(map[string]*WsdlMessage)
	for _, message := range parsed.Messages {
		fmt.Println("Message", message.Name)
		messageMap[message.Name] = message
	}

	var elementMap = make(map[string]*XsElement)
	for _, element := range parsed.Types.XsSchema.Elements {
		fmt.Println("Element", element.Name)
		elementMap[element.Name] = element
	}

	var complexTypeMap = make(map[string]*XsComplexType)
	for _, typ := range parsed.Types.XsSchema.ComplexTypes {
		fmt.Println("ComplexType", typ.Name)
		complexTypeMap[typ.Name] = typ
	}

	sort.Slice(parsed.Types.XsSchema.ComplexTypes, func(i, j int) bool {
		return parsed.Types.XsSchema.ComplexTypes[i].Name < parsed.Types.XsSchema.ComplexTypes[j].Name
	})

	var attributeTypeMap = map[string]string{
		"xs:boolean":      "bool",
		"xs:short":        "int64",
		"xs:long":         "int64",
		"xs:decimal":      "int64",
		"xs:int":          "int64",
		"xs:unsignedByte": "int64",
		"xs:string":       "string",
		"xs:dateTime":     "string",
	}
	for _, typ := range parsed.Types.XsSchema.ComplexTypes {
		attributeTypeMap[typ.Name] = typ.Name
	}

	err = packageTemplate.Execute(f, struct {
		Timestamp        time.Time
		Wsdl             Wsdl
		ServiceMap       map[string]*WsdlService
		MessageMap       map[string]*WsdlMessage
		ComplexTypeMap   map[string]*XsComplexType
		AttributeTypeMap map[string]string
		ElementMap       map[string]*XsElement
	}{
		Timestamp:        time.Now(),
		Wsdl:             parsed,
		ServiceMap:       serviceMap,
		MessageMap:       messageMap,
		ComplexTypeMap:   complexTypeMap,
		AttributeTypeMap: attributeTypeMap,
		ElementMap:       elementMap,
	})
	die(err)
}

func die(err error) {
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}

var packageTemplate = template.Must(template.ParseFiles("template.txt"))
