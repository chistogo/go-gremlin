package main


import (
    "fmt"
    "log"
    "net"
    "os/exec"
    "os"
    "io"
    // "bufio"
)


func main() {
    
    gremlinIP := 127.0.0.1
    gremlinPort := 8888
    
    clear()
    fmt.Println("ʕ◔ϖ◔ʔ  Welcome to the GO Gremlin Client Process!!!  ʕ◔ϖ◔ʔ")
    
    //Open File
    file, err := os.Open("file.txt")
    checkError(err)
    
    //File to Write to
    newFile , err := os.Create("newfile.html")
    checkError(err)
    
    //Copy over the data. DataReceive <-- DataSend
    io.Copy(newFile,file)
    
    //Connect to Gremlin
    conn, err := net.Dial("tcp", "127.0.0.1:8888")
    checkError(err)
    //Sends to file to Gremlin    
    status, err := io.Copy(newFile,conn)
    checkError(err)
    fmt.Println(status)
    
   
}

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