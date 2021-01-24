package configs

import (
	"MailConfigHandler/ErrorPkg"
	"log"
	"net/http"
)

func getHml(url string) (*http.Response, error) {

	var client http.Client

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:84.0) Gecko/20100101 Firefox/84.0")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		return nil, ErrorPkg.New("Error request")
	}
	return resp, nil
}