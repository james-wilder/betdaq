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

var modelTemplate = template.Must(template.ParseFiles("model_template.txt"))
var apiTemplate = template.Must(template.ParseFiles("api_template.txt"))

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
var functions []*Function

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
	Name         string        `xml:"name,attr"`
	ComplexType  XsComplexType `xml:"http://www.w3.org/2001/XMLSchema complexType"`
	Type         string        `xml:"type,attr"`
	XsSimpleType XsSimpleType  `xml:"http://www.w3.org/2001/XMLSchema simpleType"`
}

type XsSequence struct {
	XsElements []*XsElement `xml:"http://www.w3.org/2001/XMLSchema element"`
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

type Function struct {
	Name                      string
	ParameterStruct           string
	ReturnType                string
	ReturnStatusContainerName string
	Service                   string
}

type Parameter struct {
	Name string
	Type string
}

func main() {
	wsdl, err := ioutil.ReadFile("api/betdaq-api.wsdl")
	die(err)

	fModel, err := os.Create("api/generated_model.go")
	die(err)
	defer fModel.Close()

	fApi, err := os.Create("api/generated_api.go")
	die(err)
	defer fApi.Close()

	var parsed Wsdl
	err = xml.Unmarshal(wsdl, &parsed)
	die(err)

	var serviceMap = make(map[string]*WsdlService)
	for _, service := range parsed.Services {
		fmt.Println("Service", service.Name)
		serviceMap[service.Name] = service
	}

	for _, p := range parsed.PortTypes {
		fmt.Println(p.Name)
		for _, o := range p.Operations {
			fmt.Println("  ", o.Name)
			if o.XsDocumentation != "" {
				fmt.Println("    ", o.XsDocumentation)
			}

			fmt.Println("    ", o.Input.Message)
			inputMessage := getMessage(parsed, o.Input.Message)
			if len(inputMessage.Parts) > 1 {
				panic("I can't handle multiple input parts")
			}
			inputPart := inputMessage.Parts[0]
			fmt.Println("      ", inputPart.Name, inputPart.Element)
			buildStructFromElementByName(parsed, inputPart.Element)

			fmt.Println("    ", o.Output.Message)
			outputMessage := getMessage(parsed, o.Output.Message)
			if len(outputMessage.Parts) > 1 {
				panic("I can't handle multiple output parts")
			}
			outputPart := outputMessage.Parts[0]
			fmt.Println("      ", outputPart.Name, outputPart.Element)
			buildStructFromElementByName(parsed, outputPart.Element)

			buildFunction(parsed, o.Name, inputPart.Element, outputPart.Element, p.Name)

			fmt.Println()
		}
	}

	sort.Slice(betdaqStructs, func(i, j int) bool {
		return betdaqStructs[i].Name < betdaqStructs[j].Name
	})

	sort.Slice(functions, func(i, j int) bool {
		return functions[i].Name < functions[j].Name
	})

	err = modelTemplate.Execute(fModel, struct {
		Timestamp     time.Time
		ServiceMap    map[string]*WsdlService
		BetdaqStructs []*BetdaqStruct
	}{
		Timestamp:     time.Now(),
		ServiceMap:    serviceMap,
		BetdaqStructs: betdaqStructs,
	})
	die(err)

	err = apiTemplate.Execute(fApi, struct {
		Timestamp time.Time
		Functions []*Function
	}{
		Timestamp: time.Now(),
		Functions: functions,
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

func buildStructFromElementByName(parsed Wsdl, name string) {
	//fmt.Println("buildStructFromElement", name)

	var betdaqAttributes []*BetdaqAttribute
	element := getElement(parsed, name)
	//fmt.Println("Element found for", name)

	for _, attr := range element.ComplexType.XsSequence.XsElements {
		var newName = attr.Type
		if attr.Type == name {
			newName = attr.Name
		}

		buildStructFromComplexType(parsed, newName, name)

		betdaqAttribute := BetdaqAttribute{
			Name: capitalize(attr.Name),
			Type: newName,
			Xml:  "`xml:\"" + attr.Name + "\"`",
		}
		betdaqAttributes = append(betdaqAttributes, &betdaqAttribute)
	}
	namespaceAttribute := &BetdaqAttribute{
		Name: "XMLName",
		Type: "struct{}",
		Xml:  " `xml:\"http://www.GlobalBettingExchange.com/ExternalAPI/ " + name + "\"`",
	}
	betdaqAttributes = append([]*BetdaqAttribute{namespaceAttribute}, betdaqAttributes...)

	betdaqStruct := BetdaqStruct{
		Name:       name,
		Attributes: betdaqAttributes,
	}
	betdaqStructs = append(betdaqStructs, &betdaqStruct)
}

func buildStructFromElement(parsed Wsdl, name string, element XsElement) {
	//fmt.Println("buildStructFromElement", name)

	var betdaqAttributes []*BetdaqAttribute
	//fmt.Println("Element found for", name)

	for _, el := range element.ComplexType.XsSequence.XsElements {
		newName := mapType(el.Type)
		if newName == "" {
			newName = name + el.Name
			buildStructFromElement(parsed, newName, *el)
		}

		buildStructFromComplexType(parsed, newName, name)

		betdaqAttribute := BetdaqAttribute{
			Name: el.Name,
			Type: "[]" + newName,
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

	// TODO: do this better
	if name == "" || name == "int64" || name == "string" {
		return
	}

	var betdaqAttributes []*BetdaqAttribute

	_, found := attributeTypeMap[name]
	if found {
		return
	}
	_, baseFound := getBetdaqStructByName(name)
	if baseFound {
		return
	}

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
				Name:             capitalize(attr.Name),
				Type:             mapType(attr.Type),
				Comment:          attr.XsAnnotation.XsDocumentation,
				CommentMultiLine: strings.Contains(attr.XsAnnotation.XsDocumentation, "\n"),
				Xml:              "`xml:\"" + attr.Name + ",attr\"`",
			}
			betdaqAttributes = append(betdaqAttributes, &betdaqAttribute)
		}

		// extension attributes
		for _, attr := range typ.XsComplexContent.XsExtension.XsAttributes {
			betdaqAttribute := BetdaqAttribute{
				Name:             capitalize(attr.Name),
				Type:             mapType(attr.Type),
				Comment:          attr.XsAnnotation.XsDocumentation,
				CommentMultiLine: strings.Contains(attr.XsAnnotation.XsDocumentation, "\n"),
				Xml:              "`xml:\"" + attr.Name + ",attr\"`",
			}
			betdaqAttributes = append(betdaqAttributes, &betdaqAttribute)
		}

		// sequence attributes
		for _, el := range typ.XsSequence.XsElements {
			var attrType = mapType(el.Type)
			if attrType == "" {
				attrType = name + el.Name
				buildStructFromElement(parsed, attrType, *el)
			}

			betdaqAttribute := BetdaqAttribute{
				Name: capitalize(el.Name),
				Type: "[]" + attrType,
				Xml:  "`xml:\"" + el.Name + "\"`",
			}
			if el.Type != name {
				buildStructFromComplexType(parsed, el.Type, el.Type)
			}
			betdaqAttributes = append(betdaqAttributes, &betdaqAttribute)
		}

		// extension attributes
		for _, el := range typ.XsComplexContent.XsExtension.XsSequence.XsElements {
			var attrType = mapType(el.Type)
			if attrType == "" {
				attrType = name + el.Name
				buildStructFromElement(parsed, attrType, *el)
			}

			betdaqAttribute := BetdaqAttribute{
				Name: capitalize(el.Name),
				Type: "[]" + attrType,
				Xml:  "`xml:\"" + el.Name + "\"`",
			}
			if el.Type != name {
				buildStructFromComplexType(parsed, el.Type, el.Type)
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

func capitalize(s string) string {
	return strings.Title(s)
}

func buildFunction(parsed Wsdl, name string, inputElement string, outputElement string, service string) {
	element := getElement(parsed, outputElement)
	innerElementName := element.ComplexType.XsSequence.XsElements[0].Name

	function := Function{
		Name:                      name,
		ParameterStruct:           inputElement,
		ReturnType:                outputElement,
		ReturnStatusContainerName: innerElementName,
		Service:                   service,
	}
	functions = append(functions, &function)
}
