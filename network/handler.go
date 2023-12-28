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
	pidLength     = 100
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

func PidToBytes(cmd string) []byte {
	var bytes [pidLength]byte

	for i, c := range cmd {
		bytes[i] = byte(c)
	}

	return bytes[:]
}
func BytesToPid(bytes []byte) string {
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
	//fmt.Println("Write data successfully")
	return err
}

func readBytes(rw *bufio.ReadWriter) ([]byte, error) {
	var result []byte
	buffer := make([]byte, 2048) // Kích thước buffer có thể điều chỉnh

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
func HandleGetChain(address string) []byte {
	return blockchain.SerializeBase(address)
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
		chain := blockchain.GetChain()
		//fmt.Println("get chain", chain)
		block := HandleGetBlock(data)
		block.PrintBlock()
		chain.AppendBlock(block)
		blockchain.SetChain(chain)

	case "getChain":
		pid := BytesToPid(data[commandLength:])
		blockchain.WriteChain(GetHost().ID().String())
		chainbytes := HandleGetChain(GetHost().ID().String())
		SendGetChainResponse(pid, chainbytes)

	case "download":
		HandleDownload(GetHost().ID().String(), data)
		blockchain.ReadChain(GetHost().ID().String())

	case "version":
		//HandleVersion(req, chain)

	default:
		fmt.Println("Unknown command")
	}
	fmt.Print("blockchain -> ")
}
