### Ethereum Middle client Concept, Modular Ethereum With Extensions; Offloading specialized, semi-critical logics from the Execution Layer client.

A standard interface for adding custom logics/extensions to Ethereum nodes with no concensus layer change, no hard fork!


it's a protocol-adjacent layer introducing a new type of ethereum clients with a new set of API endponts along side a custom transaction type with a unified p2p sub-pool that interacts directly with ethereum main public mempool each node on this network of new clients can have itsown set of Extensions.
I define Extensions as protocol-aware customizable components , that can be used to perform complex computations and validations that would be costly for the EVM and on the other hand not reasonable to do a hardfork for, an example of such computations could be ERC-4337 Bundlers.

![Untitled Diagram drawio (1)](https://github.com/user-attachments/assets/a76e0ebf-96c4-487f-95b5-303b1ce4fe00)


# How Does It Work?
in a ZKP manner, Middle-Nodes act as verifiers and Extensios are provers, middle nodes maintain a routing tables of their neighbours and their respective Extensions, middle nodes regardless of the extension they have, can maintain their mempool to have all types of  operaions(a special txn). 

we introduce a two new transaction types. in regard to ERC-4337 we call them Operaions. 
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




### challenges : 

### Canonicality and Re-org Safety :
Before creating an Operation, the user's wallet queries an Execution node to get the hash of a recent, finalized block, it then gets included into the Operation.
After the Extension finishes its computation, it fetches the current block hash from the mainnet and includes it in the PostOp.
When a Middle Node receives and verifies the PostOp, it MUST performs a critical freshness check:

 1. It compares the ProcessedBlockHash against its own view of the blockchain.

 2. It enforces a rule: (RULE No. 1) the PostOp is only valid if its ProcessedBlockHash is very recent (for example, within the last 5-10 blocks).


If a re-org happens and the referenced block is orphaned (no longer part of the canonical chain), the Operation becomes instantly invalid. Middle Nodes can simply discard it because the previos state no longer exists. This prevents Extensions from processing operations based on a stale or reverted chain state.

the ProcessedBlockHash and RULE 1 also provide a defense against replay attacks, a malicious actor cannot take a valid PostOp from a week ago and submit it today, because the old ProcessedBlockHash would cause it to be immediately rejected as stale. 

* middle nodes CAN define a PROCESS_WINDOW variable, it's an interval of slots in which an Operation is deemded valid.
  

### Incentive Mechanisms (TBD) :
middle nodes run Extenstions, since the Extensions work as provers in this architecture, they need to be paid fairly.

users also want their Operations processed for a predictable fee.
 
the gas fee can be paid by the user directly or be sponsered by another entity, all we need is an ERC4337 incentive flow and users commitment to a fee, this brings up a need for a singleton entrypoint-like contract.

### Proof Formats and Trust Minimization : 

 
## Verifiability of Extensions Work:
   Validity Proofs and validity proof packets(VPP) :
   
## security model

## Specifications: 
  ### P2P stack: 
  we will utilize the ethereum execution client p2p stack, devp2p and Kademlia tables to manage our decentralized network of middle nodes.
  RLPx, DiscV5 and ENR are completely utilized.

 Ethereum middle(module) node record (EMNR) -> a new scheme and name record for modules to identify each other in the network. it's the same thing as ENR just with EMNR: prefix
  ### 

## Implementation: 

## Middle node Specs : 
  

## Extension Specs :
  Extensions could be written in any language they have to follow these roles:
  1. Generate zk-proofs on demand
  2. utilizing json-rpc API to interact with middle nodes and JWT authentication mechanism
  3. 




NOTICE : 
reducing congestion on base layer in not the direct intent of this design but it can be utilized to act as a L2 rollup 
Data Availability Proofs are optional. 


there'd could be two types of middleware clients 1: those that only act as rpc nodes and route the request to its respective Extension node(middleware node) 2. Extension client: these are nodes that has implemented the Extension interface 
each extension is like a microservice, each middle node can run as many number of extensions as they want. 

how can we ensure that all extensions of the same ExtensionID run the same logic?
 The Core Principle: Code Commitment + Proof of Execution
 social smart contract that 

 there's a registerExtension function 
step 1 : compile code into a arithmetic circuit(provable format)
step 2 : ZK Coprocessor

validity attack prevention : 

