package lib

import (
	"regexp"
)

func IsValidDNA(input string) bool {
	var regex *regexp.Regexp = regexp.MustCompile("[^AGCT]")
	return len(input) > 0 && !(regex.MatchString(input))
}

func IsValidQuery(input string) int {

	var regexDate31 *regexp.Regexp = regexp.MustCompile(`^(0?[1-9]|[12][1-9]|[123]0|31)[-/](0?[13578]|10|12)[-/]((202[0-2])|(20[01])[0-9]|(1[0-9]{1,3})|([0-9]{1,3}))$`)
	var regexDate30 *regexp.Regexp = regexp.MustCompile(`^(0?[1-9]|[12][1-9]|[123]0)[-/](0?[469]|11)[-/]((202[0-2])|(20[01])[0-9]|(1[0-9]{1,3})|([0-9]{1,3}))$`)
	var regexDate29 *regexp.Regexp = regexp.MustCompile(`^(0?[1-9]|[12][1-9]|[12]0)[-/](0?2)[-/]((202[0-2])|(20[01])[0-9]|(1[0-9]{1,3})|([0-9]{1,3}))$`)
	var regexDisease *regexp.Regexp = regexp.MustCompile(`^[a-zA-Z0-9\s]+$`)
	var regexFull *regexp.Regexp = regexp.MustCompile(`^(((0?[1-9]|[12][1-9]|[123]0|31)[-/](0?[13578]|10|12)[-/]((202[0-2])|(20[01])[0-9]|(1[0-9]{1,3})|([0-9]{1,3})))|((0?[1-9]|[12][1-9]|[123]0)[-/](0?[469]|11)[-/]((202[0-2])|(20[01])[0-9]|(1[0-9]{1,3})|([0-9]{1,3})))|((0?[1-9]|[12][1-9]|[12]0)[-/](0?2)[-/]((202[0-2])|(20[01])[0-9]|(1[0-9]{1,3})|([0-9]{1,3}))))\s[a-zA-Z0-9\s]+$`)

	var flag int = -1

	// fmt.Println(regexDate31.FindAllString(input, -1))
	// fmt.Println(regexDate30.FindAllString(input, -1))
	// fmt.Println(regexDate29.FindAllString(input, -1))
	// fmt.Println(regexDisease.FindAllString(input, -1))
	// fmt.Println(regexFull.FindAllString(input, -1))

	if regexDate31.MatchString(input) || regexDate30.MatchString(input) || regexDate29.MatchString(input) {
		flag = 0 // Date
	} else if regexDisease.MatchString(input) {
		flag = 1 // Disease
	} else if regexFull.MatchString(input) {
		flag = 2 // Both
	}

	return flag

}
