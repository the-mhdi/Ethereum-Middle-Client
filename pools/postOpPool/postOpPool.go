package postOpPool

// PostOpPool is a pool for storing post-operations that have been processed by extensions.

type PostOperation struct {
	ExtensionID        string
	OperationHash      string     // hash(Operation)
	Operation          *Operation // The original operation that was processed
	Proof              *Proof
	Data               []byte
	ProcessedBlockHash []byte //block hash at the time of processing
}

type Pool struct {
}
