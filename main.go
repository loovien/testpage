package main

import (
	"github.com/vvotm/testpage"
	"log"
	"os"
	"time"
)

var logger *log.Logger

func init() {
	panic("show stack")
	nowString := time.Now().Format("2006-01-02 15:04:05")
	logger = log.New(os.Stdout, "[DEBUG]["+nowString+"]", 0)
}

func main() {
	logger.Println("----------HttpSVR START------------")
	testpage.RunHttpSvr()
	logger.Println("----------HttpSVR END------------")
}
