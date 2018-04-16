package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
	"text/template"
	"time"
)

//go:generate go run gen.go

var codeTemplate = template.Must(template.ParseFiles("template.txt"))

var attributeTypeMap = map[string]string{
	"xs:boolean":      "bool",
	"xs:short":        "int64",
	"xs:long":         "int64",
	"xs:decimal":      "string",
	"xs:int":          "int64",
	"xs:unsignedByte": "int64",
	"xs:string":       "string",
	"xs:dateTime":     "string",
}

var betdaqStructs []*BetdaqStruct

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
	XsSequence   XsSequence     `xml:"http://www.w3.org/2001/XMLSchema sequence"`
}

type XsAttribute struct {
	Name         string       `xml:"name,attr"`
	Type         string       `xml:"type,attr"`
	Use          string       `xml:"use,attr"`
	XsAnnotation XsAnnotation `xml:"http://www.w3.org/2001/XMLSchema annotation"`
}

type XsAnnotation struct {
	XsDocumentation string `xml:"http://www.w3.org/2001/XMLSchema documentation"`
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
	Name            string     `xml:"name,attr"`
	Input           WsdlInput  `xml:"http://schemas.xmlsoap.org/wsdl/ input"`
	Output          WsdlOutput `xml:"http://schemas.xmlsoap.org/wsdl/ output"`
	XsDocumentation string     `xml:"http://www.w3.org/2001/XMLSchema documentation"`
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

type BetdaqStruct struct {
	Name       string
	Attributes []*BetdaqAttribute
}

type BetdaqAttribute struct {
	Name             string
	Type             string
	Xml              string
	Comment          string
	CommentMultiLine bool
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

	sort.Slice(parsed.Types.XsSchema.ComplexTypes, func(i, j int) bool {
		return parsed.Types.XsSchema.ComplexTypes[i].Name < parsed.Types.XsSchema.ComplexTypes[j].Name
	})

	for _, p := range parsed.PortTypes {
		fmt.Println(p.Name)
		for _, o := range p.Operations {
			fmt.Println("  ", o.Name)
			if o.XsDocumentation != "" {
				fmt.Println("    ", o.XsDocumentation)
			}

			fmt.Println("    ", o.Input.Message)
			inputMessage := getMessage(parsed, o.Input.Message)
			for _, part := range inputMessage.Parts {
				fmt.Println("      ", part.Name, part.Element)
				buildStructFromElement(parsed, part.Element)
			}

			fmt.Println("    ", o.Output.Message)
			outputMessage := getMessage(parsed, o.Output.Message)
			for _, part := range outputMessage.Parts {
				fmt.Println("      ", part.Name, part.Element)
				buildStructFromElement(parsed, part.Element)
			}
			fmt.Println()
		}
	}

	err = codeTemplate.Execute(f, struct {
		Timestamp     time.Time
		ServiceMap    map[string]*WsdlService
		BetdaqStructs []*BetdaqStruct
	}{
		Timestamp:     time.Now(),
		ServiceMap:    serviceMap,
		BetdaqStructs: betdaqStructs,
	})
	die(err)
}

func die(err error) {
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}

func mapType(wsdlType string) string {
	typ, found := attributeTypeMap[wsdlType]
	if found {
		return typ
	}
	return wsdlType
}

func buildStructFromElement(parsed Wsdl, name string) {
	//fmt.Println("buildStructFromElement", name)

	var betdaqAttributes []*BetdaqAttribute
	element := getElement(parsed, name)
	//fmt.Println("Element found for", name)

	for _, attr := range element.ComplexType.XsSequence.XsSequenceElements {
		var newName = attr.Type
		if attr.Type == name {
			newName = attr.Name
		}

		buildStructFromComplexType(parsed, newName, name)

		betdaqAttribute := BetdaqAttribute{
			Name: attr.Name,
			Type: newName,
		}
		betdaqAttributes = append(betdaqAttributes, &betdaqAttribute)
	}
	betdaqStruct := BetdaqStruct{
		Name:       name,
		Attributes: betdaqAttributes,
	}
	betdaqStructs = append(betdaqStructs, &betdaqStruct)
}

func buildStructFromComplexType(parsed Wsdl, name string, usingDataFromName string) {
	fmt.Println("buildStructFromComplexType", name)

	var betdaqAttributes []*BetdaqAttribute

	typ, found := getComplexType(parsed, usingDataFromName)
	if !found {
		typ, found = getComplexType(parsed, name)
	}
	if found {
		// base class
		if typ.XsComplexContent.XsExtension.Base != "" {
			_, baseFound := getBetdaqStructByName(typ.XsComplexContent.XsExtension.Base)
			if !baseFound {
				buildStructFromComplexType(parsed, typ.XsComplexContent.XsExtension.Base, typ.XsComplexContent.XsExtension.Base)
			}

			betdaqAttribute := BetdaqAttribute{
				Name: "*" + typ.XsComplexContent.XsExtension.Base,
			}
			betdaqAttributes = append(betdaqAttributes, &betdaqAttribute)
		}

		// simple attributes
		for _, attr := range typ.XsAttributes {
			betdaqAttribute := BetdaqAttribute{
				Name:             attr.Name,
				Type:             mapType(attr.Type),
				Comment:          attr.XsAnnotation.XsDocumentation,
				CommentMultiLine: strings.Contains(attr.XsAnnotation.XsDocumentation, "\n"),
			}
			betdaqAttributes = append(betdaqAttributes, &betdaqAttribute)
		}

		// extension attributes
		for _, attr := range typ.XsComplexContent.XsExtension.XsAttributes {
			betdaqAttribute := BetdaqAttribute{
				Name:             attr.Name,
				Type:             mapType(attr.Type),
				Comment:          attr.XsAnnotation.XsDocumentation,
				CommentMultiLine: strings.Contains(attr.XsAnnotation.XsDocumentation, "\n"),
			}
			betdaqAttributes = append(betdaqAttributes, &betdaqAttribute)
		}

		// sequence attributes
		for _, el := range typ.XsSequence.XsSequenceElements {
			betdaqAttribute := BetdaqAttribute{
				Name: el.Name,
				Type: "[]" + mapType(el.Type),
			}
			_, found := attributeTypeMap[el.Type]
			if !found {
				_, baseFound := getBetdaqStructByName(el.Type)
				if !baseFound {
					if el.Type != name {
						buildStructFromComplexType(parsed, el.Type, el.Type)
					}
				}
			}
			betdaqAttributes = append(betdaqAttributes, &betdaqAttribute)
		}

		// extension attributes
		for _, el := range typ.XsComplexContent.XsExtension.XsSequence.XsSequenceElements {
			betdaqAttribute := BetdaqAttribute{
				Name: el.Name,
				Type: "[]" + mapType(el.Type),
			}
			_, found := attributeTypeMap[el.Type]
			if !found {
				_, baseFound := getBetdaqStructByName(el.Type)
				if !baseFound {
					if el.Type != name {
						buildStructFromComplexType(parsed, el.Type, el.Type)
					}
				}
			}
			betdaqAttributes = append(betdaqAttributes, &betdaqAttribute)
		}
	} else {
		fmt.Println("getComplexType not found", name)
	}

	betdaqStruct := BetdaqStruct{
		Name:       name,
		Attributes: betdaqAttributes,
	}
	betdaqStructs = append(betdaqStructs, &betdaqStruct)
}

func getBetdaqStructByName(name string) (*BetdaqStruct, bool) {
	for _, betdaqStruct := range betdaqStructs {
		if betdaqStruct.Name == name {
			return betdaqStruct, true
		}
	}
	return nil, false
}

func getMessage(parsed Wsdl, name string) *WsdlMessage {
	for _, message := range parsed.Messages {
		if message.Name == name {
			return message
		}
	}
	panic("uh-oh getMessage")
}

func getElement(parsed Wsdl, name string) *XsElement {
	for _, element := range parsed.Types.XsSchema.Elements {
		if element.Name == name {
			return element
		}
	}
	panic("uh-oh getElement")
}

func getComplexType(parsed Wsdl, name string) (*XsComplexType, bool) {
	for _, ct := range parsed.Types.XsSchema.ComplexTypes {
		if ct.Name == name {
			return ct, true
		}
	}
	return nil, false
}
