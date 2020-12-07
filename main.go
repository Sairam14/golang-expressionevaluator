package main

import (
	"flag"
	"fmt"
	"os"
)

func main(){
	expression := expressionFromFlags()
	if expression == "" {
		expression = expressionFromEnv()
	}
	fmt.Println(expression)
	var iterator = getTokenIterator(expression)
	token := iterator.current()
	fmt.Println(token.evaluateCondition(iterator))
}

func expressionFromFlags() string {
	expression := flag.String("exp","","expression to evaluate")
	flag.Parse()
	return *expression
}

func expressionFromEnv() string {
	return os.Getenv("expression")
}