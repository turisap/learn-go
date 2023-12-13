package main

import "fmt"

func main() {
	//assignment1()
	//assignment2()
	//assignment3()
	//assignment4()
	//assignment5()
	assignment6()
}

func assignment1() {
	mySlice := []float64{1.2, 5.6}

	mySlice = append(mySlice, float64(6))

	a := 10
	mySlice[0] = float64(a)

	mySlice = append(mySlice, 10.10)

	mySlice = append(mySlice, float64(a))

	fmt.Println(mySlice)

}

func assignment2() {
	nums := []float64{1.1, 2.2, 3, 3}

	nums = append(nums, 10.1)
	nums = append(nums, 4.1, 5.5, 6.6)

	fmt.Println(nums)

	n := []float64{99.9, 100.0}

	n = append(nums, n...)

	fmt.Println(nums)
}

func assignment3() {
	nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}
	var sum int = 0

	for _, v := range nums[1 : len(nums)-2] {
		sum += v
		fmt.Println(v)
	}

	fmt.Printf("The sum is %d\n", sum)
}

func assignment4() {
	friends := []string{"Marry", "John", "Paul", "Diana"}
	friendsNew := make([]string, len(friends))

	count := copy(friendsNew, friends)
	fmt.Println("Copied items count:", count)

	friendsNew[1] = "KISKIS"

	fmt.Println(friends)
	fmt.Println(friendsNew)
}

func assignment5() {
	friends := []string{"Marry", "John", "Paul", "Diana"}
	friendsNew := append([]string{}, friends...)

	friendsNew[1] = "Oh MY Budda"

	fmt.Println(friends)
	fmt.Println(friendsNew)
}

func assignment6() {
	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
	newYears := []int{}
	newYears = append(years[:3], years[len(years)-3:]...)

	years[0] = 9999999
	fmt.Println(newYears)
}
