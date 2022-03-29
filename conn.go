package main

import (
	"context"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	maddr "github.com/multiformats/go-multiaddr"
)

func makeConn(ctx context.Context, listenAddresses []maddr.Multiaddr) (host.Host, error) {
	return libp2p.New(libp2p.ListenAddrs(listenAddresses...))
}

/*
	// Set your own keypair
	var priv, _, err = crypto.GenerateKeyPair(
		crypto.Ed25519, // Select your key type. Ed25519 are nice short
		-1,             // Select key length when possible (i.e. RSA).
	)
	if err != nil {
		return nil, err
	}

	var idht *dht.IpfsDHT

	return libp2p.New(
		// Use the keypair we generated
		libp2p.Identity(priv),

		// Multiple listen addresses
		libp2p.ListenAddrs(listenAddresses...), //"/ip4/0.0.0.0/udp/9000/quic", // a UDP endpoint for the QUIC transport

		// support QUIC
		libp2p.Transport(libp2pquic.NewTransport),

		// Let's prevent our peer from having too many
		// connections by attaching a connection manager.
		//libp2p.ConnectionManager(connmgr.NewConnManager(

		// Attempt to open ports using uPNP for NATed hosts.
		libp2p.NATPortMap(),

		// Let this host use the DHT to find other hosts
		libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {
			idht, err = dht.New(ctx, h)
			return idht, err
		}),

		// Let this host use relays and advertise itself on relays if
		// it finds it is behind NAT. Use libp2p.Relay(options...) to
		// enable active relays and more.
		libp2p.EnableAutoRelay(),
	)
}
*/
