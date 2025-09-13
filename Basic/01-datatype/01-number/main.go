/*
Signed Integers :Can store both positive and negative values

Examples: int, int8, int16, int32, int64

Unsigned Integers: Can only store zero or positive values

Examples: uint, uint8, uint16, uint32, uint64


| Type    | Size                                                                | Minimum Value                  | Maximum Value                 |
| ------- | ------------------------------------------------------------------- | ------------------------------ | ----------------------------- |
| `int8`  | 8-bit                                                               | **-128**                       | **127**                       |
| `int16` | 16-bit                                                              | **-32,768**                    | **32,767**                    |
| `int32` | 32-bit                                                              | **-2,147,483,648**             | **2,147,483,647**             |
| `int64` | 64-bit                                                              | **-9,223,372,036,854,775,808** | **9,223,372,036,854,775,807** |
| `int`   | 32 or 64-bit depending on system (usually 64-bit on modern systems) |                                |                               |


| Type     | Size                             | Minimum Value | Maximum Value                  |
| -------- | -------------------------------- | ------------- | ------------------------------ |
| `uint8`  | 8-bit                            | **0**         | **255**                        |
| `uint16` | 16-bit                           | **0**         | **65,535**                     |
| `uint32` | 32-bit                           | **0**         | **4,294,967,295**              |
| `uint64` | 64-bit                           | **0**         | **18,446,744,073,709,551,615** |
| `uint`   | 32 or 64-bit depending on system |               |                                |


*/

package main

func main() {
	// Integer type
	var a int8 = 127
	var b int16 = 32767
	var c int32 = 2147483647
	var d int64 = 9223372036854775807

	// Unsigned integer type
	var ua uint8 = 255
	var ub uint16 = 65535
	var uc uint32 = 4294967295
	var ud uint64 = 18446744073709551615

	// Print the values to avoid unused variable error
	println(a, b, c, d, ua, ub, uc, ud)

}
