package report

import "github.com/NicolasDeveloper/tracker-udp-server/interfaces/tracker-tcp/command"

type parseStub struct {
	command.Parser
	head     string
	deviceID string
	version  string
	end      string
	ipHex    string
	port     string
}

func (stub parseStub) GetHead(command []byte) (string, error) { return stub.head, nil }

func (stub parseStub) GetHexDeviceID(command []byte) (string, error) { return stub.deviceID, nil }

func (stub parseStub) GetVersion(command []byte) (string, error) { return stub.version, nil }

func (stub parseStub) GetEnd() (string, error) { return stub.end, nil }

func (stub parseStub) GetIPHex() string { return stub.ipHex }

func (stub parseStub) GetPort() string { return stub.port }

//NewParseStub contructor
func NewParseStub(head string, deviceID string, version string, ipHex string, port string, end string) command.Parser {
	return &parseStub{
		head:     head,
		deviceID: deviceID,
		version:  version,
		ipHex:    ipHex,
		port:     port,
		end:      end,
	}
}
