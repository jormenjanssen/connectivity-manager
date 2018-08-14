package main

import (
	"fmt"

	"github.com/jormenjanssen/connectivity-manager/wireless"
)

func main() {
	var info = wireless.GetInfo()
	fmt.Printf(info.Bssid)
}
