package node

import "time"

type Extension struct {
	ExtensionID      string            // Unique identifier for the extension (e.g., "erc4337-bundler-v1")
	CircuitHash      common.Hash       // Commitment to the ZK circuit definition
	VerifyingKeyHash common.Hash       // Commitment to the verifying key
	ProofType        string            // e.g., "Groth16"
	LocalDirectory   string            // Local directory where the extension is stored
	Endpoint         string            // API endpoint for the extension
	VerifierMetadata map[string]string // Curve info, SNARK flavor, etc.
	RegisteredAt     time.Time         // When this extension was registered locally
	Active           bool              // Whether the extension is currently active
}
