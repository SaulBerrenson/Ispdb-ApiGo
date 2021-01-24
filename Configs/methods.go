package configs

import (
	"MailConfigHandler/ErrorPkg"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/cheggaaa/pb/v3"
	"github.com/gammazero/workerpool"
)

func updateFromNetwork(needUpdate *bool, threads *int) (*ConfigDB, error) {
	log.Println("Trying to get list of domains...")
	resonse, err := getHml(hostISPDB)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer resonse.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resonse.Body)
	if err != nil {
		log.Fatal(err)
	}

	domains := make([]string, 0)

	doc.Find("body > table > tbody").Find("td:nth-child(2) > a").Each(func(i int, s *goquery.Selection) {
		if len(s.Nodes) <= 0 {
			return
		}
		if s.Nodes[0].Attr[0].Val == "/" {
			return
		}
		domains = append(domains, s.Text())
	})

	if len(domains) <= 0 {
		return nil, ErrorPkg.New("Not Found Domains!")
	}

	log.Println("Count for updating: ", strconv.Itoa(len(domains)))

	prgBar := pb.StartNew(len(domains))
	wp := workerpool.New(*threads)

	var mutex sync.Mutex
	c := ConfigDB{}
	for _, r := range domains {
		r := r
		wp.Submit(func() {
			//log.Println("Element ", strconv.Itoa(index), " is updating..")
			provider, err := FindConfig(r)
			if err != nil {
				return
			}
			mutex.Lock()
			c.db = append(c.db, *provider)
			prgBar.Increment()
			mutex.Unlock()
		})
	}

	wp.StopWait()
	prgBar.Finish()

	if len(c.db) <= 0 {
		return nil, ErrorPkg.New("Not Found Configs!")
	}

	log.Println("List of domain: ", strconv.Itoa(len(domains)))
	return &c, nil
}
func updateFromCache() (*ConfigDB, error) {

	file, err := os.Open(fileNameConfig)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	bytes, err := ioutil.ReadAll(file)

	config := CreateConfigDB()
	err = json.Unmarshal(bytes, &config.db)

	// = data
	return config, err
}

func (c ConfigDB) UpdateCache() (bool, error) {

	if len(c.db) <= 0 {
		return false, nil
	}

	var cache []EmailProvider
	cache = c.db
	jsonStr, err := json.Marshal(&cache)

	if err != nil {
		return false, err
	}
	if len(jsonStr) <= 0 {
		return false, nil
	}

	if err = os.Remove(fileNameConfig); err != nil {
		fmt.Println("CONFIG DELETE: ", err)
	}

	f, err := os.OpenFile(fileNameConfig,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(string(jsonStr)); err != nil {
		return false, err
	}
	return true, nil
}

func (c ConfigDB) Find(prediction string) (*EmailProvider, bool) {

	//c.locker.Lock()
	conf, exist := c.ApiDict[prediction]
	//c.locker.Unlock()
	return conf, exist
}

func (c ConfigDB) PrepareDict() (*ConfigDB, error) {

	for index, _ := range c.db {

		for i, _ := range c.db[index].EmailProvider.Domain {
			c.ApiDict[c.db[index].EmailProvider.Domain[i]] = &c.db[index]
		}
	}

	return &c, nil
}

func (c ConfigDB) Update(needUpdate *bool, threads *int) (*ConfigDB, error) {

	switch *needUpdate {
	case true:
		return updateFromNetwork(needUpdate, threads)
	case false:
		return updateFromCache()
	}

	return nil, nil
}

func (c ConfigDB) GetCount() int {
	if c.db == nil {
		return -1
	}
	return len(c.db)
}
