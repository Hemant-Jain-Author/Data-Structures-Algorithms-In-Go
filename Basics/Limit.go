package main

import (
	"fmt"
	"math"
)
const PI2  = 3.1428

func main() {
	maxInt8 := math.MaxInt8
	minInt8 := math.MinInt8

	maxInt16 := math.MaxInt16
	minInt16 := math.MinInt16

	maxInt32 := math.MaxInt32
	minInt32 := math.MinInt32

	maxInt64 := math.MaxInt64
	minInt64 := math.MinInt64

	maxUint8 := math.MaxUint8
	maxUint16 := math.MaxUint16
	maxUint32 := math.MaxUint32
	var maxUint64 uint64 = math.MaxUint64
	maxFloat32 := math.MaxFloat32
	maxFloat64 := math.MaxFloat64

	fmt.Println("Range of Int8 :: ", minInt8, " to ", maxInt8)
	fmt.Println("Range of Int16 :: ", minInt16, " to ", maxInt16)
	fmt.Println("Range of Int32 :: ", minInt32, " to ", maxInt32)
	fmt.Println("Range of Int64 :: ", minInt64, " to ", maxInt64)
	fmt.Println("Max Uint8 :: ", maxUint8)
	fmt.Println("Max Uint16 :: ", maxUint16)
	fmt.Println("Max Uint32 :: ", maxUint32)
	fmt.Println("Max Uint64 :: ", maxUint64)
	fmt.Println("Max Float32 :: ", maxFloat32)
	fmt.Println("Max Float64 :: ", maxFloat64)
}
/*

Range of Int8 ::  -128  to  127
Range of Int16 ::  -32768  to  32767
Range of Int32 ::  -2147483648  to  2147483647
Range of Int64 ::  -9223372036854775808  to  9223372036854775807
Max Uint8 ::  255
Max Uint16 ::  65535
Max Uint32 ::  4294967295
Max Uint64 ::  18446744073709551615
Max Float32 ::  3.4028234663852886e+38
Max Float64 ::  1.7976931348623157e+308

*/