package law

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type XML struct {
	XMLName xml.Name `xml:"result"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version"`
	Address string   `xml:"address"`
	Cor     struct {
		Text   string `xml:",chardata"`
		Lat    string `xml:"lat"`
		Lng    string `xml:"lng"`
		LatDms string `xml:"lat_dms"`
		LngDms string `xml:"lng_dms"`
	} `xml:"coordinate"`
	OpenLocationCode string `xml:"open_location_code"`
	URL              string `xml:"url"`
	Gm               string `xml:"google_maps"`
}

func HttpGet(url string) string {
	response, _ := http.Get(url)
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	return string(body)
}

func XmltoD(s string) XML {
	result := XML{}
	err := xml.Unmarshal([]byte(s), &result)
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)

	}

	return result

}
