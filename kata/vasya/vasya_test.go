package kata

import (
	"testing"
)

func Test_Ticketyes(t *testing.T) {
	people := []int{25, 25, 50}
	res := Tickets(people)
	if res != "YES" {
		t.Fatalf("want YES, got %v ", res)
	}
}

func Test_Ticketsno(t *testing.T) {
	people := []int{25, 100}
	res := Tickets(people)
	if res != "NO" {
		t.Fatalf("want NO, got %v ", res)
	}
}

func Test_Ticketslong(t *testing.T) {
	people := []int{25, 25, 50, 50, 100}
	res := Tickets(people)
	if res != "NO" {
		t.Fatalf("want NO, got %v ", res)
	}

}
