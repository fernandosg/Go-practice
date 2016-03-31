//Maps
//With this structure is possible to implement HashTables or dictionary
//FIll the map with 5 values and iterate over this and print the keys a value inside map 

package main

import "fmt"

func main(){
	departaments:=make(map[string]int)
	//Fill the map
	departaments["Devs"]=25
	departaments["Marketing"]=50 
	departaments["Executives"]=4
	departaments["Sellers"]=60
	departaments["Support"]=8

	//Iterate over the map and print every Key and value
	for key,value :=range departaments{
		fmt.Printf("Dept %s persons %d \n",key,value)
	}
}