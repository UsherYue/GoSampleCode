package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func testFileLog() {
	f, _ := os.OpenFile("log.txt", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0777)
	defer f.Close()
	fileLogger := log.New(f, "Filelog ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile)
	fileLogger.Println("File log message1.....")
	fileLogger.Println("File log message2.....")
	fileLogger.Println("File log message2.....")
	///fileLogger.Panic(22222333334445555)
	//fileLogger.Fatalln("file error log message......")
}

func testBufferLog() {
	var buf bytes.Buffer
	bufLogger := log.New(&buf, "Buflog ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile)
	bufLogger.Println("Buf log message1.....")
	bufLogger.Println("Buf log message2.....")
	bufLogger.Println("Buf log message2.....")
	//bufLogger.Fatalln("buf error log message......")
	fmt.Println(string(buf.Bytes()))
}

func main() {
	testFileLog()
	testBufferLog()
}
