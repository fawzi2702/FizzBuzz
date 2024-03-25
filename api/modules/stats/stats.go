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

func GetTopRequest() (*fizzbuzzParameters.Params, int, error) {
	topRequests, err := redisdb.Client.ZRevRangeWithScores(redisdb.Context, statsKey, 0, 0).Result()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get the top request: %w", err)
	}

	if len(topRequests) == 0 {
		return nil, 0, nil
	}

	topRequest := topRequests[0]

	score := int(topRequest.Score)
	key := topRequest.Member.(string)

	p, err := fizzbuzzParameters.ParseKey(key)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to parse the top request: %w", err)
	}

	return p, score, nil
}
