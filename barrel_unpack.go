package barrel

import (
	"fmt"
	"reflect"
	"strconv"
)

func (b *Barrel) Unpack(value reflect.Value) error {
	switch value.Kind() {
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			valueType := value.Type().Field(i)

			if !b.Object.Check(valueType.Name) {
				continue
			}

			if valueType.Tag.Get("endian") == "0" {
				b.processor.SetEndian(BigEndian)
			} else if valueType.Tag.Get("endian") == "1" {
				b.processor.SetEndian(LittleEndian)
			}

			b.numField = i

			err := b.Pack(value.Field(i))
			if err != nil {
				return err
			}
		}
	case reflect.Uint8:
		var valueFrom uint8
		b.processor.ReadUint8(&valueFrom)

		value.SetUint(uint64(valueFrom))
	case reflect.Int8:
		var valueFrom int8
		b.processor.ReadInt8(&valueFrom)

		value.SetInt(int64(valueFrom))
	case reflect.Uint16:

		var valueFrom uint16
		b.processor.ReadUint16(&valueFrom)

		value.SetUint(uint64(valueFrom))
	case reflect.Int16:
		var valueFrom int16
		b.processor.ReadInt16(&valueFrom)

		value.SetInt(int64(valueFrom))
	case reflect.Uint, reflect.Uint32:
		var valueFrom uint32
		b.processor.ReadUint32(&valueFrom)

		value.SetUint(uint64(valueFrom))
	case reflect.Int, reflect.Int32:
		var valueFrom int32
		b.processor.ReadInt32(&valueFrom)

		value.SetInt(int64(valueFrom))
	case reflect.Uint64:
		var valueFrom uint64
		b.processor.ReadUint64(&valueFrom)

		value.SetUint(valueFrom)
	case reflect.Int64:
		var valueFrom int64
		b.processor.ReadInt64(&valueFrom)

		value.SetInt(valueFrom)
	case reflect.String:
		var valueFrom string
		b.processor.ReadString(&valueFrom)

		value.SetString(valueFrom)
	case reflect.Slice:
		getLen, err := strconv.Atoi(value.Type().Field(b.numField).Tag.Get("len"))
		if err != nil {
			return err
		}

		var valueFrom []byte
		b.processor.ReadBytes(&valueFrom, getLen)

		value.SetBytes(valueFrom)
	case reflect.Bool:
		var valueFrom bool
		b.processor.ReadBool(&valueFrom)

		value.SetBool(valueFrom)
	case reflect.Interface:
		err := b.Pack(value.Elem())
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("Type of object is incorrect.. It is - %v", value.Kind())
	}

	return nil
}
