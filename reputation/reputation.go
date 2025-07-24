package reputation

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"sync"
	"time"
)

const (
	OK        = 0
	THROTTLED = 1
	BANNED    = 2
)

type ReputationParams interface {
	EntityType() string // Returns the type of entity (e.g., "node", "extension") just a marker method to unify the types.

}

type ReputaionManager struct {
	Reputation map[string]ReputationParams //entity-> reputaion

	db ethdb.Database //ethdb.Database                       //local db
	mu sync.RWMutex
}

type ExtensionReputationParams struct {
	//ValidProofRate
	// ValidProofCount / (ValidProofCount+InvalidProofCount) MUST BE > 0.8, too sensitive to initial failures, two defaults: prior_valid_proofs and prior_invalid_proofs
	// Score = (ValidProofCount + prior_valid_proofs) / (ValidProofCount + InvalidProofCount + prior_valid_proofs + prior_invalid_proofs)

	ValidProofCount   int // Number of valid proofs submitted
	InvalidProofCount int // Number of invalid proofs submitted

	// OperationAcceptanceRate
	//if OperationAcceptanceRate < 0.6, then the extension is considered unreliable and socket connection is closed
	OperationAcceptanceCount int // Number of Operations accepted by the extension
	OperationRejectionCount  int // Number of Operations rejected by the extension

	// latency is relative to the specific ExtensionID node decides this on local registration,
	// if latency > OperationExecutionLatency, then the extension is considered unreliable and therefore throttled.
	OperationExecutionLatency int // Time taken to process an Operation into PostOp

	Staked          bool
	StakeBalance    uint64 // Amount of stake held by the node or extension
	NegativeSlashes int    // Number of times the node or extension was penalized for malicious behavior

	LastActiveTimestamp time.Time // Last time the node or extension was active

	Blacklisted      bool
	BlacklistedUntil time.Time // If blacklisted, the time until which it is blacklisted

}

type PeerReputationParams struct {
	ValidProofCount   int // Number of valid proofs submitted to p2p mempool my the peer
	InvalidProofCount int // Number of invalid proofs submitted to p2p mempool my the peer

	// OperationAcceptanceRate
	//if OperationAcceptanceRate < 0.6, then the peer is considered unreliable and socket connection is closed
	OperationAcceptanceCount int // Number of Operations accepted by the peer
	OperationRejectionCount  int // Number of Operations rejected by the peer

	ProofVerificationLatency time.Duration // Time taken to verify a proof //middle node only
	AvailabilityScore        float64       // % uptime responding to requests // middle node only // > 0.8 okay, < 0.8 throttled, < 0.5 banned

	// DisputeOutcomeScore > 0.8 okay // DisputeOutcomeScore < 0.8 throttled , // DisputeOutcomeScore < 0.5 banned
	DisputeOutcomeScore float64 // % of times a node was challenged and proved correct vs incorrect

	staked          bool
	StakeBalance    uint64 // Amount of stake held by the node or extension
	NegativeSlashes int    // Number of times the node or extension was penalized for malicious behavior

	PeerEndorsements []string // List of peer endorsements or challenges

	LastActiveTimestamp time.Time // Last time the node or extension was active

	Blacklisted      bool
	BlacklistedUntil time.Time // If blacklisted, the time until which it is blacklisted

}

// erc-4337 style reputation manager
type WalletReputationParams struct {
	WalletAddress string

	// Operation success/failure tracking
	SuccessfulOps     int
	FailedOps         int
	InvalidSignatures int
	RejectionRate     float64
	LastSubmitted     time.Time

	// Gas behavior
	TotalFeesPaid  uint64
	AverageTip     uint64
	GasUnderpriced int     // # of times ops failed due to low gas (griefing attempts)
	GasGriefScore  float64 // 0.0–1.0; >0.7 indicates intentional gas grief

	// Paymaster-related tracking
	UsedPaymasters        map[string]*PaymasterStats
	MisbehavingPaymasters int // Count of failed paymaster validations
	PaymasterGriefScore   float64

	// Offchain endorsements
	ZkReputationScore    float64  // Derived from zkRep or similar systems [0.0 - 1.0]
	OffchainEndorsements []string // e.g. DIDs, signature proofs, reputation anchors

	// Behavior control
	SpamScore           float64
	BlacklistedUntil    time.Time
	RateLimitPenaltyEnd time.Time

	// Optional priority signal
	Staked         bool
	DelegatedStake uint64
}

func (er *ExtensionReputationParams) EntityType() string {
	return "extension"
}

func (pr *PeerReputationParams) EntityType() string {
	return "Peer"
}

func (wr *WalletReputationParams) EntityType() string {
	return "Address"
}

func (rm *ReputaionManager) GetEntity(entity string) (ReputationParams, error) {
	rep := rm.Reputation[entity]

	if rep == nil {
		return nil, fmt.Errorf("entity %s not found", entity)
	}

	return nil, fmt.Errorf("unknown entity type for %s", entity)
}

func (rm *ReputaionManager) ValidProofUpdate(entity string) error {
	e, err := rm.GetEntity(entity)
	if err != nil {
		return err
	}

	switch rep := e.(type) {
	case *ExtensionReputationParams:
		rep.ValidProofCount++
	case *PeerReputationParams:
		rep.ValidProofCount++
	default:
		return fmt.Errorf("unknown entity type for %s", entity)
	}

	rm.mu.Lock()
	defer rm.mu.Unlock()
	rm.Reputation[entity] = e

	return nil
}

func (rm *ReputaionManager) ValidProofRate(entity string) error {
	e, err := rm.GetEntity(entity)
	score := (ValidProofCount + prior_valid_proofs) / (ValidProofCount + InvalidProofCount + prior_valid_proofs + prior_invalid_proofs)
}

func (rm *ReputaionManager) OperationAcceptanceUpdate(entity string) error {

}

func (rm *ReputaionManager) OperationAcceptanceRate(entity string) error {}

func decode(rp []byte) *ReputationParams {
	buffer := bytes.NewBuffer(rp)

	dec := gob.NewDecoder(buffer)
	rep := new(ReputationParams)
	dec.Decode(rep)

	return rep
}
