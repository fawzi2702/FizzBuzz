package stats

import (
	"fmt"

	redisdb "github.com/fawzi2702/FizzBuzz/api/helpers/redis-db"
	fizzbuzzParameters "github.com/fawzi2702/FizzBuzz/api/modules/fizzbuzz/parameters"
)

const statsKey = "fizzbuzz:stats-set"

func SaveRequestStat(p *fizzbuzzParameters.Params) error {
	k := p.GetKey()

	// Increment the number of hits for this key
	_, err := redisdb.Client.ZIncrBy(redisdb.Context, statsKey, 1, k).Result()
	if err != nil {
		return fmt.Errorf("failed to increment the number of hits for key %s: %w", k, err)
	}

	return nil
}
