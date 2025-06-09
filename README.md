# Modular Ethereum Extensions; Offloading specialized, semi-critical logics from the Execution Layer client.

A standard interface for adding custom logics/extensions to ethereum clinets with no concensus layer change and no hard fork required, 
it's a protocol-adjacent layer that enables a group of ethereum clients to define a new set of API endponts along with custom a transaction type with a unified p2p sub-pool that interacts directly with ethereum main public mempool.
in other words this package is to enable devs to develope trusted, protocol-aware customizable components , that can be used to perform complex computations and validations that would be costly for the EVM and on the other hand not reasonable to do a hardfork for, an example of such computations could be ERC-4337 and Bundlers.
