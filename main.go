package main 

import (
    "os"
    "fmt"
    "os/signal"
    "strconv"
    "time"
    "github.com/KenmyZhang/experts-spider/app"
    l4g "github.com/alecthomas/log4go"
)

var URL = "http://ypk.39.net/AllCategory"

func main() {
    l4g.AddFilter("stdout", l4g.ERROR, l4g.NewConsoleLogWriter())
    cleanupDone := make(chan bool)

   // numStrChan := make(chan string, 1000)
    //go app.GetExpertId(numStrChan)
    //go app.GetExpertDetail(numStrChan)

    //go app.GetAllExpertId()

    go app.GetExpertFromGuaHao()

    Stop(cleanupDone)

}

func rangeDrugNum(drugNumChan chan string) {
    /*
    for medicine_manuals
    //for i := 500000; i <= 900000; i ++ {
    //for i := 500000; i >= 0; i-- {
    //  for i := 1000000000; i <= 1000100000; i++ {  
    */

    /*for      product
    //for i := 0; i <= 600000; i++ { 
    //for i := 600000; i <= 1229408 ; i++ { 
    */    
    //for i := 0; i <= 341815; i++ { 
    for i := 157772; i <= 1000000; i++ { 
        time.Sleep(200 * time.Millisecond)
        drugNumChan <- strconv.Itoa(i)
    }
}

func Stop(cleanupDone chan bool) {
    signalChan := make(chan os.Signal, 1)
    signal.Notify(signalChan, os.Interrupt)
    go func() {
        for _ = range signalChan {
            cleanUp()
            cleanupDone <- true
        }
    }()
    <-cleanupDone
}

func cleanUp() {
    app.MgoSession.Close()
    fmt.Println("清理...\n")
}
