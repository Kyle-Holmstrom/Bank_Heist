package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	isHeistOn := true
	eludedGuards := rand.Intn(100)

	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	} else {
		isHeistOn = false
		fmt.Println("Plan a better disguise next time?")
	}

	openedVault := rand.Intn(100)

	if isHeistOn == true && openedVault >= 60 {
		fmt.Println("Grab and GO!")
	} else if isHeistOn {
		isHeistOn = false
		fmt.Println("The vault can't be opened.")
	}

	leftSafely := rand.Intn(6)

	if isHeistOn == true {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("Looks like you tripped an alarm... run?")
		case 1:
			isHeistOn = false
			fmt.Println("Turns out vault doors don't open from the inside...")
		case 2:
			isHeistOn = false
			fmt.Println("You forgot the tools...")
		case 3:
			isHeistOn = false
			fmt.Println("Heist is a bust, doors won't open.")
		default:
			fmt.Println("Start the getaway car!")
		}
	}

	if isHeistOn {
		amtStolen := 10000 + rand.Intn(1000000)
		fmt.Println("The heist was a success and you escaped with : $", amtStolen)
	}

	//fmt.Println(isHeistOn)
} // end main
