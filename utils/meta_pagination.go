package utils

import "strconv"

type QueryParams struct {
	Page    string `query:"page"`
	PerPage string `query:"perPage"`
	Sort    string `query:"sort"`
	Search  string `query:"search"`
	Status  string `query:"status"`
}

type Meta struct {
	CurrentPage      int `json:"currentPage"`
	PerPage          int `json:"perPage"`
	TotalCurrentPage int `json:"totalCurrentPage"`
	TotalPage        int `json:"totalPage"`
	TotalData        int `json:"totalData"`
	RangeStart       int `json:"rangeStart"`
	RangeEnd         int `json:"rangeEnd"`
}

func MetaPagination(
	page int,
	perPage int,
	totalCurrentPage int,
	total int,
) Meta {
	totalPage := (total + perPage - 1) / perPage

	rangeStart := 1
	if page > 1 {
		rangeStart = (page-1)*perPage + 1
	}

	rangeEnd := total
	if totalCurrentPage == perPage {
		rangeEnd = page * totalCurrentPage
	}

	return Meta{
		CurrentPage:      page,
		PerPage:          perPage,
		TotalCurrentPage: totalCurrentPage,
		TotalPage:        totalPage,
		TotalData:        total,
		RangeStart:       rangeStart,
		RangeEnd:         rangeEnd,
	}
}

func GetPaginationParams(pageStr, perPageStr string) (int, int) {
	const defaultPage = 1
	const defaultPerPage = 10

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = defaultPage
	}

	perPage, err := strconv.Atoi(perPageStr)
	if err != nil || perPage <= 0 {
		perPage = defaultPerPage
	}

	return page, perPage
}
