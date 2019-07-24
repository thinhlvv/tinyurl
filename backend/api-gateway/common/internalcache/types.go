package internalcache

import (
	"strconv"
)

const ORDER_NUMBER_KEY string = "order_number"

type CacheValueType struct {
	value []byte
}

func (cv *CacheValueType) String() string {
	return string(cv.value)
}

func (cv *CacheValueType) Int() (int, error) {
	return strconv.Atoi(string(cv.value))
}
