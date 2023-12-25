package main

import (
	"Blockchain/cli"
	"fmt"
	"os"
)

func printCmd() {
	fmt.Println("Command:")
	fmt.Println("\topen <port> \t\t\t\tOpen port to accept incomming connection")
	fmt.Println("\tconnect <port> \t\t\t\tConnect to a new peer")
	fmt.Println("\tpeers \t \t\t\t\tGet list of peers")
	fmt.Println("\tcreateblockchain \t\t\t\t\t Create a blockchain")
	fmt.Println("\tblockchain \t\t\t\tSee the current state of the blockchain")
	fmt.Println("\tblock <index> \t\t\t\tSee a specific block")
	fmt.Println("\tverify_block <index> <hash> \t\tVerify a block")
	fmt.Println("\tcheck_exist_transcation <data> \t\tCheck if a transaction exist")
	fmt.Println("\tmine <data> \t\t\t\tMine a new block")
	fmt.Println("\texit \t \t\t\t\tExit program")
}

func main() {
	defer os.Exit(0)
	//
	cmd := cli.CommandLine{}
	cmd.Run()
	// printCmd()
	// for {
	// 	fmt.Print("blockchain -> ")

	// 	scanner := bufio.NewScanner(os.Stdin)
	// 	scanner.Scan()
	// 	cmd := strings.TrimSpace(scanner.Text())

	// 	switch {
	// 	case strings.HasPrefix(cmd, "open"):
	// 		port, _ := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(cmd, "open")))
	// 		//fmt.Printf("Opening port %s\n", port)
	// 		network.OpenNode("127.0.0.1", port)
	// 	case strings.HasPrefix(cmd, "connect"):
	// 		addr := strings.TrimSpace(strings.TrimPrefix(cmd, "connect"))
	// 		//fmt.Printf("Connecting to port %s\n", port)
	// 		network.ConnectNode(addr)
	// 	case strings.HasPrefix(cmd, "createblockchain"):
	// 		blockchain.InitBlockchain()
	// 		fmt.Printf("New blockchain had been created!\n")
	// 	case cmd == "exit":
	// 		//fmt.Println("Exiting...")
	// 		//return
	// 	default:
	// 		fmt.Println("Invalid cmd. Please enter a valid option.")
	// 	}
	// }
}
