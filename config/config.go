package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	"github.com/auriou/godebridaria/models"
)

func New() *ClientConfig {
	client := &ClientConfig{File: "data.json"}
	client.Config = &models.StoreConfig{
		Activate:    &models.ActivePinResponse{},
		AskPin:      &models.PinResponse{},
		Aria2Config: &models.Aria2Config{},
	}
	client.Read()
	return client
}

func (c *ClientConfig) Save() {
	jsonString, _ := json.MarshalIndent(&c.Config, "", "\t")
	err := ioutil.WriteFile("data.json", jsonString, os.ModePerm)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func (c *ClientConfig) Read() error {
	file, err := ioutil.ReadFile(c.File)
	if err == nil {
		err = json.Unmarshal(file, &c.Config)
		if err != nil {
			fmt.Println("Error: ", err)
			return err
		}
	} else {
		return err
	}
	return nil
}

func (c *ClientConfig) SaveActivePin(activate *models.ActivePinResponse) {
	c.Config.Activate = activate
	c.Save()
}

func (c *ClientConfig) SaveHosts(hosts *models.DebridDomains) {
	c.Hosts = hosts
}

func (c *ClientConfig) GetHostType(link string) int {
	urlAsk, _ := url.Parse(link)
	for _, host := range c.Hosts.Data.Hosts {
		if host == urlAsk.Host {
			return 2
		}
	}
	for _, host := range c.Hosts.Data.Redirectors {
		if host == urlAsk.Host {
			return 1
		}
	}
	return 0
}

func (c *ClientConfig) SavePin(pin *models.PinResponse) {
	c.Config.AskPin = pin
	c.Save()
}

func (c *ClientConfig) SaveAria2Url(url string) {
	c.Config.Aria2Config.Aria2Url = url
	c.Save()
}

func (c *ClientConfig) SaveAria2Secret(secret string) {
	c.Config.Aria2Config.Aria2Secret = secret
	c.Save()
}

func (c *ClientConfig) SaveAddress(address string) {
	c.Config.Address = address
	c.Save()
}

func (c *ClientConfig) IsActivated() bool {
	return c.Config.Activate.Activated
}

func (c *ClientConfig) GetDebridToken() string {
	return c.Config.Activate.Apikey
}

func (c *ClientConfig) GetCheckURL() string {
	return c.Config.AskPin.CheckURL
}

func (c *ClientConfig) GetAria2Url() string {
	return c.Config.Aria2Config.Aria2Url
}

func (c *ClientConfig) GetAria2Secret() string {
	return c.Config.Aria2Config.Aria2Secret
}

func (c *ClientConfig) GetAddress() string {
	return c.Config.Address
}

func (c *ClientConfig) GetToken() string {
	return c.Config.Token
}
