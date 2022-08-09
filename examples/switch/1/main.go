package main

import "log"

func main() {

	log.Println(getNrOfRoutines(3))
	log.Println(getNrOfRoutines(5))
	log.Println(getNrOfRoutines(9))
	log.Println(getNrOfRoutines(12))
	log.Println(getNrOfRoutines(100))
	log.Println(getNrOfRoutines(101))
}

func getNrOfRoutines(nrOfConnections uint64) uint16 {
	routines := 0
	switch conn := nrOfConnections; {
	case conn <= 5:
		routines = 1
	case 6 <= conn && conn <= 10:
		routines = 2
	case 11 <= conn && conn <= 50:
		routines = 5
	case 51 <= conn && conn <= 100:
		routines = 10
	case 101 <= conn && conn <= 500:
		routines = 20
	case 501 <= conn && conn <= 1000:
		routines = 40
	case 1001 <= conn && conn <= 5000:
		routines = 80
	case 5001 <= conn && conn <= 10000:
		routines = 320
	case 10001 <= conn && conn <= 50000:
		routines = 400
	case 50001 <= conn:
		routines = 500
	default:
		routines = 1
	}
	return uint16(routines)
}
