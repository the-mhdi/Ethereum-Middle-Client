### Ethereum Middle client Concept, Modular Ethereum With Extensions; Offloading specialized, semi-critical logics from the Execution Layer client.

A standard interface for adding custom logics/extensions to Ethereum nodes with no concensus layer change, no hard fork, no L2!


it's a protocol-adjacent layer introducing a new type of ethereum clients with a new set of API endponts along side a custom transaction type with a unified p2p sub-pool that interacts directly with ethereum main public mempool each node on this network of new clients can have itsown set of Extensions.
I define Extensions as protocol-aware customizable components , that can be used to perform complex computations and validations that would be costly for the EVM and on the other hand not reasonable to do a hardfork for, an example of such computations could be ERC-4337 Bundlers.

![Untitled Diagram drawio (1)](https://github.com/user-attachments/assets/a76e0ebf-96c4-487f-95b5-303b1ce4fe00) Fig. 1


# How Does It Work?
in a ZKP manner, Middle-Nodes act as verifiers and Extensios are provers, middle nodes maintain a routing tables of their neighbours and their respective Extensions, middle nodes regardless of the extension they have, can maintain their mempool to have all types of  operaions(a special txn). 
### Core Concepts : 
 #### * Middle-Client: new type of Ethereum node that sits between standard Execution clients and end-users
  * Maintains two mempools:
     1. Operation Mempool: unprocessed Operations
     2. PostOp Mempool: processed and validated Operations
  * Acts as a verifier of Extension outputs
  * Manages routing tables of peer Middle-Clients and their supported Extensions
  * Can enforce stake, fee, and reputation policies to prevent spam and maintain trust
  * Optionally participates in staking and slashing mechanisms to secure Operations economically
    
 #### * Extensions: a protocol-aware, customizable module that performs specialized computation/validation outside the EVM
  * Each Extension has a unique ExtensionID
  * Operates as a prover that processes Operations and generates validity proofs or post-processed outputs
  * Can be independently developed and deployed by any party
  * May require Middle-Clients to stake collateral, enabling slashing if the Extension produces invalid results
  * Makes Ethereum more modular, allowing new transaction types and processing logic without modifying consensus
  * Examples of Extension functionality: ERC-4337 bundler logic , Specialized compliance checks , Advanced signature schemes , ZKP circuit execution



 
we introduce two new (semi)transaction types. in regard to ERC-4337 we're calling them Operaions. 
 ### Operation Struct:
    type Operation struct {
      ExtensionID	string
      To       Address
      Gas      uint
      Data     []byte
      Sig	     []byte
      BlockHash  []byte // block hash upon op submission to extension
    }
 ### PostOperation Struct:
    type PostOp struct {
     OperationHash  string        // hash(Operation)
     ExtensionID   string
     Proof              Proof 
     Data               []byte
 
     ProcessedBlockHash []byte //block hash at the time of processing
    }

Middle RPC Nodes receive operations from users, they perform simple verifcation and reputaions management then the Operation would be submitted to the public mempool, middle nodes that have respective Extension to proccess that operation will pick it up, the operation would be processed and the post-prossessed operation will be sumbitted to another mempool called post-mempool.


so as shown on the diagram below the middle nodes manage and maintain to public p2p mempools: Operation p2p Mempool and PostOp p2p mempool

![Untitled Diagram drawio (2)](https://github.com/user-attachments/assets/d82a8de5-e489-4b50-adbe-1893887afba0)




 # CHALLENGES

## Canonicality and Re-org Safety :
Before creating an Operation, the user's wallet queries an Execution node to get the hash of a recent, finalized block, it then gets included into the Operation.
After the Extension finishes its computation, it fetches the current block hash from the mainnet and includes it in the PostOp.
When a Middle Node receives and verifies the PostOp, it MUST performs a critical freshness check:

 1. It compares the ProcessedBlockHash against its own view of the blockchain.

 2. It enforces a rule: (RULE No. 1) the PostOp is only valid if its ProcessedBlockHash is very recent (for example, within the last 5-10 blocks).


If a re-org happens and the referenced block is orphaned (no longer part of the canonical chain), the Operation becomes instantly invalid. Middle Nodes can simply discard it because the previos state no longer exists. This prevents Extensions from processing operations based on a stale or reverted chain state.

the ProcessedBlockHash and RULE 1 also provide a defense against replay attacks, a malicious actor cannot take a valid PostOp from a week ago and submit it today, because the old ProcessedBlockHash would cause it to be immediately rejected as stale. 

* middle nodes CAN define a configurable PROCESS_WINDOW variable, it's an interval of slots in which an Operation is deemded valid.
  

## Incentive Mechanisms (TBD) :
middle nodes run Extenstions, since the Extensions work as provers in this architecture, they need to be paid fairly.

users also want their Operations processed for a predictable fee.
 
the gas fee can be paid by the user directly or be sponsered by another entity, all we need is an ERC4337 incentive flow and users commitment to a fee, this brings up a need for a singleton entrypoint-like contract.

## Proof Formats and Trust Minimization : 
 the proof system must:

* Allow Extensions (Provers) to prove correct processing of an Operation.

* Enable Middle Nodes (Verifiers) to validate proofs deterministically.

* Avoid reliance on central trust.

* Be general enough to handle: ZK proofs (snarks/starks) and other cryptographic attestations (Merkle proofs, signatures).

* Allow efficient verification without heavy resource demands.

* Allow Middle nodes to send validity proof packets to other nodes and receives proof responses. (this is to maintain a vaild reputation system and prevent middle nodes from altering an extension functionality)

   #### standard Proof object: 
      type Proof struct {
      ProofType     string            // e.g., "Groth16"
      ExtensionID   string            // erc4337
      Inputs        map[string][]byte // Public inputs
      Output        []byte            // Post-processed result data (e.g., calldata)
      ProofData     []byte            // The proof itself (binary blob)
      Metadata      map[string]string // Optional metadata (versioning, etc.) MUST include "CircuitHash"
      }
    


* Verifiability of Extensions Work: (nodes cross-verifying one another’s proofs)

   a peer challenge-response protocol (distributed attestation) 

  ### validity proof packets :

 We’ll define two main packet types:
 
   #### ProofVerificationRequest
  
      type ProofVerificationRequest struct {
        RequestID     string        // Unique ID for deduplication
        SenderNodeID  string        // Node issuing the request
        OperationID   string        // Hash of the original Operation
        PostOp        PostOp        // Full PostOp struct incl. Proof
        Timestamp     int64         // Unix timestamp
        Metadata      map[string]string // Optional context
      }

   #### ProofVerificationResponse
      type ProofVerificationResponse struct {
      RequestID     string        // Echoed from the request
      ResponderNodeID string      // Who verified
      Verdict       VerificationVerdict // Enum: VALID / INVALID / ERROR
      Signature     []byte        // Signature over {RequestID, Verdict}
      Diagnostics   map[string]string // Optional error details
      Timestamp     int64
      }

   #### ProofVerificationReceipt (no sure if this one is needed)


### Proof System Flow:

1. Middle Node (A) produces a PostOp with attached Proof.

2. Before submitting on-chain, (A) broadcasts a ProofVerificationRequest to neighbors.

3. Neighboring nodes (Middle Nodes B, C, D) with the relevant Extensions re-verify the Proof.

4. Each neighbor returns a ProofVerificationResponse

* This creates a decentralized consensus over proof validity.
* This is the basis of middle nodes reputation system.

Nodes SHOULD rate limit verification requests per peer.

  ### Validity Attack protection : 
   how can we guarantee that: 1. The Extension logic is exactly the same logic other nodes expect for that ExtensionID? 2. The output and proof are produced by an approved implementation, not a malicious or buggy variant?

   we'll be using ZK Circuit Commitment committing each ExtensionID to:
* One specific ZK circuit definition 
* One verifying key 

This ensures all proofs are generated only with that circuit. Every Middle Node can deterministically verify them and that there is no ambiguity about what code was executed.
 ##### these would be a registry contract and a registry mapping like below : 
    ExtensionID → {CircuitHash, VerifyingKeyHash, ProofType, VerifierMetadata}

Middle nodes MUST only accept Operations referencing a known ExtensionID
Middle nodes MUST Only accept proofs that declare the circuitHash and are verified against the verifying key hash
Middle Nodes MUST check : proof.Metadata["CircuitHash"] == registry[proof.ExtensionID].circuitHash (before verifying)

#### Extension Registry Flow : 
two parties involved : Extension dev and Registry Smart Contract

1. Developer Prepares the Extension, the extension should be compilable to an arithmetic circuit (R1CS/QAP)
2. serializeed the constraint system to a canonical blob.
      * what do i mean by canonical blob ?
        A byte-serialized representation of a build artifact that every Middle Node can hash to the same value.
        we MIGHT define a VerifierBinary struct in our design :

            type VerifierBinary struct {
            BinaryFormat string // e.g., "wasm", "solidity", "elf"
            BinaryData []byte
            BinaryHash []byte
             }



3. Trusted Setup (if needed): run trusted setup to generate proving/verifying keys.

4. Hashing

       circuitHash = keccak256(r1cs_blob)
       verifyingKeyHash = keccak256(serialized_verifying_key)
   
6. Verifier Metadata curve info, SNARK flavor ..
7. ExtensionID: it should be a unique string 

8. calling the registerExtension function of Registry Contract
   
       function registerExtension(
       string calldata extensionId,
       bytes32 circuitHash,
       bytes32 verifyingKeyHash,
       string calldata proofType,
       bytes calldata verifierMetadata
       ) external;
   
Registry Contract MUST store a mapping of ExtensionID to ExtensionMetadata 
               
      mapping(string => ExtensionMetadata)
      
Extension Metadata struct COULD be : 

       ExtensionMetadata {
       bytes32 circuitHash;
       bytes32 verifyingKeyHash;
       string proofType;
       bytes verifierMetadata;
     }

******* After this point *********

Any node or user can query the registry on-chain to get the canonical circuit commitment and verifying key commitment for any ExtensionID.

### Middle Node Adding an Extension (Middle Node Onboarding)
  Imagine you are operating a Middle Node and want to add a new Extension to your already working set of Extensions (Fig. 1)
 1. query the on-chain registry -> GET ExtensionMetadata(extensionID)
 2. obtain verifier Code (could be a smart contract address) & Verifying Key
 3. Store Extension Locally
 
        json {
        "ExtensionID": "erc4337-bundler-v1",
         "CircuitHash": "0xabc123...",
         "VerifyingKeyHash": "0xdef456...",
         "ProofType": "Groth16",
        "VerifierMetadata": {...},
        "VerifierPath": "/extensions/erc4337-bundler-v1/verifier", or address : "0x...."
        "VerifyingKeyPath": "/extensions/erc4337-bundler-v1/vk"
        }

#### Runtime Verification Flow : what happens when your middle node receives a PostOP 
(TBD)
  
(work in process) : 1. Define gossip strategies for distributing verification requests 2. build reputation scoring algorithms.

## Reputation System :
 reputation needs to be managed in two main scopes: 1. Extension Reputation 2. Middle-Node Reputation
 
 | Metric                    | Applies To             | Description                                                        |
| ------------------------- | ---------------------- | ------------------------------------------------------------------ |
| ValidProofCount           | Middle Node, Extension | How many proofs this node/extension generated or verified as valid |
| InvalidProofCount         | Middle Node, Extension | How many invalid proofs this node/extension generated or verified  |
| OperationAcceptanceRate   | Middle Node            | % of Operations accepted vs rejected                               |
| OperationExecutionLatency | Middle Node, Extension | Time taken to process an Operation into PostOp                     |
| ProofVerificationLatency  | Middle Node            | Time to respond to proof verification requests                     |
| AvailabilityScore         | Middle Node            | % uptime responding to requests                                    |
| DisputeOutcomeScore       | Middle Node, Extension | % of times a node was challenged and proved correct vs incorrect   |
| StakeBalance              | Middle Node, Extension  |                                                                   |
| Peer Endorsements         | Middle Node, Extension | Positive attestations signed by reputable peers                    |
| Negative Slashes          | Middle Node, Extension | Slashing events due to malicious or faulty behavior                |
| Recent Activity Timestamp | Middle Node, Extension | Last time this node/extension was seen active                      |


### Data Models : 
#### Extension Reputation
    type ExtensionReputation struct {
    ExtensionID            string
    ValidProofCount        uint64
    InvalidProofCount      uint64
    DisputeOutcomeScore    float64
    AvgExecutionLatencyMs  float64
    LastActiveTimestamp    int64
    StakeBalance           uint64   // New: How much stake this Extension has
    UnstakeDelaySeconds    uint64   // New: Cooldown period
    }
#### Middle-Node Reputation
    type MiddleNodeReputation struct {
    NodeID                      string
    ValidProofsProduced         uint64
    InvalidProofsProduced       uint64
    OperationAcceptanceRate     float64
    AvgProofVerificationLatency float64
    AvailabilityScore           float64
    StakeBalance                uint64
    ExtensionStakes             []ExtensionReputation
    Endorsements                []string
    NegativeSlashes             uint64
    LastActiveTimestamp         int64
    }
### ReputationScore formula + Design on-chain dispute resolution -> TBD 


## Extension Registry Smart Contract :


## Specifications: 
  ### P2P stack: 
  we will utilize the ethereum execution client p2p stack, devp2p and Kademlia tables to manage our decentralized network of middle nodes.
  RLPx, DiscV5 and ENR are completely utilized.
   node MUST broadcast a Capability Advertisement Packet upon peer connection so they can advertise SupportedExtensions, SupportedProofTypes , MaxProofSize , FeeSchedule


NOTICE : 
reducing txn congestion on base layer in not the direct intent of this design but it can be utilized to also act as an L2 rollup 

Data Availability Proofs are optional. 



