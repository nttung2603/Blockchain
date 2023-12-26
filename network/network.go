package network

import (
	"context"
	"fmt"
	"log"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	_ "github.com/libp2p/go-libp2p/core/network"
	peerstore "github.com/libp2p/go-libp2p/core/peer"
	_ "github.com/libp2p/go-libp2p/p2p/protocol/ping"
	"github.com/multiformats/go-multiaddr"
)

var (
	node host.Host
)

func GetHost() host.Host {
	return node
}

func OpenNode(host string, port int) {
	opts := []libp2p.Option{
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/%s/tcp/%d", host, port)),
		//libp2p.Identity(priv),
	}
	//node, err = libp2p.New(libp2p.ListenAddrs(multiaddr))
	node, _ = libp2p.New(opts...)
	//print the node's PeerInfo in multiaddr format
	peerInfo := peerstore.AddrInfo{
		ID:    node.ID(),
		Addrs: node.Addrs(),
	}
	addrs, err := peerstore.AddrInfoToP2pAddrs(&peerInfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("libp2p node address:", addrs[0])
	node.SetStreamHandler("/p2p/1.0.0", handleStream)
}

func ConnectNode(addr string) {
	multiaddr, err := multiaddr.NewMultiaddr(addr)
	if err != nil {
		log.Panic(err)
	}
	peerAddrInfo, err := peerstore.AddrInfoFromP2pAddr(multiaddr)
	if err != nil {
		panic(err)
	}

	err = node.Connect(context.Background(), *peerAddrInfo)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Connected to", peerAddrInfo.String())

}

func BroadcastData(data []byte) {
	connections := node.Network().Conns()
	for _, conn := range connections {
		connectedPeerID := conn.RemotePeer()
		stream, _ := node.NewStream(context.Background(), connectedPeerID, "/p2p/1.0.0")
		request := append(CmdToBytes("getBlock"), data...)
		go writeBytes(stream, request)
		fmt.Println("Broadcast data to pid", connectedPeerID, "successfully")
	}
}
