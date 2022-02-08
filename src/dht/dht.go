package dht

// Based off http://bittorrent.org/beps/bep_0005.html
type Node struct {
	NodeID string
	// Routing table for some other Nodes
}

type NodeStatus int64

const (
	Unknown NodeStatus = iota
	Good
)

type Bucket struct {
	Nodes []Node
}

/*
Upon inserting the first node into its routing table and when starting up thereafter, the node should attempt to find
the closest nodes in the DHT to itself. It does this by issuing find_node messages to closer and closer nodes until it
cannot find any closer. The routing table should be saved between invocations of the client software.
*/
