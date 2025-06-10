### Modular Ethereum Extensions; Offloading specialized, semi-critical logics from the Execution Layer client.

A standard interface for adding custom logics/extensions to Ethereum nodes with no concensus layer change and no hard fork required.

it's a protocol-adjacent layer that enables a group of ethereum clients to define a new set of API endponts along with custom a transaction type with a unified p2p sub-pool that interacts directly with ethereum main public mempool.
in other words this package is to enable devs to develope trusted, protocol-aware customizable components , that can be used to perform complex computations and validations that would be costly for the EVM and on the other hand not reasonable to do a hardfork for, an example of such computations could be ERC-4337 and Bundlers.

# How Does It Work?
it's a node that sits between Execution node/layer and application layer of Ethereum.
## Overview
![workflow](https://github.com/user-attachments/assets/afaf7f66-fdf6-4436-b64f-8b59cd1a2da1)

## Specifications: 
  ### P2P stack: 
  we will utilize the ethereum execution client p2p stack, devp2p and Kademlia tables to manage our decentralized network of nodes.
  RLPx, DiscV5 and ENR are completely utilized.

 Ethereum middle(module) node record (EMNR) -> a new scheme and name record for modules to identify each other in the network. it's the same thing as ENR just with EMNR: prefix
  ###
  
