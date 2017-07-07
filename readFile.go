package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
	"path/filepath"
	"strings"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/go-stomp/stomp"
)

func main() {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	checkErr(err)
	dataFile := os.Args[1]
	fmt.Println(dir+"/"+dataFile)
	file, err := os.Open(dir+"/"+dataFile)
	checkErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stringProcessor(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func stringProcessor(str string){

	c1 := make(chan string)
	c2 := make(chan string)
	var dataArray = strings.Split(str, "|")
	if(len(dataArray[4]) < 2){
		if dataArray[4] == "C" || dataArray[4] == "G" {
			go insertToDatabase(c1, dataArray)
			} else {
				go insertTOActiveMQ(c2, dataArray)
			}

			select {
			case msg1 := <-c1:
				fmt.Println("received", msg1)
			case msg2 := <-c2:
				fmt.Println("received", msg2)
			}
		}
	}

	func insertToDatabase(c1 chan string, data []string){
		mySqlUser := os.Args[2]
		mySqlPasswood := os.Args[3]
		mySqlDatabase := os.Args[4]
		mySqlConnectionParm := mySqlUser+":"+mySqlPasswood+"@/"+mySqlDatabase
		con, err := sql.Open("mysql", mySqlConnectionParm)
		_, err = con.Exec("insert into data (column1, column2, column3, column4, column5, column6, column7, column8, column9, column10, column11, column12) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", data[0], data[1], data[2], data[3], data[4], data[5], data[6], data[7], data[8], data[9], data[10], data[11])
		checkErr(err)
		defer con.Close()
		c1 <- "value inserted to database"
	}

	func insertTOActiveMQ(c2 chan string, data []string){
		conn, err := stomp.Dial("tcp", "localhost:61613")
		checkErr(err)
		go Producer(c2, conn, data)
	}

	func Producer(c2 chan string, conn *stomp.Conn, data []string) {
		stringValue := strings.Join(data, "|")
		err := conn.Send(
			"/queue/greenshoe", // destination
			"text/plain", // content-type
			[]byte(stringValue)) // body
			checkErr(err)
			c2 <- "value inserted to activeMQ"
		}

		func checkErr(err error) {
			if err != nil {
				panic(err)
			}
		}
