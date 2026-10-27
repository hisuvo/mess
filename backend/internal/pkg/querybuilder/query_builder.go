package querybuilder

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
	TotalPage int `json:"total_page"`
}

// Default and maximum pagination limits prevent excessively large queries.
const (
	DefaultPage = 1
	DefaultLimit = 10
	MaxLimit = 100
)

