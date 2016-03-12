package barrel

import (
	"log"
	"testing"
)

type Packet struct {
	ID    int32
	Level uint32
	HP    uint32
}

func (p *Packet) Default() {
	p.ID = 2
	p.Level = 3
	p.HP = 4
}

func (p Packet) Check(stats *Stats) bool {
	switch stats.NameField {
	case "ID":
		stats.Endian = BigEndian
	case "Level":
	case "HP":
	}

	return true
}

func TestBarrelPack(t *testing.T) {
	barrel := NewBarrel()
	packet := &Packet{}
	load := barrel.Load(packet, []byte{}, true)

	err := barrel.Pack(load)
	if err != nil {
		t.Error(err)
	}

	log.Printf("Buffer result: % x\n", barrel.Bytes())
}

func TestBarrelUnpack(t *testing.T) {
	barrel := NewBarrel()
	packet := &Packet{}
	load := barrel.Load(packet, []byte{0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x04}, false)

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
	load := barrel.Load(packet, []byte{0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x04}, false)

	for i := 0; i < b.N; i++ {
		barrel.Unpack(load)
	}
}
