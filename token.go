package main

import "fmt"

type token struct {
	tokenKind tokenKind
	text string
	floatValue float64
}

type tokenCollection struct {
	tokens []token
	index int
}

type tokenIterator interface {
	moveNext() token
	current() token
	hasNext() bool
}

func getTokenIterator(expression string) tokenIterator{
	return &tokenCollection {tokens:tokenize(expression)}
}

func (tokens *tokenCollection) moveNext() token {
	currentIndex := tokens.index + 1
	if currentIndex < len(tokens.tokens) {
		tokens.index = currentIndex
		return tokens.tokens[currentIndex]
	}

	fmt.Println("endddd")
	panic("end of collection")
}

func (tokens *tokenCollection) current() token {
	if tokens.index < len(tokens.tokens) {
		return tokens.tokens[tokens.index]
	}

	panic("end of collection")
}

func (tokens *tokenCollection) hasNext() bool {
	currentIndex := tokens.index + 1
	return currentIndex < len(tokens.tokens)
}
