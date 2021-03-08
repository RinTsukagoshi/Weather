package law

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

type Re struct {
	XMLName xml.Name `xml:"response"`
	Text    string   `xml:",chardata"`
	Err     string   `xml:"error"`

	Location []struct {
		Text     string `xml:",chardata"`
		City     string `xml:"city"`
		CityKana string `xml:"city-kana"`
		Town     string `xml:"town"`
		TownKana string `xml:"town-kana"`
		X        struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
		} `xml:"x"`
		Y struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
		} `xml:"y"`
		Distance struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
		} `xml:"distance"`
		Prefecture string `xml:"prefecture"`
		Postal     string `xml:"postal"`
	} `xml:"location"`
}

func Getadd(x string, y string) Re {
	res := HttpGet("http://geoapi.heartrails.com/api/xml?method=searchByGeoLocation&x=" + x + "&y=" + y)
	result := Re{}
	print(result.Err)
	err := xml.Unmarshal([]byte(res), &result)
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)

	}
	if strings.Contains(result.Err, "Cities") {
		print("えら")
		os.Exit(1)
	}

	return result
}
