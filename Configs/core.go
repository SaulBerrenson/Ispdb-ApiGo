package configs

import "sync"

//////////////////////////////////////////////////////////////
const hostISPDB = "https://autoconfig.thunderbird.net/v1.1/"
const fileNameConfig = "ispdb.json"
//////////////////////////////////////////////////////////////
type configs []EmailProvider
type dictApi map[string]*EmailProvider

type ConfigDB struct {
	db      []EmailProvider
	ApiDict map[string]*EmailProvider
	locker 	sync.Mutex
}

//////////////////////////////////////////
type EmailProvider struct {
	EmailProvider struct {
		Domain         []string `xml:"domain" json:"Domain"`
		ImapPopServers []struct {
			Type           string `xml:"type,attr" json:"Type"`
			Hostname       string `xml:"hostname" json:"Hostname"`
			Port           string `xml:"port" json:"Port"`
			SocketType     string `xml:"socketType" json:"SocketType"`
			Username       string `xml:"username" json:"Username"`
			Authentication string `xml:"authentication" json:"Authentication"`
		} `xml:"incomingServer" json:"ImapPopServers"`
		SMTP struct {
			Type           string `xml:"type,attr" json:"Type"`
			Hostname       string `xml:"hostname" json:"Hostname"`
			Port           string `xml:"port" json:"Port"`
			SocketType     string `xml:"socketType" json:"SocketType"`
			Username       string `xml:"username" json:"Username"`
			Authentication string `xml:"authentication" json:"Authentication"`
		} `xml:"outgoingServer" json:"SMTP"`
		Enable struct {
			Instruction string `xml:"instruction" json:"Instruction"`
		} `xml:"enable" json:"Enable"`
	} `xml:"emailProvider" json:"EmailProvider"`
}

/////////////////////////////////////////////////////////////////

type Ispdb interface {
	Update(needUpdate bool,threads *int) (*ConfigDB, error)
	GetCount() int
}


