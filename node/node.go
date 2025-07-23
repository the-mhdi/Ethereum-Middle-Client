package node

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	geth "github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
)

// Node represents a middle client node that interacts with the Ethereum network and other middle clients, it's to maintain the p2p network and API endpoints.
type Node struct {
	eth     *geth.Node
	chainID *big.Int
	config  *Config
	Event   *event.Feed // Internal event bus (Operation received, Proof verified)

	supportedConsensusContracts        []common.Address
	supportedExtensionRegistryContract common.Address

	wallet *accounts.Manager // for tx signing and account management

	dirLock *flock.Flock // prevents concurrent use of instance directory

	// ExtensionPool map[string]*Extension // Registered Extensions (ExtensionID -> Extension)

	Monitor    *Monitor                      // constantly monitors the node's state and health.
	Reputation *reputation.ReputationManager // Tracks node + extension reputation + participating in peer challenge-response protocol

	P2P       *p2p.Server // Handles gossip and discovery
	endpoints *servers    // Node endpoints for rep requset handling

	db Database

	state    int // Tracks state of node lifecycle
	Shutdown chan struct{}
}
