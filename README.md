### Modular Ethereum Extensions; Offloading specialized, semi-critical logics from the Execution Layer client.

A standard interface for adding custom logics/extensions to Ethereum nodes with no concensus layer change, no hard fork and no rollup required.

it's a protocol-adjacent layer that enables a group of ethereum clients to define a new set of API endponts along with custom a transaction type with a unified p2p sub-pool that interacts directly with ethereum main public mempool.
in other words this package is to enable devs to develope trusted, protocol-aware customizable components , that can be used to perform complex computations and validations that would be costly for the EVM and on the other hand not reasonable to do a hardfork for, an example of such computations could be ERC-4337 and Bundlers.

# How Does It Work?
it's a node that sits between Execution node/layer and application layer of Ethereum.
## Overview
![workflow](https://github.com/user-attachments/assets/afaf7f66-fdf6-4436-b64f-8b59cd1a2da1)

## security model
 ### Economic Security:

 ### Verifiability of Work:
   Validity Proofs and validity proof packets(VPP) :
   
   
## Specifications: 
  ### P2P stack: 
  we will utilize the ethereum execution client p2p stack, devp2p and Kademlia tables to manage our decentralized network of nodes.
  RLPx, DiscV5 and ENR are completely utilized.

 Ethereum middle(module) node record (EMNR) -> a new scheme and name record for modules to identify each other in the network. it's the same thing as ENR just with EMNR: prefix
  ### 

## Implementation: 
  all developers have to do is to implement the methods below, use the middleware package provided by the execution client team (being worked on by me for Geth), spin up an ethereun node and that's it.
  #### methods:
    RequestToTxn() -> maps the custom txn fields to a standard ethereum txn 
    API() -> methods avaliable to be called by users, it statrs with MOD_ for all the extentions 
    Endpoint() -> port and network interfaces / if not set -> do default 
    Contract() -> smartcontract associated with this Extention -> default nul 
    SecurityModel() -> stake/slash + Verifiability of Work
    ExtensionID()
    PoolID()

* all middle clients ,regardless of their job and logic, participate in a p2p network of middleWare nodes and use the same kademlia table for routing and node discovery, what separates them is ExtentionID and PoolID

* execution clients must implement a new handshake scheme and maintain a designated distributed hash table keeking track of reputable middle client/node (Extensions), sending periodical VVPs to filter out non-responsive and dishonest nodes, dishonest nodes will be blocked and cant interact with execution nodes no more
*  
