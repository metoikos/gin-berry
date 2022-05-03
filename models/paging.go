package models

import (
	"gin-berry/lib"
	"math"
)

type Paging struct {
	count        int64
	pages        []int
	currentPage  int64
	previousPage int64
	nextPage     int64
	lastPage     int64
	pageSize     int64
	offset       int64
	enabled      bool
}

func BuildPaging(currentPage int64, pageSize int64, totalItems int64) *Paging {
	var offset, page int64
	// initial current page
	page = 1

	totalPages := int64(math.Ceil(float64(totalItems) / float64(pageSize)))

	// calculate an offset only if the total page size more than 1
	if totalPages > 1 {
		offset = pageSize * (currentPage - 1)
	}

	// check total available pages is greater than we received as parameter
	if totalPages >= currentPage {
		page = currentPage
	}

	previousPage, nextPage := getNextPrevPages(totalPages, currentPage)

	return &Paging{
		currentPage:  page,
		nextPage:     nextPage,
		previousPage: previousPage,
		pages:        lib.MakeRange(1, int(totalPages)),
		count:        totalItems,
		offset:       offset,
		lastPage:     totalPages,
		pageSize:     pageSize,
		enabled:      totalItems > pageSize,
	}
}

func getNextPrevPages(totalPages int64, currentPage int64) (int64, int64) {
	var previousPage, nextPage int64

	if currentPage > totalPages || currentPage < 1 {
		return previousPage, nextPage
	}

	if currentPage > 1 {
		previousPage = currentPage - 1
	}

	if currentPage+1 <= totalPages {
		nextPage = currentPage + 1
	}

	return previousPage, nextPage
}
