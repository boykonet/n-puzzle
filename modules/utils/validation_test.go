package utils

import (
	"errors"
	"testing"
)

func TestValidateInputData(t *testing.T) {
	firstTest := []string{"3", "3 2 6", "1 4 0", "8 7 5"}
	err := ValidateInputData(firstTest)
	if err == nil {
		t.Log("first test OK")
	} else {
		t.Fatalf("first test: unexpected error: %v", err)
	}

	secondTest := []string{"3", "3  2 6  ", "  1 4 0", "8   7    5"}
	err = ValidateInputData(secondTest)
	if err == nil {
		t.Log("second test OK")
	} else {
		t.Fatalf("second test: unexpected error: %v", err)
	}

	thirdTest := []string{"4", "3 2 6", "1 4 0", "8 7 5"}
	err = ValidateInputData(thirdTest)
	if err == nil {
		t.Fatalf("third test: unhandled error")
	} else {
		if errors.Is(err, ErrorVDIncorrectPuzzleSize) == true {
			t.Log("third test OK")
		} else {
			t.Logf("third test: unexpected error: %v", err)
		}
	}

	forthTest := []string{"3", "3 2 6", "1 4 0 8", "8 7 5"}
	err = ValidateInputData(forthTest)
	if err == nil {
		t.Fatalf("forth test: unhandled error")
	} else {
		if errors.Is(err, ErrorVDIncorrectCountElemsInRow) == true {
			t.Log("forth test OK")
		} else {
			t.Logf("forth test: unexpected error: %v", err)
		}
	}

	fifthTest := []string{"3", "3 2 6", "1 4 -1", "8 7 5"}
	err = ValidateInputData(fifthTest)
	if err == nil {
		t.Fatalf("fifth test: unhandled error")
	} else {
		if errors.Is(err, ErrorVDIncorrectPuzzleNumber) == true {
			t.Log("fifth test OK")
		} else {
			t.Logf("fifth test: unexpected error: %v", err)
		}
	}

	sixthTest := []string{"3", "3 2 6", "1 4 9", "8 7 5"}
	err = ValidateInputData(sixthTest)
	if err == nil {
		t.Fatalf("sixth test: unhandled error")
	} else {
		if errors.Is(err, ErrorVDIncorrectPuzzleNumber) == true {
			t.Log("sixth test OK")
		} else {
			t.Logf("sixth test: unexpected error: %v", err)
		}
	}

	seventhTest := []string{"3", "3 2 6", "1 4 0", "8 7 3"}
	err = ValidateInputData(seventhTest)
	if err == nil {
		t.Fatalf("seventh test: unhandled error")
	} else {
		if errors.Is(err, ErrorVDRepeatedNumber) == true {
			t.Log("seventh test OK")
		} else {
			t.Logf("seventh test: unexpected error: %v", err)
		}
	}

	eighthTest := []string{"2", "1 2", "0 3"}
	err = ValidateInputData(eighthTest)
	if err == nil {
		t.Fatalf("eighth test: unhandled error")
	} else {
		if errors.Is(err, ErrorVDNotEnoughInfo) == true {
			t.Log("eighth test OK")
		} else {
			t.Logf("eighth test: unexpected error: %v", err)
		}
	}
}
