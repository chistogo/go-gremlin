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
    clear()
    fmt.Println("ʕ◔ϖ◔ʔ  Welcome to the GO Gremlin Client Process!!!  ʕ◔ϖ◔ʔ")
    
    //Open File
    file, err := os.Open("file.txt")
    checkError(err)
    
    //File to Write to
    newFile , err := os.Create("newfile.txt")
    checkError(err)
    
    //Copy over the data. DataReceive <-- DataSend
    io.Copy(newFile,file)
    
    //Connect to Gremlin
    conn, err := net.Dial("tcp", "google.com:80")
    checkError(err)
    fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
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