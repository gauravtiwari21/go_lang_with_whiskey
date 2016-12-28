package myGo

import ("testing"
	"sort"
)

// Following Tests are to check requirement 1 which is to identify Null string
func TestIsNullWithNullStr_Requirement1(t *testing.T) {
	if !isNull(""){
		t.Error("TEST FAILED: Expected True but got False")
	}
}

// Following test also to check requirement 1 which is to return false for valid string
func TestIsNullWithFullStr_Requirement1(t *testing.T) {
	if isNull("FOO") {
		t.Error("TEST FAILED: Expected True but got False")
	}
}

// Following test also to check requirement 1 which is to return false for valid string
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

// Following test also to check requirement 2 but returns error due to non-positive input
func TestFindDivisorsErrorNegativeInput_Requirement2(t *testing.T) {
	_, e  := findDivisors(-90)
	if e == nil{
		t.Error("TEST FAILED: Non positive value accepted")
	}
}

// Following test also to check requirement 2 but returns error due to 0 as input
func TestFindDivisorsErrorZeroInput_Requirement2(t *testing.T) {
	_, e  := findDivisors(0)
	if e == nil{
		t.Error("TEST FAILED: 0 accepted as valid value ")
	}
}

// Following test checks requirement 3 happy path with valid inputs
func TestFindAreaOfTriangle_Requirement3(t *testing.T) {
	desiredResult := float64(9.797958971132712)
	area,_ := getTriangleArea(5,4,7)
	if area != desiredResult{
		t.Error("TEST FAILED: Area calculation is faulty", area, desiredResult)
	}
}

// Following test checks requirement 3 with negative inputs
func TestFindAreaOfTriangleNegativeInput_Requirement3(t *testing.T) {
	_,e := getTriangleArea(-5,4,7)
	if e == nil{
		t.Error("TEST FAILED: Negative value accepted")
	}
}

// Following test checks requirement 3 with invalid sides that cannot form a triangle
func TestFindAreaOfTriangleInvaidTriangle_Requirement3(t *testing.T) {
	_,e := getTriangleArea(5,4,12)
	if e == nil{
		t.Error("TEST FAILED: Invalid input accepted as a side of triangle")
	}
}


// Following test checks requirement 4 when input has repetitions
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


// Following test checks requirement 4 when the input has all unique elements
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

// Test for a helper function used to remove duplicates in a list
func TestRemoveDuplicates(t *testing.T){
	inputLIst := []int {1,9,9,41}
	outputSampleList := []int {1,9,41}
	realOutput := []int (removeDuplicates(inputLIst))
	for i := range outputSampleList {
		if realOutput[i] != outputSampleList[i] {
			t.Error("TEST FAILED: Expected output mismatch", realOutput)
			break
		}
	}
}

// Test for a helper function used to count occurences of elements in a list
func TestCountOccurences(t *testing.T){
	inputLIst := []int {1,9,9,41}
	outputSample := map[int]int {1:1,9:2,41:1}
	realOutput := map[int]int (countOccurences(inputLIst))
	for i := range outputSample {
		if realOutput[i] != outputSample[i] {
			t.Error("TEST FAILED: Expected output mismatch", realOutput, outputSample)
			break
		}
	}
}