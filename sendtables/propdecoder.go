package sendtables

import (
	"fmt"
	"math"

	bs "github.com/ghostanalysis/demoinfocs-golang/bitread"
	"github.com/golang/geo/r3"
)

const (
	coordFractionalBitsMp             = 5
	coordFractionalBitsMpLowPrecision = 3
	coordDenominator                  = 1 << coordFractionalBitsMp
	coordResolution                   = 1.0 / coordDenominator
	coordDenominatorLowPrecision      = 1 << coordFractionalBitsMpLowPrecision
	coordResolutionLowPrecision       = 1.0 / coordDenominatorLowPrecision
	coordIntegerBitsMp                = 11
	coordIntegerBits                  = 14
)

const (
	normalFractBits   = 11
	normalDenominator = 1 << (normalFractBits - 1)
	normalResolution  = 1.0 / normalDenominator
)

const specialFloatFlags = SPF_NoScale | SPF_Coord | SPF_CellCoord | SPF_Normal | SPF_CoordMp | SPF_CoordMpLowPrecision | SPF_CoordMpIntegral | SPF_CellCoordLowPrecision | SPF_CellCoordIntegral

var propDecoder propertyDecoder

type PropValue struct {
	VectorVal r3.Vector
	IntVal    int
	ArrayVal  []PropValue
	StringVal string
	FloatVal  float32
}

type propertyDecoder struct{}

func (propertyDecoder) decodeProp(fProp *FlattenedPropEntry, reader *bs.BitReader) PropValue {
	switch fProp.prop.RawType {
	case SPT_Float:
		return PropValue{FloatVal: propDecoder.decodeFloat(fProp.prop, reader)}

	case SPT_Int:
		return PropValue{IntVal: propDecoder.decodeInt(fProp.prop, reader)}

	case SPT_VectorXY:
		return PropValue{VectorVal: propDecoder.decodeVectorXY(fProp.prop, reader)}

	case SPT_Vector:
		return PropValue{VectorVal: propDecoder.decodeVector(fProp.prop, reader)}

	case SPT_Array:
		return PropValue{ArrayVal: propDecoder.decodeArray(fProp, reader)}

	case SPT_String:
		return PropValue{StringVal: propDecoder.decodeString(fProp.prop, reader)}

	default:
		panic(fmt.Sprintf("Unknown prop type %d", fProp.prop.RawType))
	}
}

func (propertyDecoder) decodeInt(prop *SendTableProperty, reader *bs.BitReader) int {
	if prop.Flags.HasFlagSet(SPF_VarInt) {
		if prop.Flags.HasFlagSet(SPF_Unsigned) {
			return int(reader.ReadVarInt32())
		}
		return int(reader.ReadSignedVarInt32())
	}
	if prop.Flags.HasFlagSet(SPF_Unsigned) {
		return int(reader.ReadInt(uint(prop.NumberOfBits)))
	}
	return reader.ReadSignedInt(uint(prop.NumberOfBits))
}

func (propertyDecoder) decodeFloat(prop *SendTableProperty, reader *bs.BitReader) float32 {
	if prop.Flags&specialFloatFlags != 0 {
		return propDecoder.decodeSpecialFloat(prop, reader)
	}

	dwInterp := reader.ReadInt(uint(prop.NumberOfBits))
	return prop.LowValue + ((prop.HighValue - prop.LowValue) * (float32(dwInterp) / float32((int(1)<<uint(prop.NumberOfBits))-1)))
}

func (propertyDecoder) decodeSpecialFloat(prop *SendTableProperty, reader *bs.BitReader) float32 {
	// Because multiple flags can be set this order is fixed for now (priorities).
	// TODO: Would be more efficient to first check the most common ones tho.
	if prop.Flags.HasFlagSet(SPF_Coord) {
		return propDecoder.readBitCoord(reader)
	} else if prop.Flags.HasFlagSet(SPF_CoordMp) {
		return propDecoder.readBitCoordMp(reader, false, false)
	} else if prop.Flags.HasFlagSet(SPF_CoordMpLowPrecision) {
		return propDecoder.readBitCoordMp(reader, false, true)
	} else if prop.Flags.HasFlagSet(SPF_CoordMpIntegral) {
		return propDecoder.readBitCoordMp(reader, true, false)
	} else if prop.Flags.HasFlagSet(SPF_NoScale) {
		return reader.ReadFloat()
	} else if prop.Flags.HasFlagSet(SPF_Normal) {
		return propDecoder.readBitNormal(reader)
	} else if prop.Flags.HasFlagSet(SPF_CellCoord) {
		return propDecoder.readBitCellCoord(reader, uint(prop.NumberOfBits), false, false)
	} else if prop.Flags.HasFlagSet(SPF_CellCoordLowPrecision) {
		return propDecoder.readBitCellCoord(reader, uint(prop.NumberOfBits), true, false)
	} else if prop.Flags.HasFlagSet(SPF_CellCoordIntegral) {
		return propDecoder.readBitCellCoord(reader, uint(prop.NumberOfBits), false, true)
	}
	panic(fmt.Sprintf("Unexpected special float flag (Flags %v)", prop.Flags))
}

func (propertyDecoder) readBitCoord(reader *bs.BitReader) float32 {
	var intVal, fractVal int
	var res float32
	isNegative := false

	intVal = int(reader.ReadInt(1))
	fractVal = int(reader.ReadInt(1))

	if intVal|fractVal != 0 {
		isNegative = reader.ReadBit()
		if intVal == 1 {
			intVal = int(reader.ReadInt(coordIntegerBits) + 1)
		}

		if fractVal == 1 {
			fractVal = int(reader.ReadInt(coordFractionalBitsMp))
		}

		res = float32(intVal) + (float32(fractVal) * coordResolution)
	}

	if isNegative {
		res *= -1
	}

	return res
}

func (propertyDecoder) readBitCoordMp(reader *bs.BitReader, isIntegral bool, isLowPrecision bool) float32 {
	var res float32
	isNegative := false

	inBounds := reader.ReadBit()
	if isIntegral {
		if reader.ReadBit() {
			isNegative = reader.ReadBit()
			if inBounds {
				res = float32(reader.ReadInt(coordIntegerBitsMp) + 1)
			} else {
				res = float32(reader.ReadInt(coordIntegerBits) + 1)
			}
		}
	} else {
		readIntVal := reader.ReadBit()
		isNegative = reader.ReadBit()

		var intVal int
		if readIntVal {
			if inBounds {
				intVal = int(reader.ReadInt(coordIntegerBitsMp)) + 1
			} else {
				intVal = int(reader.ReadInt(coordIntegerBits)) + 1
			}
		}
		if isLowPrecision {
			res = float32(intVal) + (float32(reader.ReadInt(coordFractionalBitsMpLowPrecision)) * coordResolutionLowPrecision)
		} else {
			res = float32(intVal) + (float32(reader.ReadInt(coordFractionalBitsMp)) * coordResolution)
		}
	}

	if isNegative {
		res *= -1
	}

	return res
}

func (propertyDecoder) readBitNormal(reader *bs.BitReader) float32 {
	isNegative := reader.ReadBit()

	fractVal := reader.ReadInt(normalFractBits)

	res := float32(fractVal) * normalResolution

	if isNegative {
		res *= -1
	}

	return res
}

func (propertyDecoder) readBitCellCoord(reader *bs.BitReader, bits uint, isIntegral bool, isLowPrecision bool) float32 {
	var intVal, fractVal int
	var res float32

	if isIntegral {
		res = float32(reader.ReadInt(bits))
	} else {
		intVal = int(reader.ReadInt(bits))
		if isLowPrecision {
			fractVal = int(reader.ReadInt(coordFractionalBitsMpLowPrecision))

			res = float32(intVal) + (float32(fractVal) * (coordResolutionLowPrecision))
		} else {
			fractVal = int(reader.ReadInt(coordFractionalBitsMp))

			res = float32(intVal) + (float32(fractVal) * (coordResolution))
		}
	}

	return res
}

func (propertyDecoder) decodeVector(prop *SendTableProperty, reader *bs.BitReader) r3.Vector {
	res := r3.Vector{
		X: float64(propDecoder.decodeFloat(prop, reader)),
		Y: float64(propDecoder.decodeFloat(prop, reader)),
	}

	if !prop.Flags.HasFlagSet(SPF_Normal) {
		res.Z = float64(propDecoder.decodeFloat(prop, reader))
	} else {
		absolute := res.X*res.X + res.Y*res.Y
		if absolute < 1.0 {
			res.Z = math.Sqrt(1 - absolute)
		} else {
			res.Z = 0
		}

		if reader.ReadBit() {
			res.Z *= -1
		}
	}

	return res
}

func (propertyDecoder) decodeArray(fProp *FlattenedPropEntry, reader *bs.BitReader) []PropValue {
	numElement := fProp.prop.NumberOfElements

	var numBits uint = 1

	for maxElements := (numElement >> 1); maxElements != 0; maxElements = maxElements >> 1 {
		numBits++
	}

	nElements := int(reader.ReadInt(numBits))

	res := make([]PropValue, 0, nElements)

	tmp := &FlattenedPropEntry{prop: fProp.arrayElementProp}

	for i := 0; i < nElements; i++ {
		res = append(res, propDecoder.decodeProp(tmp, reader))
	}

	return res
}

func (propertyDecoder) decodeString(fProp *SendTableProperty, reader *bs.BitReader) string {
	length := int(reader.ReadInt(DT_MaxStringBits))
	if length > DT_MaxStringLength {
		length = DT_MaxStringLength
	}
	return reader.ReadCString(length)
}

func (propertyDecoder) decodeVectorXY(prop *SendTableProperty, reader *bs.BitReader) r3.Vector {
	return r3.Vector{
		X: float64(propDecoder.decodeFloat(prop, reader)),
		Y: float64(propDecoder.decodeFloat(prop, reader)),
	}
}
