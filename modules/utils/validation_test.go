package utils

import (
	"errors"
	"testing"
)

func TestValidateInputDataCorrectMap(t *testing.T) {
	testMap := []string{"3", "3 2 6", "1 4 0", "8 7 5"}
	err := ValidateInputData(testMap)
	if err == nil {
		t.Log("first test OK")
	} else {
		t.Fatalf("first test: unexpected error: %v", err)
	}

	testMap = []string{"3", "3  2 6  ", "  1 4 0", "8   7    5"}
	err = ValidateInputData(testMap)
	if err == nil {
		t.Log("second test OK")
	} else {
		t.Fatalf("second test: unexpected error: %v", err)
	}
}

func TestValidateInputDataIncorrectAmountOfRows(t *testing.T) {
	testMap := []string{"4", "3 2 6", "1 4 0", "8 7 5"}
	err := ValidateInputData(testMap)
	if err == nil {
		t.Fatalf("third test: unhandled error")
	} else {
		if errors.Is(err, ErrorIncorrectAmountOfRows) == true {
			t.Log("third test OK")
		} else {
			t.Logf("third test: unexpected error: %v", err)
		}
	}
}

func TestValidateInputDataIncorrectAmountOfColumns(t *testing.T) {
	testMap := []string{"3", "3 2 6", "1 4 0 8", "8 7 5"}
	err := ValidateInputData(testMap)
	if err == nil {
		t.Fatalf("forth test: unhandled error")
	} else {
		if errors.Is(err, ErrorIncorrectAmountOfColumns) == true {
			t.Log("forth test OK")
		} else {
			t.Logf("forth test: unexpected error: %v", err)
		}
	}
}

func TestValidateInputDataNegativeNumber(t *testing.T) {
	testMap := []string{"3", "3 2 6", "1 4 -1", "8 7 5"}
	err := ValidateInputData(testMap)
	if err == nil {
		t.Fatalf("fifth test: unhandled error")
	} else {
		if errors.Is(err, ErrorIncorrectMap) == true {
			t.Log("fifth test OK")
		} else {
			t.Logf("fifth test: unexpected error: %v", err)
		}
	}
}

func TestValidateInputDataBigNumber(t *testing.T) {
	testMap := []string{"3", "3 2 6", "1 4 9", "8 7 5"}
	err := ValidateInputData(testMap)
	if err == nil {
		t.Fatalf("sixth test: unhandled error")
	} else {
		if errors.Is(err, ErrorIncorrectMap) == true {
			t.Log("sixth test OK")
		} else {
			t.Logf("sixth test: unexpected error: %v", err)
		}
	}
}

func TestValidateInputDataRepeatedNumber(t *testing.T) {
	testMap := []string{"3", "3 2 6", "1 4 0", "8 7 3"}
	err := ValidateInputData(testMap)
	if err == nil {
		t.Fatalf("seventh test: unhandled error")
	} else {
		if errors.Is(err, ErrorIncorrectMap) == true {
			t.Log("seventh test OK")
		} else {
			t.Logf("seventh test: unexpected error: %v", err)
		}
	}
}

func TestValidateInputDataLittleMap(t *testing.T) {
	testMap := []string{"2", "1 2", "0 3"}
	err := ValidateInputData(testMap)
	if err == nil {
		t.Fatalf("eighth test: unhandled error")
	} else {
		if errors.Is(err, ErrorIncorrectMap) == true {
			t.Log("eighth test OK")
		} else {
			t.Logf("eighth test: unexpected error: %v", err)
		}
	}
}
