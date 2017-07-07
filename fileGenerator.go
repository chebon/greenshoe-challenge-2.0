package main

import (
	"fmt"
	"os"
	"math/rand"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main(){
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	var headers string = "column1|column2|column3|column4|column5|column6|column7|column8|column9|column10|column11|column12\n"
	values := make([]string, 0)
	values = append(values,
		"A",
		"G",
		"C",
		"T")

	valuesLimited := make([]string, 0)
	valuesLimited = append(valuesLimited,
		"A",
		"T")

	for i :=0; i<100; i++{
		var columnValue string = ""
		for j :=0; j<12; j++{
			if i<10 && j==4  {
				var myValue string =  "G|"
				columnValue +=myValue
				}else if i>=10 && i<20 && j==4{
					var myValue string =  "C|"
					columnValue +=myValue
					}else if j == 4{
						var myValue string = getValues(valuesLimited, 2) + "|"
						columnValue +=myValue
						} else {
							var myValue string = getValues(values, 4) + "|"
							columnValue +=myValue
						}
					}
					headers +=columnValue+"\n"
				}
				f, err := os.Create(dir+"/data.csv")
				check(err)
				n2, err := f.WriteString(headers)
				fmt.Printf("wrote %d string\n", n2)
				f.Sync()
				check(err)

			}

			func getValues(values []string, permutation int) string{
				list := rand.Perm(permutation)
				return values[list[0]]
			}
