package main

import (
	"github.com/jormenjanssen/connectivity-manager/protocol"
	"github.com/jormenjanssen/connectivity-manager/wireless/protocol/supplicant"
)

func main() {

}

// GetProtocol return a protocol from a protocol name
func GetProtocol(name string) protocol.ConnectionProtocol {
	return supplicant.Protocol{}
}
