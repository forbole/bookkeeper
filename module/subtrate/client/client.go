package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/forbole/bookkeeper/module/subtrate/types"
	"github.com/rs/zerolog/log"
)

type SubscanClient struct {
	api    string
	client *http.Client
}

func NewSubscanClient(api string) *SubscanClient {
	c := http.DefaultClient
	return &SubscanClient{
		api:    api,
		client: c,
	}
}

func (client SubscanClient) CallApi(url string, payload interface{}, v types.SubscanApi) error {
	log.Trace().Str("module", "subtrate").Msg("Query api")

	req,err:=client.makeRequest(url,payload)
	if err!=nil{
		return err
	}

	bz,err:=client.doRequest(req)
	
	for strings.Contains(string(bz),"rate limit")||err!=nil{
		log.Info().Str("module", "subtrate").Msg("API rate limit exceeded")
		time.Sleep(time.Minute)
		req,err:=client.makeRequest(url,payload)
		if err!=nil{
			return err
		}

		bz,err=client.doRequest(req)
		

	}
	err = json.Unmarshal(bz, &v)
	if err != nil {
		return fmt.Errorf("Fail to marshal:%s,bz:%s", err,string(bz))
	}

	return nil
}

func (client SubscanClient) doRequest(req *http.Request) ([]byte,error) {
	resp, err := client.client.Do(req)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()

	if resp.StatusCode!=200{
		return nil,fmt.Errorf("Status Code:%d,Status:%s",resp.StatusCode,resp.Status)
	}
	var bz []byte
	bz, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	return bz,nil
}

func (client SubscanClient) makeRequest(url string,payload interface{}) (*http.Request,error) {
	data := payload

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return nil,err
	}
	body := bytes.NewReader(payloadBytes)

	api := fmt.Sprintf("https://%s.api.subscan.io/%s", client.api, url)

	req, err := http.NewRequest("POST", api, body)
	if err != nil {
		return nil,err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Key", "YOUR_KEY")
	return req,nil
}
