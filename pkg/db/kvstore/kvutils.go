/*
 * Copyright 2018-present Open Networking Foundation

 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at

 * http://www.apache.org/licenses/LICENSE-2.0

 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package kvstore

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"time"
)

// GetDuration converts a timeout value from int to duration.  If the timeout value is
// either not set of -ve then we default KV timeout (configurable) is used.
func GetDuration(timeout int) time.Duration {
	if timeout <= 0 {
		return defaultKVGetTimeout * time.Second
	}
	return time.Duration(timeout) * time.Second
}

// ToString converts an interface value to a string.  The interface should either be of
// a string type or []byte.  Otherwise, an error is returned.
func ToString(value interface{}) (string, error) {
	switch t := value.(type) {
	case []byte:
		return string(value.([]byte)), nil
	case string:
		return value.(string), nil
	default:
		return "", fmt.Errorf("unexpected-type-%T", t)
	}
}

// ToByte converts an interface value to a []byte.  The interface should either be of
// a string type or []byte.  Otherwise, an error is returned.
func ToByte(value interface{}) ([]byte, error) {
	switch t := value.(type) {
	case []byte:
		return value.([]byte), nil
	case string:
		return []byte(value.(string)), nil
	default:
		return nil, fmt.Errorf("unexpected-type-%T", t)
	}
}

// GetAddress concatenates the Host and Port as single arguement.
func GetAddress(host string, port int) string {
	addr := host + ":" + strconv.Itoa(port)
	return addr
}

// ValidateAddress validates the host and port values
func ValidateAddress(address string) error {
	host, port, err := net.SplitHostPort(address)
	if err != nil {
		return errors.New("Invalid Format of address")
	}

	ip, err := net.LookupIP(host)
	if err != nil || len(ip) == 0 {
		return errors.New("Unknown host")
	}

	portInt, errValue := strconv.Atoi(port)
	if errValue != nil {
		return errors.New("Invalid address port")
	}

	if !(portInt <= 65535 && portInt >= 0) {
		return errors.New("Invalid address port range")
	}
	return nil
}
