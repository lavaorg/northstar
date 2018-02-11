/*
Copyright (C) 2017 Verizon. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package portalglobal

import (
	"fmt"
	"github.com/lavaorg/lrt/env"
	"github.com/lavaorg/lrt/x/mlog"
	"net"
)

const (
	ServiceName = "nsportal"
	// AccessTokenKeyName defines the name of the Access Token key.
	AccessTokenKeyName = "Ns.Auth.AccessToken"
)

var Config configuration

type configuration struct {
	ServiceName          string `default:"nsportal"`
	ServiceHostPort      string
	Port                 string `default:"8080"`
	NorthstarAPIProtocol string `default:"http"`
	NorthstarAPIHostPort string `require:"true"`
	AcctProtocol         string `default:"http"`
	AcctUserHostPort     string `require:"true"`
	AcctAuthHostPort     string `require:"true"`
	AcctClientID         string `default:"nsclient"`
	AcctClientSecret     string `default:"nssecret"`
	AcctUserScopes       string `default:"ts.user ts.user.ro ts.transformation ts.transformation.ro ts.notebook ts.notebook.ro ts.model.ro ts.nsobject.ro"`
	AcctClientScopes     string `default:"ts.configuration"`
	ConnectionBufferSize int    `default:"1024"`
}

// Load loads the configuration from the environment variables.
func Load() error {

	err := env.Load(ServiceName, &Config)
	if err != nil {
		return err
	}

	// Get the interface IP address to generate service host and port.
	if Config.ServiceHostPort == "" {
		if interfaceIP := getInterfaceIP(); interfaceIP != "" {
			Config.ServiceHostPort = fmt.Sprintf("%s:%s", interfaceIP, Config.Port)
		}
	}

	mlog.Info("environment settings:%v", Config)

	return nil
}

// getInterfaceIp is a helper method used to get interface ip address.
func getInterfaceIP() string {
	mlog.Debug("getInterfaceIp")

	// Determine the service IP address from the eth interface.
	interfaces, err := net.Interfaces()

	if err != nil {
		mlog.Error("Net interfaces returned error: %+v", err)
		return ""
	}

	for _, ethInterface := range interfaces {
		mlog.Debug("Interface: %+v", ethInterface)

		// Check if interface is the loop back.
		if ethInterface.Flags&net.FlagLoopback == 0 {

			// Get the address.
			if addresses, err := ethInterface.Addrs(); err == nil {
				for _, address := range addresses {
					// Our IPv4 address will have a `.` in it.
					if ip, _, err := net.ParseCIDR(address.String()); err == nil {
						if ip4 := ip.To4(); ip4 != nil {
							return ip4.String()
						}
					}
				}
			}
		}
	}

	return ""
}
