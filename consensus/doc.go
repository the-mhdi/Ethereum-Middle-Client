package consensus

/*
the consensus machanism for the Ethereum Middle Client :

1. middle nodes need to resigter themselves in the Consensus Contract
	to be able to participate in the consensus process.
	they could have somewhere between 0 to ~ ETH staked.
	more stake means more chance to be selected as a submitter.

2. the consensus contract keeps track of a subbmitter set

3. committee of submitters :
	in each epoch, a committee of submitters gets elected from the subbmitter set.
	they get to submit final postOps(txs) on main ethereum pool and.

	to disincentivize arbitrary middle nodes from submitting the final txns, only chosen submitters get rewarded and compensated.
**************************
another aspect of the consensus mechanism is guaranteeing that the extension logic is exactly the same across all middle nodes.
 this is done in two ways :
1. the extension logic is stored in a ZK circuit, which is committed to the consensus(or Entension regisrty) contract.
2. the extension logic is verified by the middle node verifier using validity proof packets.


*/
