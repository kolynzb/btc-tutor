package arr_example

import (
	"fmt"
) 
func arr_example() { 
 var array [10]string = [10]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"} 

 for i := 0; i < 9; i++ { 
 	fmt.Printf("%s, ", array[i]) 
 } 

 fmt.Printf("%s\n", array[9]) 
 //Output: one, two, three, four, five, six, seven, eight, nine, ten
} 