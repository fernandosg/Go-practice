package main 

import "fmt"

func main(){
	var names [5]string

	friends:=[5] string {"Luis","Eduardo","Martin","Luis Fernando","Carlos"}
	names = friends
	for i, named := range names {
		fmt.Println(named, &names[i]) // the & is for get the direction
	}
	fmt.Println(names)


}