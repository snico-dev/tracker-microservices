package models

//Message main package
type Message struct {
	Head       string
	Length     int
	Version    string
	DeviceID   string
	ProtocolID string
	CRC        string
	Tail       string
}
