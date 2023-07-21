package dhcpd

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateICMPv6RAPacket(t *testing.T) {
	wantData := []byte{
		0x86, 0x00, 0x00, 0x00, 0x40, 0x40, 0x07, 0x08,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x03, 0x04, 0x40, 0xc0, 0x00, 0x00, 0x0e, 0x10,
		0x00, 0x00, 0x0e, 0x10, 0x00, 0x00, 0x00, 0x00,
		0x12, 0x34, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x05, 0x01, 0x00, 0x00, 0x00, 0x00, 0x05, 0xdc,
		0x01, 0x01, 0x0a, 0x00, 0x27, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x19, 0x03, 0x00, 0x00, 0x00, 0x00,
		0x0e, 0x10, 0xfe, 0x80, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x08, 0x00, 0x27, 0xff, 0xfe, 0x00,
		0x00, 0x00,
	}

	gotData, err := createICMPv6RAPacket(icmpv6RA{
		managedAddressConfiguration: false,
		otherConfiguration:          true,
		mtu:                         1500,
		prefix:                      net.ParseIP("1234::"),
		prefixLen:                   64,
		recursiveDNSServer:          net.ParseIP("fe80::800:27ff:fe00:0"),
		sourceLinkLayerAddress:      []byte{0x0a, 0x00, 0x27, 0x00, 0x00, 0x00},
	})

	assert.NoError(t, err)
	assert.Equal(t, wantData, gotData)
}
