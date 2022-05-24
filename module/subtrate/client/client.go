package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type SubscanClient struct{
	api string
	client *http.Client
}

func NewSubscanClient(api string) *SubscanClient{
	c:=http.DefaultClient
	return &SubscanClient{
		api:api,
		client:c,
	}
}


func (client SubscanClient) CallApi(url string,payload interface{},v interface{})error{

	data := payload

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	api:=fmt.Sprintf("https://%s.api.subscan.io/%s",client.api,url)

	req, err := http.NewRequest("POST", api , body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Key", "YOUR_KEY")

	resp, err := client.client.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	var bz []byte
	bz, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

	err = json.Unmarshal(bz, &v)
		if err != nil {
			return fmt.Errorf("Fail to marshal:%s", err)
		}

	return nil
}