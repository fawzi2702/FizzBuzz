package handlers

import (
	"net/http"

	"github.com/fawzi2702/FizzBuzz/api/helpers/response"
	"github.com/fawzi2702/FizzBuzz/api/modules/fizzbuzz"
	fizzbuzzParameters "github.com/fawzi2702/FizzBuzz/api/modules/fizzbuzz/parameters"
	"github.com/fawzi2702/FizzBuzz/api/modules/stats"
	"github.com/gin-gonic/gin"
)

func Fizzbuzz(c *gin.Context) {
	p, err := fizzbuzzParameters.ParseParams(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}

	res := fizzbuzz.Fizzbuzz(p)

	// Save stats
	if err := stats.SaveRequestStat(p); err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.Ok(c, res)
}
