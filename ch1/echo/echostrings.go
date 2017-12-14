package echo

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	//echo0()
	echo1()
	//echo2()
	//echo3()
	fmt.Printf("%v elapsed\n", time.Since(start).Nanoseconds())
}

func echo0() {

	fmt.Println(os.Args[0:])
}

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
		fmt.Printf("Indice[%v] - Conteúdo: %s\r\n", i, os.Args[i])
	}
}

func echo2() {
	//execio 1, 2
	for i, arg := range os.Args[0:] {
		s, sep := "", ""
		s += sep + arg
		fmt.Printf("Indice[%v] - Conteúdo: %s\r\n", i, s)
		sep = " "
	}
}

func echo3() {
	fmt.Println(strings.Join(os.Args[0:], "\n "))
}
