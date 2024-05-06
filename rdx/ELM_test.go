package rdx

import (
	"testing"

	"github.com/drpcorg/chotki/protocol"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	var correct = []byte{
		0x6D, 0x15,
		0x36, 0x03, 0x00, 0xaf, 0x00, 0x0b, 0x0b,
		0x73, 0x04, 0x30, 0x4b, 0x65, 0x79,
		0x73, 0x06, 0x30, 0x56, 0x61, 0x6c, 0x75, 0x65,
	}
	id := IDFromString("b0b-af0-3")
	env2 := protocol.Record('M',
		protocol.TinyRecord('I', id.ZipBytes()),
		protocol.Record('S', protocol.TinyRecord('T', ZipIntUint64Pair(0, 0)), []byte("Key")),
		protocol.Record('S', protocol.TinyRecord('T', ZipIntUint64Pair(0, 0)), []byte("Value")),
	)
	assert.Equal(t, correct, env2)
}

func TestEmerge(t *testing.T) {
	tlv1 := Eparse("{1, 2, \"four\"}")
	assert.Equal(t, "{1,2,\"four\"}", Estring(tlv1))
	tlv2 := Eparse("{3, \"four\", 5}")
	tlv12 := Emerge(protocol.Records{tlv1, tlv2})
	str12 := Estring(tlv12)
	assert.Equal(t, "{1,2,3,\"four\",5}", str12)

	del := Itlve(-1, 0, 1)
	tlv12d := Emerge(protocol.Records{tlv1, tlv2, del})
	str12d := Estring(tlv12d)
	assert.Equal(t, "{2,3,\"four\",5}", str12d)
}

func TestMmerge(t *testing.T) {
	tlv1 := Mparse("{1: 2,  5:6, 3: 4}")
	assert.Equal(t, "{1:2,3:4,5:6}", Mstring(tlv1))
	tlv2 := Mparse("{ 7:8, 3:4, 5:6}")
	tlv12 := Mmerge(protocol.Records{tlv1, tlv2})
	str12 := Mstring(tlv12)
	assert.Equal(t, "{1:2,3:4,5:6,7:8}", str12)

	del := protocol.Concat(
		Itlve(-1, 0, 5),
		Itlve(-1, 0, 6), // todo T
	)
	tlv12d := Mmerge(protocol.Records{tlv1, tlv2, del})
	str12d := Mstring(tlv12d)
	assert.Equal(t, "{1:2,3:4,7:8}", str12d)
}

func TestLmerge(t *testing.T) {
	tlv1 := Lparse("[1, 2, 3,  5]")
	assert.Equal(t, "[1,2,3,5]", Lstring(tlv1))
	patch1 := protocol.Record('B',
		protocol.TinyRecord('T', ZipIntUint64Pair(3, 0)),
		Itlve(5, 0, 4),
	)
	tlv2 := Lmerge(protocol.Records{tlv1, patch1})
	assert.Equal(t, "[1,2,3,4,5]", Lstring(tlv2))
}
