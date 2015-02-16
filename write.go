// Writing files in Go follows similar patterns to the
// ones we saw earlier for reading.

package main

import (
	// "bufio"
	"fmt"
	"io/ioutil"
	// "os"
	//"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	var a int
	var b int
	var c int
	var d int
	var sum int
	var i int
	var j int
	var k int
	var ave float64
	j = 0
	k = 0
	sum = 0

	a = 3
	b = 5
	c = 1
	d = 1000

	for h := c; h < d; h++ {
		if h%a == 0 || h%b == 0 {
			k++
		}
	}

	x := make([]int, k)

	for i = c; i < d; i++ {
		if i%a == 0 || i%b == 0 {
			x[j] = i
			j++
			sum += i
		}
	}

	ave = float64(sum) / float64(k)

	fmt.Println(x)
	//fmt.Println(k)
	fmt.Println("Sum =", sum)
	fmt.Println("Average =", ave)

	// d1 := a.Write([]byte(x))
	d1 := []byte(string(sum))

	// To start, here's how to dump a string (or just
	// bytes) into a file.
	// d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("dat1.txt", d1, 0644)
	check(err)

}
