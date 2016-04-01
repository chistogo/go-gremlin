package main


import (
    "fmt"
    "log"
    "net"
    "os/exec"
    "os"
    "io"
    "time"
    // "bufio"
)


func main() {
    
    // gremlinIP := "127.0.0.1"
    // gremlinPort := "8002"
    
    clear()
    fmt.Println("ʕ◔ϖ◔ʔ  Welcome to the GO Gremlin Client Process!!!  ʕ◔ϖ◔ʔ")
    
    //Open File
    file, err := os.Open("go.png")

    
    checkError(err)   
   
    //Listen for Grimlin for dropped Packets
    // GremlinAddr, err := net.ResolveUDPAddr("udp",":9001") //Creates the Grimlim Address Object for listening
    // checkError(err)
    // ListeningConn, err := net.ListenUDP("udp", GremlinAddr) //Creates the connection for listening from the Grimlim
    // checkError(err)
    
    //Create Sending Object
    serverAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8001") //Address to Send to
    checkError(err)
    localAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0") //LocalAddress
    checkError(err)
    
    //Create the Connection for Sending 
    SendingConn, err := net.DialUDP("udp", localAddr, serverAddr)
    checkError(err)
    //defer SendingConn.Close()
    
    
    
    fmt.Println("Sending File ...")
    //Sends to file to Gremlin
    //Copy over the data. DataReceive <-- DataSend    
    //Making 32 byte buffer
    buff := make([]byte , 32)
    
    
    var offSet int64
    var sequenceNum byte;
    sequenceNum = 0
    seqBuff := make([]byte, 1)
    offSet = 0
    
    for{
             
        numberOfBytes, err := file.ReadAt(buff, offSet) // Read file with our buffer
        SendingConn.Write(append(buff[:numberOfBytes], sequenceNum)) // Places sequenceNum at the end
        
       if (err == io.EOF) {
           break;
       }else{
           checkError(err)
       }
       //Time for timeout
       t:=time.Now()
       SendingConn.SetDeadline(t.Add(time.Second))
       //Recieves ACK packet
       _,_,err = SendingConn.ReadFromUDP(seqBuff)
       //If time out occures
       
       if(err != nil && err.(net.Error).Timeout()){
            SendingConn.Write(append(buff[:numberOfBytes], sequenceNum))
            fmt.Println("Resending")
       }else{
       checkError(err)
       
       
       offSet = offSet + 32
       
       if(sequenceNum == 255){ //If the sequenceNum is at its max value restart
            sequenceNum = 0
        } else {
            sequenceNum++
        }
       }//end of else
        
       
    }
  
    checkError(err)
    
    fmt.Println("Done Sending")
   
}

//General Error Catching 
func checkError(err error)  {
    if err != nil {
        log.Fatal(err)
    }
}

func clear(){

	cmd := exec.Command("clear")
	stdout, err := cmd.Output()


	if err != nil {
		println(err.Error())
		return
	}

	print(string(stdout))

}