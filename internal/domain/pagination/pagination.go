package pagination

// Pagination represents pagination parameters and metadata
type Pagination struct {
	Page       int   // Current page number (1-indexed)
	PageSize   int   // Number of items per page
	TotalItems int64 // Total number of items across all pages
	TotalPages int   // Total number of pages
	HasMore    bool  // Whether there are more pages after current page
}

// NewPagination creates a new pagination instance with sensible defaults
func NewPagination(page, pageSize int) *Pagination {
	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	// Limit maximum page size
	if pageSize > 100 {
		pageSize = 100
	}

	return &Pagination{
		Page:     page,
		PageSize: pageSize,
	}
}

// SetTotal sets the total count and calculates derived values
func (p *Pagination) SetTotal(totalItems int64) {
	p.TotalItems = totalItems
	p.TotalPages = int((totalItems + int64(p.PageSize) - 1) / int64(p.PageSize))
	p.HasMore = p.Page < p.TotalPages
}

// GetOffset returns the SQL offset value for the current page
func (p *Pagination) GetOffset() int {
	return (p.Page - 1) * p.PageSize
}

// GetLimit returns the SQL limit value
func (p *Pagination) GetLimit() int {
	return p.PageSize
}
