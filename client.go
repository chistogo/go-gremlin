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
    
    gremlinIP := "127.0.0.1"
    gremlinPort := "8888"
    
    clear()
    fmt.Println("ʕ◔ϖ◔ʔ  Welcome to the GO Gremlin Client Process!!!  ʕ◔ϖ◔ʔ")
    
    //Open File
    file, err := os.Open("go.png")
    testFile, _ := os.Create("testfile.png.")
    
    
    //Making 32 byte buffer
    buff := make([]byte , 32)
    
    
    var offSet int64
    offSet = 0
    for{
        _, err := file.ReadAt(buff, offSet)
        testFile.Write(buff)
       if (err == io.EOF) {
           break;
       }else{
           checkError(err)
       }
       offSet = offSet + 32
    }
    
    
    
    
    
    
    checkError(err)   
   
    //Connect to Gremlin
    conn, err := net.Dial("tcp", gremlinIP+":"+gremlinPort)
    checkError(err)
    fmt.Println("Sending File ...")
    //Sends to file to Gremlin
    //Copy over the data. DataReceive <-- DataSend    
    status, err := io.Copy(conn,file)
    checkError(err)
    fmt.Println(status)
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