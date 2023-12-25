package cli

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
)

type CommandLine struct{}

func (cli *CommandLine) printUsage() {
	fmt.Println("Commands:")
	fmt.Println("open <port>")
	fmt.Println("connect <port>")
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

func (cli *CommandLine) Run() {
	cli.validateArgs()

	openNodeCmd := flag.NewFlagSet("open", flag.ExitOnError)
	connectNodeCmd := flag.NewFlagSet("connect", flag.ExitOnError)

	openNode := openNodeCmd.String("port", "", "Open port to accept incomming connection. Eg: open 5000")
	//connectNode := connectNodeCmd.String("port", "", "Connect port to accept incomming connection. Eg: open 5000")

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

	case "sendBlock":


	default:
		cli.printUsage()
		runtime.Goexit()
	}

	if openNodeCmd.Parsed() {
		cli.OpenNode(*openNode)
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
