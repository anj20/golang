package main

import (
	"fmt"
	"sort"
)

func main() {
	// Slices:- the length is of variable
	fmt.Println("Slices")
	var fruitList = []string{"Apple", "Tomato", "Banana"}
	fmt.Println("The fruit list is: ", fruitList)
	fmt.Printf("The type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Orange")
	fmt.Println("The fruit list is: ", fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println("The fruit list is: ", fruitList)
	

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	//highScores[4] = 976

	fmt.Println("The high scores are: ", highScores)
	sort.Ints(highScores)
	fmt.Println("The sorted high scores are: ", highScores)
	fmt.Println("Is Sorted already? ", sort.IntsAreSorted(highScores))
	
	
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println("Courses: ", courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Courses after removal is : ", courses)
}
