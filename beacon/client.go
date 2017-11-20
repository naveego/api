package beacon

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// Client provides access to the Beacon API. Obtain an instance by calling NewClient.
type Client struct {
	c     *http.Client
	url   string
	token string
}

// NewClient returns a new Beacon API client which will make requests
// to the provided beaconURL using the given token.
func NewClient(beaconURL, token string) (Client, error) {

	_, err := url.Parse(beaconURL)

	if err != nil {
		return Client{}, fmt.Errorf("beaconUrl was invalid: %s", err)
	}

	return Client{
		c:     &http.Client{},
		url:   beaconURL,
		token: token,
	}, nil
}

// GetConfig returns the config with the provided ID.
func (c Client) GetConfig(id string) (Config, error) {

	config := Config{}

	var path string
	if id == "" {
		path = "/api/configs"
	} else {
		path = "/api/configs/" + id
	}

	bytes, err := c.get(path)
	if err != nil {
		return Config{}, err
	}

	err = json.Unmarshal(bytes, &config)

	return config, err
}

func (c Client) get(path string) ([]byte, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("%s%s", c.url, path), nil)
	if err != nil {
		log.Panicf("couldn't create request")
	}

	request.Header = http.Header{
		"Authorization": []string{"Bearer " + c.token},
	}

	response, err := c.c.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)

	return bytes, err
}

// GetConfig gets a Config instance from a full path.
// The path can be a path to a config in the Beacon service (like 'http://beacon/api/configs/xyz')
// or a file path (like 'file:///var/lib/naveego/agent.json').
// The token parameter is only needed if the config is secured.
func GetConfig(fullUrl, token string) (Config, error) {

	config := Config{}

	if strings.HasPrefix(fullUrl, "file://") {
		path := strings.Replace(fullUrl, "file://", "", 1)
		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			return config, err
		}
		err = json.Unmarshal(bytes, &config)
		return config, err
	}

	if strings.HasPrefix(fullUrl, "http") {
		parsedURL, err := url.Parse(fullUrl)
		if err != nil {
			return config, err
		}

		beaconUrl := fmt.Sprintf("%s://%s", parsedURL.Scheme, parsedURL.Host)

		client, err := NewClient(beaconUrl, token)
		if err != nil {
			return config, err
		}

		pathSegments := strings.Split(parsedURL.Path, "/")
		configID := pathSegments[len(pathSegments)]

		config, err = client.GetConfig(configID)
		return config, err
	}

	return config, errors.New("unrecognized protocol for fullUrl")

}
