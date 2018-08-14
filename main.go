package main

import (
	"context"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	pf := protocolFactory{}

	protocol := pf.GetProtocol("supplicant")
	protocol.Initialize(ctx)

	cancel()

}
