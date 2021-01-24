package configs

import (
	"MailConfigHandler/ErrorPkg"
	"encoding/xml"
	"io/ioutil"
	"log"
)

func FindConfig(domain string) (*EmailProvider, error) {

	var requestUrl = hostISPDB + domain
	resp, err := getHml(requestUrl)

	if err != nil {
		log.Println("ERROR: ", err)
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return getConfigFromXml(&bodyBytes)
}

func getConfigFromXml(xmlStr *[]byte) (*EmailProvider, error) {

	if xmlStr == nil {
		return nil, ErrorPkg.New("String null or empty")
	}

	config := new(EmailProvider)

	errXml := xml.Unmarshal(*xmlStr, &config)
	if errXml != nil {
		return nil, errXml
	}

	return config, nil
}
