 
	// structure :
	// package (node) manages the middle client node, which interacts with the Ethereum network and other middle clients.
	// It maintains the p2p network, API endpoints, registers extensions, reputation management, and monitors node state.

	// THE simlpe FLOW :
	// op -> middle node -> opPool -> respective extension -> postOpPool + proof -> middle node verifier -> eth txn -> submit to L1
	// <----------------------------------reputation management through all these steps------------------------------------------->