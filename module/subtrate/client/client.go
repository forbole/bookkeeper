package subtrate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type SubtrateClient struct{
	api string
	client *http.Client
}

func NewSubtrateClient(api string) *SubtrateClient{
	c:=http.DefaultClient
	return &SubtrateClient{
		api:api,
		client:c,
	}
}


func (v SubtrateClient) CallApi(url string,payload interface{})error{

data := payload

payloadBytes, err := json.Marshal(data)
if err != nil {
	// handle err
}
body := bytes.NewReader(payloadBytes)

api:=fmt.Sprintf("https://%s.api.subscan.io/%s",v.api,url)

req, err := http.NewRequest("POST", api , body)
if err != nil {
	// handle err
}
req.Header.Set("Content-Type", "application/json")
req.Header.Set("X-Api-Key", "YOUR_KEY")

resp, err := http.DefaultClient.Do(req)
if err != nil {
	// handle err
}
defer resp.Body.Close()

var bz []byte
bz, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

err = json.Unmarshal(bz, &payload)
	if err != nil {
		return fmt.Errorf("Fail to marshal:%s", err)
	}

return nil
}