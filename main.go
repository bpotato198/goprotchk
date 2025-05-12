package main

import (
	"fmt"
)

func main() {
	var protinput string
	fmt.Printf("Goprotchk, a simple radio protocol checker\n")
	fmt.Scan(&protinput)
	switch protinput {
	case "bandplan":
		bandplan()
		break
	case "aprs":
		aprs()
		break
	}
}
func aprs() {
	fmt.Printf("APRS (Automatic Packet Reporting System) is a mode used by amateur radio operators\n")
	main()
}
