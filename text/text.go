package text

// This consists of a bunch of text utilities to help in generating pretty and meaningful 
// help messages, error messages, etc.

import (

	"strconv"
	"strings"

	"path/filepath"

	"charm/token"
)

const (
	VERSION = "0.2.1"
	BULLET = " ▪ "
	PROMPT = "→ ")

func ToEscapedText(s string) string {
	result := "\""
	for _, ch := range(s) {
		switch ch {
		case '\n' :
			result = result + "\n"
		case '\r' :
			result = result + "\r"
		case '\t' :
			result = result + "\t"
		default : result = result + string(ch)
		}
	}
	return result + "\""
}

func FlattenedFilename(s string) string {
	s = filepath.Base(s)
	s = strings.Replace(s, ".", "_", -1)
	return s
}

func Cyan(s string) string {
	return CYAN + s + RESET
}

func Emph(s string) string { 
	return Cyan("'" + s + "'");
}

func EmphType(s string) string { 
	return Cyan("<" + s + ">");
}

func Red(s string) string {
	return RED + s + RESET;
}

func Green(s string) string {
	return GREEN + s + RESET;
}

func Yellow(s string) string {
	return YELLOW + s + RESET;
}

func Logo() string {
	var padding string
	if len(VERSION) % 2 == 0 {padding = ","}
	titleText := " Charm" + padding + " version " + VERSION + " "
	loveHeart := Red("♥")
	leftMargin := "  "
	bar := strings.Repeat("═", len(titleText) / 2)
	logoString := "\n" + 
		leftMargin + "╔" + bar + loveHeart + bar + "╗\n" +
		leftMargin + "║"       + titleText +       "║\n" +
	    leftMargin + "╚" + bar + loveHeart + bar + "╝\n\n"
	return logoString
}

func DescribePos(token token.Token) string {
	prettySource := token.Source
	if prettySource != "REPL input" {
		prettySource = "'" + prettySource + "'"
	}
	if token.Line > 0 {
		result := strconv.Itoa(token.Line) + ":" + strconv.Itoa(token.ChStart)
		if token.ChStart != token.ChEnd { result = result + "-" + strconv.Itoa(token.ChEnd) }
		result = " at line" + "@" + result + "@"
	
		return result + "of " + prettySource
	}
	return " in " + prettySource
}

// Describes a token for the purposes of error messages etc.
//
func DescribeTok(tok token.Token) string {
	switch tok.Type {
		case token.LPAREN :
			if tok.Literal == "|->" { return "indent" }
		case token.RPAREN :
			if tok.Literal == "<-|" { return "outdent" }
		case token.NEWLINE :
			if tok.Literal == "\n" { return "newline" }
		case token.EOF :
			return "end of line"
		case token.STRING :
			return "<string>"
		case token.INT :
			return "<int>"
		case token.FLOAT :
			return "<float>"
		case token.TRUE :
			return "<bool>"
		case token.FALSE :
			return "<bool>"
		case token.IDENT :
			return "'" + tok.Literal + "'"	
	}
	return "'" + tok.Literal + "'"
}



func DescribeOpposite(tok token.Token) string {
	switch tok.Literal {
	case "<-|" : { return "indent" }
	case ")" : { return "'('" }
	case "]" : { return "[" }
	case "}" : { return "{" }
	}
	return "You goofed, that doesn't have an opposite."
}

var (
	RESET  = "\033[0m"
	RED    = "\033[31m"
	GREEN  = "\033[32m"
	YELLOW = "\033[33m"
	BLUE  = "\033[34m"
	PURPLE = "\033[35m"
	CYAN   = "\033[36m"
	GRAY   = "\033[37m"
	WHITE  = "\033[97m"

	ERROR = "$Error$" 
	RT_ERROR = "$Runtime error$"
	HUB_ERROR = "$Hub error$"
	OK = Green("ok")
)


func HighlightLine (plainLine string, highlighter rune) string {
	// Now we highlight the line. The rules are: anything enclosed in '   ' is code and is 
	// therefore highlighted, i.e. 'foo' serves the same function as writing foo in a monotype 
	// font would in a textbook or manual.

	// Because it looks kind of odd and redundant to write '"foo"' and '<foo>',  these are also
	// highlighted without requiring '.

	// The ' doesn't trigger the highlighting unless it follows a line beginning or space etc, because it
	// might be an apostrophe.
	
	highlitLine := ""
	prevCh := ' '
	if highlighter != ' ' {
		highlitLine = CYAN
	}

	for _, ch := range(plainLine) {
		if highlighter == ' ' && (((prevCh == ' ' || prevCh == '\n' || prevCh == '$') &&
				/**/(ch == '\'' || ch == '"' || ch == '<' || ch == '$') || ch == '@')) {
			highlighter = ch
			if highlighter == '<' { highlighter = '>' }
			if highlighter == '$' {
				highlitLine = highlitLine + RED
				continue
			}
			if highlighter == '@' {
				highlitLine = highlitLine + " " + YELLOW
				continue
			}
			highlitLine = highlitLine + CYAN	
			} else {
				if ch == highlighter {
					prevCh = ch
					highlighter = ' '

					if ch == '$' {
						highlitLine = highlitLine + RESET + ": " 
						continue
					}
					if ch == '@' {
						highlitLine = highlitLine + " " + RESET
						continue
					}
					highlitLine = highlitLine + string(ch) + RESET
					continue
				}
			}
			prevCh = ch
			highlitLine = highlitLine + string(ch)
		}
	return highlitLine
}