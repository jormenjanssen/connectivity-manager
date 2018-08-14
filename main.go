package main

import (
	"context"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	pf := protocolFactory{}

	protocol := pf.GetProtocol("supplicant")
	protocol.Initialize(ctx)

	time.Sleep(10 * time.Second)

	cancel()

	time.Sleep(5 * time.Second)

}
