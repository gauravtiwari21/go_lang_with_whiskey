package myGo

import ("testing"
	"sort"
)

// Following test is to check requirement 1 which is to identify Null string, passed a null string and expecting True
func TestIsNullWithNullStr_Requirement1(t *testing.T) {
	if !isNull(""){
		t.Error("TEST FAILED: Expected True but got False")
	}
}

// Following test also checks requirement 1 which is to return false for valid string passed as FOO
func TestIsNullWithFullStr_Requirement1(t *testing.T) {
	if isNull("FOO") {
		t.Error("TEST FAILED: Expected True but got False")
	}
}

// Following test is to check requirement 2 which is to return divisors for a number. HAPPYPATH test with valid input
func TestFindDivisors_Requirement2(t *testing.T) {
	outputDivisors, _  := findDivisors(60)
	desiredDivisors := []int { 1, 2, 3, 4, 5, 6, 10, 12, 15, 20, 30, 60, 30}
	for i := range outputDivisors{
		if desiredDivisors[i] != outputDivisors[i]{
			t.Error("TEST FAILED: Expected output mismatch", outputDivisors, desiredDivisors)
			break
		}
	}

}

// Following test checks requirement 2 which is to return divisors for a number but expects error due to non-positive input
func TestFindDivisorsErrorNegativeInput_Requirement2(t *testing.T) {
	_, e  := findDivisors(-90)
	if e == nil{
		t.Error("TEST FAILED: Non positive value accepted")
	}
}

// Following test checks requirement 2 which is to return divisors for a number but expects error due to 0 as input
func TestFindDivisorsErrorZeroInput_Requirement2(t *testing.T) {
	_, e  := findDivisors(0)
	if e == nil{
		t.Error("TEST FAILED: 0 accepted as valid value ")
	}
}

// Following test checks requirement 3 happy path with valid inputs to return area of a triangle
func TestFindAreaOfTriangle_Requirement3(t *testing.T) {
	desiredResult := float64(9.797958971132712)
	area,_ := getTriangleArea(5,4,7)
	if area != desiredResult{
		t.Error("TEST FAILED: Area calculation is faulty", area, desiredResult)
	}
}

// Following test checks requirement 3 with negative inputs as sides of triangle and expects error
func TestFindAreaOfTriangleNegativeInput_Requirement3(t *testing.T) {
	_,e := getTriangleArea(-5,4,7)
	if e == nil{
		t.Error("TEST FAILED: Negative value accepted")
	}
}

// Following test checks requirement 3 with invalid sides that cannot form a triangle and expects error
func TestFindAreaOfTriangleInvaidTriangle_Requirement3(t *testing.T) {
	_,e := getTriangleArea(5,4,12)
	if e == nil{
		t.Error("TEST FAILED: Invalid input accepted as a side of triangle")
	}
}


// Following test checks requirement 4 when input array has repetitions
func TestFindDuplicates_Requirement4(t *testing.T) {
	input := []int{1,4,2,5,6,1,6,2,9}
	outputArray := getDuplicates(input)
	desiredResult := []int{1,2,6}
	sort.Ints(outputArray)
	sort.Ints(desiredResult)
	for i := range desiredResult {
		if outputArray[i] != desiredResult[i] {
			t.Error("TEST FAILED: Expected output mismatch", outputArray, desiredResult)
			break
		}
	}
}


// Following test checks requirement 4 when the input array has all unique elements
func TestFindDuplicatesNoRepetitions_Requirement4(t *testing.T) {
	var input = []int{1,4,2,5,6,9}
	outputArray := getDuplicates(input)
	desiredResult := []int{1,2,4,5,6,9}
	sort.Ints(outputArray)
	sort.Ints(desiredResult)
	for i := range desiredResult {
		if outputArray[i] != desiredResult[i] {
			t.Error("TEST FAILED: Expected output mismatch", outputArray, desiredResult)
			break
		}
	}
}

/* Following Tests are to check requirement 5 which is to add startswith and endswith methods to a type
This test should test returning true for main string found to start with substring
 */
func TestStartsWith_Requirement5(t *testing.T) {
	mainStr := String("hang the dj han")
	findStr := String("hang")
	if !mainStr.startsWith(findStr){
		t.Error("TEST FAILED: valid substring not found in main string: ", mainStr, findStr)
	}
	findStr = ""
	if !mainStr.startsWith(findStr){
		t.Error("TEST FAILED: valid substring not found in main string: ", mainStr, findStr)
	}
}

/* Following Tests are to check requirement 5 which is to add startswith and endswith methods to a type
This test should test returning true for main string found to end with substring
 */
func TestEndsWith_Requirement5(t *testing.T) {
	mainStr := String("hang the dj hand")
	findStr := String("hand")
	if !mainStr.endsWith(findStr){
		t.Error("TEST FAILED: valid substring not found in main string: ", mainStr, findStr)
	}
	findStr = ""
	if !mainStr.endsWith(findStr){
		t.Error("TEST FAILED: valid substring not found in main string: ", mainStr, findStr)
	}
}

/* Following Tests are to check requirement 5 which is to add startswith and endswith methods to a type
This test should test returning false for main string not found to end with substring
 */
func TestEndsWithMissing_Requirement5(t *testing.T) {
	mainStr := String("hang the dj hand")
	findStr := String("handle")
	if mainStr.endsWith(findStr) {
		t.Error("TEST FAILED: valid substring not found in main string: ", mainStr, findStr)
	}
}

/* Following Tests are to check requirement 5 which is to add startswith and endswith methods to a type
This test should test returing false for main string not found to start with substring
 */
func TestStartsWithMissing_Requirement5(t *testing.T) {
	mainStr := String("hang the dj hand")
	findStr := String("handle")
	if mainStr.endsWith(findStr) {
		t.Error("TEST FAILED: valid substring not found in main string: ", mainStr, findStr)
	}
}
