package main
//The goal of this process is to create a buffered reader and Writer to the network and drop a percentage of the buffered

import (
    "fmt"
    "log"
    "net" 
    "os/exec"
    //"os"
    // "io"
   // "io/ioutil"
    // "bytes"
    "os"
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
    
    testFile , err := os.Create("dump.png")
    checkError(err)
    
    seqBuff := make([]byte, 1)
    buf := make([]byte, 33)
    var currentSequenceNum byte
    currentSequenceNum = 0
    var oldSequenceNum byte 
    oldSequenceNum = 0 
    
    rand.Seed(time.Now().Unix())
    fail := 95
    for{
        //Read from Client
        fmt.Println("Reading from client...")
        numOfBytesReceived, addressFrom , err := clientConn.ReadFromUDP(buf)
        fmt.Printf("Received %d bytes\n", numOfBytesReceived)
        checkError(err)
        //The sequenceNum on the packet matches the anticipated sequenceNum
        fmt.Printf("Current SeqNum:%d\n",buf[numOfBytesReceived -1])
        if(currentSequenceNum == buf[numOfBytesReceived -1]) {
            _, err = testFile.Write(buf[:numOfBytesReceived-1])
            seqBuff[0] = currentSequenceNum 
            
            
            
            if(rand.Intn(100)>fail){
                clientConn.WriteToUDP(seqBuff,addressFrom)
            }
            
            
            
            checkError(err)
            oldSequenceNum = currentSequenceNum
            if(currentSequenceNum == 255){
                currentSequenceNum = 0
            }else{
                currentSequenceNum++
                fmt.Println(currentSequenceNum)
            }       
        //Client didn't get AWK of the old packet
        } else if (buf[numOfBytesReceived -1] == oldSequenceNum){
            fmt.Println("Sequence Number Does Not Match")
            seqBuff[0] = oldSequenceNum
            
            if(rand.Intn(100)>fail){
                clientConn.WriteToUDP(seqBuff,addressFrom)
            }
            
            
        }else {
            fmt.Println("Sequence Number Does Not Match and is not old sequenceNum")
        }
        
        
        
        checkError(err)
        
    }

}


//TODO: Make sure this is READ and WRITTEN using a buffer
// func handleConnection(conn net.Conn){
    
//     receiverIP := "127.0.0.1"
//     receiverPort := "8887"
//     connToReceive, err := net.Dial("tcp", receiverIP+":"+receiverPort)
//     checkError(err)
//     fmt.Println("Sending File ...")
    
//     //file, err := os.Create("newfile.png")
    
    
//     //We can change the buffer size if needed
//    // buff := make([]byte, 32) //Creates a buffer for 64 Bytes
//     reader, err := ioutil.ReadAll(conn)
//     checkError(err)
//     //writer := 
        
//     fmt.Println("Connection:"+conn.RemoteAddr().String())
    
//     //  Newfile<==OLD
//     //conn.Write
// //     for i := 0;  i < len(reader); i++ { //while i is less than the length of the reader
// //         if(i % 32 ==  0 && i != 0){     //If 
// //                connToReceive.Write(reader[i-32:i])
// //         }
// //         if(i % 32 ==  0 && i + 32 > len(reader)){
// //             connToReceive.Write(reader[i:len(reader)])
// //             connToReceive.Close()
// //         }
// //     }
    
// // }



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