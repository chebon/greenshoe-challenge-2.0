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

./readFile  data.csv root *rtpwd893:) greenshoe localhost:61613
```
This program will read file then connect to mysql and activemq and insert values concurrently

Reading from database and message broker (part 2):
```
./readDatabase <mysqluser mysqlpassword mysqldatabase activeMQAddress>

./readDatabase root *rtpwd893:) greenshoe localhost:61613
```

This program will connect to mysql and activemq concurrently then generate files:
  "data-2.csv" containg data fetched from mysql
  "data-3.csv" containg data fetched from activeMQ
  
  
Best challange ever 
