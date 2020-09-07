package main

import (
	"testing"
)

func TestStringExpression(t *testing.T){
	input := "'AAAAA' != 'BBBBB'"
	var iterator = getTokenIterator(input)
	token := iterator.current()
	want := true
	got := token.evaluateCondition(iterator)
	if want != got{
		t.Fatalf("wanted %t but got %t", want, got)
	}
}

func TestIntegerExpression(t *testing.T){
	input := "121221 == 121221"
	var iterator = getTokenIterator(input)
	token := iterator.current()
	want := true
	got := token.evaluateCondition(iterator)
	if want != got{
		t.Fatalf("wanted %t but got %t", want, got)
	}

	input = "121221 == 121221"
	iterator = getTokenIterator(input)
	token = iterator.current()
	want = true
	got = token.evaluateCondition(iterator)
	if want != got{
		t.Fatalf("wanted %t but got %t", want, got)
	}
}

func TestInvalidTokenPanic(t *testing.T){
	defer func(){
		if r := recover(); r != nil{
			t.Logf("recovering fron panic %v.", r)
		}
	}()

	input := "121221 != aasadasd"
	iterator := getTokenIterator(input)
	token := iterator.current()
	token.evaluateCondition(iterator)
}

func TestAndExpression(t *testing.T){
	input := "A==A AND b==b"
	var iterator = getTokenIterator(input)
	token := iterator.current()
	want := true
	got := token.evaluateCondition(iterator)
	if want != got{
		t.Fatalf("wanted %t but got %t", want, got)
	}

	input = "c==c AND c!=c"
	iterator = getTokenIterator(input)
	token = iterator.current()
	want = false
	got = token.evaluateCondition(iterator)
	if want != got{
		t.Fatalf("wanted %t but got %t", want, got)
	}
}

func TestOrExpression(t *testing.T){
	input := "A==A OR b==b"
	var iterator = getTokenIterator(input)
	token := iterator.current()
	want := true
	got := token.evaluateCondition(iterator)
	if want != got{
		t.Fatalf("wanted %t but got %t", want, got)
	}

	input = "C!=c OR b==c"
	iterator = getTokenIterator(input)
	token = iterator.current()
	want = false
	got = token.evaluateCondition(iterator)
	if want != got{
		t.Fatalf("wanted %t but got %t", want, got)
	}
}

func TestNegationExpression(t *testing.T){
	input := "!(A==A)"
	var iterator = getTokenIterator(input)
	token := iterator.current()
	want := false
	got := token.evaluateCondition(iterator)
	if want != got{
		t.Fatalf("wanted %t but got %t", want, got)
	}

/*	input = ""
	iterator = getTokenIterator(input)
	token = iterator.current()
	want = true
	got = token.evaluateConjunction(iterator)
	if want != got{
		t.Fatalf("wanted %t but got %t", want, got)
	}*/
}