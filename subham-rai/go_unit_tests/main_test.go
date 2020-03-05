package main

import "testing"

func  TestCalculate(t *testing.T) {
	var test =[]struct {
	input int
	output int
	}{	
		{1,3},
		{5,7},
		{8,10},
		{1000,1002},
	}

	for _,i:=range test{
		if out:=Calculate(i.input);out!=i.output{
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", i.input, i.output)
		}
	}

}
