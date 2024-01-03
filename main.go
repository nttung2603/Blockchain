package main

//import (
//	"Blockchain/transaction"
//	"Blockchain/network"
//	"bufio"
//	"fmt"
//	"os"
//	"strconv"
//	"strings"
//)
//
//func printCmd() {
//	fmt.Println("Command:")
//	fmt.Println("\topen <port> \t\t\t\tOpen port to accept incomming connection")
//	fmt.Println("\tconnect <address> \t\t\tConnect to a new peer")
//	fmt.Println("\tpeers \t\t\t\t\tGet list of peers")
//	fmt.Println("\tcreate_blockchain \t\t\tCreate a transaction")
//	fmt.Println("\tload_blockchain <file> \t\t\tLoad a transaction from a file")
//	fmt.Println("\tsave_blockchain \t\t\tSave a transaction to a file")
//	fmt.Println("\ttransaction \t\t\t\tSee the current state of the transaction")
//	fmt.Println("\tblock <index> \t\t\t\tSee a block in transaction at index")
//	fmt.Println("\tmine <tx_1>,<tx_2>,...,<tx_n>  \t\tMine a new block")
//	fmt.Println("\tclone <pid> \t\t\t\tClone transaction from a peer")
//	fmt.Println("\texit \t \t\t\t\tExit program")
//}
//
//func main() {
//	printCmd()
//	for {
//		fmt.Print("transaction -> ")
//
//		scanner := bufio.NewScanner(os.Stdin)
//		scanner.Scan()
//		cmd := strings.TrimSpace(scanner.Text())
//
//		switch {
//		case strings.HasPrefix(cmd, "open"):
//			port, _ := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(cmd, "open")))
//			network.OpenNode("127.0.0.1", port)
//
//		case strings.HasPrefix(cmd, "connect"):
//			addr := strings.TrimSpace(strings.TrimPrefix(cmd, "connect"))
//			network.ConnectNode(addr)
//
//		case strings.HasPrefix(cmd, "peers"):
//			node := network.GetHost()
//			connections := node.Network().Conns()
//			fmt.Println("Address: ", node.Addrs()[0])
//			fmt.Println("ID: ", node.ID().String())
//			fmt.Println("Number of nodes connected: ", len(connections))
//			for _, conn := range connections {
//				fmt.Println("- Node address: ", conn.RemoteMultiaddr())
//				fmt.Println("- Node ID: ", conn.RemotePeer())
//			}
//		case strings.HasPrefix(cmd, "create_blockchain"):
//			transaction.InitBlockchain()
//
//		case strings.HasPrefix(cmd, "load_blockchain"):
//			fileName := strings.TrimSpace(strings.TrimPrefix(cmd, "load_blockchain"))
//			transaction.ReadChain(fileName)
//
//		case strings.HasPrefix(cmd, "save_blockchain"):
//			pid := network.GetHost().ID().String()
//			transaction.WriteChain(pid)
//
//		case strings.HasPrefix(cmd, "transaction"):
//			//chain := transaction.GetChain(network.GetHost().ID().String())
//			//chain := transaction.ReadChain("base_chain")
//			chain := transaction.GetChain()
//			chain.PrintChain()
//
//		case strings.HasPrefix(cmd, "block"):
//			//chain := transaction.ReadChain("base_chain")
//			chain := transaction.GetChain()
//			index, _ := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(cmd, "block")))
//			chain.PrintBlock(index)
//
//		case strings.HasPrefix(cmd, "mine"):
//			// Create a new transaction with the genesis block
//			listTransactionStr := strings.TrimSpace(strings.TrimPrefix(cmd, "mine"))
//			txList := strings.Split(listTransactionStr, ",")
//			var transactions []*transaction.Transaction
//			for _, tx := range txList {
//				fmt.Println(tx)
//				transactions = append(transactions, &transaction.Transaction{Data: []byte(tx)})
//			}
//			chain := transaction.GetChain()
//			newBlock := transaction.CreateBlock(transactions, chain.GetPrevHash())
//			data, _ := newBlock.Serialize()
//			network.BroadcastData(data)
//
//		case strings.HasPrefix(cmd, "clone"):
//			pidClone := strings.TrimSpace(strings.TrimPrefix(cmd, "clone"))
//			network.SendGetChainRequest(pidClone)
//
//		case cmd == "exit":
//			fmt.Println("Exiting...")
//			return
//
//		default:
//			fmt.Println("Invalid cmd. Please enter a valid option.")
//		}
//	}
//}
