package main

import "testing"

func TestShould_add_two_digits(t *testing.T) {
	got := Calculate("2+2")
	want := 4

	if got != want {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}


func TestShould_add_two_digits_with_spaces(t *testing.T) {
	got := Calculate("   2   +  2    ")
	want := 4

	if got != want {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}

func TestShould_add_digit_and_number(t *testing.T) {
	got := Calculate("42+2")
	want := 44

	if got != want {
		t.Errorf("got '%d' want '%d'", got, want)
	}

	got = Calculate("2+42")
	want = 44

	if got != want {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}
//func TestGet_stack_of_numbers(t *testing.T) {
//	got := Get_stack_of_numbers("3-5-5/2-5")
//	want := []int {3, 5, 5, 2, 5}
//
//	for i, d := range got {
//		if d != want[i] {
//			t.Errorf("got '%d' want '%d'", d, want[i])
//		}
//	}
//}
//
//func TestAddition(t *testing.T) {
//	got := Calculate("2+2")
//	want := 4
//
//	if got != want {
//		t.Errorf("got '%d' want '%d'", got, want)
//	}
//}
//
//func TestCalculate(t *testing.T) {
//	got := Calculate("2+2")
//	want := 4
//
//	if got != want {
//		t.Errorf("got '%d' want '%d'", got, want)
//	}
//}
