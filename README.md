### Ethereum Middle client Concept, Modular Ethereum With Extensions; Offloading specialized, semi-critical logics from the Execution Layer client.

A standard interface for adding custom logics/extensions to Ethereum nodes with no concensus layer change, no hard fork!


it's a protocol-adjacent layer introducing a new type of ethereum clients with a new set of API endponts along side a custom transaction type with a unified p2p sub-pool that interacts directly with ethereum main public mempool each node on this network of new clients can have itsown set of Extensions.
I define Extensions as protocol-aware customizable components , that can be used to perform complex computations and validations that would be costly for the EVM and on the other hand not reasonable to do a hardfork for, an example of such computations could be ERC-4337 Bundlers.

![Untitled Diagram drawio (1)](https://github.com/user-attachments/assets/a76e0ebf-96c4-487f-95b5-303b1ce4fe00)


# How Does It Work?
in a ZKP manner, Middle-Nodes act as verifiers and Extensios are provers, middle nodes maintain a routing tables of their neighbours and their respective Extensions, middle nodes regardless of the extension they have, can maintain their mempool to have all types of  operaions(a special txn). 

we introduce a new type of transaction. in regard to ERC-4337 we call them Operaions. 
 ### Operation Struct:
    type Operation struct {
     ExtensionID	string
     To       Address
     Data []byte
     Sig	 []byte
    }
Middle RPC Nodes receive operations from users, they perform simple verifcation and reputaions management then the Operation would be submitted to the public mempool, middle nodes that have respective Extension to proccess that operation will pick it up, the operation would be processed and the post-prossessed operation will be sumbitted to another mempool called post-mempool.


so as shown on the diagram below the middle nodes manage and maintain to public p2p mempools: Operation p2p Mempool and PostOp p2p mempool

![Untitled Diagram drawio (2)](https://github.com/user-attachments/assets/d82a8de5-e489-4b50-adbe-1893887afba0)






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

