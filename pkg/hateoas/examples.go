package hateoas

/*
Examples of using the HATEOAS package

This file contains didactic examples of how to use the HATEOAS package.
It is not directly compiled into the project, it serves only as documentation.

1. Global initialization

To initialize the HATEOAS generator globally and register resources:

```go
func init() {
    // Initialize the generator with the API base path
    hateoas.GlobalInstance.Setup("/api/v1")

    // Register a user resource
    hateoas.GlobalInstance.RegisterResource("user", hateoas.ResourceConfig{
        ResourceName: "users",
        DefaultLinkTypes: []string{"self", "collection", "create", "update", "delete"},
        CustomLinks: map[string]string{
            "profile": "{baseURL}/profiles/{id}",
            "avatar":  "{baseURL}/users/{id}/avatar",
        },
    })

    // Register a category resource
    hateoas.GlobalInstance.RegisterResource("category", hateoas.ResourceConfig{
        ResourceName: "categories",
        DefaultLinkTypes: []string{"self", "collection"},
    })
}
```

2. Using in a Gin handler or controller

To generate links for a collection of resources:

```go
func (c *UserController) ListUsers(ctx *gin.Context) {
    page := 1 // get from query
    pageSize := 10 // get from query

    // Fetch users from service
    users, err := c.userService.FindAll()
    if err != nil {
        // error handling
    }

    // Convert to response
    usersResponse := convertToUserResponse(users)

    // Generate links for the collection
    links := hateoas.GlobalInstance.GetLinksForCollection("user", ctx, page, pageSize)
    hateoasLinks := hateoas.ToLinks(links)

    // Create response with HATEOAS
    response := ApiResponse{
        Data: usersResponse,
        Links: hateoasLinks,
        Meta: MetaData{
            Timestamp: time.Now(),
            StatusCode: http.StatusOK,
        },
        PageInfo: &PageInfo{
            PageSize: pageSize,
            PageNumber: page,
            TotalItems: len(users),
            TotalPages: calculateTotalPages(len(users), pageSize),
        },
    }

    ctx.JSON(http.StatusOK, response)
}
```

To generate links for a single resource:

```go
func (c *UserController) GetUser(ctx *gin.Context) {
    id := ctx.Param("id")

    // Fetch user from service
    user, err := c.userService.FindByID(id)
    if err != nil {
        // error handling
    }

    // Convert to response
    userResponse := convertToUserResponse(user)

    // Generate links for the resource
    links := hateoas.GlobalInstance.GetLinksForResource("user", userResponse, ctx)
    hateoasLinks := hateoas.ToLinks(links)

    // Create response with HATEOAS
    response := ApiResponse{
        Data: userResponse,
        Links: hateoasLinks,
        Meta: MetaData{
            Timestamp: time.Now(),
            StatusCode: http.StatusOK,
        },
    }

    ctx.JSON(http.StatusOK, response)
}
```

3. Override specific links

You can override specific links for special cases:

```go
links := hateoas.GlobalInstance.GetLinksForResource("user", userResponse, ctx)

// Add or override custom links
builder := hateoas.GlobalInstance.generator.For("user", userResponse, ctx)
builder.Override("activate", fmt.Sprintf("/api/v1/users/%s/activate", userResponse.ID))
links = builder.Build()

hateoasLinks := hateoas.ToLinks(links)
```

4. Use with different frameworks

The package can be adapted for different web frameworks besides Gin.
To do this, implement the RequestContext interface for your framework.
*/
