/*********************************************************************
*        ʕ◔ϖ◔ʔ               Client.go             ʕ◔ϖ◔ʔ            *
**********************************************************************
* CS 381 - Intro to Computer Networks    Project 3                   *
* April 4th, 2016                                                    *
**********************************************************************
* Authors: Taylor Atkinson, Christopher Goulet                       *
**********************************************************************
* Input: This program asks the user to input file name into the      *
* command line including the file extention.                         *
**********************************************************************
* Output: This program outputs the sequence Number, Payload size and *
* when it is incrementing to the next packet.                        *
*********************************************************************/

package main


import (
    "fmt"
    "log"
    "net"
    "os/exec"
    "os"
    "io"
    "time"
)


func main() {
    
    
    clear()
    fmt.Println("ʕ◔ϖ◔ʔ  Welcome to the GO Gremlin Client Process!!!  ʕ◔ϖ◔ʔ")
    
    //Open File
    var fileName string
    fmt.Print("Please enter your filename with extension: ")
    _, err := fmt.Scanf("%s\n", &fileName)
    checkError(err)
    file, err := os.Open(fileName)

    
    checkError(err)   
   

    
    //Create Sending Object
    serverAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8001") //Address to Send to
    checkError(err)
    localAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0") //LocalAddress
    checkError(err)
    
    //Create the Connection for Sending 
    SendingConn, err := net.ListenUDP("udp", localAddr)
    checkError(err)
    
    
    
    fmt.Println("Sending File ...")
    //Sends to file to Gremlin
    //Copy over the data. DataReceive <-- DataSend    
    //Making 32 byte buffer
    
    
    
    var offSet int64
    var sequenceNum byte;
    sequenceNum = 0
    seqBuff := make([]byte, 1)
    offSet = 0
    
    for{
        buff := make([]byte , 32)
        //Read from file at off set and load it into our buffer
        numberOfBytes, err := file.ReadAt(buff, offSet) // Read file with our buffer
        fmt.Printf("Sending Payload: %d  Seqence Number:%d \n",offSet,sequenceNum)
        SendingConn.WriteToUDP(append(buff[:numberOfBytes], sequenceNum),serverAddr) // Places sequenceNum at the end
        
       if (err == io.EOF) {
           break;
       }else{
           checkError(err)
       }
       //Time for timeout
       
       t:=time.Now()
       SendingConn.SetDeadline(t.Add(time.Millisecond+100000000 ))
       //Recieves ACK packet
       
       
       fmt.Println("Reading from UDP...")
       _,_,err = SendingConn.ReadFromUDP(seqBuff)
       //If time out occures
       
       if(err != nil && err.(net.Error).Timeout() || seqBuff[0] != sequenceNum){
           SendingConn.Close()
           SendingConn, err = net.ListenUDP("udp", localAddr)
            //SendingConn.WriteToUDP(append(buff[:numberOfBytes], sequenceNum), serverAddr)
            fmt.Println("Resending")
       }else{
       checkError(err)
       
       
       offSet = offSet + 32
       
       if(sequenceNum == 255){ //If the sequenceNum is at its max value restart
            sequenceNum = 0
        } else {
            fmt.Println("I'm incrementing")
            sequenceNum = sequenceNum + 1
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