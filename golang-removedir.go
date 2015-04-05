package main

import (
	"fmt"
	 "github.com/alirajabi/golang-removedir/filehelper"
)

func main() {
	var user_path string
	var user_confirm string
	
	fmt.Println("This program will remove your directory, so be careful\n")
	fmt.Println("please insert absolute path :  \n")
	fmt.Scanf("%s", &user_path)
	if filehelper.Exists(user_path) {
		
		fmt.Println("are you sure remove this directory ? (yes|no) \n",user_path)
		fmt.Scanf("%s", &user_confirm)
		
		switch user_confirm {
			default:
				fmt.Println("remove directory aborted!")
				
			case "yes":
				fmt.Println("please wait removing...\n")
				if filehelper.RemoveDir(user_path){	fmt.Println(user_path," removed , good bye") } 
					
			case "no":
				fmt.Println("remove directory aborted!")
          }
          
	} else {
		fmt.Println(user_path,"\t does not exist")
	}
}
