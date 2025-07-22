package node

import (
	"crypto/ecdsa"
	"math/big"

	geth "github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/rpc"
)

type Node struct {
	eth     *geth.Node
	chainID *big.Int

	supportedConsensusContract         []common.Address
	supportedExtensionRegistryContract common.Address
	PrivateKey                         *ecdsa.PrivateKey // For signing messages and transactions

	ExtensionPool map[string]*Extension // Registered Extensions (ExtensionID -> Extension)
	dirLock       *flock.Flock          // prevents concurrent use of instance directory

	OperationPool map[string]*Operation // Hash -> Operation
	PostOpPool    map[string]*PostOp    // Hash  ->PostOp

	proofPool  map[string]*Proof  // Hash  -> Proof
	Reputation *ReputationManager // Tracks node + extension reputation
	config     *Config            // Custom config (timeouts, process window ...)
	P2P        *p2p.Server        // Handles gossip and discovery
	rpcAPIs    []rpc.API          // List of APIs currently provided by the node
	endpoints  *servers           // Node endpoints for rpc handling

	ProofVerifier *ProofVerifier // Local verifier interface
	TxSubmitter   *TxSubmitter   // For L1 tx submission logic

	state     int            // Tracks state of node lifecycle
	EventChan chan NodeEvent // Internal event bus (Operation received, Proof verified)
	databases *Databases     // Database handles for operations, proofs, etc.
	Shutdown  chan struct{}
}
