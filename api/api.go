package api

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
        "crypto/tls"
	"github.com/bitly/go-simplejson"
)

func Request(req *http.Request) (*simplejson.Json, error) {
        tr := &http.Transport{
           TLSClientConfig: &tls.Config{ InsecureSkipVerify: true },
        }
	httpclient := &http.Client{Transport: tr}
	//httpclient := &http.Client{  }
        resp, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		log.Printf("got response code %d - %s", resp.StatusCode, body)
		return nil, errors.New("api request returned non 200 status code")
	}
	data, err := simplejson.NewJson(body)
	if err != nil {
		return nil, err
	}
        log.Printf("Data response %s",data)
	return data, nil
}
