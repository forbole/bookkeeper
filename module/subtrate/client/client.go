package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
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
	c := http.Client{}
	return &SubscanClient{
		api:    api,
		client: &c,
	}
}

func (client SubscanClient) CallApi(url string, payload interface{}, v types.SubscanApi) error {
	// Prevent call limit
	time.Sleep(time.Second)

	log.Trace().Str("module", "subtrate").Msg("CallApi")

	req, err := client.makeRequest(url, payload)
	if err != nil {
		return fmt.Errorf("Cannot make request:%s", err)
	}

	resp, err := client.client.Do(req)
	if err != nil {
		return fmt.Errorf("cannot do request:%s", err)
	}
	defer resp.Body.Close()

	for resp.StatusCode == 429 || resp.StatusCode == 1015 {
		log.Info().Str("module", "subtrate").Msg(fmt.Sprintf("API too many request"))
		retry := resp.Header.Get("Retry-After")
		retryInt, err := strconv.Atoi(retry)
		if err != nil {
			return fmt.Errorf("cannot parse int:%s,%s", retry, err)
		}
		fmt.Println(retryInt)
		time.Sleep(time.Duration(retryInt) * time.Second)

		return client.CallApi(url, payload, v)
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Status Code:%d,Status:%s", resp.StatusCode, resp.Status)
	}

	var bz []byte
	bz, err = io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("cannot do request:%s", err)
	}

	fmt.Println(string(bz))

	for strings.Contains(string(bz), "rate limit") {
		log.Info().Str("module", "subtrate").Msg(fmt.Sprintf("API rate limit exceeded %s", err.Error()))
		time.Sleep(time.Minute)

		return client.CallApi(url, payload, v)
	}

	err = json.Unmarshal(bz, &v)
	if err != nil {
		return fmt.Errorf("Fail to marshal:%s,bz:%s", err, string(bz))
	}

	return nil
}

func (client SubscanClient) makeRequest(url string, payload interface{}) (*http.Request, error) {
	data := payload

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("cannot make request:%s", err)
	}
	body := bytes.NewReader(payloadBytes)

	api := fmt.Sprintf("https://%s.api.subscan.io/%s", client.api, url)

	req, err := http.NewRequest("POST", api, body)
	if err != nil {
		return nil, fmt.Errorf("cannot make request:%s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Key", os.Getenv("SUBSCAN_API_KEY"))
	return req, nil
}
