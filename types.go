package bion

import "fmt"

type Type uint8

const (
	TypeUndefined     Type = 0x00
	TypeNull          Type = 0x10
	TypeBooleanFalse  Type = 0x20
	TypeBooleanTrue   Type = 0x21
	TypeNumberInt8    Type = 0x30
	TypeNumberInt16   Type = 0x31
	TypeNumberInt32   Type = 0x32
	TypeNumberInt64   Type = 0x33
	TypeNumberUint8   Type = 0x34
	TypeNumberUint16  Type = 0x35
	TypeNumberUint32  Type = 0x36
	TypeNumberUint64  Type = 0x37
	TypeNumberFloat64 Type = 0x38
	TypeStringEmpty   Type = 0x40
	TypeStringShort   Type = 0x41
	TypeStringMedium  Type = 0x42
	TypeStringLong    Type = 0x43
	TypeArrayEmpty    Type = 0x50
	TypeArrayShort    Type = 0x51
	TypeArrayMedium   Type = 0x52
	TypeArrayLong     Type = 0x53
	TypeObjectEmpty   Type = 0x60
	TypeObjectShort   Type = 0x61
	TypeObjectMedium  Type = 0x62
	TypeObjectLong    Type = 0x63
)

func parseType(data []byte) (Type, error) {
	if len(data) == 0 {
		return TypeUndefined, fmt.Errorf("data is empty")
	}

	t := Type(data[0])

	switch t {
	case TypeUndefined,
		TypeNull,
		TypeBooleanFalse,
		TypeBooleanTrue,
		TypeNumberInt8,
		TypeNumberInt16,
		TypeNumberInt32,
		TypeNumberInt64,
		TypeNumberUint8,
		TypeNumberUint16,
		TypeNumberUint32,
		TypeNumberUint64,
		TypeNumberFloat64,
		TypeStringEmpty,
		TypeStringShort,
		TypeStringMedium,
		TypeStringLong,
		TypeArrayEmpty,
		TypeArrayShort,
		TypeArrayMedium,
		TypeArrayLong,
		TypeObjectEmpty,
		TypeObjectShort,
		TypeObjectMedium,
		TypeObjectLong:
		return t, nil
	default:
		return TypeUndefined, fmt.Errorf("invalid type: 0x%02x", t)
	}
}
