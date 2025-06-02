package types

//NodeType type
type NodeType uint8

//Constants for node types
const (
	BOOT NodeType = iota
	PEER
	CLIENT
)
