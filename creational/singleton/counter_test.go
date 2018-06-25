package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()
	if counter1 == nil {
		t.Errorf("expected pointer to Singleton after calling GetInstance(), not nil")
		return
	}

	expectedCounter := counter1
	currentCount := counter1.AddOne()
	if currentCount != 1 {
		t.Errorf("After calling for the first time to count, the count must be 1 but it is %d\n", currentCount)
		return
	}
	counter2 := GetInstance()
	if counter2 != expectedCounter {
		t.Errorf("Excepted same instance with counter 1 but it is not same")
		return
	}
	currentCount = counter2.AddOne()
	if currentCount != 2 {
		t.Errorf("After calling for the second time to count, the count must be 2 but it is %d\n", currentCount)
		return
	}
}
