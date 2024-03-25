package handlers

import (
	"net/http"

	"github.com/fawzi2702/FizzBuzz/api/helpers/response"
	"github.com/fawzi2702/FizzBuzz/api/modules/fizzbuzz"
	fizzbuzzParameters "github.com/fawzi2702/FizzBuzz/api/modules/fizzbuzz/parameters"
	"github.com/fawzi2702/FizzBuzz/api/modules/stats"
	"github.com/gin-gonic/gin"
)

// Fizzbuzz godoc
// @Summary Get the fizzbuzz result
// @Description Get the fizzbuzz result
// @Tags fizzbuzz
// @Accept json
// @Produce json
// @Param int1 query int true "Multiplicator for str1" minimum(1)
// @Param int2 query int true "Multiplicator for str2" minimum(1)
// @Param limit query int true "Limit number" minimum(1)
// @Param str1 query string true "First string, to be displayed when the number is a multiple of int1" minLength(1)
// @Param str2 query string true "Second string, to be displayed when the number is a multiple of int2" minLength(1)
// @Success 200 {object} response.OkResponse{data=[]string}
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /fizzbuzz [get]
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
