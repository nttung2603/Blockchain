package network

import (
	"bufio"
	"fmt"
	net "github.com/libp2p/go-libp2p-core/network"
	"log"
)

const (
	protocol      = "tcp"
	version       = 1
	commandLength = 12
)

func handleStream(s net.Stream) {

	log.Println("Got a new stream!")

	// Create a buffer stream for non blocking read and write.
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))
	// Read a line from the stream
	command, err := rw.ReadString('\n')
	if err != nil {
		log.Panic(err)
	}
	switch command {
	case "addr":
		//HandleAddr(req)
	case "block":
		//HandleBlock(req, chain)
	case "inv":
		//HandleInv(req, chain)
	case "getblocks":
		//HandleGetBlocks(req, chain)
	case "getdata":
		//HandleGetData(req, chain)
	case "tx":
		//HandleTx(req, chain)
	case "version":
		//HandleVersion(req, chain)
	default:
		fmt.Println("Unknown command")
	}

	// stream 's' will stay open until you close it (or the other side closes it).
}
