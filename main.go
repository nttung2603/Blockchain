package main

import (
	"Blockchain/blockchain"
	"Blockchain/network"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func printCmd() {
	fmt.Println("Command:")
	fmt.Println("\topen <port> \t\t\t\tOpen port to accept incomming connection")
	fmt.Println("\tconnect <address> \t\t\t\tConnect to a new peer")
	fmt.Println("\tpeers \t \t\t\t\tGet list of peers")
	fmt.Println("\tblockchain \t\t\t\tSee the current state of the blockchain")
	fmt.Println("\tblock <index> \t\t\t\tSee a specific block")
	fmt.Println("\tverify_block <index> <hash> \t\tVerify a block")
	fmt.Println("\tcheck_exist_transcation <data> \t\tCheck if a transaction exist")
	fmt.Println("\tmine <data> \t\t\t\tMine a new block")
	fmt.Println("\texit \t \t\t\t\tExit program")
}

func main() {
	//defer os.Exit(0)
	//
	//cmd := cli.CommandLine{}
	//cmd.Run()
	printCmd()
	for {
		fmt.Print("blockchain -> ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		cmd := strings.TrimSpace(scanner.Text())

		switch {
		case strings.HasPrefix(cmd, "open"):
			port, _ := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(cmd, "open")))
			//fmt.Printf("Opening port %s\n", port)
			network.OpenNode("127.0.0.1", port)
		case strings.HasPrefix(cmd, "connect"):
			addr := strings.TrimSpace(strings.TrimPrefix(cmd, "connect"))
			//fmt.Printf("Connecting to port %s\n", port)
			network.ConnectNode(addr)
		case strings.HasPrefix(cmd, "broadcastBlock"):
			// Create a new blockchain with the genesis block
			transactions := []*blockchain.Transaction{
				{Data: []byte("Transaction 3")},
				{Data: []byte("Transaction 4")},
			}
			newBlock := blockchain.CreateBlock(transactions, []byte("123"))
			data, _ := newBlock.Serialize()
			network.BroadcastData(data)
		case cmd == "exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid cmd. Please enter a valid option.")
		}
	}
}
