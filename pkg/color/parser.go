package color

import (
	"regexp"
)

type ParserMap map[string]func(string) string

// Regexes
var HEX_REGEX = regexp.MustCompile(`(?i)#(?:[a-f0-9]{3}){1,2}`)
var RGB_REGEX = regexp.MustCompile(`rgb\(\s*(\d+)\s*,\s*(\d+)\s*,\s*(\d+)\s*\)`)

var parsers = map[*regexp.Regexp]ParserMap{
	HEX_REGEX: HexParserMap,
	RGB_REGEX: RgbParserMap,
}

func LineParser(toFormat string, l string) string {
	if len(toFormat) <= 0 || len(l) <= 0 {
		return l
	}

	var newLine string = l

	for rgx, parserMap := range parsers {
		if !rgx.MatchString(l) {
			continue
		}

		converter := parserMap[toFormat]
		newLine = rgx.ReplaceAllStringFunc(l, func(h string) string {
			c := rgx.FindString(l)
			return converter(c)
		})
	}

	return newLine
}
