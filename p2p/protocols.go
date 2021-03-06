/*
 * @file
 * @copyright defined in aergo/LICENSE.txt
 */

package p2p

import (
	"github.com/libp2p/go-libp2p-protocol"
)


// constants of p2p protocol since v0.3
const (
	// this magic number is useful only in handshaking
	MAGICMain uint32 = 0x47416841
	MAGICTest uint32 = 0x2e415429

	P2PVersion030 uint32 = 0x00000300

	SigLength = 16
	IDLength = 16

	MaxPayloadLength = 1 << 23  // 8MB

	MaxBlockHeaderResponseCount = 10000
	MaxBlockResponseCount       = 2000
)

// context of multiaddr, as higher type of p2p message
const (
	aergoP2PSub   protocol.ID = "/aergop2p/0.3"
)


// SubProtocol identifies the lower type of p2p message
type SubProtocol uint32

// NOTE: change const of protocols_test.go
const (
	_ SubProtocol = 0x00 + iota
	StatusRequest
	PingRequest
	PingResponse
	GoAway
	AddressesRequest
	AddressesResponse
)
const (
	GetBlocksRequest SubProtocol = 0x010 + iota
	GetBlocksResponse
	GetBlockHeadersRequest
	GetBlockHeadersResponse
	GetMissingRequest  // Deprecated
	GetMissingResponse // Deprecated
	NewBlockNotice
	GetAncestorRequest
	GetAncestorResponse
	GetHashesRequest
	GetHashesResponse
	GetHashByNoRequest
	GetHashByNoResponse
)
const (
	GetTXsRequest SubProtocol = 0x020 + iota
	GetTXsResponse
	NewTxNotice
)

// subprotocols for block producers and their own trusted nodes
const (
	// BlockProducedNotice from block producer to trusted nodes and other bp nodes
	BlockProducedNotice SubProtocol = 0x030 + iota
)

//go:generate stringer -type=SubProtocol

func (i SubProtocol) Uint32() uint32 {
	return uint32(i)
}

