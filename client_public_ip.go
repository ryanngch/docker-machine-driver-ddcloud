package main

/*
 * Detect the client machine's external IPv4 address
 * -------------------------------------------------
 *
 * Uses http://ifconfig.co/json
 */

import (
	"encoding/json"
	"fmt"
	"github.com/docker/machine/libmachine/log"
	"io/ioutil"
	"net/http"
)

// A subset of the IP address information returned by ifconfig.co.
type ipInfo struct {
	IPAddress string `json:"ip"`
}

// Retrieve the client machine's public IPv4 address.
func getClientPublicIPv4Address() (string, error) {
	log.Infof("Auto-detecting client's public (external) IP address...")

	response, err := http.Get("https://ifconfig.co/json")
	if err != nil {
		return "", fmt.Errorf("Unable to connect to ifconfig.co to determine your IP address: %s", err.Error())
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	info := &ipInfo{}
	err = json.Unmarshal(responseBody, info)
	if err != nil {
		return "", err
	}

	return info.IPAddress, nil
}
