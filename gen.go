package main

//go:generate go run gen.go

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"text/template"
	"time"
)

var modelTemplate = template.Must(template.ParseFiles("templates/model_template.txt"))
var clientTemplate = template.Must(template.ParseFiles("templates/client_template.txt"))

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
var enumTypes []EnumType

var enumRegex = regexp.MustCompile("^([a-zA-Z]+\\(\\d+\\), )+[a-zA-Z]+\\(\\d+\\)$")

// Wsdl represents the root wsdl:definition element
type Wsdl struct {
	XMLName   struct{}        `xml:"http://schemas.xmlsoap.org/wsdl/ definitions"`
	Types     WsdlTypes       `xml:"http://schemas.xmlsoap.org/wsdl/ types"`
	Services  []*WsdlService  `xml:"http://schemas.xmlsoap.org/wsdl/ service"`
	PortTypes []*WsdlPortType `xml:"http://schemas.xmlsoap.org/wsdl/ portType"`
	Messages  []*WsdlMessage  `xml:"http://schemas.xmlsoap.org/wsdl/ message"`
}

// WsdlTypes represents the wsdl:types XML node
type WsdlTypes struct {
	XsSchema XsSchema `xml:"http://www.w3.org/2001/XMLSchema schema"`
}

// XsSchema represents the xs:schema XML node
type XsSchema struct {
	ComplexTypes []*XsComplexType `xml:"http://www.w3.org/2001/XMLSchema complexType"`
	Elements     []*XsElement     `xml:"http://www.w3.org/2001/XMLSchema element"`
}

// XsComplexType represents the xs:complexType XML node
type XsComplexType struct {
	XsComplexContent XsComplexContent `xml:"http://www.w3.org/2001/XMLSchema complexContent"`
	XsAttributes     []*XsAttribute   `xml:"http://www.w3.org/2001/XMLSchema attribute"`
	XsSequence       XsSequence       `xml:"http://www.w3.org/2001/XMLSchema sequence"`
	Name             string           `xml:"name,attr"`
	XsAnnotation     XsAnnotation     `xml:"http://www.w3.org/2001/XMLSchema annotation"`
}

// XsComplexContent represents the xs:complexContent XML node
type XsComplexContent struct {
	XsExtension XsExtension `xml:"http://www.w3.org/2001/XMLSchema extension"`
}

// XsExtension represents the xs:extension XML node
type XsExtension struct {
	XsAttributes []*XsAttribute `xml:"http://www.w3.org/2001/XMLSchema attribute"`
	Base         string         `xml:"base,attr"`
	XsSequence   XsSequence     `xml:"http://www.w3.org/2001/XMLSchema sequence"`
}

// XsAttribute represents the xs:attribute XML node
type XsAttribute struct {
	Name         string       `xml:"name,attr"`
	Type         string       `xml:"type,attr"`
	Use          string       `xml:"use,attr"`
	XsAnnotation XsAnnotation `xml:"http://www.w3.org/2001/XMLSchema annotation"`
}

// XsAnnotation represents the xs:annotation XML node
type XsAnnotation struct {
	XsDocumentation string `xml:"http://www.w3.org/2001/XMLSchema documentation"`
}

// XsSequence represents the xs:sequence XML node
type XsSequence struct {
	XsElements []*XsElement `xml:"http://www.w3.org/2001/XMLSchema element"`
}

// XsElement represents the xs:element XML node
type XsElement struct {
	Name         string        `xml:"name,attr"`
	ComplexType  XsComplexType `xml:"http://www.w3.org/2001/XMLSchema complexType"`
	Type         string        `xml:"type,attr"`
	XsSimpleType XsSimpleType  `xml:"http://www.w3.org/2001/XMLSchema simpleType"`
}

// XsSimpleType represents the xs:simpleType XML node
type XsSimpleType struct {
	XsRestriction XsRestriction `xml:"http://www.w3.org/2001/XMLSchema restriction"`
}

// XsRestriction represents the xs:restriction XML node
type XsRestriction struct {
	Base string `xml:"base,attr"`
}

// WsdlService represents the wsdl:service XML node
type WsdlService struct {
	Name string   `xml:"name,attr"`
	Port WsdlPort `xml:"http://schemas.xmlsoap.org/wsdl/ port"`
}

// WsdlPort represents the wsdl:port XML node
type WsdlPort struct {
	SoapAddress SoapAddress `xml:"http://schemas.xmlsoap.org/wsdl/soap/ address"`
}

// SoapAddress represents the wsdl:address XML node
type SoapAddress struct {
	Location string `xml:"location,attr"`
}

// WsdlPortType represents the wsdl:portType XML node
type WsdlPortType struct {
	Name       string          `xml:"name,attr"`
	Operations []WsdlOperation `xml:"http://schemas.xmlsoap.org/wsdl/ operation"`
}

// WsdlOperation represents the wsdl:operation XML node
type WsdlOperation struct {
	Name            string     `xml:"name,attr"`
	Input           WsdlInput  `xml:"http://schemas.xmlsoap.org/wsdl/ input"`
	Output          WsdlOutput `xml:"http://schemas.xmlsoap.org/wsdl/ output"`
	XsDocumentation string     `xml:"http://schemas.xmlsoap.org/wsdl/ documentation"`
}

// WsdlInput represents the wsdl:input XML node
type WsdlInput struct {
	Message string `xml:"message,attr"`
}

// WsdlOutput represents the wsdl:output XML node
type WsdlOutput struct {
	Message string `xml:"message,attr"`
}

// WsdlMessage represents the wsdl:message XML node
type WsdlMessage struct {
	Name  string     `xml:"name,attr"`
	Parts []WsdlPart `xml:"http://schemas.xmlsoap.org/wsdl/ part"`
}

// WsdlPart represents the wsdl:part XML node
type WsdlPart struct {
	Name    string `xml:"name,attr"`
	Element string `xml:"element,attr"`
}

// BetdaqStruct represents a single generated struct from the API
type BetdaqStruct struct {
	Name          string
	Attributes    []*BetdaqAttribute
	Documentation string
}

// BetdaqAttribute represents an attribute for a BetdaqStruct
type BetdaqAttribute struct {
	Name             string
	Type             string
	XML              string
	Comment          string
	CommentMultiLine bool
}

// Function represents a single API client func to generate
type Function struct {
	Name                      string
	ParameterStruct           string
	ReturnType                string
	ReturnStatusContainerName string
	Service                   string
	Documentation             string
}

// EnumType represnts a type used for an enumeration of values for an API call
type EnumType struct {
	Name  string
	Type  string
	Enums []Enum
}

// Enum represents one name/value pair for an EnumType
type Enum struct {
	Name  string
	Value string
}

func main() {
	wsdl, err := ioutil.ReadFile("wsdl/betdaq-api.wsdl")
	die(err)

	fModel, err := os.Create("model/generated_model.go")
	die(err)
	defer fModel.Close()

	fClient, err := os.Create("client/generated_client.go")
	die(err)
	defer fClient.Close()

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

			buildFunction(parsed, o.Name, inputPart.Element, outputPart.Element, p.Name, o.XsDocumentation)

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
		EnumTypes     []EnumType
	}{
		Timestamp:     time.Now(),
		ServiceMap:    serviceMap,
		BetdaqStructs: betdaqStructs,
		EnumTypes:     enumTypes,
	})
	die(err)

	err = clientTemplate.Execute(fClient, struct {
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
			XML:  "`xml:\"" + attr.Name + "\"`",
		}
		betdaqAttributes = append(betdaqAttributes, &betdaqAttribute)
	}
	namespaceAttribute := &BetdaqAttribute{
		Name: "XMLName",
		Type: "struct{}",
		XML:  " `xml:\"http://www.GlobalBettingExchange.com/ExternalAPI/ " + name + "\"`",
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
			var attrType = attr.Type
			enumName, created := createEnumIfNeeded(name, attr)
			if created {
				attrType = enumName
			}

			betdaqAttribute := BetdaqAttribute{
				Name:             capitalize(attr.Name),
				Type:             mapType(attrType),
				Comment:          attr.XsAnnotation.XsDocumentation,
				CommentMultiLine: strings.Contains(attr.XsAnnotation.XsDocumentation, "\n"),
				XML:              "`xml:\"" + attr.Name + ",attr\"`",
			}
			betdaqAttributes = append(betdaqAttributes, &betdaqAttribute)
		}

		// extension attributes
		for _, attr := range typ.XsComplexContent.XsExtension.XsAttributes {
			var attrType = mapType(attr.Type)
			enumName, created := createEnumIfNeeded(name, attr)
			if created {
				attrType = enumName
			}

			betdaqAttribute := BetdaqAttribute{
				Name:             capitalize(attr.Name),
				Type:             attrType,
				Comment:          attr.XsAnnotation.XsDocumentation,
				CommentMultiLine: strings.Contains(attr.XsAnnotation.XsDocumentation, "\n"),
				XML:              "`xml:\"" + attr.Name + ",attr\"`",
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
				XML:  "`xml:\"" + el.Name + "\"`",
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
				XML:  "`xml:\"" + el.Name + "\"`",
			}
			if el.Type != name {
				buildStructFromComplexType(parsed, el.Type, el.Type)
			}
			betdaqAttributes = append(betdaqAttributes, &betdaqAttribute)
		}
	} else {
		fmt.Println("getComplexType not found", name)
	}

	fmt.Println(enumTypes)

	betdaqStruct := BetdaqStruct{
		Name:          name,
		Attributes:    betdaqAttributes,
		Documentation: typ.XsAnnotation.XsDocumentation,
	}
	betdaqStructs = append(betdaqStructs, &betdaqStruct)
}

func createEnumIfNeeded(typeName string, attr *XsAttribute) (string, bool) {
	if enumRegex.MatchString(attr.XsAnnotation.XsDocumentation) {
		// CancelOrders(1), SuspendOrders(2), SuspendPunter(3)
		doc := attr.XsAnnotation.XsDocumentation
		values := strings.Split(doc, ",")
		var enums []Enum
		enumName := typeName + "_" + capitalize(attr.Name)
		for _, value := range values {
			fmt.Println("***", value)
			value = strings.Trim(value, " ")
			valueName := typeName + "_" + capitalize(attr.Name) + "_" + strings.Split(value, "(")[0]
			value := strings.Split(value, "(")[1]
			value = strings.Trim(value, ")")
			enums = append(enums, Enum{
				Name:  valueName,
				Value: value,
			})
		}
		enumTypes = append(enumTypes, EnumType{
			Name:  enumName,
			Type:  mapType(attr.Type),
			Enums: enums,
		})
		return enumName, true
	}
	return "", false
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

func buildFunction(parsed Wsdl, name string, inputElement string, outputElement string, service string, doc string) {
	element := getElement(parsed, outputElement)
	innerElementName := element.ComplexType.XsSequence.XsElements[0].Name

	function := Function{
		Name:                      name,
		ParameterStruct:           inputElement,
		ReturnType:                outputElement,
		ReturnStatusContainerName: innerElementName,
		Service:                   service,
		Documentation:             doc,
	}
	functions = append(functions, &function)
}
