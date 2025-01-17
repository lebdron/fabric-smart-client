/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package common

import (
	"context"
	"crypto/tls"
	"github.com/hyperledger-labs/fabric-smart-client/platform/fabric/core/generic/peer"
	grpc2 "google.golang.org/grpc"

	"github.com/hyperledger/fabric-protos-go/discovery"
	pb "github.com/hyperledger/fabric-protos-go/peer"
	"github.com/pkg/errors"

	"github.com/hyperledger-labs/fabric-smart-client/platform/view/services/grpc"
	discovery2 "github.com/hyperledger/fabric/discovery/client"
)

// PeerClient represents a client for communicating with a peer
type PeerClient struct {
	CommonClient
	Signer discovery2.Signer
}

func (pc *PeerClient) Close() {
	pc.CommonClient.Client.Close()
}

func (pc *PeerClient) Connection() (*grpc2.ClientConn, error) {
	conn, err := pc.CommonClient.NewConnection(pc.Address(), grpc.ServerNameOverride(pc.Sn))
	if err != nil {
		return nil, errors.WithMessagef(err, "endorser client failed to connect to %s", pc.Address())
	}
	return conn, nil
}

// Endorser returns a client for the Endorser service
func (pc *PeerClient) Endorser() (pb.EndorserClient, error) {
	conn, err := pc.CommonClient.NewConnection(pc.Address(), grpc.ServerNameOverride(pc.Sn))
	if err != nil {
		return nil, errors.WithMessagef(err, "endorser client failed to connect to %s", pc.Address())
	}
	return pb.NewEndorserClient(conn), nil
}

func (pc *PeerClient) Discovery() (discovery.DiscoveryClient, error) {
	conn, err := pc.CommonClient.NewConnection(pc.Address(), grpc.ServerNameOverride(pc.Sn))
	if err != nil {
		return nil, errors.WithMessagef(err, "discovery client failed to connect to %s", pc.Address())
	}
	return discovery.NewDiscoveryClient(conn), nil
}

func (pc *PeerClient) DiscoveryClient() (peer.DiscoveryClient, error) {
	return discovery2.NewClient(
		func() (*grpc2.ClientConn, error) {
			conn, err := pc.CommonClient.NewConnection(pc.Address(), grpc.ServerNameOverride(pc.Sn))
			if err != nil {
				return nil, errors.WithMessagef(err, "discovery client failed to connect to %s", pc.Address())
			}
			return conn, nil
		},
		pc.Signer,
		1), nil
}

func (pc *PeerClient) DeliverClient() (pb.DeliverClient, error) {
	conn, err := pc.CommonClient.NewConnection(pc.Address(), grpc.ServerNameOverride(pc.Sn))
	if err != nil {
		return nil, errors.WithMessagef(err, "endorser client failed to connect to %s", pc.Address())
	}
	return pb.NewDeliverClient(conn), nil
}

// Deliver returns a client for the Deliver service
func (pc *PeerClient) Deliver() (pb.Deliver_DeliverClient, error) {
	conn, err := pc.CommonClient.NewConnection(pc.Address(), grpc.ServerNameOverride(pc.Sn))
	if err != nil {
		return nil, errors.WithMessagef(err, "deliver client failed to connect to %s", pc.Address())
	}
	return pb.NewDeliverClient(conn).Deliver(context.TODO())
}

// Certificate returns the TLS client certificate (if available)
func (pc *PeerClient) Certificate() tls.Certificate {
	return pc.CommonClient.Certificate()
}

func (pc *PeerClient) Address() string {
	return pc.CommonClient.Address
}
