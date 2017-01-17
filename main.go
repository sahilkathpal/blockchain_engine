package main

import (
	"flag"

	. "github.com/tendermint/go-common"
	"github.com/sahilkathpal/blockchain_engine/middleware"
	"github.com/tendermint/abci/server"
	"github.com/tendermint/abci/types"
)

func main() {

	addrPtr := flag.String("addr", "tcp://0.0.0.0:46658", "Listen address")
	tmspPtr := flag.String("tmsp", "socket", "socket | grpc")
	urlPtr := flag.String("url", "http://localhost:3000", "Url of the smart contracts engine")
	flag.Parse()

	// Create the application - in memory or persisted to disk
	var app types.Application
	app = middleware.NewMiddlewareApplication(*urlPtr)

	// Start the listener
	srv, err := server.NewServer(*addrPtr, *tmspPtr, app)
	if err != nil {
		Exit(err.Error())
	}

	// Wait forever
	TrapSignal(func() {
		// Cleanup
		srv.Stop()
	})

}
