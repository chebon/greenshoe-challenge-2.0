Final Solution
=======================
This solution was build with Golang. Mysql database is my prefred database choice

To generate a datafile (part 0):
```
./fileGenerator
```
This will generate a data.csv file 

Reading from file:
```
./readFile   <filename mysqluser mysqlpassword mysqldatabase activeMQAddress>
```
This program will read file then connect to mysql and activemq and insert values concurrently

Reading from database and message broker:
```
./readDatabase <mysqluser mysqlpassword mysqldatabase activeMQAddress>
```

This program will connect to mysql and activemq concurrently then generate files:
  "data-2.csv" containg data fetched from mysql
  "data-3.csv" containg data fetched from activeMQ
