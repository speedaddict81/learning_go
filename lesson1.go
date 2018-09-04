package main

import (
        "log"
        "bufio"

        "github.com/tarm/serial"
)

func main() {
        c := &serial.Config{Name: "/dev/ttyS0", Baud: 9600}
        s, err := serial.OpenPort(c)
        if err != nil {
                log.Fatal(err)
        }
        
        lines := gpsSerialReader()
        n = len(lines)
        log.Printf("%q", lines[:n])
}


func gpsSerialReader() {
	defer serialPort.Close()
  lines = []string
	i := 0 //debug monitor
	scanner := bufio.NewScanner(serialPort)
	for scanner.Scan() && i<10 {
		i++

		s := scanner.Text()
		
	}
	return lines
}
