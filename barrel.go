package barrel

import "reflect"

type (
	Unit interface {
		Default()
		Check(string) bool
	}

	Barrel struct {
		Object    Unit
		numField  int
		processor *Processor
	}
)

func NewBarrel() *Barrel {
	return &Barrel{}
}

func (b *Barrel) Load(object Unit, buffer []byte, def bool) reflect.Value {
	if def {
		object.Default()
	}

	b.Object = object
	b.processor = NewProcessor(buffer)

	return reflect.Indirect(reflect.ValueOf(object))
}

func (b *Barrel) Bytes() []byte {
	return b.processor.Bytes()
}
