package fizzbuzzParameters

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const Separator = "%%___%%"

type Params struct {
	Int1      int
	Int2      int
	Limit     int
	Str1      string
	Str2      string
	Separator string
}

func ParseParams(c *gin.Context) (*Params, error) {
	params := &Params{}

	var err error

	if params.Int1, err = strconv.Atoi(c.Query("int1")); err != nil {
		return nil, errors.New("Int1 must be a valid integer")
	}

	if params.Int2, err = strconv.Atoi(c.Query("int2")); err != nil {
		return nil, errors.New("Int2 must be a valid integer")
	}

	if params.Limit, err = strconv.Atoi(c.Query("limit")); err != nil {
		return nil, errors.New("limit must be a valid integer")
	}

	params.Str1 = c.Query("str1")
	params.Str2 = c.Query("str2")

	if ok, err := params.validate(); !ok {
		return nil, err
	}

	return params, nil
}

func (p *Params) validate() (bool, error) {
	if err := p.validateInt1(); err != nil {
		return false, err
	}

	if err := p.validateInt2(); err != nil {
		return false, err
	}

	if err := p.validateLimit(); err != nil {
		return false, err
	}

	if err := p.validateStr1(); err != nil {
		return false, err
	}

	if err := p.validateStr2(); err != nil {
		return false, err
	}

	return true, nil
}

func (p *Params) validateInt1() error {
	if p.Int1 <= 0 {
		return errors.New("Int1 must be greater than 0")
	}

	return nil
}

func (p *Params) validateInt2() error {
	if p.Int2 <= 0 {
		return errors.New("int2 must be greater than 0")
	} else if p.Int2 == p.Int1 {
		return errors.New("int2 must be different from Int1")
	}

	return nil
}

func (p *Params) validateLimit() error {
	if p.Limit <= 0 {
		return errors.New("limit must be greater than 0")
	}

	return nil
}

func (p *Params) validateStr1() error {
	if p.Str1 == "" {
		return errors.New("str1 must not be empty")
	} else if p.Str1 == Separator {
		return errors.New("str1 must not be equal to " + Separator)
	}

	return nil
}

func (p *Params) validateStr2() error {
	if p.Str2 == "" {
		return errors.New("str2 must not be empty")
	} else if p.Str2 == Separator {
		return errors.New("str2 must not be equal to " + Separator)
	} else if strings.Contains(p.Str1+p.Str2, Separator) {
		return errors.New("str1 and str2 must not contain " + Separator)
	}

	return nil
}

func (p *Params) GetKey() string {
	return strings.Join([]string{
		strconv.Itoa(p.Int1),
		strconv.Itoa(p.Int2),
		strconv.Itoa(p.Limit),
		p.Str1,
		p.Str2,
	}, Separator)
}