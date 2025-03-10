package hateoas

import (
	"time"
)

// MetaData contains metadata about the response
type MetaData struct {
	Timestamp  time.Time `json:"timestamp"`
	StatusCode int       `json:"statusCode"`
}

// PageInfo contains pagination information
type PageInfo struct {
	PageSize   int `json:"pageSize"`
	PageNumber int `json:"pageNumber"`
	TotalItems int `json:"totalItems"`
	TotalPages int `json:"totalPages"`
}

// Response is a generic structure for API responses with HATEOAS
type Response struct {
	Data     interface{} `json:"data,omitempty"`
	Links    Links       `json:"_links,omitempty"`
	Meta     MetaData    `json:"meta"`
	PageInfo *PageInfo   `json:"pageInfo,omitempty"`
}

// NewResponse creates a new HATEOAS response
func NewResponse(data interface{}, statusCode int) *Response {
	return &Response{
		Data: data,
		Meta: MetaData{
			Timestamp:  time.Now(),
			StatusCode: statusCode,
		},
	}
}

// WithLinks adds HATEOAS links to the response
func (r *Response) WithLinks(links Links) *Response {
	r.Links = links
	return r
}

// WithLinksMap adds HATEOAS links to the response from a map of strings
func (r *Response) WithLinksMap(links map[string]string) *Response {
	r.Links = ToLinks(links)
	return r
}

// WithPageInfo adds pagination information to the response
func (r *Response) WithPageInfo(pageSize, pageNumber, totalItems int) *Response {
	totalPages := totalItems / pageSize
	if totalItems%pageSize > 0 {
		totalPages++
	}

	r.PageInfo = &PageInfo{
		PageSize:   pageSize,
		PageNumber: pageNumber,
		TotalItems: totalItems,
		TotalPages: totalPages,
	}
	return r
}

// Single generates a HATEOAS response for a single resource
func Single(resourceType string, resource interface{}, ctx interface{}, statusCode int) *Response {
	// For a single resource, we generate resource-specific links
	// This includes self, update, delete, but not collection-wide links with pagination
	links := GlobalInstance.GetLinksForResource(resourceType, resource, ctx)

	// Create response
	return NewResponse(resource, statusCode).WithLinksMap(links)
}

// Collection generates a HATEOAS response for a collection of resources
func Collection(resourceType string, resources interface{}, ctx interface{}, page, pageSize, totalItems, statusCode int) *Response {
	// For a collection, we generate collection-wide links with pagination support
	links := GlobalInstance.GetLinksForCollection(resourceType, ctx, page, pageSize)

	// Create response
	return NewResponse(resources, statusCode).
		WithLinksMap(links).
		WithPageInfo(pageSize, page, totalItems)
}
