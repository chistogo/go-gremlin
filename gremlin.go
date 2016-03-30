//The goal of this process is to create a buffered reader and Writer to the network and drop a percentage of the buffered

package main

import (
    "fmt"
    "log"
    "net"
    "os/exec"
    "os"
    
    "bufio"
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
    
    file, err := os.Create("newfile.txt")
    checkError(err)
    
    //We can change the buffer size if needed
    reader := bufio.NewReader(conn)
    writer := bufio.NewWriter(file)
        
    fmt.Println("Connection:"+conn.RemoteAddr().String())
    
    //  Newfile<==OLD
    
    buffer := make([]byte, 32)
    
    numberOfBytesReadIn , err := reader.Read(buffer)
    checkError(err)
    _, err = writer.Write(buffer)
    checkError(err)
    writer.Flush()
    
    for(numberOfBytesReadIn != 0){
        numberOfBytesReadIn , err = reader.Read(buffer)
        checkError(err)
        _, err = writer.Write(buffer)
        checkError(err)
        
    
    }
    writer.Flush()
    
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