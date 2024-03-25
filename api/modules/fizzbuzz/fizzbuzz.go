package fizzbuzz

import (
	"strconv"
	"strings"

	fizzbuzzParameters "github.com/fawzi2702/FizzBuzz/api/modules/fizzbuzz/parameters"
)

const startNb = 1

func Fizzbuzz(p *fizzbuzzParameters.Params) []string {
	var result string

	for i := startNb; i <= p.Limit; i++ {
		if i != startNb {
			result += p.Separator
		}

		if i%p.Int1 == 0 && i%p.Int2 == 0 {
			result += p.Str1 + p.Str2
		} else if i%p.Int1 == 0 {
			result += p.Str1
		} else if i%p.Int2 == 0 {
			result += p.Str2
		} else {
			result += strconv.Itoa(i)
		}
	}

	return strings.Split(result, p.Separator)
}
