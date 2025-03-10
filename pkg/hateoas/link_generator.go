package hateoas

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

// ResourceConfig stores the link configuration for a resource type
type ResourceConfig struct {
	ResourceName     string            // Resource name (e.g., "transactions", "users")
	IDExtractor      func(any) string  // Function to extract the resource ID
	DefaultLinkTypes []string          // Default link types to be generated (e.g., "self", "collection")
	CustomLinks      map[string]string // Resource-specific custom links (key: type, value: URL pattern)
	PaginationLinks  []string          // Link types that should include pagination parameters
}

// LinkGenerator is responsible for generating HATEOAS links
type LinkGenerator struct {
	configs      map[string]ResourceConfig                   // Map of configurations by resource type
	apiBasePath  string                                      // API base path (e.g., "/api/v1")
	defaultLinks map[string]func(string, string, any) string // Default links available for all resources
}

// NewLinkGenerator creates a new link generator
func NewLinkGenerator(apiBasePath string) *LinkGenerator {
	lg := &LinkGenerator{
		configs:     make(map[string]ResourceConfig),
		apiBasePath: apiBasePath,
		defaultLinks: map[string]func(string, string, any) string{
			"self":       selfLinkFunc,
			"collection": collectionLinkFunc,
			"create":     createLinkFunc,
			"show":       showLinkFunc,
			"update":     updateLinkFunc,
			"delete":     deleteLinkFunc,
		},
	}
	return lg
}

// RegisterResource registers a new resource type with the generator
func (lg *LinkGenerator) RegisterResource(resourceType string, config ResourceConfig) {
	// Set default pagination links if not provided
	if len(config.PaginationLinks) == 0 {
		config.PaginationLinks = []string{"self", "collection"}
	}
	lg.configs[resourceType] = config
}

// LinkBuilder builds HATEOAS links for a specific resource
type LinkBuilder struct {
	generator    *LinkGenerator
	resourceType string
	resource     any
	baseURL      string
	ctx          *gin.Context
	links        map[string]string
	overrides    map[string]string
	queryParams  map[string]string
}

// For starts building links for a specific resource
func (lg *LinkGenerator) For(resourceType string, resource any, ctx *gin.Context) *LinkBuilder {
	baseURL := getBaseURL(ctx, lg.apiBasePath)
	return &LinkBuilder{
		generator:    lg,
		resourceType: resourceType,
		resource:     resource,
		baseURL:      baseURL,
		ctx:          ctx,
		links:        make(map[string]string),
		overrides:    make(map[string]string),
		queryParams:  make(map[string]string),
	}
}

// WithQueryParams adds query parameters for all links
func (lb *LinkBuilder) WithQueryParams(params map[string]string) *LinkBuilder {
	for k, v := range params {
		lb.queryParams[k] = v
	}
	return lb
}

// Override allows replacing a specific link
func (lb *LinkBuilder) Override(linkType string, url string) *LinkBuilder {
	lb.overrides[linkType] = url
	return lb
}

// Build constructs all links for the resource
func (lb *LinkBuilder) Build() map[string]string {
	config, exists := lb.generator.configs[lb.resourceType]
	if !exists {
		return lb.links
	}

	// Generate default links for the resource type
	for _, linkType := range config.DefaultLinkTypes {
		if linkFunc, ok := lb.generator.defaultLinks[linkType]; ok {
			url := linkFunc(lb.baseURL, config.ResourceName, lb.resource)
			// Only save non-empty URLs
			if url != "" {
				lb.links[linkType] = url
			}
		}
	}

	// Add custom links for the resource type
	for linkType, pattern := range config.CustomLinks {
		id := ""
		if config.IDExtractor != nil && lb.resource != nil {
			id = config.IDExtractor(lb.resource)
		}
		url := pattern
		url = strings.Replace(url, "{baseURL}", lb.baseURL, -1)
		url = strings.Replace(url, "{resourceName}", config.ResourceName, -1)
		url = strings.Replace(url, "{id}", id, -1)
		lb.links[linkType] = url
	}

	// Apply overrides
	for linkType, url := range lb.overrides {
		lb.links[linkType] = url
	}

	// Add query parameters only to pagination-enabled links
	if len(lb.queryParams) > 0 {
		queryString := "?"
		first := true
		for k, v := range lb.queryParams {
			if !first {
				queryString += "&"
			}
			queryString += fmt.Sprintf("%s=%s", k, v)
			first = false
		}

		// Add only to pagination-enabled links
		for linkType, url := range lb.links {
			if shouldApplyPagination(linkType, config.PaginationLinks) && !strings.Contains(url, "?") {
				lb.links[linkType] = url + queryString
			}
		}
	}

	return lb.links
}

// Helper function to determine if pagination should be applied to a link
func shouldApplyPagination(linkType string, paginationLinks []string) bool {
	for _, paginationType := range paginationLinks {
		if paginationType == linkType {
			return true
		}
	}
	return false
}

// Standard functions to generate links

func selfLinkFunc(baseURL string, resourceName string, resource any) string {
	if resource == nil {
		return fmt.Sprintf("%s/%s", baseURL, resourceName)
	}

	id := extractID(resource)
	if id != "" {
		return fmt.Sprintf("%s/%s/%s", baseURL, resourceName, id)
	}
	return fmt.Sprintf("%s/%s", baseURL, resourceName)
}

func collectionLinkFunc(baseURL string, resourceName string, _ any) string {
	return fmt.Sprintf("%s/%s", baseURL, resourceName)
}

func showLinkFunc(baseURL string, resourceName string, resource any) string {
	if resource == nil {
		return ""
	}

	id := extractID(resource)
	if id != "" {
		return fmt.Sprintf("%s/%s/%s", baseURL, resourceName, id)
	}
	return ""
}

func createLinkFunc(baseURL string, resourceName string, _ any) string {
	return fmt.Sprintf("%s/%s", baseURL, resourceName)
}

func updateLinkFunc(baseURL string, resourceName string, resource any) string {
	// Only generate update links for specific resources
	if resource == nil {
		return ""
	}

	id := extractID(resource)
	if id != "" {
		return fmt.Sprintf("%s/%s/%s", baseURL, resourceName, id)
	}
	return ""
}

func deleteLinkFunc(baseURL string, resourceName string, resource any) string {
	// Only generate delete links for specific resources
	if resource == nil {
		return ""
	}

	id := extractID(resource)
	if id != "" {
		return fmt.Sprintf("%s/%s/%s", baseURL, resourceName, id)
	}
	return ""
}

// Generic ID extractor - tries several common methods
func extractID(resource any) string {
	if resource == nil {
		return ""
	}

	val := reflect.ValueOf(resource)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// Try various common approaches to extract ID

	// 1. Try ID() method
	idMethod := val.MethodByName("ID")
	if idMethod.IsValid() {
		result := idMethod.Call(nil)
		if len(result) > 0 {
			return fmt.Sprintf("%v", result[0].Interface())
		}
	}

	// 2. Look for ID field
	if val.Kind() == reflect.Struct {
		idField := val.FieldByName("ID")
		if idField.IsValid() {
			return fmt.Sprintf("%v", idField.Interface())
		}
	}

	// 3. Try Id or id as well
	otherNames := []string{"Id", "id"}
	for _, name := range otherNames {
		// Try method
		idMethod := val.MethodByName(name)
		if idMethod.IsValid() {
			result := idMethod.Call(nil)
			if len(result) > 0 {
				return fmt.Sprintf("%v", result[0].Interface())
			}
		}

		// Try field
		if val.Kind() == reflect.Struct {
			idField := val.FieldByName(name)
			if idField.IsValid() {
				return fmt.Sprintf("%v", idField.Interface())
			}
		}
	}

	return ""
}

// getBaseURL returns the base URL for HATEOAS links
func getBaseURL(ctx *gin.Context, apiBasePath string) string {
	scheme := "http"
	if ctx.Request.TLS != nil {
		scheme = "https"
	}
	return fmt.Sprintf("%s://%s%s", scheme, ctx.Request.Host, apiBasePath)
}
