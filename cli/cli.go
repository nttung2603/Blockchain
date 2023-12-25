package cli

import (
	"Blockchain/blockchain"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
)

type CommandLine struct{}

func (cli *CommandLine) printUsage() {
	fmt.Println("Commands:")
	fmt.Println("\topen <port> \t\t\t\tOpen port to accept incomming connection")
	fmt.Println("\tconnect <port> \t\t\t\tConnect to a new peer")
	fmt.Println("\tpeers \t\t\t\t\tGet list of peers")
	fmt.Println("\tcreateblockchain <address> \t\t\t Create a blockchain")
	fmt.Println("\tmine <address>\t\t\t\t Mine a block")
}

func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}
}

func (cli *CommandLine) OpenNode(port string) {
	addr := fmt.Sprintf("/ip4/%s/tcp/%s", "127.0.0.1", port)
	//network.StartServer(addr)
	fmt.Printf("Open Node on address %s\n", addr)
}

func (cli *CommandLine) createBlockChain(address string) {
	blockchain.InitBlockchain(address)
	fmt.Printf("Finished. New Blockchain had been successfully created by %s!\n", address)
}

func (cli *CommandLine) addBlock(address string) {
	transData := blockchain.FakeTransactionData()
	fmt.Println("Creating fake Transaction data! Mining...")
	blockchain.MineBlock(address, transData)
}

func (cli *CommandLine) Run() {
	cli.validateArgs()
	// nodeID := os.Getenv("NODE_ID")
	// if nodeID == "" {
	// 	fmt.Printf("NODE_ID env. var is not set!")
	// 	os.Exit(1)
	// }
	// fmt.Println(nodeID)
	openNodeCmd := flag.NewFlagSet("open", flag.ExitOnError)
	connectNodeCmd := flag.NewFlagSet("connect", flag.ExitOnError)
	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	mineCmd := flag.NewFlagSet("mine", flag.ExitOnError)

	openNode := openNodeCmd.String("port", "", "Open port to accept incomming connection. Eg: open 5000")
	connectNode := connectNodeCmd.String("port", "", "Connect port to accept incomming connection. Eg: open 5000")
	createBlockchainAddress := createBlockchainCmd.String("address", "", "hihi")
	mineAddress := mineCmd.String("address", "", "hihi")

	switch os.Args[1] {
	case "open":
		err := openNodeCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "connect":
		err := connectNodeCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	case "createblockchain":
		err := createBlockchainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	case "mine":
		err := mineCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	case "sendBlock":

	default:
		cli.printUsage()
		runtime.Goexit()
	}

	if openNodeCmd.Parsed() {
		cli.OpenNode(*openNode)
	}

	if openNodeCmd.Parsed() {
		cli.OpenNode(*connectNode)
	}
	if createBlockchainCmd.Parsed() {
		cli.createBlockChain(*createBlockchainAddress)
	}
	if mineCmd.Parsed() {
		cli.addBlock(*mineAddress)
	}

	//if connectNodeCmd.Parsed() {
	//	nodeID := os.Getenv("NODE_ID")
	//	if nodeID == "" {
	//		openNodeCmd.Usage()
	//		runtime.Goexit()
	//	}
	//	cli.OpenNode(nodeID, *openNode)
	//}
}
