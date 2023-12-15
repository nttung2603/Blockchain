package network

import (
	"context"
	"fmt"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	peerstore "github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
	"log"
)

var (
	node host.Host
)

func OpenNode(host string, port int) {
	addr := fmt.Sprintf("/ip4/%s/tcp/%d", host, port)
	multiaddr, err := multiaddr.NewMultiaddr(addr)
	fmt.Println(multiaddr)
	if err != nil {
		log.Panic(err)
	}
	node, err = libp2p.New(libp2p.ListenAddrs(multiaddr))
	if err != nil {
		log.Panic(err)
	}
	// print the node's PeerInfo in multiaddr format
	peerInfo := peerstore.AddrInfo{
		ID:    node.ID(),
		Addrs: node.Addrs(),
	}
	addrs, err := peerstore.AddrInfoToP2pAddrs(&peerInfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("libp2p node address:", addrs[0])
	//defer node.Close()
	fmt.Printf("Open Node on address %s\n", node.ID())
	fmt.Printf("Open Node on address %s\n", node.Addrs())
}

func ConnectNode(addr string) {
	//addr := fmt.Sprintf("/ip4/%s/tcp/%d", host, port)
	multiaddr, err := multiaddr.NewMultiaddr(addr)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(multiaddr)
	peer, err := peerstore.AddrInfoFromP2pAddr(multiaddr)
	if err != nil {
		panic(err)
	}

	err = node.Connect(context.Background(), *peer)
	if err != nil {
		log.Panic(err)
	}
}
