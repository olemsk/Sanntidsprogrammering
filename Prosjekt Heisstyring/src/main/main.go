package main

import ( 
	"fmt"
	"udp"
	"driver"
	"control"
	"runtime" 
	//"net"
	//"os"
)

func main() {
	
	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println(udp.GetID())	
	floorChan := make(chan int)
	//var Status udp.Status
	var Data udp.Data	
	//Data := make(map[int]udp.Status)
	//var PrimaryQ [3]string

	udp.UdpInit(30169, 39998, 1024, &Data)
	//Status.ID = udp.GetID()	
	fmt.Println("Getfloor", driver.GetFloorSensorSignal())	
	//PrintStatus(Status)
	//Data[Status.ID] = Status;
	//Data[154] = Status;
	//fmt.Println("Hera er Data")
	//PrintStatus(Data[153])
	//fmt.Println("Data hadde lengden: ", len(Data))

	if driver.InitElevator() == 0 {
		fmt.Println("Unable to initialize elevator hardware!")
		return
	}
		
	


	fmt.Println("Press STOP button to stop elevator and exit program.")
	
	//if Status.Primary == true {
	//	go udp.Send()
	//} else {
	//	go udp.Listen()
	//}	
		
	go control.GoToFloor(2,floorChan,&Data)
	
	for {
		_, temp := control.GetCommand()
		floorChan<- temp
		//PrintStatus(Data.Status)
		fmt.Println("Stop signal pressed ", driver.GetStopSignal())
		/*if driver.GetStopSignal() != 0 {
			fmt.Println("Stop signal pressed ", driver.GetStopSignal())			
			driver.SetMotorDirection(driver.DIRN_STOP)
			break
		}
	*/
	}
}		 

func PrintStatus(Status udp.Status) {
	fmt.Println("Running: ", Status.Running)
	fmt.Println("CurrentFloor: ", Status.CurrentFloor)
	fmt.Println("NextFloor: ", Status.NextFloor)
	fmt.Println("Primary: ", Status.Primary)
	fmt.Println("ID: ", Status.ID)
}
