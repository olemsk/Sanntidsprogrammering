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

func GoToFloor(floor int, data *udp.Data) { // Lamper for command buttons må leggas til, kall list noe annet
	//fmt.Println("control 82: går til floor floor:",floor)
	//fmt.Println("control 82: er i floor:",status.CurrentFloor)
	for { 

		driver.SetFloorIndicator(driver.GetFloorSensorSignal())
		if floor == driver.GetFloorSensorSignal() {
				
				driver.SetFloorIndicator(floor)
				
				driver.SetButtonLamp(2,floor,0)
				driver.SetMotorDirection(driver.DIRN_STOP)
				driver.SetDoorOpenLamp(true)				
				time.Sleep(1500*time.Millisecond)
				driver.SetDoorOpenLamp(false)
				fmt.Println("ButtonList??: ", data.ButtonList)
				data.Statuses[udp.GetIndex(udp.GetID(),data)].CurrentFloor = floor
				if((floor == 0 || floor == 3) || len(data.Statuses[udp.GetIndex(udp.GetID(), data)].OrderList) == 0){
					data.Statuses[udp.GetIndex(udp.GetID(),data)].Running = 0
					//driver.SetButtonLamp(0,floor,0)
					//driver.SetButtonLamp(1,floor,0)
					
					if(floor == 3){
						data.ButtonList[5] = 0	
					}else if(floor == 0){
						data.ButtonList[0] = 0
					}else{
						data.ButtonList[floor]=0
						data.ButtonList[floor+2]=0 
					}
				}else if(data.Statuses[udp.GetIndex(udp.GetID(),data)].Running == 1){
					//driver.SetButtonLamp(0, floor, 0)
					data.ButtonList[floor] = 0
				}else if(data.Statuses[udp.GetIndex(udp.GetID(),data)].Running == -1){
					//driver.SetButtonLamp(1,floor,0)
					data.ButtonList[floor+2] = 0
				}

				//if list == 1 {
				//	driver.SetButtonLamp((*status).ButtonList[0], floor, 0)
				//	(*status).ButtonList = functions.UpdateList((*status).ButtonList,0)
				//}
				fmt.Println("Heisen er framme på floor:", floor)
				udp.PrintData(*data)
				
				break
		} else if floor > driver.GetFloorSensorSignal() && driver.GetFloorSensorSignal() != -1 && floor != -1 {   
			driver.SetMotorDirection(driver.DIRN_UP)

			data.Statuses[udp.GetIndex(udp.GetID(),data)].Running = 1
			

		} else if floor < driver.GetFloorSensorSignal() && driver.GetFloorSensorSignal() != -1 && floor != -1{
			driver.SetMotorDirection(driver.DIRN_DOWN)

			data.Statuses[udp.GetIndex(udp.GetID(),data)].Running = -1
			

		}/*else if data.Statuses[udp.GetIndex(udp.GetID(),data)].CurrentFloor == -1{
			
			driver.SetMotorDirection(driver.DIRN_DOWN)
		}*/
		if driver.GetFloorSensorSignal() != -1{
			data.Statuses[udp.GetIndex(udp.GetID(),data)].CurrentFloor = driver.GetFloorSensorSignal()
			driver.SetButtonLamp(2,data.Statuses[udp.GetIndex(udp.GetID(),data)].CurrentFloor,0)
		}	
	}
}


func ElevatorControl(data *udp.Data) {
	//time.Sleep(1*time.Second)
	//var data.Statuses[udp.GetIndex(udp.GetID(),data)] *udp.data.Statuses[udp.GetIndex(udp.GetID(),data)]
	temp := 0
	//temp = temp + 0
	
	for {


		if len(data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList) > 0{
				fmt.Println("OrderList: ", data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList)
				// 
				if data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList[0] != data.Statuses[udp.GetIndex(udp.GetID(),data)].CurrentFloor  {
					//fmt.Println(data.data.Statuses[udp.GetIndex(udp.GetID(),data)]es[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
					// Sjekker om heisens ordreliste
					temp = data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList[0]
					data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList = functions.UpdateList(data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList,0)
					GoToFloor(temp,data) // vurdere å kjøre commandbuttons inni gotofloor
					temp = 0
				/*
				}else if data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList[0] < data.Statuses[udp.GetIndex(udp.GetID(),data)].CurrentFloor{
					temp = data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList[0]
					data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList = functions.UpdateList(data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList,0)
					GoToFloor(temp,data)
					temp = 0
				*/		
				}else if data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList[0] == driver.GetFloorSensorSignal() {
					data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList=functions.UpdateList(data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList,0)
					GoToFloor(driver.GetFloorSensorSignal(), data)
				}

		}	
	}
}
	
		
func GetDestination(data *udp.Data) { //returnerer bare button, orderlist oppdateres
	//time.Sleep(1*time.Second)
	for {
		time.Sleep(50*time.Millisecond) // Polling rate, mby change	
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
						fmt.Println("er")
						data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList = append(data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList, floor)
						data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList = functions.SortUp(data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList)
						driver.SetButtonLamp(2,floor,1)
					} else if data.Statuses[udp.GetIndex(udp.GetID(),data)].CurrentFloor > floor && data.Statuses[udp.GetIndex(udp.GetID(),data)].Running == -1{
						fmt.Println("er eg")
						data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList = append(data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList, floor)
						data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList = functions.SortDown(data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList)
						driver.SetButtonLamp(2,floor,1)
					} else if data.Statuses[udp.GetIndex(udp.GetID(),data)].CurrentFloor < floor && data.Statuses[udp.GetIndex(udp.GetID(),data)].Running == 0{
						fmt.Println("er eg her")
						data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList = append(data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList, floor)
						data.Statuses[udp.GetIndex(udp.GetID(),data)].Running = 1
						driver.SetButtonLamp(2,floor,1)
					} else if data.Statuses[udp.GetIndex(udp.GetID(),data)].CurrentFloor > floor && data.Statuses[udp.GetIndex(udp.GetID(),data)].Running == 0{
						fmt.Println("er eg her?")
						data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList = append(data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList, floor)
						data.Statuses[udp.GetIndex(udp.GetID(),data)].Running = -1
						driver.SetButtonLamp(2,floor,1)
					} else if data.Statuses[udp.GetIndex(udp.GetID(),data)].CurrentFloor == floor && data.Statuses[udp.GetIndex(udp.GetID(),data)].Running == 0{
						fmt.Println("er eg her!!!!!")
						data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList = append(data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList, floor)
						data.Statuses[udp.GetIndex(udp.GetID(),data)].Running = 0
						driver.SetButtonLamp(2,floor,1)
					}
				} else if driver.GetButtonSignal(2,floor) == 1 && functions.CheckList(data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList, floor) {
					driver.SetButtonLamp(2,floor,1)
				}
					
		}
	}
}
func CostFunction(in chan *udp.Data, out chan *udp.Data) {
	handled := 0
	var DownList []int
	var UpList []int
	var data *udp.Data
	data = <-in
	for {
		//fmt.Println("control 243, handled: ",handled)
		handled = 0
		//fmt.Println("status.UpList i CostFunction: ",(*data).Statuses[udp.GetIndex((*data).PrimaryQ[0], data)].UpList)
		//fmt.Println("Lengden til statuses: ", len(data.Statuses))
		fmt.Println("PrimaryQ: ", data.PrimaryQ)
		fmt.Println("Her er lengden til primaryQ",len(data.Statuses))

		if len(UpList) == 0 && len(DownList) == 0{
			if handled == 0{
				out<- data
				data = <-in
				//fmt.Println("control: 383. Er inne i costfunction")
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
				//fmt.Printf("PrimaryQ: %v Lengde Statuses: %v\n", data.PrimaryQ, len(data.Statuses))
				//fmt.Println("Har den samme info om status.Uplist her: ", data.Statuses[k].UpList)
				DownList = append(DownList,data.Statuses[k].DownList...)
				//data.Statuses[udp.GetIndex(data.PrimaryQ[k], data)].DownList = data.Statuses[udp.GetIndex(data.PrimaryQ[k], data)].DownList[:0]

				UpList = append(UpList,data.Statuses[k].UpList...)
				//data.Statuses[udp.GetIndex(data.PrimaryQ[k], data)].UpList = data.Statuses[udp.GetIndex(data.PrimaryQ[k], data)].UpList[:0]
			}
		}

	
	if len(UpList) > 0 {
		UpList = functions.SortUp(UpList)
	}
	if len(DownList) > 0 {
		DownList = functions.SortDown(DownList)
	}
	

	for down:=0; down<len(DownList);down++ { // Kanskje feil å sette running inni her? 
		
		if handled == 1{
			handled = 0
			down = 0	
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ { // Heis i samme floor, og står stille
			if DownList[down] == data.Statuses[i].CurrentFloor && data.Statuses[i].Running == 0 {
				data.ButtonList[2+DownList[down]] =1
				data.Statuses[i].OrderList = append(data.Statuses[i].OrderList,DownList[down])
				for j := 0; j < len(data.PrimaryQ); j++{
					for k := 0; k < len(data.Statuses[j].DownList); k++{
						if data.Statuses[j].DownList[k] == DownList[down]{
							data.Statuses[j].DownList = functions.UpdateList(data.Statuses[j].DownList, k)
							break
						}
					}	
				}
				DownList = functions.UpdateList(DownList,down) //Må modifiseres
				data.Statuses[i].OrderList = functions.SortDown(data.Statuses[i].OrderList)
				data.Statuses[i].Running = 0
				fmt.Println("control 280: Heis i samma floor og står stille. Downlist:", DownList)				
				
				udp.SendOrderlist(data,i)

				handled = 1
				break
				//pluss noe mer, som å åpne døra
			}
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ { // Heis i etasjen over og på veg nedover
			if data.Statuses[i].CurrentFloor == DownList[down]+1 && data.Statuses[i].Running == -1 && handled != 1 {
				data.ButtonList[2+DownList[down]] =1
				data.Statuses[i].OrderList = append(data.Statuses[i].OrderList,DownList[down])
				fmt.Println("control 370: Heis i etasjen over og på vei nedover. Downlist:", DownList)
				data.Statuses[i].OrderList = functions.SortDown(data.Statuses[i].OrderList)
				for j := 0; j < len(data.PrimaryQ); j++{
					for k := 0; k < len(data.Statuses[j].DownList); k++{
						if data.Statuses[j].DownList[k] == DownList[down]{
							data.Statuses[j].DownList = functions.UpdateList(data.Statuses[j].DownList, k)
							break
						}
					}	
				}
				DownList = functions.UpdateList(DownList,down)
				
				if i != 0 {
					udp.SendOrderlist(data, i) // , i)
				}
				handled = 1
				break 
			}
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ { // Heis i etasjen over og står stille
			if data.Statuses[i].CurrentFloor == DownList[down]+1 && data.Statuses[i].Running == 0 && handled != 1 {
				data.ButtonList[2+DownList[down]] =1
				data.Statuses[i].OrderList = append(data.Statuses[i].OrderList,DownList[down])
				data.Statuses[i].OrderList = functions.SortDown(data.Statuses[i].OrderList)
				
				data.Statuses[i].Running = -1
				fmt.Println("control 385: Heis i etasjen over og står stille")
				for j := 0; j < len(data.PrimaryQ); j++{
					for k := 0; k < len(data.Statuses[j].DownList); k++{
						if data.Statuses[j].DownList[k] == DownList[down]{
							data.Statuses[j].DownList = functions.UpdateList(data.Statuses[j].DownList, k)
							break
						}
					}	
				}
				DownList = functions.UpdateList(DownList,down)
				
				udp.SendOrderlist(data,i)
				handled = 1
				break 
			}
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ { // Heis på veg nedover
			if data.Statuses[i].CurrentFloor > DownList[down] && data.Statuses[i].Running == -1  && handled != 1{
				data.ButtonList[2+DownList[down]] =1
				data.Statuses[i].OrderList = append(data.Statuses[i].OrderList,DownList[down])
				for j := 0; j < len(data.PrimaryQ); j++{
					for k := 0; k < len(data.Statuses[j].DownList); k++{
						if data.Statuses[j].DownList[k] == DownList[down]{
							data.Statuses[j].DownList = functions.UpdateList(data.Statuses[j].DownList, k)
							break
						}
					}	
				}
				DownList = functions.UpdateList(DownList,down)
				fmt.Println("control 398: Heis på vei nedover og er over floor. Downlist:", DownList)
				data.Statuses[i].OrderList = functions.SortDown(data.Statuses[i].OrderList)
			
				
				udp.SendOrderlist(data,i)
				handled = 1
				break 
			}
		}

		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ {
			if data.Statuses[i].Running == 0 {
				data.Statuses[i].OrderList = append(data.Statuses[i].OrderList,DownList[down])
				fmt.Println("control 437: heis står stille generelt. Downlist:", DownList)
				data.ButtonList[2+DownList[down]] =1
				if DownList[down] > data.Statuses[i].CurrentFloor{
					data.Statuses[i].OrderList = functions.SortUp(data.Statuses[i].OrderList)
					data.Statuses[i].Running = 1
				}else if DownList[down] < data.Statuses[i].CurrentFloor{
					data.Statuses[i].OrderList = functions.SortDown(data.Statuses[i].OrderList)
					data.Statuses[i].Running = -1
				}
				
				for j := 0; j < len(data.PrimaryQ); j++{
					for k := 0; k < len(data.Statuses[j].DownList); k++{
						if data.Statuses[j].DownList[k] == DownList[down]{
							data.Statuses[j].DownList = functions.UpdateList(data.Statuses[j].DownList, k)
							break
						}
					}	
				}
				DownList = functions.UpdateList(DownList,down)
				udp.SendOrderlist(data,i)
				handled = 1
				break 
			}
		}


	}

for up:=0; up<len(UpList);up++ {
		//fmt.Println("Up: ",up)
		//fmt.Println(data.PrimaryQ)
		if handled == 1{
			handled = 0
			up = 0	
		}
		for i:=0;i<len(data.PrimaryQ) && handled==0;i++{
			if UpList[up] == data.Statuses[i].CurrentFloor && data.Statuses[i].Running == 0 {
				data.ButtonList[UpList[up]]=1
				data.Statuses[i].OrderList = append(data.Statuses[i].OrderList,UpList[up])
				data.Statuses[i].OrderList = functions.SortUp(data.Statuses[i].OrderList)
				data.Statuses[i].NextFloor = UpList[up]
				fmt.Println("control 387: heis i samme etasjen og står stille. UpList:", UpList)
				for j := 0; j < len(data.PrimaryQ); j++{
					for k := 0; k < len(data.Statuses[j].UpList); k++{
						if data.Statuses[j].UpList[k] == UpList[up]{
							data.Statuses[j].UpList = functions.UpdateList(data.Statuses[j].UpList, k)
							break
						}
					}	
				}
				UpList = functions.UpdateList(UpList,up)
				
				udp.SendOrderlist(data,i)
				handled = 1
				break 
				//pluss noe mer, som å åpne døra
			}
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ {
			if data.Statuses[i].CurrentFloor == UpList[up]-1 && data.Statuses[i].Running == 1 && handled != 1 {
				data.ButtonList[UpList[up]]=1
				data.Statuses[i].OrderList = append(data.Statuses[i].OrderList,UpList[up])
				data.Statuses[i].OrderList = functions.SortUp(data.Statuses[i].OrderList)
				data.Statuses[i].NextFloor = UpList[up]
				fmt.Println("control 402: heis i etasjen under og på vei oppover. UpList:", UpList)
				for j := 0; j < len(data.PrimaryQ); j++{
					for k := 0; k < len(data.Statuses[j].UpList); k++{
						if data.Statuses[j].UpList[k] == UpList[up]{
							data.Statuses[j].UpList = functions.UpdateList(data.Statuses[j].UpList, k)
							break
						}
					}	
				}
				UpList = functions.UpdateList(UpList,up)
				//data.Statuses[i].UpList = functions.UpdateList(data.Statuses[i].UpList, up)
				udp.SendOrderlist(data,i)
				handled = 1
				break 
			}
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ {
			if data.Statuses[i].CurrentFloor == UpList[up]-1 && data.Statuses[i].Running == 0 && handled != 1 {
				data.ButtonList[UpList[up]]=1
				data.Statuses[i].OrderList = append(data.Statuses[i].OrderList,UpList[up])
				data.Statuses[i].OrderList = functions.SortUp(data.Statuses[i].OrderList)
				data.Statuses[i].NextFloor = UpList[up]
				data.Statuses[i].Running = 1
				fmt.Println("control 417: heis i etasjen under og står stille. UpList:", UpList)
				for j := 0; j < len(data.PrimaryQ); j++{
					for k := 0; k < len(data.Statuses[j].UpList); k++{
						if data.Statuses[j].UpList[k] == UpList[up]{
							data.Statuses[j].UpList = functions.UpdateList(data.Statuses[j].UpList, k)
							break
						}
					}	
				}
				UpList = functions.UpdateList(UpList,up)				
				udp.SendOrderlist(data,i)
				handled = 1
				break 
			}
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ {
			if data.Statuses[i].CurrentFloor < UpList[up] && data.Statuses[i].Running == 1  && handled != 1{
				data.ButtonList[UpList[up]]=1
				data.Statuses[i].OrderList = append(data.Statuses[i].OrderList,UpList[up])
				data.Statuses[i].OrderList = functions.SortUp(data.Statuses[i].OrderList)
				fmt.Println("control 430: floor over heis.currentfloor og på vei oppover. UpList:", UpList)
				data.Statuses[i].NextFloor = UpList[up]
				for j := 0; j < len(data.PrimaryQ); j++{
					for k := 0; k < len(data.Statuses[j].UpList); k++{
						if data.Statuses[j].UpList[k] == UpList[up]{
							data.Statuses[j].UpList = functions.UpdateList(data.Statuses[j].UpList, k)
							break
						}
					}	
				}
				UpList = functions.UpdateList(UpList,up)
				
				udp.SendOrderlist(data,i)
				handled = 1
				break 
			}
		}

		fmt.Println("Her er handled: ",handled)		
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ {
			fmt.Println("RUNNING RUNNING RUNNING",data.Statuses[i].Running  )
			if data.Statuses[i].Running == 0 {
				data.Statuses[i].OrderList = append(data.Statuses[i].OrderList,UpList[up])
				
				fmt.Println("control 473: heisen står stille. UpList:", UpList)
				data.ButtonList[UpList[up]]=1
				if UpList[up] > data.Statuses[i].CurrentFloor{
					data.Statuses[i].OrderList = functions.SortUp(data.Statuses[i].OrderList)
					data.Statuses[i].Running = 1
				}else if UpList[up] < data.Statuses[i].CurrentFloor{
					data.Statuses[i].OrderList = functions.SortDown(data.Statuses[i].OrderList)
					data.Statuses[i].Running = -1
				}
				for j := 0; j < len(data.PrimaryQ); j++{
					for k := 0; k < len(data.Statuses[j].UpList); k++{
						if data.Statuses[j].UpList[k] == UpList[up]{
							data.Statuses[j].UpList = functions.UpdateList(data.Statuses[j].UpList, k)
							break
						}
					}	
				}
				UpList = functions.UpdateList(UpList,up)
				
				udp.SendOrderlist(data,i)
				
				handled = 1
				break 
			}
		}

	}
	}
}
func LampControl(data *udp.Data){
	for{
		for i:=0;i<3;i++{
			if(data.ButtonList[i] == 0){
				driver.SetButtonLamp(0,i,0)
			}else if(data.ButtonList[i] == 1){
				driver.SetButtonLamp(0,i,1)
			}
			if(data.ButtonList[i+3] == 0){
				driver.SetButtonLamp(1,i+1,0)
			}else if(data.ButtonList[i+3] == 1){
				driver.SetButtonLamp(1,i+1,1)
			}
		}
	}
}







