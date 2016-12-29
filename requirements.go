package myGo

import (
	"math"
	"sort"
	"strings"
)

/* REQUIREMENT 1
Function for requirement 1. In go extended methods are not applicable so implementing this way,
in C# there would have been an extended method 
Input: String
Output: Boolean, true if not null, false if null
*/ 
func isNull(input string)  bool{
	if input == ""{
		return true
	} else{
		return false
	}
}


/* REQUIREMENT 2
Function for requirement 2 to find all divisors
This fuction works on following logic to get all the divisors
# Calculate all the divisors that are less than the square root of the number
# Use the divisors obtained above to divide the number and find other divisors that are above the square root
# Remove any duplicates in the above set of divisors and return the result
INPUT: Positive Integar whose divisors need to be calculated
Output: Integar array of divisors
 */
func findDivisors(num int) ([]int, error) {
	divisors := []int{}
	// reject if negative value is passed and raise error for them
	if !isPositive(num){
		return nil, MyError{"Only positive values are allowed"}
	}
	/* First find all divisors less than square root*/
	for i := 1; i <= int(math.Sqrt(float64(num))) + 1; i++ {
		if num % i == 0 {
			divisors = append(divisors, i)
		}
	}
	/* Divide the given number by the divisors above to find other divisors. */
	higherDivisors := []int{}
	for _,v := range divisors{
		if (num/v) != v{
			higherDivisors = append(higherDivisors, (num/v))
		}
	}

	/* Append all the higher divisors to the initial list of divisors lower thatn square-root to get single lit of
	 all divisors */
	for _,v := range higherDivisors{
		divisors = append(divisors, v)
	}

	/*Remove duplicates in the array of divisors and sort the array*/
	divisors = removeDuplicates(divisors)
	sort.Ints(divisors)
	return divisors, nil
}

/*
REQUIREMENT 4
This method finds the most common elements of the array and returns an array containing them
It follows the following logic
1. Iterates through input integar array and stores the the occurrence of the elements in a  map with unique element
as key
2. Iterates through the hash map created and populates another final  map with only one key that is the maximum
 count pointing to array of integers having that count
3. Returns the array refered to by the only key in the result hash map
 */
func getDuplicates(list []int) []int {
	result_map := make(map[int][]int)
	result_map[0] = []int{0}   // Initalise an empty map
	var max int
	duplicatesCounts := make(map[int]int)
	duplicatesCounts = countOccurences(list)

	/* Populate the result map that needs to have only one key (maximum occurences)
	which will point to an array of elements */
	for k, v := range duplicatesCounts {
		for result_key := range result_map{
			if v < result_key{
				continue
			}else if v > result_key{
				/* This ensures that the length of the map will always be 1 which will
				hold the array of most common elements that were in the initally passed array.
				This also ensures the size of the result map stays low. */
				delete(result_map, result_key)
			}
			result_map[v]=append(result_map[v], k)
			max = v
		}
	}
	return result_map[max]
}



/* REQUIREMENT 3
Function for requirement 3. Calculate area of a Triangle given its three sides are passed as parameters,
Input: Length of the three sides as integer
Returns: Boolean, true if not null, false if null
*/
func getTriangleArea(side1, side2, side3 int) (float64, error){
	sides := [] int {side1, side2, side3}   // Create and array of sides for convenience
	sort.Ints(sides)

	// Enforce positive input for sides only
	for item := range sides {
		if !isPositive(sides[item]){
			return -1, MyError{"Only positive values are allowed"}
		}
	}

	/* Validate if a triangle can be formed from the given sides
	Logic is that the largest side will always smaller than the sum of other two sides
	Last element of array is longest as it was already sorted
	*/
	if sides[2] >= sides[0] + sides[1] {
		return -2, MyError{"Invalid side lengths, cannot draw a traiangle from these"}
	}

	halfPerimeter := float64((side1 + side2 + side3)/2)   // s=(a+b+c)/2  for a,b,c as sides

	// Area = sqrt(s * (s-a) * (s-b) * (s-c)) for a,b,c as sides
	area := (math.Sqrt(
		halfPerimeter * (
			halfPerimeter-float64(side1)) * (
			halfPerimeter-float64(side2)) * (
			halfPerimeter-float64(side3))))

        return area, nil
}


/*
REQUIREMENT 5: Adding methods to a type as startswith and endswith
 */

type String string  // Created a type and added the below methods to it so that they can be called using . <dot> notation

/*
Method to check if a text in the type "String" starts with a substring or ""
Input: Main text to be searched for substring (Type: String), substring text to be searched(Type: String)
Returns: bool (true if found or substring "", false if not found)
 */
func (fullString String) startsWith(subString String) bool {
	if subString == ""{
		return true
	}
	return strings.HasPrefix(string(fullString), string(subString))
}

/*
Method to check if a text in the type "String" ends with a substring or ""
Input: Main text to be searched for substring (Type: String), substring text to be searched(Type: String)
Returns: bool (true if found or substring "", false if not found)
 */
func (fullString String) endsWith(findStr String) bool {
	if findStr == ""{
		return true
	}
	return strings.HasSuffix(string(fullString), string(findStr))
}
