package main

import (
	"github.com/jormenjanssen/connectivity-manager/protocol"
	"github.com/jormenjanssen/connectivity-manager/wireless/protocol/supplicant"
)

type protocolFactory struct {
}

// GetProtocol return a protocol from a protocol name
func (pf protocolFactory) GetProtocol(name string) protocol.ConnectionProtocol {

	switch name {
	case "supplicant":
		return supplicant.Protocol{}
	default:
		return nil
	}
}
