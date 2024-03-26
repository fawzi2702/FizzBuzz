package fizzbuzz

import (
	"strconv"

	fizzbuzzParameters "github.com/fawzi2702/FizzBuzz/api/modules/fizzbuzz/parameters"
)

const startNb = 1

func Fizzbuzz(p *fizzbuzzParameters.Params) []string {
	size := p.Limit - startNb + 1
	result := make([]string, size)

	for i := startNb; i <= p.Limit; i++ {
		idx := i - startNb
		if i%p.Int1 == 0 && i%p.Int2 == 0 {
			result[idx] = p.Str1 + p.Str2
		} else if i%p.Int1 == 0 {
			result[idx] = p.Str1
		} else if i%p.Int2 == 0 {
			result[idx] = p.Str2
		} else {
			result[idx] = strconv.Itoa(i)
		}
	}

	return result
}
