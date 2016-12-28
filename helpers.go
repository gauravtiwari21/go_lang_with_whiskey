/*
This is a set of helper functions that will be used by the code in main.go
 */
package myGo

import "fmt"

// MyError is an error raising that includes a message.
type MyError struct {
	message string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v", e.message,)
}

/*
Follwing fucntion removes duplicates from the map
 */
func removeDuplicates(elements []int) []int {
	// Use map to record aduplicates as we find them.
	encountered := map[int]bool{}
	result := []int{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

/*
Follwing fucntion checks if a goven input is positive integar
 */
func isPositive(input int) bool {
	if input <= 0{
		return false
	}
	return true
}

/*
Funtion to find count of occurrences elements in a given array of integers
input: Int array
returns: map with elements of above array as keys and occurence as value
 */
func countOccurences(list []int) map[int]int{
	duplicatesCounts := make(map[int]int)
	for _, item := range list {
		// check if the item/element exist in the duplicate_frequency map
		_, exist := duplicatesCounts[item]
		if exist {
			duplicatesCounts[item] += 1 // increase counter by 1 if already in the map
		} else {
			duplicatesCounts[item] = 1 // else start counting from 1
		}
	}
	return duplicatesCounts
}
