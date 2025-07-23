package middleNode

type MiddleNode struct {
	nodeID    enode.ID // Unique identifier for the middle node, used for Op/postOp/proof propagation topology
	networkID uint64

	config *MiddleNodeConfig

	opPool     *opPool.Pool
	postOpPool *postOpPool.Pool
	proofpool  *proofpool.Pool

	APIs *MiddleNodeAPIs
}
