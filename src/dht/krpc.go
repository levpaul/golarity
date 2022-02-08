package dht

type Message struct {
	TransactionID string // "t" - go with 2 bytes as recommended
	Type          string // "y"
	ClientVersion string // "v" - 2 chars (GT?) and then 2 chars for version num
}

type MessageType string

const (
	QueryMessage    MessageType = "q"
	ResponseMessage MessageType = "r"
	ErrorMessage    MessageType = "e"
)
