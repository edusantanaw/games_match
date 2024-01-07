package utils

type IPagination struct {
	Take   int `form:"take"`
	Offset int `form:"offset"`
}

type PaginationResponse[T comparable] struct {
	Data  T
	Total int64
}
