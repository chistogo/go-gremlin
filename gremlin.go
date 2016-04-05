/*********************************************************************
*        ʕ◔ϖ◔ʔ               Gremlin.go             ʕ◔ϖ◔ʔ           *
**********************************************************************
* CS 381 - Intro to Computer Networks    Project 3                   *
* April 4th, 2016                                                    *
**********************************************************************
* Authors: Taylor Atkinson, Christopher Goulet                       *
**********************************************************************
* Input: This program asks the user to input the drop rate into the  *
* command line.                                                      *
**********************************************************************
* Output: This program outputs when the Grimlin reads something from *
* the client, It displays the amount recieved, the drop rate, and if *
* it is dropping a data packet or an AWK packet.                     *
*********************************************************************/

package main

import (
    "fmt"
    "log"
    "net" 
    "os/exec"
    "math/rand"
    "time"
)


func main(){
    clear()
    fmt.Println("ʕ◔ϖ◔ʔ  Welcome to the GO Gremlin, Gremlin Process!!!  ʕ◔ϖ◔ʔ")
    
    //Receive from sender port
    
    clientAddr , err := net.ResolveUDPAddr("udp","127.0.0.1:8001")
    checkError(err)
    
    clientConn, err := net.ListenUDP("udp", clientAddr)
    checkError(err)
    
    //Create a Dynamic Port to send to the server.go process
    localAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0") //LocalAddress
    checkError(err)
    
    //Creates the Address to send to the Server.go process
    serverAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8008") //Address to Send to
    checkError(err)
    serverConn, err := net.ListenUDP("udp", localAddr)
    checkError(err)
    
    buf := make([]byte, 33)

    
    rand.Seed(time.Now().Unix())
    var fail int
    fmt.Print("Please enter your drop rate: ")
    _, err = fmt.Scanf("%d\n", &fail)
    checkError(err)
    
    for{
        //Read from Client
        fmt.Printf("Reading from client... Drop Rate is: %d%c\n", fail,'%')
        numOfBytesReceived, receivedAddr , err := clientConn.ReadFromUDP(buf)
        fmt.Printf("Received %d bytes\n", numOfBytesReceived)
        checkError(err)
        //The sequenceNum on the packet matches the anticipated sequenceNum
            if(numOfBytesReceived>=1){
                if (rand.Intn(100)+1 >= fail){
                    serverConn.WriteToUDP(buf[:numOfBytesReceived],serverAddr)
                    numOfBytesReceived, _ , err = serverConn.ReadFromUDP(buf)
                    
                    checkError(err)
                        if (rand.Intn(100)+1 >= fail) {
                            serverConn.WriteToUDP(buf[:numOfBytesReceived],receivedAddr)
                        } else {
                            fmt.Println("Dropping AWK Packet")
                        }
                
                } else {
                    fmt.Println("Dropping Data Packet")
                }
                numOfBytesReceived = 0
            }
            
            
            checkError(err)

        
        
        
    }

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