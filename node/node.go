package node

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/rpc"
)

type Node struct {
	NodeID     string
	PrivateKey *ecdsa.PrivateKey // For signing messages and transactions

	Extensions map[string]*Extension // Registered Extensions (ExtensionID -> Extension)
	Peers      map[string]*PeerInfo  // Peer routing table

	OperationPool map[string]*Operation // Hash -> Operation
	PostOpPool    map[string]*PostOp    // Hash  ->PostOp

	proofPool  map[string]*Proof  // Hash  -> Proof
	Reputation *ReputationManager // Tracks node + extension reputation
	config     *Config            // Custom config (timeouts, process window, etc.)
	P2P        *p2p.Server        // Handles gossip and discovery
	rpcAPIs    []rpc.API          // List of APIs currently provided by the node
	endpoints  *servers           // Node endpoints for rpc handling

	ProofVerifier *ProofVerifier // Local verifier interface
	TxSubmitter   *TxSubmitter   // For L1 tx submission logic

	EventChan chan NodeEvent // Internal event bus (Operation received, Proof verified)
	Shutdown  chan struct{}
}
