Final Solution
=======================
This solution was build with Golang. Mysql database is my preferred database choice

To generate a datafile (part 0):
```
./fileGenerator
```
This will generate a data.csv file 

Reading from file (part 1):
```
./readFile   <filename mysqluser mysqlpassword mysqldatabase activeMQAddress>

example:
./readFile  data.csv root *rtpwd893:) greenshoe localhost:61613
```
This program will read file then connect to mysql and activemq and insert values concurrently

Reading from database and message broker (part 2):
```
./readDatabase <mysqluser mysqlpassword mysqldatabase activeMQAddress>

example:
./readDatabase root *rtpwd893:) greenshoe localhost:61613
```

This program will connect to mysql and activemq concurrently then generate files:
  "data-2.csv" with entry records from from mysql
  "data-3.csv" with entry records from from activeMQ
  
  
Wow this challange was fun. Best so far ;)
