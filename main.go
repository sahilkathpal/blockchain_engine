package main

import (
	"flag"

	. "github.com/tendermint/go-common"
	"github.com/sahilkathpal/blockchain_engine"
	"github.com/tendermint/tmsp/server"
	"github.com/tendermint/tmsp/types"
)

func main() {

	addrPtr := flag.String("addr", "tcp://0.0.0.0:46658", "Listen address")
	tmspPtr := flag.String("tmsp", "socket", "socket | grpc")
	persistencePtr := flag.String("persist", "", "directory to use for a database")
	flag.Parse()

	// Create the application - in memory or persisted to disk
	var app types.Application
	app = middleware.NewMiddlewareApplication()

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
