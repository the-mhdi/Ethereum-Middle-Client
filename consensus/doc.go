package consensus

/*
the consensus machanism for the Ethereum Middle Client :

1. middle nodes need to resigter themselves in the Consensus Contract
	to be able to participate in the consensus process.
	they could have somewhere between 0 to ~ ETH staked which can bring them reputaion.

2. the consensus contract keeps track of subbmitters set

3. committee of submitters :
	in each epoch, a committee of submitters is selected from the registered middle nodes.
	they get to submit final postops(txs) on main ethereum pool and to the execution client

	to disincentivize middle nodes from submitting txns, only chosen submitters get rewarded and compensated.




*/
