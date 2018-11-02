package main
import "fmt"

func twoSum(nums []int, target int) []int {
	found := make(map[int]int)
	for index, element := range nums {
        
		if val, ok := found[element]; ok {
			temp1 := []int{val, index}
			return temp1
		}
        
		temp := target - element
		found[temp] = index
	}
	res := []int{-1, -1}
	return res
}

func main() {
	array := []int{-3, 4, 3, 90}
	a := twoSum(array, 0)
	fmt.Printf("%v\n", a)
}

