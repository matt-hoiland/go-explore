package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pborman/getopt"
)

func main() {
	getopt.Parse()
	for _, arg := range getopt.Args() {
		fp, err := os.Open(arg)
		if _, ok := err.(*os.PathError); ok {
			fmt.Printf("%s: %s: No such file or directory\n", os.Args[0], arg)
			continue
		} else if err != nil {
			fmt.Println("Error unknown:", err)
			continue
		}

		data, err := ioutil.ReadAll(fp)
		if err != nil {
			fmt.Println("Error reading file:", err)
			continue
		}

		for line := 0; line < len(data)/0x10+1; line++ {
			fmt.Printf("%07x ", line*0x10)

			start := line * 0x10
			end := start + 0x10
			if end > len(data) {
				end = len(data)
			}

			for i := start; i < end-1; i++ {
				fmt.Printf("%02x ", data[i])
			}
			fmt.Printf("%02x\n", data[end-1])
		}
		fmt.Printf("%07x\n", len(data))
	}
}
