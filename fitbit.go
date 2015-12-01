// Package fitbit overview
package fitbit

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"golang.org/x/oauth2"
)

// Config get *oauth2.Config from env variables
func Config() *oauth2.Config {
	c := &oauth2.Config{
		ClientID:     os.Getenv("FITBIT_CLIENT_ID"),
		ClientSecret: os.Getenv("FITBIT_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("FITBIT_REDIRECT_URL"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  os.Getenv("FITBIT_AUTH_URL"),
			TokenURL: os.Getenv("FITBIT_TOKEN_URL"),
		},
		Scopes: strings.Split(os.Getenv("FITBIT_SCOPE"), ","),
	}
	return c
}

// Fitbit hogehoge
type Fitbit struct {
	config *oauth2.Config
	token  *oauth2.Token
}

// Client hogehoge
type Client struct {
	httpClient *http.Client
	Activity   *Activity
}

// SetConfig set *oauth2.Config
func (f *Fitbit) SetConfig(config *oauth2.Config) {
	f.config = config
}

// AuthURL return authorize url
func (f *Fitbit) AuthURL() (string, error) {
	return f.config.AuthCodeURL(""), nil
}

// SetTokenFromFile Read Token file And Set
func (f *Fitbit) SetTokenFromFile(filename string) error {
	_, err := os.Stat(filename)
	if err != nil {
		return err
	}
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	text, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	token := &oauth2.Token{}
	err = json.Unmarshal(text, token)
	if err != nil {
		return err
	}
	f.token = token
	return nil
}

// GetToken get *oauth2.Token
func (f *Fitbit) GetToken() (*oauth2.Token, error) {
	if f.token == nil {
		return nil, errors.New("tokenがnilです")
	}
	return f.token, nil
}

// ExchangeToken oauth token exchange code
func (f *Fitbit) ExchangeToken(code string) error {
	if f.config == nil {
		return errors.New("configがnilです")
	}

	token, err := f.config.Exchange(oauth2.NoContext, code)
	if err != nil {
		return err
	}
	f.token = token
	return nil
}

// Client return fitbit http client
func (f *Fitbit) Client() (*Client, error) {
	if f.config == nil || f.token == nil {
		return nil, errors.New("configがtokenのいずれかがnil")
	}
	client := &Client{httpClient: f.config.Client(oauth2.NoContext, f.token)}
	client.Activity = &Activity{c: client}
	return client, nil
}

// Get do GetRequest specific url
func (c *Client) Get(url string) ([]byte, error) {
	result, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()

	responseByteArray, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}
	return responseByteArray, nil
}
