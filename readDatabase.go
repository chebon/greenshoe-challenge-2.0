package main

import (
  "fmt"
  "github.com/go-stomp/stomp"
  _ "github.com/go-sql-driver/mysql"
  "database/sql"
  "os"
  "path/filepath"
)
var activequeueEntry = ""

func main() {
  c1 := make(chan string)
  c2 := make(chan string)
  go fetchQueue(c1)
  go fetchMysql(c2)

  select {
  case msg1 := <-c1:
    fmt.Println("received", msg1)
  case msg2 := <-c2:
    fmt.Println("received", msg2)
  }
}

//Connect to ActiveMQ and listen for messages
func fetchQueue(c1 chan string) {
  conn, err := stomp.Dial("tcp", "localhost:61613")
  if err != nil {
    fmt.Println(err)
  }

  sub, err := conn.Subscribe("/queue/greenshoe", stomp.AckAuto)
  checkErr(err)
  for n := 0; n <= 79; n++ {
    fmt.Println(n)
    msg := <-sub.C
    entry := string(msg.Body)+"\n"
    fmt.Println(entry)
    activequeueEntry = activequeueEntry + entry
  }
  fileGenerator(activequeueEntry, "data-3")
  err = sub.Unsubscribe()
  checkErr(err)
  defer conn.Disconnect()
  fmt.Println("activeMQ loop done")
  c1 <- "Fetching from activeMQ completed"
}

//processing mysql data
func fetchMysql(c2 chan string) {
  mySqlUser := os.Args[1]
  mySqlPasswood := os.Args[2]
  mySqlDatabase := os.Args[3]
  mySqlConnectionParm := mySqlUser+":"+mySqlPasswood+"@/"+mySqlDatabase
  con, err := sql.Open("mysql", mySqlConnectionParm)
  rows, err := con.Query("select * from data")
  if err != nil { }
  var column1, column2, column3, column4, column5, column6, column7, column8, column9, column10, column11, column12 string
  amendedString := ""
  for rows.Next() {
    err = rows.Scan(&column1, &column2, &column3, &column4, &column5, &column6, &column7, &column8, &column9, &column10, &column11, &column12)
    var entry string = column1+"|"+column2+"|"+column3+"|"+column4+"|"+column5+"|"+column6+"|"+column7+"|"+column8+"|"+column9+"|"+column10+"|"+column11+"|"+column12
    amendedString += entry+"\n"
  }
  fileGenerator(amendedString, "data-2")
  c2 <- "value copied from database"

}

//writting to file
func fileGenerator(data string, filename string){
  dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
  var headers string = "column1|column2|column3|column4|column5|column6|column7|column8|column9|column10|column11|column12\n"
  headers +=data+"\n"
  f, err := os.Create(dir+"/"+filename+".csv")
  checkErr(err)
  n2, err := f.WriteString(headers)
  fmt.Printf("wrote %d string\n", n2)
  f.Sync()
  checkErr(err)
  fmt.Println("writting done")
  return
}

//error checker
func checkErr(err error) {
  if err != nil {
    panic(err)
  }
}
