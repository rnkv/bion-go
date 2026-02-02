package bion

import (
	"fmt"
	"math"
)

func Unmarshal(data []byte, value *any) error {
	var err error
	data, err = unmarshal(data, value)
	if err != nil {
		return err
	}

	if len(data) != 0 {
		return fmt.Errorf("data is not empty after unmarshalling")
	}

	return nil
}

func unmarshal(data []byte, value *any) ([]byte, error) {
	dataType, err := parseType(data)
	if err != nil {
		return data, fmt.Errorf("failed to parse type: %w", err)
	}

	data = data[1:]

	switch dataType {
	case TypeUndefined:
		return unmarshalUndefined(data, value)
	case TypeNull:
		return unmarshalNull(data, value)
	case TypeBooleanFalse:
		return unmarshalBooleanFalse(data, value)
	case TypeBooleanTrue:
		return unmarshalBooleanTrue(data, value)
	case TypeNumberInt8:
		return unmarshalInt8(data, value)
	case TypeNumberInt16:
		return unmarshalInt16(data, value)
	case TypeNumberInt32:
		return unmarshalInt32(data, value)
	case TypeNumberInt64:
		return unmarshalInt64(data, value)
	case TypeNumberUint8:
		return unmarshalUint8(data, value)
	case TypeNumberUint16:
		return unmarshalUint16(data, value)
	case TypeNumberUint32:
		return unmarshalUint32(data, value)
	case TypeNumberUint64:
		return unmarshalUint64(data, value)
	case TypeNumberFloat64:
		return unmarshalFloat64(data, value)
	case TypeStringEmpty:
		return unmarshalStringEmpty(data, value)
	case TypeStringShort:
		return unmarshalStringShort(data, value)
	case TypeStringMedium:
		return unmarshalStringMedium(data, value)
	case TypeStringLong:
		return unmarshalStringLong(data, value)
	case TypeArrayEmpty:
		return unmarshalArrayEmpty(data, value)
	case TypeArrayShort:
		return unmarshalArrayShort(data, value)
	case TypeArrayMedium:
		return unmarshalArrayMedium(data, value)
	case TypeArrayLong:
		return unmarshalArrayLong(data, value)
	case TypeObjectEmpty:
		return unmarshalObjectEmpty(data, value)
	case TypeObjectShort:
		return unmarshalObjectShort(data, value)
	case TypeObjectMedium:
		return unmarshalObjectMedium(data, value)
	case TypeObjectLong:
		return unmarshalObjectLong(data, value)
	default:
		return data, fmt.Errorf("invalid type: 0x%02x", dataType)
	}
}

func unmarshalUndefined(data []byte, value *any) ([]byte, error) {
	*value = "\x00"
	return data, nil
}

func unmarshalNull(data []byte, value *any) ([]byte, error) {
	*value = nil
	return data, nil
}

func unmarshalBooleanFalse(data []byte, value *any) ([]byte, error) {
	*value = false
	return data, nil
}

func unmarshalBooleanTrue(data []byte, value *any) ([]byte, error) {
	*value = true
	return data, nil
}

func unmarshalInt8(data []byte, value *any) ([]byte, error) {
	if len(data) < 1 {
		return data, fmt.Errorf("data is too short")
	}

	*value = float64(data[1])
	data = data[1:]
	return data, nil
}

func unmarshalInt16(data []byte, value *any) ([]byte, error) {
	if len(data) < 2 {
		return data, fmt.Errorf("data is too short")
	}

	*value = float64(int16(data[1])<<8 | int16(data[2]))
	data = data[2:]
	return data, nil
}

func unmarshalInt32(data []byte, value *any) ([]byte, error) {
	if len(data) < 4 {
		return data, fmt.Errorf("data is too short")
	}

	*value = float64(int32(data[1])<<24 | int32(data[2])<<16 | int32(data[3])<<8 | int32(data[4]))
	data = data[4:]
	return data, nil
}

func unmarshalInt64(data []byte, value *any) ([]byte, error) {
	if len(data) < 8 {
		return data, fmt.Errorf("data is too short")
	}

	*value = float64(
		int64(data[1])<<56 | int64(data[2])<<48 | int64(data[3])<<40 | int64(data[4])<<32 |
			int64(data[5])<<24 | int64(data[6])<<16 | int64(data[7])<<8 | int64(data[8]),
	)

	return data, nil
}

func unmarshalUint8(data []byte, value *any) ([]byte, error) {
	if len(data) < 1 {
		return data, fmt.Errorf("data is too short")
	}

	*value = float64(uint8(data[0]))
	data = data[1:]
	return data, nil
}

func unmarshalUint16(data []byte, value *any) ([]byte, error) {
	if len(data) < 2 {
		return data, fmt.Errorf("data is too short")
	}

	*value = float64(uint16(data[0])<<8 | uint16(data[1]))
	data = data[2:]
	return data, nil
}

func unmarshalUint32(data []byte, value *any) ([]byte, error) {
	if len(data) < 4 {
		return data, fmt.Errorf("data is too short")
	}

	*value = float64(uint32(data[0])<<24 | uint32(data[1])<<16 | uint32(data[2])<<8 | uint32(data[3]))
	data = data[4:]
	return data, nil
}

func unmarshalUint64(data []byte, value *any) ([]byte, error) {
	if len(data) < 8 {
		return data, fmt.Errorf("data is too short")
	}

	*value = float64(
		uint64(data[0])<<56 | uint64(data[1])<<48 | uint64(data[2])<<40 | uint64(data[3])<<32 |
			uint64(data[4])<<24 | uint64(data[5])<<16 | uint64(data[6])<<8 | uint64(data[7]),
	)

	data = data[8:]
	return data, nil
}

func unmarshalFloat64(data []byte, value *any) ([]byte, error) {
	if len(data) < 8 {
		return data, fmt.Errorf("data is too short")
	}

	*value = float64(math.Float64frombits(
		uint64(data[0])<<56 | uint64(data[1])<<48 | uint64(data[2])<<40 | uint64(data[3])<<32 |
			uint64(data[4])<<24 | uint64(data[5])<<16 | uint64(data[6])<<8 | uint64(data[7]),
	))

	data = data[8:]
	return data, nil
}

func unmarshalString(data []byte, dataType Type, value *any) ([]byte, error) {
	switch dataType {
	case TypeStringEmpty:
		return unmarshalStringEmpty(data, value)
	case TypeStringShort:
		return unmarshalStringShort(data, value)
	case TypeStringMedium:
		return unmarshalStringMedium(data, value)
	case TypeStringLong:
		return unmarshalStringLong(data, value)
	default:
		return data, fmt.Errorf("invalid string type: 0x%02x", dataType)
	}
}

func unmarshalStringEmpty(data []byte, value *any) ([]byte, error) {
	if len(data) != 0 {
		return data, fmt.Errorf("data is not empty")
	}

	*value = ""
	return data, nil
}

func unmarshalStringShort(data []byte, value *any) ([]byte, error) {
	if len(data) < 1 {
		return data, fmt.Errorf("data is too short")
	}

	stringLength := uint8(data[0])
	data = data[1:]

	if len(data) < int(stringLength) {
		return data, fmt.Errorf("data is too short")
	}

	*value = string(data[:stringLength])
	data = data[stringLength:]
	return data, nil
}

func unmarshalStringMedium(data []byte, value *any) ([]byte, error) {
	if len(data) < 2 {
		return data, fmt.Errorf("data is too short")
	}

	stringLength := uint16(data[0])<<8 | uint16(data[1])
	data = data[2:]

	if len(data) < int(stringLength) {
		return data, fmt.Errorf("data is too short")
	}

	*value = string(data[:stringLength])
	data = data[stringLength:]
	return data, nil
}

func unmarshalStringLong(data []byte, value *any) ([]byte, error) {
	if len(data) < 4 {
		return data, fmt.Errorf("data is too short")
	}

	stringLength := uint32(data[0])<<24 | uint32(data[1])<<16 | uint32(data[2])<<8 | uint32(data[3])
	data = data[4:]

	if len(data) < int(stringLength) {
		return data, fmt.Errorf("data is too short")
	}

	*value = string(data[:stringLength])
	data = data[stringLength:]
	return data, nil
}

func unmarshalArrayEmpty(data []byte, value *any) ([]byte, error) {
	*value = []any{}
	return data, nil
}

func unmarshalArrayShort(data []byte, value *any) ([]byte, error) {
	if len(data) < 1 {
		return data, fmt.Errorf("data is too short")
	}

	arrayLength := uint8(data[0])
	data = data[1:]
	return unmarshalArrayBody(data, int(arrayLength), value)
}

func unmarshalArrayMedium(data []byte, value *any) ([]byte, error) {
	if len(data) < 2 {
		return data, fmt.Errorf("data is too short")
	}

	arrayLength := uint16(data[0])<<8 | uint16(data[1])
	data = data[2:]
	return unmarshalArrayBody(data, int(arrayLength), value)
}

func unmarshalArrayLong(data []byte, value *any) ([]byte, error) {
	if len(data) < 4 {
		return data, fmt.Errorf("data is too short")
	}

	arrayLength := uint32(data[0])<<24 | uint32(data[1])<<16 | uint32(data[2])<<8 | uint32(data[3])
	data = data[4:]
	return unmarshalArrayBody(data, int(arrayLength), value)
}

func unmarshalArrayBody(data []byte, arrayLength int, value *any) ([]byte, error) {
	array := make([]any, arrayLength)

	for itemIndex := 0; itemIndex < int(arrayLength); itemIndex++ {
		var err error
		data, err = unmarshal(data, &array[itemIndex])
		if err != nil {
			return data, fmt.Errorf("failed to unmarshal item %d: %w", itemIndex, err)
		}
	}

	*value = array
	return data, nil
}

func unmarshalObjectEmpty(data []byte, value *any) ([]byte, error) {
	*value = map[string]any{}
	return data, nil
}

func unmarshalObjectShort(data []byte, value *any) ([]byte, error) {
	if len(data) < 1 {
		return data, fmt.Errorf("data is too short")
	}

	objectLength := uint8(data[0])
	data = data[1:]
	return unmarshalObjectBody(data, int(objectLength), value)
}

func unmarshalObjectMedium(data []byte, value *any) ([]byte, error) {
	if len(data) < 2 {
		return data, fmt.Errorf("data is too short")
	}

	objectLength := uint16(data[0])<<8 | uint16(data[1])
	data = data[2:]
	return unmarshalObjectBody(data, int(objectLength), value)
}

func unmarshalObjectLong(data []byte, value *any) ([]byte, error) {
	if len(data) < 4 {
		return data, fmt.Errorf("data is too short")
	}

	objectLength := uint32(data[0])<<24 | uint32(data[1])<<16 | uint32(data[2])<<8 | uint32(data[3])
	data = data[4:]
	return unmarshalObjectBody(data, int(objectLength), value)
}

func unmarshalObjectBody(data []byte, objectLength int, value *any) ([]byte, error) {
	object := make(map[string]any, objectLength)

	for propertyIndex := 0; propertyIndex < objectLength; propertyIndex++ {
		dataType, err := parseType(data)
		if err != nil {
			return data, fmt.Errorf("failed to parse type: %w", err)
		}
		data = data[1:]

		var propertyKey any
		data, err = unmarshalString(data, dataType, &propertyKey)
		if err != nil {
			return data, fmt.Errorf("failed to unmarshal property key: %w", err)
		}

		var propertyValue any
		data, err = unmarshal(data, &propertyValue)
		if err != nil {
			return data, fmt.Errorf("failed to unmarshal property value: %w", err)
		}

		object[propertyKey.(string)] = propertyValue
	}

	*value = object
	return data, nil
}
