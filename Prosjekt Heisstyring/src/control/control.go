//Sanntidsprogrammering!!
package control

import ( 
	"fmt"
	"udp"
	//"runtime"
	"time"
	"driver"
	//"control"
	//"os"
	"functions"
	
	
)
/*
func GoToFloor(button int,  floorChan chan int,data *udp.Data) {
	floor := <-floorChan
	if driver.GetFloorSensorSignal() == -1 {
		driver.SetMotorDirection(driver.DIRN_DOWN)
	}
	var done int
	temp:= floor	
	//polse:
	for {		
		select {
		
		case temp = <-floorChan:
			//fmt.Println("Her er temp: %d", temp)
			//fmt.Println("Her er DONE: %d", done)
			//if done == 1{
									
			//	floor = temp
			//	done = 0
			//}

		default:			

			driver.SetFloorIndicator(driver.GetFloorSensorSignal())	
			if done == 1{
				//fmt.Printf("GAA IN EHFE")				
				floor = temp
				done = 0
				
				
			}	
			//fmt.Printf("Hva er done? %d\n",done)
			driver.SetButtonLamp(button,floor,1)
			//fmt.Printf("Her er flooooooooor: %d\n", floor)
				
			if floor == driver.GetFloorSensorSignal()  {
				
				fmt.Println("Framme på:", floor)
				udp.SetStatus(data,0,floor)

				driver.SetDoorOpenLamp(true)				
				driver.SetMotorDirection(driver.DIRN_STOP)
				time.Sleep(1*time.Second)
				driver.SetDoorOpenLamp(false)
				driver.SetFloorIndicator(floor)
				driver.SetButtonLamp(button,floor,0)
				done = 1
				
				//temp = -1
				//driver.SetDoorOpenLamp(false)	
				//fmt.Println("Done: %d", done)
				break
		
			} else if floor > driver.GetFloorSensorSignal() && driver.GetFloorSensorSignal() != -1 && floor != -1 {   
			udp.SetStatus(data,2, floor)
			driver.SetMotorDirection(driver.DIRN_UP) 
		
			} else if floor < driver.GetFloorSensorSignal() && driver.GetFloorSensorSignal() != -1 && floor != -1{
			udp.SetStatus(data ,1, floor)
			driver.SetMotorDirection(driver.DIRN_DOWN)
			}

		}
	
	}
}
*/
func GoToFloor(floor int, status *udp.Status, list int) { // Lamper for command buttons må leggas til, kall list noe annet
	fmt.Println("control 82: går til floor floor:",floor)
	fmt.Println("control 82: går til floor floor:",status.CurrentFloor)
	for {
		driver.SetFloorIndicator(driver.GetFloorSensorSignal())
		if floor == driver.GetFloorSensorSignal() {
				driver.SetFloorIndicator(floor)
				driver.SetMotorDirection(driver.DIRN_STOP)
				driver.SetDoorOpenLamp(true)				
				time.Sleep(2*time.Second)
				driver.SetDoorOpenLamp(false)
				if floor == 0 || floor == 3 {
					(*status).Running = 0
				}
				//if list == 1 {
				//	driver.SetButtonLamp((*status).ButtonList[0], floor, 0)
				//	(*status).ButtonList = functions.UpdateList((*status).ButtonList,0)
				//}
				fmt.Println("Heisen er framme på floor:", floor)
				break
		} else if floor > driver.GetFloorSensorSignal() && driver.GetFloorSensorSignal() != -1 && floor != -1 {   
			driver.SetMotorDirection(driver.DIRN_UP) 
		} else if floor < driver.GetFloorSensorSignal() && driver.GetFloorSensorSignal() != -1 && floor != -1{
			driver.SetMotorDirection(driver.DIRN_DOWN)
		}else if status.CurrentFloor == -1{
			
			driver.SetMotorDirection(driver.DIRN_DOWN)
		}
		if driver.GetFloorSensorSignal() != -1{
			(*status).CurrentFloor = driver.GetFloorSensorSignal()
		}	
	}
}


func ElevatorControl(status *udp.Status) {
	//time.Sleep(1*time.Second)
	//var status *udp.Status
	temp := 0
	//temp = temp + 0
	
	for {
		//&&status = <-statusIn
	
		if driver.GetFloorSensorSignal() != -1 {
			status.CurrentFloor = driver.GetFloorSensorSignal()
		}
		//fmt.Println(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CurrentFloor) 
		//fmt.Println("control 109: OrderList",status.OrderList)
		//if len((*status).OrderList)==0 {
		//	(*status).OrderList = append((*status).OrderList, -1)
		//}
		if len(status.OrderList) == 0{
			status.Running = 0
		//	statusOut<-status
		} 
		//for i:=0;i<len(status.ButtonList);i++ {
		//	fmt.Println("ButtonList(i): ",status.ButtonList[i])
		//	fmt.Println("OrderList(i): ",status.OrderList[i])
		//	driver.SetButtonLamp(status.ButtonList[i], status.OrderList[i], 1)
		//}
		//ButtonList = ButtonList[:0]
						
		//fmt.Printf("OrderList: %d CommandList[0]: CurrentFloor: %d ID: %d \n",(*status).OrderList, (*status).CurrentFloor, (*status).ID)
		
		if len(status.OrderList) > 0{
				fmt.Println("OrderList: ", status.OrderList)
				// 
				if status.OrderList[0] > status.CurrentFloor  {
					//fmt.Println(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
					// Sjekker om heisens ordreliste
					temp = status.OrderList[0]
					status.OrderList = functions.UpdateList(status.OrderList,0)
					GoToFloor(temp, status, 1) // vurdere å kjøre commandbuttons inni gotofloor
					temp = 0
				}else if status.OrderList[0] < status.CurrentFloor{
					temp = status.OrderList[0]
					status.OrderList = functions.UpdateList(status.OrderList,0)
					GoToFloor(temp, status,1)
					temp = 0
						
				}else if 	status.OrderList[0] == driver.GetFloorSensorSignal() {
						status.OrderList=functions.UpdateList(status.OrderList,0)
						GoToFloor(driver.GetFloorSensorSignal(), status,1)
				}
				/*
			if !((*status).OrderList[0] == -1 && (*status).CommandList[0] ==-1){
				fmt.Println("OrderList: ", (*status).CurrentFloor)
				// 
				if (*status).OrderList[0] > (*status).CurrentFloor  {
					//status.Running =1
					//fmt.Println(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
					// Sjekker om heisens ordreliste
					if (*status).CommandList[0] == -1{
						temp = (*status).OrderList[0]
						(*status).OrderList = functions.UpdateList((*status).OrderList,0)
						
						GoToFloor(temp, status,1)
						temp = 0
					}else if (*status).OrderList[0] == -1{
						temp = (*status).CommandList[0]
						(*status).CommandList = functions.UpdateList((*status).CommandList,0)
						GoToFloor(temp, status,0)
						temp = 0
					}else if (*status).OrderList[0]>(*status).CommandList[0]{
						temp = (*status).CommandList[0]
						(*status).CommandList = functions.UpdateList((*status).CommandList,0)	
						GoToFloor(temp, status,0)
						temp = 0
					}else if (*status).CommandList[0]>(*status).OrderList[0]{
						temp = (*status).OrderList[0]
						(*status).OrderList = functions.UpdateList((*status).OrderList,0)
						GoToFloor(temp, status,1)
						temp = 0
					}else if (*status).OrderList[0] == (*status).CommandList[0]{
						temp = (*status).OrderList[0]
						(*status).CommandList=functions.UpdateList((*status).CommandList,0)
						(*status).OrderList=functions.UpdateList((*status).OrderList,0)
						GoToFloor(temp, status,1)
						temp = 0
					}
				}else if (*status).OrderList[0] < (*status).CurrentFloor{
					//status.Running = -1
					if (*status).CommandList[0] == -1 {
						temp = (*status).OrderList[0]
						(*status).OrderList = functions.UpdateList((*status).OrderList,0)
						GoToFloor(temp, status,1)
						temp = 0
					}else if (*status).OrderList[0] == -1 {
						temp = (*status).CommandList[0] 
						(*status).CommandList = functions.UpdateList((*status).CommandList,0)
						GoToFloor(temp, status,0)
						temp = 0
					}else if (*status).OrderList[0] < (*status).CommandList[0]{
						temp = (*status).CommandList[0]
						(*status).CommandList = functions.UpdateList((*status).CommandList,0)
						GoToFloor(temp, status,0)
						temp = 0
					}else if (*status).CommandList[0] < (*status).OrderList[0]{
						temp = (*status).OrderList[0]
						(*status).OrderList = functions.UpdateList((*status).OrderList,0)
						GoToFloor(temp, status,1)
						temp = 0
					}else if (*status).OrderList[0] == (*status).CommandList[0]{
						temp = (*status).OrderList[0]
						(*status).CommandList=functions.UpdateList((*status).CommandList,0)
						(*status).OrderList=functions.UpdateList((*status).OrderList,0)
						GoToFloor(temp, status,1)
						temp = 0						
					}						
				}else if 	(*status).OrderList[0] == driver.GetFloorSensorSignal() {
						(*status).OrderList=functions.UpdateList((*status).OrderList,0)
						GoToFloor(driver.GetFloorSensorSignal(), status,1)						
				}
			}
		*/
		}	
	}
}
	
		
func GetDestination(data.Statuses[udp.GetIndex(udp.GetID(),data)] *udp.data.Statuses[udp.GetIndex(udp.GetID(),data)]) { //returnerer bare button, orderlist oppdateres
	//time.Sleep(1*time.Second)
	for {
		time.Sleep(2*time.Millisecond) // Polling rate, mby change	
		for floor := 0; floor < driver.N_FLOORS; floor++ {
				if driver.GetButtonSignal(0,floor) == 1 && len(data.Statuses[udp.GetIndex(udp.GetID(),data)].UpList) == 0 {
					data.Statuses[udp.GetIndex(udp.GetID(),data)].UpList = append(data.Statuses[udp.GetIndex(udp.GetID(),data)].UpList, floor)
					fmt.Println("control: 250, data.Statuses[udp.GetIndex(udp.GetID(),data)].UpList: ", data.Statuses[udp.GetIndex(udp.GetID(),data)].UpList) 
				}else if driver.GetButtonSignal(0,floor) == 1 && len(data.Statuses[udp.GetIndex(udp.GetID(),data)].UpList) > 0 {
					if functions.CheckList(data.Statuses[udp.GetIndex(udp.GetID(),data)].UpList,floor) == false {
						data.Statuses[udp.GetIndex(udp.GetID(),data)].UpList = append(data.Statuses[udp.GetIndex(udp.GetID(),data)].UpList,floor)
						fmt.Println("control: 254, data.Statuses[udp.GetIndex(udp.GetID(),data)].UpList: ", data.Statuses[udp.GetIndex(udp.GetID(),data)].UpList) 
					}				
				}else if driver.GetButtonSignal(1,floor) == 1 && len(data.Statuses[udp.GetIndex(udp.GetID(),data)].DownList)==0 {	
					data.Statuses[udp.GetIndex(udp.GetID(),data)].DownList = append(data.Statuses[udp.GetIndex(udp.GetID(),data)].DownList, floor)
					fmt.Println("control: 257, data.Statuses[udp.GetIndex(udp.GetID(),data)].DownList: ", data.Statuses[udp.GetIndex(udp.GetID(),data)].DownList)
				}else if driver.GetButtonSignal(1,floor) == 1 && len(data.Statuses[udp.GetIndex(udp.GetID(),data)].DownList) > 0 {
					if functions.CheckList(data.Statuses[udp.GetIndex(udp.GetID(),data)].DownList,floor) == false {
						data.Statuses[udp.GetIndex(udp.GetID(),data)].DownList = append(data.Statuses[udp.GetIndex(udp.GetID(),data)].DownList,floor)
						fmt.Println("control: 260, data.Statuses[udp.GetIndex(udp.GetID(),data)].DownList: ", data.Statuses[udp.GetIndex(udp.GetID(),data)].DownList)
					}	
				}else if driver.GetButtonSignal(2,floor) == 1 && !functions.CheckList(data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList, floor){
					//if data.Statuses[udp.GetIndex(udp.GetID(),data)].Running == 0{
					//	
					//}
					if data.Statuses[udp.GetIndex(udp.GetID(),data)].CurrentFloor < floor && data.Statuses[udp.GetIndex(udp.GetID(),data)].Running == 1{
						data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList = append(data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList, floor)
						data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList = functions.SortUp(data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList)
					} else if data.Statuses[udp.GetIndex(udp.GetID(),data)].CurrentFloor > floor && data.Statuses[udp.GetIndex(udp.GetID(),data)].Running == -1{
						data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList = append(data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList, floor)
						data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList = functions.SortDown(data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList)
					}
				}
				/*
				}else if driver.GetButtonSignal(2,floor) == 1 && len(data.Statuses[udp.GetIndex(udp.GetID(),data)].CommandList) == 0{
						if data.Statuses[udp.GetIndex(udp.GetID(),data)].Running == 0 {
							if data.Statuses[udp.GetIndex(udp.GetID(),data)].CurrentFloor < floor{
								data.Statuses[udp.GetIndex(udp.GetID(),data)].Running = 1
								data.Statuses[udp.GetIndex(udp.GetID(),data)].CommandList = append(data.Statuses[udp.GetIndex(udp.GetID(),data)].CommandList,floor)
							}else if data.Statuses[udp.GetIndex(udp.GetID(),data)].CurrentFloor > floor{
								data.Statuses[udp.GetIndex(udp.GetID(),data)].Running = -1
								data.Statuses[udp.GetIndex(udp.GetID(),data)].CommandList = append(data.Statuses[udp.GetIndex(udp.GetID(),data)].CommandList,floor) 
							}else{
								data.Statuses[udp.GetIndex(udp.GetID(),data)].Running = 0
							}
						}else{
							data.Statuses[udp.GetIndex(udp.GetID(),data)].CommandList = append(data.Statuses[udp.GetIndex(udp.GetID(),data)].CommandList, floor)
						}
				}else if driver.GetButtonSignal(2,floor) == 1  && data.Statuses[udp.GetIndex(udp.GetID(),data)].CommandList[0] == -1 {
						if data.Statuses[udp.GetIndex(udp.GetID(),data)].Running == 0{
							if data.Statuses[udp.GetIndex(udp.GetID(),data)].CurrentFloor < floor{
								data.Statuses[udp.GetIndex(udp.GetID(),data)].Running = 1
								data.Statuses[udp.GetIndex(udp.GetID(),data)].CommandList[0] = floor
							}else if data.Statuses[udp.GetIndex(udp.GetID(),data)].CurrentFloor > floor{
								data.Statuses[udp.GetIndex(udp.GetID(),data)].Running = -1
								data.Statuses[udp.GetIndex(udp.GetID(),data)].CommandList[0] = floor
							}else{
								data.Statuses[udp.GetIndex(udp.GetID(),data)].Running = 0
							}
						}else{
							data.Statuses[udp.GetIndex(udp.GetID(),data)].CommandList[0] = floor
						}
						
				}else if driver.GetButtonSignal(2,floor) == 1 && len(data.Statuses[udp.GetIndex(udp.GetID(),data)].CommandList) > 0 {
					if data.Statuses[udp.GetIndex(udp.GetID(),data)].CommandList[len(data.Statuses[udp.GetIndex(udp.GetID(),data)].CommandList)-1] != floor {
						//data.Statuses[udp.GetIndex(udp.GetID(),data)].CommandList = append(data.Statuses[udp.GetIndex(udp.GetID(),data)].CommandList, floor)
						if data.Statuses[udp.GetIndex(udp.GetID(),data)].Running == 1{
							if floor <= data.Statuses[udp.GetIndex(udp.GetID(),data)].CurrentFloor{
								data.Statuses[udp.GetIndex(udp.GetID(),data)].CommandList = functions.SortUp(data.Statuses[udp.GetIndex(udp.GetID(),data)].CommandList)
								data.Statuses[udp.GetIndex(udp.GetID(),data)].CommandList = append(data.Statuses[udp.GetIndex(udp.GetID(),data)].CommandList, floor)
							}else{
								data.Statuses[udp.GetIndex(udp.GetID(),data)].CommandList = append(data.Statuses[udp.GetIndex(udp.GetID(),data)].CommandList, floor)
								data.Statuses[udp.GetIndex(udp.GetID(),data)].CommandList = functions.SortUp(data.Statuses[udp.GetIndex(udp.GetID(),data)].CommandList)
							}
						}else if data.Statuses[udp.GetIndex(udp.GetID(),data)].Running == -1{
							if floor >= data.Statuses[udp.GetIndex(udp.GetID(),data)].CurrentFloor{
								data.Statuses[udp.GetIndex(udp.GetID(),data)].CommandList = functions.SortDown(data.Statuses[udp.GetIndex(udp.GetID(),data)].CommandList)
								data.Statuses[udp.GetIndex(udp.GetID(),data)].CommandList = append(data.Statuses[udp.GetIndex(udp.GetID(),data)].CommandList, floor)
							}else{
								data.Statuses[udp.GetIndex(udp.GetID(),data)].CommandList = append(data.Statuses[udp.GetIndex(udp.GetID(),data)].CommandList, floor)
								data.Statuses[udp.GetIndex(udp.GetID(),data)].CommandList = functions.SortDown(data.Statuses[udp.GetIndex(udp.GetID(),data)].CommandList)
							}
						}
					}
				}				
				
				if data.Statuses[udp.GetIndex(udp.GetID(),data)].Running == 0 {
					data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList = append(data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList, floor)
					// tenne lampe?
				} else if data.Statuses[udp.GetIndex(udp.GetID(),data)].Running == 1 {
					if floor < data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList[len(data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList)-1] && floor > data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList[0] {
				}*/
					
		}
		/*
		if(driver.GetStopSignal() != 0) {
			driver.SetMotorDirection(driver.DIRN_STOP)
			break
		}
		*/
	}
}
/*
func GetCommand() (int,int) {
	button := 2	
	for {	
		for floor := 0; floor < driver.N_FLOORS; floor++ {
			if(driver.GetButtonSignal(button,floor) == 1) {
				return button,floor
			}
			
		}
		if(driver.GetStopSignal() != 0) {
			driver.SetMotorDirection(driver.DIRN_STOP)
			break
		}
	}
return -1,-1
}
*/

func CostFunction(in chan *udp.Data, out chan *udp.Data) {
	handled := 0
	var DownList []int
	var UpList []int
	var data *udp.Data
	fmt.Println("control: 369. costfunction, Ventar her")
	data = <-in
	fmt.Println("control: 371. costfunction, Går vidare")
	for {
		//fmt.Println("control 243, handled: ",handled)
		handled = 0
		//fmt.Println("status.UpList i CostFunction: ",(*data).Statuses[udp.GetIndex((*data).PrimaryQ[0], data)].UpList)
		//fmt.Println("Lengden til statuses: ", len(data.Statuses))
		//fmt.Println("PrimaryQ: ", data.PrimaryQ)
		/*if len(data.Statuses[udp.GetIndex(data.PrimaryQ[0], data)].OrderList) > 0 {
		if(data.Statuses[udp.GetIndex(data.PrimaryQ[0], data)].OrderList[0] == -1){
			data.Statuses[udp.GetIndex(data.PrimaryQ[0], data)].OrderList = UpdateList(data.Statuses[udp.GetIndex(data.PrimaryQ[0], data)].OrderList, 0)
		}
		}*/
		if len(UpList) == 0 && len(DownList) == 0{
			if handled == 0{
				out<- data
				data = <-in
				fmt.Println("control: 383. Er inne i costfunction")
			}else{
				out<-data // Tømt UpList og DownList 
				data = <-in // Venter
			}
		}else if len(UpList) > 0 || len(DownList) > 0{
			fmt.Printf("Could not handle all orders, will try again after recieving new.\nUpList: %v DownList: %v\n", UpList, DownList)
			out<-data
			data = <-in
		}
		for k := 0; k < len(data.PrimaryQ);k++ {
			if udp.GetIndex(data.PrimaryQ[k],data) != -1 {
				fmt.Printf("PrimaryQ: %v Lengde Statuses: %v\n", data.PrimaryQ, len(data.Statuses))
				fmt.Println("Har den samme info om status.Uplist her: ", data.Statuses[k].UpList)
				DownList = append(DownList,data.Statuses[udp.GetIndex(data.PrimaryQ[k], data)].DownList...)
				data.Statuses[udp.GetIndex(data.PrimaryQ[k], data)].DownList = data.Statuses[udp.GetIndex(data.PrimaryQ[k], data)].DownList[:0]

				UpList = append(UpList,data.Statuses[udp.GetIndex(data.PrimaryQ[k], data)].UpList...)
				data.Statuses[udp.GetIndex(data.PrimaryQ[k], data)].UpList = data.Statuses[udp.GetIndex(data.PrimaryQ[k], data)].UpList[:0]
			}
		}

	if len(UpList) > 0 || len(DownList) > 0{
		//fmt.Println("status.UpList i CostFunction: ",data.Statuses[udp.GetIndex(data.PrimaryQ[0], data)].UpList)
		fmt.Println("Lengden til statuses: ", len(data.Statuses))
		fmt.Println("PrimaryQ: ", data.PrimaryQ)
		fmt.Println("control 258: OppList i cost function: ", UpList)
		fmt.Println("control 259: Down List i cost function: ", DownList)
		time.Sleep(2*time.Second)
	}
	
	if len(UpList) > 0 {
		UpList = functions.SortUp(UpList)
	}
	if len(DownList) > 0 {
		DownList = functions.SortDown(DownList)
	}
	
	//fmt.Println("OrderList: ", data.Statuses[udp.GetIndex(data.PrimaryQ[0], data)].OrderList)
	/*if(len(data.PrimaryQ) > 1){
		data.Statuses[1].OrderList = append(data.Statuses[1].OrderList,3)
		
		udp.SendOrderlist(data,1)
	}*/
	//fmt.Println("Sjekk om UPLIST oppdateres riktig: ", UpList)
	//fmt.Println("Sjekk om DOWNLIST oppdateres riktig: ", DownList)
	//fmt.Println(DownList)
	for down:=0; down<len(DownList);down++ { // Kanskje feil å sette running inni her? 
		
		if handled == 1{
			handled = 0
			down = 0	
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ { // Heis i samme floor, og står stille
			if DownList[down] == data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CurrentFloor && data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == 0 {

				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,DownList[down])
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = functions.SortDown(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running = 0
				fmt.Println("control 280: Heis i samma floor og står stille. Downlist:", DownList)
				DownList = functions.UpdateList(DownList,down) //Må modifiseres
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].ButtonList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].ButtonList,1)
				if i != 0 {
					udp.SendOrderlist(data,i) // , udp.GetIndex(data.PrimaryQ[i], data))
				}else{
					
				}
				handled = 1
				break
				//pluss noe mer, som å åpne døra
			}
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ { // Heis i etasjen over og på veg nedover
			if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CurrentFloor == DownList[down]+1 && data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == -1 && handled != 1 {
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,DownList[down])
				fmt.Println("control 370: Heis i etasjen over og på vei nedover. Downlist:", DownList)
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = functions.SortDown(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
				DownList = functions.UpdateList(DownList,down)
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].ButtonList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].ButtonList,1)
				if i != 0 {
					udp.SendOrderlist(data, i) // , udp.GetIndex(data.PrimaryQ[i], data))
				}
				handled = 1
				break 
			}
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ { // Heis i etasjen over og står stille
			if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CurrentFloor == DownList[down]+1 && data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == 0 && handled != 1 {
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,DownList[down])
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = functions.SortDown(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running = -1
				fmt.Println("control 385: Heis i etasjen over og står stille")
				DownList = functions.UpdateList(DownList,down)
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].ButtonList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].ButtonList,1)
				if i != 0 {
					udp.SendOrderlist(data,i) // , udp.GetIndex(data.PrimaryQ[i], data))
				}
				handled = 1
				break 
			}
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ { // Heis på veg nedover
			if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CurrentFloor > DownList[down] && data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == -1  && handled != 1{
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,DownList[down])
				DownList = functions.UpdateList(DownList,down)
				fmt.Println("control 398: Heis på vei nedover og er over floor. Downlist:", DownList)
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = functions.SortDown(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].ButtonList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].ButtonList,1)
				if i != 0 {
					udp.SendOrderlist(data,i) // , udp.GetIndex(data.PrimaryQ[i], data))
				}
				handled = 1
				break 
			}
		}
		/*for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ { // Heis på veg oppover, men siste skal stoppe på denne etasjen
			if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == 1 && data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList[len(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)-1] == DownList[down] { 
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,DownList[down])
				fmt.Println("e")
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = SortDown(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
				DownList = UpdateList(DownList,down)
				if i != 0 {
					udp.SendOrderlist(data) // , udp.GetIndex(data.PrimaryQ[i], data))
				}
				handled = 1
				break 
			}	
		}*/
		/*for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ { // Heis på veg oppover, men siste stopp er under
			if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == 1 && data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList[len(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)-1] < DownList[down] {
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,DownList[down])
				fmt.Println("f")
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = SortDown(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
				DownList = UpdateList(DownList,down)
				if i != 0 {
					udp.SendOrderlist(data) // , udp.GetIndex(data.PrimaryQ[i], data))
				}
				handled = 1
				break 
			}
		}*/
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ {
			if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == 0 {
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,DownList[down])
				fmt.Println("control 437: heis står stille generelt. Downlist:", DownList)
				if DownList[down] > data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CurrentFloor{
					data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = functions.SortUp(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
					data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running = 1
				}else if DownList[down] < data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CurrentFloor{
					data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = functions.SortDown(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
					data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running = -1
				}
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].ButtonList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].ButtonList,1)
				DownList = functions.UpdateList(DownList,down)
				if i != 0 {
					udp.SendOrderlist(data,i) // , udp.GetIndex(data.PrimaryQ[i], data))
				}
				handled = 1
				break 
			}
		}

		/*if handled == 0 {
			data.Statuses[data.PrimaryQ[0]].OrderList = append(data.Statuses[data.PrimaryQ[0]].OrderList,DownList[down])
			data.Statuses[udp.GetIndex(data.PrimaryQ[0], data)].OrderList = SortDown(data.Statuses[udp.GetIndex(data.PrimaryQ[0], data)].OrderList)
			DownList = UpdateList(DownList,down)
			fmt.Println("h")
			handled = 1 
		}*/
	}

for up:=0; up<len(UpList);up++ {
		//fmt.Println("Up: ",up)
		//fmt.Println(data.PrimaryQ)
		if handled == 1{
			handled = 0
			up = 0	
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ {
			if UpList[up] == data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CurrentFloor && data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == 0 {
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,UpList[up])
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = functions.SortUp(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].NextFloor = UpList[up]
				fmt.Println("control 387: heis i samme etasjen og står stille. UpList:", UpList)
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].ButtonList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].ButtonList,0)
				UpList = functions.UpdateList(UpList,up) //Må modifiseres
				if i != 0 {
					udp.SendOrderlist(data,i) // , udp.GetIndex(data.PrimaryQ[i], data))
				}
				handled = 1
				break 
				//pluss noe mer, som å åpne døra
			}
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ {
			if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CurrentFloor == UpList[up]-1 && data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == 1 && handled != 1 {
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,UpList[up])
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = functions.SortUp(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].NextFloor = UpList[up]
				fmt.Println("control 402: heis i etasjen under og på vei oppover. UpList:", UpList)
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].ButtonList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].ButtonList,0)
				UpList = functions.UpdateList(UpList,up)
				if i != 0 {
					udp.SendOrderlist(data,i) // , udp.GetIndex(data.PrimaryQ[i], data))
				}
				handled = 1
				break 
			}
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ {
			if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CurrentFloor == UpList[up]-1 && data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == 0 && handled != 1 {
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,UpList[up])
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = functions.SortUp(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].NextFloor = UpList[up]
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running = 1
				fmt.Println("control 417: heis i etasjen under og står stille. UpList:", UpList)
				UpList = functions.UpdateList(UpList,up)
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].ButtonList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].ButtonList,0)
				if i != 0 {
					udp.SendOrderlist(data,i) // , udp.GetIndex(data.PrimaryQ[i], data))
				}
				handled = 1
				break 
			}
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ {
			if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CurrentFloor < UpList[up] && data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == 1  && handled != 1{
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,UpList[up])
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = functions.SortUp(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
				fmt.Println("control 430: floor over heis.currentfloor og på vei oppover. UpList:", UpList)
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].NextFloor = UpList[up]
				UpList = functions.UpdateList(UpList,up)
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].ButtonList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].ButtonList,0)
				if i != 0 {
					udp.SendOrderlist(data,i) // , udp.GetIndex(data.PrimaryQ[i], data))
				}
				handled = 1
				break 
			}
		}
		/*		
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ {
			if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == -1 && data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList[len(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)-1] == UpList[up] { 
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,UpList[up])
				fmt.Println("m")
				UpList = UpdateList(UpList,up)
				if i != 0 {
					udp.SendOrderlist(data) // , udp.GetIndex(data.PrimaryQ[i], data))
				}
				handled = 1
				break 		
			}	
		}
		*/
		/*		
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ {
			if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == -1 && data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList[len(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)-1] > UpList[up] {
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,UpList[up])
				fmt.Println("n")
				UpList = UpdateList(UpList,up)
				if i != 0 {
					udp.SendOrderlist(data) // , udp.GetIndex(data.PrimaryQ[i], data))
				}
				handled = 1
				break 
			}
		}
		*/
		fmt.Println("Her er handled: ",handled)		
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ {
			fmt.Println("RUNNING RUNNING RUNNING",data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running  )
			if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == 0 {
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,UpList[up])
				
				fmt.Println("control 473: heisen står stille. UpList:", UpList)
				
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running = 1
				fmt.Println("Sjekker GetIndex: ", udp.GetIndex(data.PrimaryQ[i], data))
				fmt.Println("Sjekker Statuses: ", len(data.Statuses))
				fmt.Println("Sjekker CurrentFloor: ", data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CurrentFloor)
				fmt.Println("Sjekker UpList: ", UpList)
				if UpList[up] > data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CurrentFloor{
					data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = functions.SortUp(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
					data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running = 1
				}else if UpList[up] < data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CurrentFloor{
					data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = functions.SortDown(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
					data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running = -1
				}
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].ButtonList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].ButtonList,0)
				if i != 0 {
					udp.SendOrderlist(data,i) // , udp.GetIndex(data.PrimaryQ[i], data))
				}
				UpList = functions.UpdateList(UpList,up)
				handled = 1
				break 
			}
		}
		/*
		if handled == 0 {
			data.Statuses[data.PrimaryQ[0]].OrderList = append(data.Statuses[data.PrimaryQ[0]].OrderList,UpList[up])
			UpList = UpdateList(UpList,up)
			handled = 1
			fmt.Println("p") 
		}*/
	}
	}
}









