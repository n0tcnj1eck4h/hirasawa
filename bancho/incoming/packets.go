package incoming

// Packet 3
type RequestStatusUpdate struct {
	PacketHeader
}

// Packet 4
type Ping struct {
	PacketHeader
}
