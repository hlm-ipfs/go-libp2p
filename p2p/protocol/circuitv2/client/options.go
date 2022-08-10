package client

import (
	"fmt"
	"strings"

	"github.com/libp2p/go-libp2p-core/peer"
)

type RelayType string

const (
	RelayTypeV1 RelayType = "relay-v1"
	RelayTypeV2 RelayType = "relay-v2"
)

type Parser func(info *peer.AddrInfo) RelayType
type Builder func(id peer.ID, info *peer.AddrInfo) string

type Options struct {
	parser  Parser
	builder Builder
}

var (
	options = &Options{
		parser:  defaultParser,
		builder: defaultBuilder,
	}
)

func defaultParser(info *peer.AddrInfo) RelayType {
	if strings.HasPrefix(info.ID.String(), "Qm") {
		return RelayTypeV1
	}
	return RelayTypeV2
}

func defaultBuilder(id peer.ID, info *peer.AddrInfo) string {
	return fmt.Sprintf("%v@%v", id.String(), defaultParser(info))
}

func WithParser(parser Parser) {
	if parser != nil {
		options.parser = parser
	}
}

func WithBuilder(builder Builder) {
	if builder != nil {
		options.builder = builder
	}
}
