package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type Lexer struct {
	curr   int
	source string
	len    int
}

func (lex *Lexer) advance() byte {

	lex.curr++
	return lex.source[lex.curr-1]
}
func (lex *Lexer) peek(n int) byte {
	return lex.source[lex.curr+n]
}

func main() {
	sum := 0
	doFlag := false
	dontFlag := false

	corruptedData, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lex := Lexer{
		curr:   0,
		source: string(corruptedData),
		len:    len(corruptedData),
	}

	for lex.curr < lex.len {
		char := lex.peek(0)
		firstVal := 0
		secondVal := 0
		switch char {
		case 'd':
			{
				word := ""
				for unicode.IsLetter(rune(lex.peek(0))) || lex.peek(0) == '\'' {
					word += string(lex.advance())
				}
				switch word {
				case "do":
					{

						if lex.peek(0) == '(' && lex.peek(1) == ')' {
							lex.advance()
							lex.advance()
							doFlag = true
							dontFlag = false
						}

					}
				case "don't":
					{
						if lex.peek(0) == '(' && lex.peek(1) == ')' {
							lex.advance()
							lex.advance()
							doFlag = false
							dontFlag = true
						}
					}
				default:
					lex.advance()
					continue

				}
			}
		case 'm':
			{
				word := ""
				for unicode.IsLetter(rune(lex.peek(0))) {
					word += string(lex.advance())
				}

				if word != "mul" {
					continue
				}

				if lex.peek(0) != '(' {
					lex.advance()
					continue
				}

				lex.advance() // skips (

				if !unicode.IsDigit(rune(lex.peek(0))) {
					lex.advance()
					continue
				}

				valStr := ""

				for unicode.IsDigit(rune(lex.peek(0))) {
					valStr += string(lex.advance())
				}

				firstVal, _ = strconv.Atoi(valStr)

				if lex.peek(0) != ',' {
					lex.advance()
					continue
				}

				lex.advance() // skips comma

				if !unicode.IsDigit(rune(lex.peek(0))) {
					lex.advance()
					continue
				}

				valStr = ""

				for unicode.IsDigit(rune(lex.peek(0))) {
					valStr += string(lex.advance())
				}

				secondVal, _ = strconv.Atoi(valStr)

				if lex.peek(0) != ')' {
					lex.advance()
					continue
				}
				lex.advance() // skips )
			}
		default:
			lex.advance()
		}

		if (!doFlag && !dontFlag) || doFlag {
			sum += (firstVal * secondVal)
		}

	}
	fmt.Println(sum)
}
