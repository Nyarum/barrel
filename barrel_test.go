package barrel

import (
	"log"
	"testing"
)

type Beril struct {
	Nates uint32
}

type Packet struct {
	Opcode uint64
	Test   string
	Beril  interface{}
	Bit    []byte
}

func (p *Packet) Default() {
	p.Opcode = 139
	p.Test = "Hello, world!"
	p.Beril = &Beril{6}
	p.Bit = []byte{0x01, 0x02, 0x03, 0x04, 0x05}
}

func (p Packet) Check(stats *Stats) bool {
	switch stats.NameField {
	case "Opcode":
		return true
	case "Test":
		return true
	case "Beril":
		return true
	case "Nates":
		return true
	case "Bit":
		stats.LenSlice = 5

		return true
	default:
		return false
	}
}

func TestBarrelPack(t *testing.T) {
	barrel := NewBarrel()
	packet := &Packet{}
	load := barrel.Load(packet, []byte{}, true)

	packet.Test = "O!"
	packet.Beril = &Beril{6}

	err := barrel.Pack(load)
	if err != nil {
		t.Error(err)
	}

	log.Printf("Buffer result: % x\n", barrel.Bytes())
}

func TestBarrelUnpack(t *testing.T) {
	barrel := NewBarrel()
	packet := &Packet{Beril: &Beril{}}
	load := barrel.Load(packet, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x8b, 0x00, 0x03, 0x4f, 0x21, 0x00, 0x00, 0x00, 0x00, 0x06, 0x01, 0x02, 0x03, 0x04, 0x05}, false)

	err := barrel.Unpack(load)
	if err != nil {
		t.Error(err)
	}

	log.Println("Struct result:", packet)
}

func BenchmarkBarrelPack(b *testing.B) {
	barrel := NewBarrel()
	packet := &Packet{}
	load := barrel.Load(packet, []byte{}, true)

	for i := 0; i < b.N; i++ {
		barrel.Pack(load)
	}
}

func BenchmarkBarrelUnpack(b *testing.B) {
	barrel := NewBarrel()
	packet := &Packet{}
	load := barrel.Load(packet, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x8b, 0x00, 0x03, 0x4f, 0x21, 0x00}, false)

	for i := 0; i < b.N; i++ {
		barrel.Unpack(load)
	}
}
