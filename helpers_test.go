package myGo

import ("testing"
)

/* TODO delete following tests if not used externally because as they may not be required now 
as they are already validated by tests for calling functions.
 
Test for a helper function used to remove duplicates in a list
*/
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
