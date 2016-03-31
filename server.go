package main

import (
    "fmt"
    "net"
    "os"
    "io/ioutil"
    "os/exec"
    "log"
)
 
func main() {
    clear()
    fmt.Println("ʕ◔ϖ◔ʔ  Welcome to the GO Gremlin, Receiving Process!!!  ʕ◔ϖ◔ʔ")
    
    recivingPort := ":8887"
    
    ln, err := net.Listen("tcp",recivingPort)
    checkError(err)
    
    for {
        conn, err := ln.Accept()
        checkError(err)
        
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    
    file, err := os.Create("newfile.png") //Creates the file to store it
    fmt.Println("pls work")
    reader, err := ioutil.ReadAll(conn) 
    checkError(err)
    fmt.Println("pls work1")
        
    fmt.Println("Connection:"+conn.RemoteAddr().String())
    
    //  Newfile<==OLD
    
    for i := 0;  i < len(reader); i++ { //while i is less than the length of the reader
        if(i % 32 ==  0 && i != 0){     //If 
               file.Write(reader[i-32:i])
        }
        if(i % 32 ==  0 && i + 32 > len(reader)){
            file.Write(reader[i:len(reader)])
            
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
    
