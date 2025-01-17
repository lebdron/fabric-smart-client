/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package driver

import (
	"github.com/hyperledger-labs/fabric-smart-client/platform/view/services/grpc"
	"github.com/hyperledger-labs/fabric-smart-client/platform/view/view"
)

type TxID struct {
	Nonce   []byte
	Creator []byte
}

// ChaincodeInvocation models a client-side chaincode invocation
type ChaincodeInvocation interface {
	Endorse() (Envelope, error)

	Query() ([]byte, error)

	Submit() (string, []byte, error)

	WithTransientEntry(k string, v interface{}) ChaincodeInvocation

	WithEndorsers(ids ...view.Identity) ChaincodeInvocation

	WithEndorsersByMSPIDs(mspIDs ...string) ChaincodeInvocation

	WithEndorsersFromMyOrg() ChaincodeInvocation

	WithSignerIdentity(id view.Identity) ChaincodeInvocation

	WithTxID(id TxID) ChaincodeInvocation

	WithEndorsersByConnConfig(ccs ...*grpc.ConnectionConfig) ChaincodeInvocation

	WithImplicitCollections(mspIDs ...string) ChaincodeInvocation
}

// ChaincodeDiscover models a client-side chaincode's endorsers discovery operation
type ChaincodeDiscover interface {
	Call() ([]view.Identity, error)
	WithFilterByMSPIDs(mspIDs ...string) ChaincodeDiscover
	WithImplicitCollections(mspIDs ...string) ChaincodeDiscover
}

// Chaincode exposes chaincode-related functions
type Chaincode interface {
	NewInvocation(function string, args ...interface{}) ChaincodeInvocation
	NewDiscover() ChaincodeDiscover
	IsAvailable() (bool, error)
	IsPrivate() bool
}

// ChaincodeManager manages chaincodes
type ChaincodeManager interface {
	// Chaincode returns a chaincode handler for the passed chaincode name
	Chaincode(name string) Chaincode
}
