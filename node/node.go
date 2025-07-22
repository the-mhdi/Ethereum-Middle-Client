package node

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	geth "github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/rpc"
)

// Node represents a middle client node that interacts with the Ethereum network and other middle clients, it's to maintain the p2p network and API endpoints.
type Node struct {
	eth     *geth.Node
	chainID *big.Int

	supportedConsensusContracts        []common.Address
	supportedExtensionRegistryContract common.Address

	wallet *accounts.Manager // for tx signing and account management

	dirLock *flock.Flock // prevents concurrent use of instance directory

	ExtensionPool map[string]*Extension // Registered Extensions (ExtensionID -> Extension)

	Monitor *Monitor // constantly monitors the node's state and health.

	Reputation *ReputationManager // Tracks node + extension reputation + participating in peer challenge-response protocol
	config     *Config            // Custom config (timeouts, process window ...)
	P2P        *p2p.Server        // Handles gossip and discovery
	rpcAPIs    []rpc.API          // List of APIs currently provided by the node
	endpoints  *servers           // Node endpoints for rpc handling

	state     int         // Tracks state of node lifecycle
	Event     *event.Feed // Internal event bus (Operation received, Proof verified)
	databases *Databases  // Database handles for operations, proofs, etc.
	Shutdown  chan struct{}
}
