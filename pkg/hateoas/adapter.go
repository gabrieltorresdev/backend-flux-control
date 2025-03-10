package hateoas

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Link is the structure for HATEOAS links
type Link struct {
	Href string `json:"href"`
	Rel  string `json:"rel,omitempty"`
}

// Links is a map of HATEOAS links
type Links map[string]Link

// ToLinks converts a map of string links to the Links format
func ToLinks(links map[string]string) Links {
	responseLinks := make(Links)
	for rel, href := range links {
		responseLinks[rel] = Link{
			Href: href,
		}
	}
	return responseLinks
}

// RequestContext is an interface to abstract the HTTP request context
type RequestContext interface {
	GetScheme() string
	GetHost() string
}

// ContextAdapter adapts a specific context to the RequestContext interface
type ContextAdapter struct {
	ctx interface{}
}

// GinContextAdapter adapts a Gin context to RequestContext
func (ca *ContextAdapter) GinAdapter() RequestContext {
	if ginCtx, ok := ca.ctx.(*gin.Context); ok {
		return &GinRequestContext{ginCtx: ginCtx}
	}
	return nil
}

// GinRequestContext implements RequestContext for Gin
type GinRequestContext struct {
	ginCtx *gin.Context
}

// GetScheme returns the HTTP scheme used (http/https)
func (c *GinRequestContext) GetScheme() string {
	if c.ginCtx.Request.TLS != nil {
		return "https"
	}
	return "http"
}

// GetHost returns the request host
func (c *GinRequestContext) GetHost() string {
	return c.ginCtx.Request.Host
}

// Global is a singleton to manage the global generator
type Global struct {
	generator *LinkGenerator
}

// GlobalInstance is the global instance of the generator
var GlobalInstance = &Global{}

// Setup initializes the global generator
func (g *Global) Setup(apiBasePath string) {
	g.generator = NewLinkGenerator(apiBasePath)
}

// RegisterResource registers a resource in the global generator
func (g *Global) RegisterResource(resourceType string, config ResourceConfig) {
	if g.generator != nil {
		g.generator.RegisterResource(resourceType, config)
	}
}

// GetLinksForCollection generates links for a collection using the global generator
func (g *Global) GetLinksForCollection(resourceType string, ctx interface{}, page, pageSize int) map[string]string {
	if g.generator == nil {
		return nil
	}

	// Verify and convert the context
	ginCtx, ok := ctx.(*gin.Context)
	if !ok {
		return nil
	}

	// For collections, we should only include pagination-related links
	// and collection-level operations (like create)
	params := map[string]string{
		"page":      fmt.Sprintf("%d", page),
		"page_size": fmt.Sprintf("%d", pageSize),
	}

	// For a collection, we don't pass any specific resource
	return g.generator.For(resourceType, nil, ginCtx).
		WithQueryParams(params).
		Build()
}

// GetLinksForResource generates links for a resource using the global generator
func (g *Global) GetLinksForResource(resourceType string, resource any, ctx interface{}) map[string]string {
	if g.generator == nil {
		return nil
	}

	ginCtx, ok := ctx.(*gin.Context)
	if !ok {
		return nil
	}

	// For a specific resource, we provide the resource so proper IDs can be extracted
	return g.generator.For(resourceType, resource, ginCtx).Build()
}
