package barrel

import (
	"log"
	"testing"
)

type Packet struct {
	Opcode uint64
	Test   string
}

func (p *Packet) Default() {
	p.Opcode = 139
	p.Test = "Hello, world!"
}

func (p Packet) Check(name string) bool {
	switch name {
	case "Opcode":
		return true
	case "Test":
		return true
	case "Be":
		return true
	default:
		return false
	}
}

func TestBarrelPack(t *testing.T) {
	barrel := NewBarrel()
	packet := &Packet{}
	load := barrel.Load(packet, []byte{})

	packet.Test = "O!"

	err := barrel.Pack(load)
	if err != nil {
		t.Error(err)
	}

	log.Printf("Buffer result: % x\n", barrel.Bytes())
}

func TestBarrelUnpack(t *testing.T) {
	barrel := NewBarrel()
	packet := &Packet{}
	load := barrel.Load(packet, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x8b, 0x00, 0x03, 0x4f, 0x21, 0x00})

	err := barrel.Unpack(load)
	if err != nil {
		t.Error(err)
	}

	log.Println("Struct result:", packet)
}

func BenchmarkBarrelPack(b *testing.B) {
	barrel := NewBarrel()
	packet := &Packet{}
	load := barrel.Load(packet, []byte{})

	for i := 0; i < b.N; i++ {
		barrel.Pack(load)
	}
}

func BenchmarkBarrelUnpack(b *testing.B) {
	barrel := NewBarrel()
	packet := &Packet{}
	load := barrel.Load(packet, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x8b, 0x00, 0x03, 0x4f, 0x21, 0x00})

	for i := 0; i < b.N; i++ {
		barrel.Unpack(load)
	}
}
