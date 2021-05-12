package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	fmt.Printf("\t,- port scanner\n\t|\n\t|\n")
	for i := 0; i < 8500; i++ {
		time.Sleep(15 * time.Millisecond)

		go func(j int) {
			address := fmt.Sprintf(os.Args[1]+":%d", j)
			conn, err := net.Dial("tcp", address)
			if err == nil {
				fmt.Printf("\t|-%d (i think its ssh)\n\t|\n", j)

				conn.Close()
			}

		}(i)

	}
	fmt.Printf("\t\\_ scan finish")

}
