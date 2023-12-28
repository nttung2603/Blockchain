package network

import (
	"Blockchain/blockchain"
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"log"

	"github.com/libp2p/go-libp2p/core/network"
)

const (
	protocol      = "tcp"
	version       = 1
	commandLength = 12
)

func CmdToBytes(cmd string) []byte {
	var bytes [commandLength]byte

	for i, c := range cmd {
		bytes[i] = byte(c)
	}

	return bytes[:]
}

func BytesToCmd(bytes []byte) string {
	var cmd []byte
	for _, b := range bytes {
		if b != 0x0 {
			cmd = append(cmd, b)
		}
	}
	return fmt.Sprintf("%s", cmd)
}
func writeBytes(stream network.Stream, data []byte) error {
	_, err := stream.Write(data)
	fmt.Println("Write data successfully")
	return err
}

func readBytes(rw *bufio.ReadWriter) ([]byte, error) {
	var result []byte
	buffer := make([]byte, 1024) // Kích thước buffer có thể điều chỉnh

	for {
		n, err := rw.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		result = append(result, buffer[:n]...)

		if n < len(buffer) {
			break
		}
	}
	return result, nil
}

func HandleGetBlock(request []byte) *blockchain.Block {
	var buff bytes.Buffer
	var block blockchain.Block

	buff.Write(request[commandLength:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}

// *blockchain.Blockchain
func HandleGetChain(request []byte) []byte {
	return blockchain.SerializeBase("base_chain")
	// blockchain.DeserializeBase(request[commandLength:], "cloning")
}

func HandleDownload(pid string, request []byte) {
	blockchain.DeserializeBase(request[commandLength:], pid)
}

func handleStream(s network.Stream) {

	fmt.Println("Got a new stream!")
	// Create a buffer stream for non blocking read and write.
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))
	data, _ := readBytes(rw)
	command := BytesToCmd(data[:commandLength])
	switch command {
	case "getBlock":
		chain := blockchain.GetChain("base_chain")
		fmt.Println("get chain", chain)
		block := HandleGetBlock(data)
		block.PrintBlock()
		chain.AppendBlock(block)
		fmt.Println("prepare to write", chain)
		blockchain.SetChain(chain, "base_chain")
		fmt.Println("done")
	case "getChain":
		chainbytes := HandleGetChain(data)
		fmt.Println("done get chain data")
		SendGetChainResponse("hihi", chainbytes)
		fmt.Println("send data back")
	case "download":
		HandleDownload(GetHost().ID().String(), data)
		fmt.Println("done cloning")
		//HandleTx(req, chain)
	case "version":
		//HandleVersion(req, chain)
	default:
		fmt.Println("Unknown command")
	}
	// stream 's' will stay open until you close it (or the other side closes it).
}
