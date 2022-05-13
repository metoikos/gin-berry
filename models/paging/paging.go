// Package paging creates pagination object based on the number of records and other configs.
package paging

import (
	"gin-berry/utils"
	"math"
)

// Paging struct defines the pagination variables.
type Paging struct {
	Count        int64 `json:"count"`
	Pages        []int `json:"pages"`
	CurrentPage  int64 `json:"currentPage"`
	PreviousPage int64 `json:"previousPage"`
	NextPage     int64 `json:"nextPage"`
	LastPage     int64 `json:"lastPage"`
	PageSize     int64 `json:"pageSize"`
	Offset       int64 `json:"offset"`
	Enabled      bool  `json:"enabled"`
}

// New generates a pagination from number of items, total items to be displayed and the current page we're in.
func New(currentPage int64, pageSize int64, totalItems int64) Paging {
	var offset, page int64
	// initial current page
	page = 1

	totalPages := int64(math.Ceil(float64(totalItems) / float64(pageSize)))

	// calculate an offset only if the total page size more than 1
	if totalPages > 1 {
		offset = pageSize * (currentPage - 1)
	}

	// check total available pages if greater than we received as parameter
	if totalPages >= currentPage {
		page = currentPage
	}

	previousPage, nextPage := getNextPrevPages(totalPages, currentPage)

	return Paging{
		CurrentPage:  page,
		NextPage:     nextPage,
		PreviousPage: previousPage,
		Pages:        utils.MakeRange(1, int(totalPages)),
		Count:        totalItems,
		Offset:       offset,
		LastPage:     totalPages,
		PageSize:     pageSize,
		Enabled:      totalItems > pageSize,
	}
}

// getNextPrevPages method calculates the next and the previous pages from the variables.
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
