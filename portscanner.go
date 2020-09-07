package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"sync"
)

var (
	targetIP  string
	startPort int
	endPort   int
	showClose bool
)

func ipCheck(wg *sync.WaitGroup, port string) {

	// connection, error
	_, err := net.Dial("tcp", targetIP+":"+port)

	if err != nil {
		if showClose == true {
			fmt.Printf("Port: %s\n", port+"\tClose")
		}
	} else {
		fmt.Printf("Port: %s\n", port+"\tOpen")
	}

	defer wg.Done()
}

func main() {
	var ports []int
	var wg sync.WaitGroup

	flag.StringVar(&targetIP, "targetIP", "127.0.0.1", "Target IP address")
	flag.IntVar(&startPort, "startPort", 0, "Start Port Range")
	flag.IntVar(&endPort, "endPort", 1024, "End Port Range")
	flag.BoolVar(&showClose, "showClose", false, "Show Close Port")
	flag.Parse()

	for i := startPort; i <= endPort; i++ {
		ports = append(ports, i)
	}

	for _, port := range ports {
		wg.Add(1)
		go ipCheck(&wg, strconv.Itoa(port))
		wg.Wait()
	}

}
