package main
//The goal of this process is to create a buffered reader and Writer to the network and drop a percentage of the buffered

import (
    "fmt"
    "log"
    "net" 
    "os/exec"
    //"os"
    // "io"
    "io/ioutil"
    // "bytes"
)


func main(){
    clear()
    fmt.Println("ʕ◔ϖ◔ʔ  Welcome to the GO Gremlin, Gremlin Process!!!  ʕ◔ϖ◔ʔ")
    
    gremlinPort := "8888"
    
    ln, err := net.Listen("tcp", ":"+gremlinPort)
    
    if err != nil {
        // handle error
    }
    
    for {
        conn, err := ln.Accept()
        if err != nil {
            // handle error
        }
        go handleConnection(conn)
    }

}


//TODO: Make sure this is READ and WRITTEN using a buffer
func handleConnection(conn net.Conn){
    
    receiverIP := "127.0.0.1"
    receiverPort := "8887"
    connToReceive, err := net.Dial("tcp", receiverIP+":"+receiverPort)
    checkError(err)
    fmt.Println("Sending File ...")
    
    //file, err := os.Create("newfile.png")
    
    
    //We can change the buffer size if needed
   // buff := make([]byte, 32) //Creates a buffer for 64 Bytes
    reader, err := ioutil.ReadAll(conn)
    checkError(err)
    //writer := 
        
    fmt.Println("Connection:"+conn.RemoteAddr().String())
    
    //  Newfile<==OLD
    //conn.Write
    for i := 0;  i < len(reader); i++ { //while i is less than the length of the reader
        if(i % 32 ==  0 && i != 0){     //If 
               connToReceive.Write(reader[i-32:i])
        }
        if(i % 32 ==  0 && i + 32 > len(reader)){
            connToReceive.Write(reader[i:len(reader)])
            connToReceive.Close()
        }
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