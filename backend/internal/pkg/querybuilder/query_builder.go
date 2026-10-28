package querybuilder

import (
	"fmt"
	"math"
	"net/url"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

// Options contains reuseables query parameters for list endpoint
type Options struct {
	Page      int
	Limit     int
	SortBy    string
	SortOrder string
	Search    string
}

// Pagination contains metadate for paginated API request
type Pagination struct {
	Page int `json:"page"`
	Limit int `json:"limit"`
	Total int 	`json:"total"`
	TotalPages int `json:"total_pagess"`
}

// Default and maximum pagination limits prevent excessively large queries.
const (
	DefaultPage = 1
	DefaultLimit = 10
	MaxLimit = 100
)

func Parse(values url.Values, allowedSortFields []string) Options {
	page, err := strconv.Atoi(values.Get("page"))
	if err != nil  || page < 1 {
		page = DefaultPage
	}

	limit, err := strconv.Atoi(values.Get("limit"))

	if err != nil || limit < 1 {
		limit = DefaultLimit
	}

	if limit > MaxLimit {
		limit = MaxLimit
	}


	search := strings.TrimSpace(values.Get("search"))

	return Options{
		Page: page,
		Limit: limit,
		Search: search,
	}
}

// ApplySorting applies sorting only to an approved column.
// The column name must come from the server-side allowlist.
func ApplySorting(
	db *gorm.DB,
	options Options,
	defaultSort string,
	allowedSortFields []string,
) *gorm.DB {
	sortBy := options.SortBy

	if !contains(allowedSortFields, sortBy) {
		sortBy = defaultSort
	}

	if !contains(allowedSortFields, sortBy) {
		return db
	}

	// Quote the column identifier using GORM's database dialect.
	column := db.Statement.Quote(sortBy)
	direction := strings.ToUpper(options.SortOrder)

	if direction != "ASC" && direction != "DESC" {
		direction = "DESC"
	}

	return db.Order(column + " " + direction)
}

// ApplySearch searches across approved columns using parametarized values
func ApplySearch(db *gorm.DB, search string, searchableFields []string) *gorm.DB {
	search = strings.TrimSpace(search)

	if search == "" || len(searchableFields) == 0 {
		return db
	}

	pattern := "%" + search + "%"
	conditions := make([]string, 0, len(searchableFields))
	args := make([]any, 0, len(searchableFields))

	for _, field := range searchableFields {
		if field == "" {
			continue
		}

		// Quote each server-defined column identifier.
		conditions = append(conditions,db.Statement.Quote(field)+" LIKE ?")
		args = append(args, pattern)
	}


	if len(conditions) == 0 {
		return db
	}

	return db.Where("("+strings.Join(conditions, " OR ")+")",args...)
}

// ApplyPagination adds LIMIT and OFFSET to the query
func ApplyPagination(db *gorm.DB, options Options) *gorm.DB {
	page := options.Page
	limit := options.Limit

	if page < 1 {
		page = DefaultPage
	}

	if limit < 1 {
		limit = DefaultLimit
	}

	if limit > MaxLimit {
		limit = MaxLimit
	}

	offset := (page - 1)*limit

	return db.Limit(limit).Offset(offset)
}

// contains checks whether a string is present in an allowlist.
func contains(fields []string, value string) bool {
	for _, field := range fields {
		if field == value {
			return true
		}
	}
	return false
}

// ApplyFilter adds an equality filter using a server-defined column.
func ApplyFilter(db *gorm.DB,column string,value interface{},allowedFields []string) *gorm.DB {
	if !contains(allowedFields, column) {
		return db
	}

	return db.Where(
		fmt.Sprintf("%s = ?", db.Statement.Quote(column)),
		value,
	)
}

// BuildPagination creates response metadata from the total record count.
func BuildPagination(page, limit int, total int64) Pagination {
	if page < 1 {
		page = DefaultPage
	}
	
	if limit < 1 {
		limit = DefaultLimit
	}

	totalPages := 0
	if total > 0 {
		totalPages = int(math.Ceil(float64(total) / float64(limit)))
	}

	return Pagination{
		Page:       page,
		Limit:      limit,
		Total:      int(total),
		TotalPages: int(totalPages),
	}
}