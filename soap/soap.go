package soap

import (
	"encoding/xml"
)

type Envelope struct {
	XMLName struct{} `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  Header
	Body    Body
}

type Header struct {
	XMLName           struct{}          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`
	ExternalApiHeader ExternalApiHeader `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ExternalApiHeader"`
}

type Body struct {
	XMLName  struct{} `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	Contents []byte   `xml:",innerxml"`
}

type ExternalApiHeader struct {
	Version      string `xml:"version,attr"`
	LanguageCode string `xml:"languageCode,attr"`
	Username     string `xml:"username,attr"`
	Password     string `xml:"password,attr"`
}

func Encode(request interface{}, username, password string) ([]byte, error) {
	data, err := xml.MarshalIndent(request, "    ", "  ")
	if err != nil {
		return nil, err
	}
	data = append([]byte("\n"), data...)
	data = append(data, '\n')
	env := Envelope{
		Header: Header{
			ExternalApiHeader: ExternalApiHeader{
				Version:      "2",
				LanguageCode: "en",
				Username:     username,
				Password:     password,
			},
		},
		Body: Body{
			Contents: data,
		},
	}
	soapEncoded, err := xml.MarshalIndent(&env, "", "  ")
	if err != nil {

	}
	encoded := append([]byte(xml.Header), soapEncoded...)

	return encoded, nil
}

func Decode(data []byte, response interface{}) error {
	env := Envelope{
		Body: Body{},
	}
	err := xml.Unmarshal(data, &env)
	if err != nil {
		return err
	}
	return xml.Unmarshal(env.Body.Contents, response)
}
