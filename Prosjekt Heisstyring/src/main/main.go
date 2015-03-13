package main

import ( 
	"fmt"
	//"udp"
	"driver"
	"control"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	floorChan := make(chan int)	

	if driver.InitElevator() == 0 {
		fmt.Println("Unable to initialize elevator hardware!")
		return
	}
		

	fmt.Println("Press STOP button to stop elevator and exit program.")
	
	go control.GoToFloor(2,floorChan)
	
	for {
		_, temp := control.GetCommand()
		floorChan<- temp
	
		
		if driver.GetStopSignal() != 0 {
			driver.SetMotorDirection(driver.DIRN_STOP)
			break
		}
	
	}
}		 
