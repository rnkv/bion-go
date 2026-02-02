package bion

import (
	"fmt"
	"math"
)

func Marshal(value any) ([]byte, error) {
	switch value := value.(type) {
	case nil:
		return []byte{byte(TypeNull)}, nil
	case bool:
		return marshalBoolean(value)
	case int:
		return marshalInt(value)
	case float64:
		return marshalFloat64(value)
	case string:
		if value == "\x00" {
			return []byte{byte(TypeUndefined)}, nil
		}

		return marshalString(value)
	case map[string]any:
		return marshalMap(value)
	case []any:
		return marshalArray(value)
	default:
		panic(fmt.Sprintf("type [%T] is not supported", value))
	}
}

func marshalBoolean(value bool) ([]byte, error) {
	if value {
		return []byte{byte(TypeBooleanTrue)}, nil
	}

	return []byte{byte(TypeBooleanFalse)}, nil
}

func marshalInt(value int) ([]byte, error) {
	if value < 0 {
		if value >= -128 {
			return []byte{
				byte(TypeNumberInt8),
				byte(value),
			}, nil
		}

		if value >= -32768 {
			return []byte{
				byte(TypeNumberInt16),
				byte(value >> 8),
				byte(value & 0xFF),
			}, nil
		}

		if value >= -2147483648 {
			return []byte{
				byte(TypeNumberInt32),
				byte(value >> 24),
				byte(value >> 16),
				byte(value >> 8),
				byte(value & 0xFF),
			}, nil
		}

		return []byte{
			byte(TypeNumberInt64),
			byte(value >> 56),
			byte(value >> 48),
			byte(value >> 40),
			byte(value >> 32),
			byte(value >> 24),
			byte(value >> 16),
			byte(value >> 8),
			byte(value & 0xFF),
		}, nil
	}

	if value < 256 {
		return []byte{
			byte(TypeNumberUint8),
			byte(value),
		}, nil
	}

	if value < 65536 {
		return []byte{
			byte(TypeNumberUint16),
			byte(value >> 8),
			byte(value & 0xFF),
		}, nil
	}

	if value < 16777216 {
		return []byte{
			byte(TypeNumberUint32),
			byte(value >> 24),
			byte(value >> 16),
			byte(value >> 8),
			byte(value & 0xFF),
		}, nil
	}

	return []byte{
		byte(TypeNumberUint64),
		byte(value >> 56),
		byte(value >> 48),
		byte(value >> 40),
		byte(value >> 32),
		byte(value >> 24),
		byte(value >> 16),
		byte(value >> 8),
		byte(value & 0xFF),
	}, nil
}

func marshalFloat64(value float64) ([]byte, error) {
	intPart, fracPart := math.Modf(value)

	if fracPart == 0 &&
		intPart >= math.MinInt &&
		intPart <= math.MaxInt {
		return marshalInt(int(intPart))
	}

	bits := math.Float64bits(value)

	return []byte{
		byte(TypeNumberFloat64),
		byte(bits >> 56),
		byte(bits >> 48),
		byte(bits >> 40),
		byte(bits >> 32),
		byte(bits >> 24),
		byte(bits >> 16),
		byte(bits >> 8),
		byte(bits),
	}, nil
}

func marshalString(value string) ([]byte, error) {
	if len(value) == 0 {
		return []byte{byte(TypeStringEmpty)}, nil
	}

	if len(value) < 256 {
		data := make([]byte, 2+len(value))
		data[0] = byte(TypeStringShort)
		data[1] = byte(len(value))
		copy(data[2:], value)
		return data, nil
	}

	if len(value) < 65536 {
		data := make([]byte, 3+len(value))
		data[0] = byte(TypeStringMedium)
		data[1] = byte(len(value) >> 8)
		data[2] = byte(len(value) & 0xFF)
		copy(data[3:], value)
		return data, nil
	}

	data := make([]byte, 4+len(value))
	data[0] = byte(TypeStringLong)
	data[1] = byte(len(value) >> 24)
	data[2] = byte(len(value) >> 16)
	data[3] = byte(len(value) >> 8)
	data[4] = byte(len(value) & 0xFF)
	copy(data[5:], value)
	return data, nil
}

func marshalMap(value map[string]any) ([]byte, error) {
	if len(value) == 0 {
		return []byte{byte(TypeObjectEmpty)}, nil
	}

	data := make([]byte, 2, 2+len(value)*32)

	if len(value) < 256 {
		data[0] = byte(TypeObjectShort)
		data[1] = byte(len(value))
	} else if len(value) < 65536 {
		data[0] = byte(TypeObjectMedium)
		data[1] = byte(len(value) >> 8)
		data[2] = byte(len(value) & 0xFF)
	} else {
		data[0] = byte(TypeObjectLong)
		data[1] = byte(len(value) >> 24)
		data[2] = byte(len(value) >> 16)
		data[3] = byte(len(value) >> 8)
		data[4] = byte(len(value) & 0xFF)
	}

	for itemKey, itemValue := range value {
		itemKeyData, err := marshalString(itemKey)
		if err != nil {
			return nil, err
		}

		itemValueData, err := Marshal(itemValue)
		if err != nil {
			return nil, err
		}

		data = append(data, itemKeyData...)
		data = append(data, itemValueData...)
	}

	return data, nil
}

func marshalArray(value []any) ([]byte, error) {
	if len(value) == 0 {
		return []byte{byte(TypeArrayEmpty)}, nil
	}

	data := make([]byte, 2, 2+len(value)*32)

	if len(value) < 256 {
		data[0] = byte(TypeArrayShort)
		data[1] = byte(len(value))
	} else if len(value) < 65536 {
		data[0] = byte(TypeArrayMedium)
		data[1] = byte(len(value) >> 8)
		data[2] = byte(len(value) & 0xFF)
	} else {
		data[0] = byte(TypeArrayLong)
		data[1] = byte(len(value) >> 24)
		data[2] = byte(len(value) >> 16)
		data[3] = byte(len(value) >> 8)
		data[4] = byte(len(value) & 0xFF)
	}

	for _, itemValue := range value {
		itemValueData, err := Marshal(itemValue)
		if err != nil {
			return nil, err
		}

		data = append(data, itemValueData...)
	}

	return data, nil
}
