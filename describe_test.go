package main

import "testing"

var (
	boolean bool
	number ChosenNumber
	describer Describer
	value string
    )

func TestNotZero(t *testing.T) {
    number = ChosenNumber(2.0)
    describer = number
    boolean = describer.NotZero()
    if !boolean {
	    t.Errorf("Expected true, got %v\n", boolean)
    }

    number = ChosenNumber(0.0)
    describer = number
    boolean = describer.NotZero()
    if boolean {
	    t.Errorf("Expected false, got %v\n", boolean)
    }
}

func TestIsFactor(t *testing.T) {
	number = ChosenNumber(17.0)
	describer = number
	boolean = describer.IsFactor(2)
	if boolean {
		t.Errorf("Expected false, got %v\n", boolean)
	}

	number = ChosenNumber(12.0)
	describer = number
	boolean = describer.IsFactor(4)
	if !boolean {
		t.Errorf("Expected true, got %v\n", boolean)
	}
}

func TestDescribe(t *testing.T) {
	number = ChosenNumber(3.0)
	describer = number
	value = describer.Describe()
        if value != "Fizz" {
		t.Errorf("Expected Fizz, got %v\n", value)
	}
	number = ChosenNumber(5.0)
	describer = number
	value = describer.Describe()
	if value != "Buzz" {
		t.Errorf("Expected Buzz, got %v\n", value)
	}
	number = ChosenNumber(15.0)
	describer = number
	value = describer.Describe()
	if value != "FizzBuzz" {
		t.Errorf("Expected FizzBuzz, got %v\n", value)
	}
	number = ChosenNumber(0.0)
	describer = number
	value = describer.Describe()
	if value != "0" {
		t.Errorf("Expected 0, got %v\n", value)
	}
	number = ChosenNumber(16.0)
	describer = number
	value = describer.Describe()
	if value != "16" {
		t.Errorf("Expected 16, got %v\n", value)
	}
}

func ToString(t *testing.T) {
	number = ChosenNumber(10)
	describer = number
	value = describer.ToString()
	if value != "10" {
		t.Errorf("Expected string \"10\", got %v\n", value)
	}
}
