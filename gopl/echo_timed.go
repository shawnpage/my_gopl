// test different implementations of echo
package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

func main() {
	echo1()
	echo2()
	echo3()
	echo4()
}

func echo1() {
	start := time.Now()

	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)

	fmt.Printf("echo1: %.20fs elapsed\n", time.Since(start).Seconds())

}

func echo2() {
	start := time.Now()
	s, sep := "", ""

	for i, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
		fmt.Printf("%d\t%s\n", i, arg)
		//fmt.Println(i, arg)
	}

	//fmt.Println(s)
	fmt.Printf("echo2: %.20fs elapsed\n", time.Since(start).Seconds())

}

func echo3() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("echo3: %.20fs elapsed\n", time.Since(start).Seconds())
}

func echo4() {
	start := time.Now()

	fmt.Println(os.Args[1:])

	fmt.Printf("echo4: %.20fs elapsed\n", time.Since(start).Seconds())
}
