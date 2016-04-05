/*********************************************************************
*        ʕ◔ϖ◔ʔ               Server.go              ʕ◔ϖ◔ʔ           *
**********************************************************************
* CS 381 - Intro to Computer Networks    Project 3                   *
* April 4th, 2016                                                    *
**********************************************************************
* Authors: Taylor Atkinson, Christopher Goulet                       *
**********************************************************************
* Input: This program asks the user to input the file name and       *
* extension of the output file.
**********************************************************************
* Output: This program outputs when the server reads something from  *
* the Grimlin, It displays the amount recieved and the current       *
* sequence Number                                                    *
*********************************************************************/

package main

import (
    "fmt"
    "net"
    "os"
    "os/exec"
    "log"
)
 
func main() {
    clear()
    fmt.Println("ʕ◔ϖ◔ʔ  Welcome to the GO Gremlin, Receiving Process!!!  ʕ◔ϖ◔ʔ")
    
    //Receive from sender port
    
    clientAddr , err := net.ResolveUDPAddr("udp","127.0.0.1:8008")
    checkError(err)
    
    gremlinConn, err := net.ListenUDP("udp", clientAddr)
    checkError(err)
    
    
    var fileName string
    fmt.Print("Please enter your filename with extension: ")
    _, err = fmt.Scanf("%s\n", &fileName)
    checkError(err)
    
    testFile , err := os.Create(fileName)
    checkError(err)
    
    seqBuff := make([]byte, 1)
    buf := make([]byte, 33)
    var currentSequenceNum byte
    currentSequenceNum = 0
    var oldSequenceNum byte 
    oldSequenceNum = 0 
    
    //rand.Seed(time.Now().Unix())
    
    
    //fail := 95
    for{
        //Read from Client
        fmt.Println("Reading from Gremlin...")
        numOfBytesReceived, addressFrom , err :=    gremlinConn.ReadFromUDP(buf)
        fmt.Printf("Received %d bytes\n", numOfBytesReceived)
        checkError(err)
        //The sequenceNum on the packet matches the anticipated sequenceNum
        fmt.Printf("Current SeqNum:%d\n",buf[numOfBytesReceived -1])
        if(currentSequenceNum == buf[numOfBytesReceived -1]) {
            _, err = testFile.Write(buf[:numOfBytesReceived-1])
            seqBuff[0] = currentSequenceNum 
            
            
            gremlinConn.WriteToUDP(seqBuff,addressFrom)
            checkError(err)

            oldSequenceNum = currentSequenceNum
            if(currentSequenceNum == 255){
                currentSequenceNum = 0
            }else{
                currentSequenceNum++
            }       
        //Client didn't get AWK of the old packet
        } else if (buf[numOfBytesReceived -1] == oldSequenceNum){
            fmt.Println("Sequence Number Does Not Match")
            seqBuff[0] = oldSequenceNum
            

                gremlinConn.WriteToUDP(seqBuff,addressFrom)

            
            
        }else {
            fmt.Println("Sequence Number Does Not Match and is not old sequenceNum")
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
    
