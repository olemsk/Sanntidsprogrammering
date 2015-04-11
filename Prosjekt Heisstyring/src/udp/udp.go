// go run networkUDP.go
package udp
import (. "fmt" // Using '.' to avoid prefixing functions with their package names
		// This is probably not a good idea for large projects...
	//"runtime"
	"time"
	"net"
	//"bufio"
	"os"
	"strconv"
	"driver"
	//"sort"
	"encoding/json"
)


type Status struct {
	Running int
	CurrentFloor int
	NextFloor int
	Primary bool
	ID int
}

type Data struct {
	//Running [3]int
	//CurrentFloor [3]int
	//NextFloor [3]int
	PrimaryQ [3]int
	
}


func SetStatus(Status *Status, running int, NextFloor int) {
	(*Status).Running = running
	(*Status).CurrentFloor = driver.GetFloorSensorSignal()
	(*Status).NextFloor = NextFloor
	
	(*Status).ID = GetID()
	//Println(" id i func:", (*Status).ID)
}
func GetID() int {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}
 	var ipAddr string
	for _, a := range addrs {
		
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipAddr = ipnet.IP.String()
			}
		}
	} 
	ut,_ := strconv.Atoi(ipAddr[12:15])
	return ut
}


func PrimaryListen(bconn *net.UDPConn, Data *Data) {
	buffer := make([]byte, 1024)
	//udpAddr, err := net.ResolveUDPAddr("udp", ":39998")
	//conn, err := net.ListenUDP("udp", udpAddr)
	//checkError(err)
	for {
		Println("Hører")
		bconn.SetReadDeadline(time.Now().Add(5*time.Second))		
		n, err := bconn.Read(buffer)
		checkError(err)
		//Data = buffer
		err = json.Unmarshal(buffer[0:n], Data)		
		Println("her er primaryQen:", Data.PrimaryQ)		
		// Printf("Rcv %d bytes: %s\n",n, buffer)
		
		
	}	
}


func PrimaryBroadcast(baddr *net.UDPAddr, Data *Data) { // data []byte
	//udpAddr, err := net.ResolveUDPAddr("udp", "129.241.187.255:39998")
	//checkError(err)
	bconn, err := net.DialUDP("udp", nil, baddr)
	checkError(err)
	for {
		Println("SENDER")
		//buffer = nil
		time.Sleep(2500*time.Millisecond)
		
		// WRITE
		//Println("Er du der server??")
		b,_ := json.Marshal(Data)
		bconn.Write(b) // \x00
		
		checkError(err)
	}

}

//func SlaveUpdate(

// send_ch, receive_ch chan Udp_message
func UdpInit(localListenPort int, broadcastListenPort int, message_size int, Status *Status, PrimaryQ *[]string, Data *Data) (err error) {
	buffer := make([]byte, message_size)
	

	(*Status).Primary = false
	SetStatus(Status,0,driver.GetFloorSensorSignal())	
	//InitStatus(*Status)
	//Println("SE HER::::: ", (Status).ID)
	
	//Generating broadcast address
	baddr, err = net.ResolveUDPAddr("udp4", "129.241.187.255:"+strconv.Itoa(broadcastListenPort))
	if err != nil {
		return err
	}

	//Generating localaddress
	tempConn, err := net.DialUDP("udp4", nil, baddr)
	defer tempConn.Close()
	tempAddr := tempConn.LocalAddr()
	laddr, err = net.ResolveUDPAddr("udp4", tempAddr.String())
	laddr.Port = localListenPort

	//Creating local listening connections
	localListenConn, err := net.ListenUDP("udp4", laddr)
	if err != nil {
		return err
	}

	//Creating listener on broadcast connection
	broadcastListenConn, err := net.ListenUDP("udp", baddr)
	if err != nil {
		localListenConn.Close()
		return err
	}

	//go udp_receive_server(localListenConn, broadcastListenConn, message_size receive_ch)
	//go udp_transmit_server(localListenConn, broadcastListenConn ,send_ch)

	//Setting first primary
	broadcastListenConn.SetReadDeadline(time.Now().Add(3*time.Second))
	_, err = broadcastListenConn.Read(buffer)
	if err != nil {
		Println("Tar over som primary!")
		(*Status).Primary = true
		*Data.PrimaryQ = append(*PrimaryQ, strconv.Itoa(Status.ID)) 
		go PrimaryBroadcast(baddr,Data)	
	} else {
		//*PrimaryQ = append(*PrimaryQ, string(buffer))
		go PrimaryListen(broadcastListenConn, Data)
	}
	


	//	fmt.Printf("Generating local address: \t Network(): %s \t String(): %s \n", laddr.Network(), laddr.String())
	//	fmt.Printf("Generating broadcast address: \t Network(): %s \t String(): %s \n", baddr.Network(), baddr.String())
	return err
}

/*
func SendCommandList() { // Bare sende siste tal for simplicity
	udpAddr, err := net.ResolveUDPAddr("udp", "129.241.187.255:30169") // Broadcast (endre ip nettverket du sitter på)
	checkError(err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)
	currentStruct := TellerStruct{teller}

	for {
		b,_ := json.Marshal(currentStruct)
		conn.Write(b)	
		Println("Sent: ",currentStruct.Teller) 		
		currentStruct.Teller = currentStruct.Teller + 1
		time.Sleep(1*time.Second)
	}
}*/





/*
func SendCommand(floorChan chan int) {
	udpAddr, err := net.ResolveUDPAddr("udp", "129.241.187.255:30169") // Broadcast (endre ip nettverket du sitter på)
	checkError(err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)
	currentStruct := TellerStruct{teller}

	for {
		b,_ := json.Marshal(currentStruct)
		conn.Write(b)	
		Println("Sent: ",currentStruct.Teller) 		
		currentStruct.Teller = currentStruct.Teller + 1
		time.Sleep(1*time.Second)
	}

}*/

func checkError(err error) {
	if err != nil {
		Println("Noe gikk galt %v", err)
		return
	}
}












