package common

import (
	"math"

	"github.com/gin-gonic/gin"
)

// Paginate handles the pagination logic directly from the request context.
func Paginate(total int, data []interface{}, currentPage int, perPage int) interface{} {
	totalPages := int(math.Ceil(float64(total) / float64(perPage)))

	if currentPage > totalPages {
		currentPage = totalPages
	}

	nextPage := currentPage + 1
	if nextPage > totalPages {
		nextPage = totalPages
	}

	prevPage := currentPage - 1
	if prevPage < 1 {
		prevPage = 1
	}
	// Create pagination result
	return gin.H{
		"perPage":     perPage,
		"currentPage": currentPage,
		"nextPage":    nextPage,
		"prevPage":    prevPage,
		"totalPages":  totalPages,
		"total":       total,
		"data":        data,
	}

}
