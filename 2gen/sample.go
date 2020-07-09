package main

import "os"
import "syscall"

func writeln(s string) {
	//var s2 string = s + "\n"
	var s2 string = "\n"
	write(s)
	write(s2)
}

func write(s string) {
	var slc []uint8 = []uint8(s)
	syscall.Write(1, slc)
}

func exit(x int) {
	os.Exit(x)
}

func testArgAssign(x int) int {
	x = 13
	return x
}

func testMinus() int {
	var x int = -1
	x = x * -5
	return x
}

func sum(x int, y int) int {
	return x + y
}

func add1(x int) int {
	return x + 1
}

var globalint int
var globalint2 int
var globaluint8 uint8
var globaluint16 uint16
var globalarray [9]uint8
var globaluintptr uintptr

func testString() {
	var s string = "hello string"
	writeln(s)
}

func nop() {

}

func testIf() {
	var tr bool = true
	//if tr {
	//	writeln("ok true")
	//}
}


var globalslice []uint8

func testChar() {
	globalarray[0] = 'A'
	globalarray[1] = 'B'
	globalarray[2] = globalarray[0]
	globalarray[3] = 100 / 10 // '\n'
	globalarray[1] = 'B'
	var chars []uint8 = globalarray[0:4]
	write(string(chars))
	globalslice = chars
	write(string(globalarray[0:4]))
}

func testMisc() {
	var i13 int = 0
	i13 = testArgAssign(i13)
	var i5 int = testMinus()
	globalint2 = sum(1, i13 * i5)

	var locali3 int
	var tmp int
	tmp = int(uint8('3' - '1'))
	tmp = tmp + int(globaluint16)
	tmp = tmp + int(globaluint8)
	tmp = tmp + int(globaluintptr)
	locali3 = add1(tmp)
	var i42 int
	i42 = sum(globalint, globalint2) + locali3
	exit(i42)
}

func test() {
	testIf()
	testChar()
	testString()
	testMisc()
}

func main() {
	test()
}
