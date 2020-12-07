package main

import (
	"regexp"
	"strconv"
	"strings"
)

func (this  token)  evaluateCondition(iterator tokenIterator) bool {
	result := this.evaluateDisjunction(iterator)
	return result
}

func (this token) evaluateDisjunction(iterator tokenIterator) bool {
	lhs  := this.evaluateConjunction(iterator)

	return lhs
}

func (this token) evaluateNegation(iterator tokenIterator) bool {
	negated := false
	for this.tokenKind == "!"{
		negated = !negated
	}

	result := this.evaluateEquality(iterator)
	if !negated {
		return result != negated
	} else {
		return result == negated
	}
}

func (this token) evaluateConjunction(iterator tokenIterator) bool {
	var lhs = this.evaluateNegation(iterator)

	if iterator.hasNext() {
		operator := iterator.moveNext()

		for operator.tokenKind == "AND" || operator.tokenKind == "OR" {
			rhsTokenKind := operator.tokenKind
			operator = iterator.moveNext()
			var rhs = operator.evaluateNegation(iterator)

			if rhsTokenKind == "AND" {
				lhs = lhs && rhs
			} else if rhsTokenKind == "OR" {
				lhs = lhs || rhs
			}
		}
	}

	return lhs
}

func (this token) evaluateString(iterator tokenIterator) string {
	result := ""
	switch this.tokenKind {
	case STRING:
		result = this.text
	case BAREWORD:
		result = this.text
	}
	return result
}

func (this token) evaluateInteger(iterator tokenIterator) float64 {
	result := 0.0
	switch  this.tokenKind {
	case FLOAT:
		result = this.floatValue
	case INTEGER:
		result = this.floatValue
	default:
		panic("Invalid token kind")

	}
	return result
}

func (this token) evaluateEquality(iterator tokenIterator) bool{
	result := false

	if this.tokenKind == "BAREWORD" || this.tokenKind == "STRING" {
		lhs := this.evaluateString(iterator)
		operator :=  iterator.moveNext()
		rhsToken := iterator.moveNext()
		rhs := rhsToken.evaluateString(iterator)

		switch operator.tokenKind {
		case EQUAL:
			result = strings.EqualFold(lhs, rhs)
			break

		case NOTEQUAL:
			result = !strings.EqualFold(lhs, rhs)
			break

		default:
			panic("IllegalArgumentError")
		}
	}

	if this.tokenKind == INTEGER || this.tokenKind == FLOAT {
		lhs := this.evaluateInteger(iterator)
		operator := iterator.moveNext()
		rhsToken := iterator.moveNext()
		rhs := rhsToken.evaluateInteger(iterator)

		switch operator.tokenKind {
		case GREATERTHAN:
			result = lhs > rhs
			break;
		case GREATERTHANANDEQUAL:
			result = lhs >= rhs
		case LESSTHAN:
			result = lhs < rhs
		case LESSTHANANDEQUAL:
			result = lhs <= rhs
		case EQUAL:
			result = lhs == rhs
		case NOTEQUAL:
			result = lhs != rhs
		default:
			panic("Invalud integer token kind")
		}

	}

	return result
}

func getToken(kind tokenKind, text string, floatValue float64) token{
	var token token
	token.tokenKind = kind
	token.text = text
	token.floatValue = floatValue
	return token
}

func tokenize(input string) []token{
	var tokens []token
	var token token

	i := 0
	runes := []rune(input)
	for i < len(runes) {
		if input[i] == 32 { // White space
			i += 1
			continue
		}

		r, _ := regexp.Compile("^\\'[^\\']*'")
		matchIndex := r.FindIndex([]byte(string(runes[i:])))
		if matchIndex != nil {
			token = getToken(tokenKind("STRING"), string(runes[i:i+matchIndex[1]]), 0)
			tokens = append(tokens, token)
			i += len(token.text)
			continue
		}

		r, _ = regexp.Compile("^\\=\\=")
		matchIndex = r.FindIndex([]byte(string(runes[i:])))
		if matchIndex != nil {
			token = getToken(tokenKind("EQUAL"), string(runes[i:i+matchIndex[1]]), 0)
			tokens = append(tokens, token)
			i += len(token.text)
			continue
		}


		r, _ = regexp.Compile("^\\!\\=")
		matchIndex = r.FindIndex([]byte(string(runes[i:])))
		if matchIndex != nil {
			token = getToken(tokenKind("NOTEQUAL"), string(runes[i:i+matchIndex[1]]), 0)
			tokens = append(tokens, token)
			i += len(token.text)
			continue
		}

		r, _ = regexp.Compile("(?i)^and")
		matchIndex = r.FindIndex([]byte(string(runes[i:])))
		if matchIndex != nil {
			token = getToken(tokenKind("AND"), string(runes[i:i+matchIndex[1]]), 0)
			tokens = append(tokens, token)
			i += len(token.text)
			continue
		}

		r, _ = regexp.Compile("(?i)^or")
		matchIndex = r.FindIndex([]byte(string(runes[i:])))
		if matchIndex != nil {
			token = getToken(tokenKind("OR"), string(runes[i:i+matchIndex[1]]), 0)
			tokens = append(tokens, token)
			i += len(token.text)
			continue
		}

		r, _ = regexp.Compile("^[+-]?([0-9]*[.])?[0-9]+")
		matchIndex = r.FindIndex([]byte(string(runes[i:])))
		if matchIndex != nil {
			value, _ := strconv.ParseFloat(string(runes[i:i+matchIndex[1]]), 64)
			token = getToken(tokenKind("FLOAT"), "", value)
			tokens = append(tokens, token)
			i += len(string(runes[i:i+matchIndex[1]]))
			continue
		}

		r, _ = regexp.Compile("^\\d+")
		matchIndex = r.FindIndex([]byte(string(runes[i:])))
		if matchIndex != nil {
			value, _ := strconv.ParseInt(string(runes[i:i+matchIndex[1]]), 8,64)
			token = getToken(tokenKind("INTEGER"), "", float64(value))
			tokens = append(tokens, token)
			i += len(string(runes[i:i+matchIndex[1]]))
			continue
		}

		r, _ = regexp.Compile(	"^\\w+")
		matchIndex = r.FindIndex([]byte(string(runes[i:])))
		if matchIndex != nil {
			token = getToken(tokenKind("BAREWORD"), string(runes[i:i+matchIndex[1]]), 0)
			tokens = append(tokens, token)
			i += len(token.text)
			continue
		}



		token = getToken(tokenKind("ENDOFINPUT"), "end of input", 0)
		tokens = append(tokens, token)
		break
	}

	return tokens
}