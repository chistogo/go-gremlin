package main


import (
    "fmt"
     "log"
    // "net"
    "os/exec"
    "os"
    "io"
    // "bufio"
)


func main() {
    clear()
    fmt.Println("ʕ◔ϖ◔ʔ  Welcome to the GO Gremlin Client Process!!!  ʕ◔ϖ◔ʔ")
    
    file, err := os.Open("file.txt")
    checkError(err)
    
    newFile , err := os.Create("newfile.txt")
    checkError(err)
    
    io.Copy(newFile,file)
   
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