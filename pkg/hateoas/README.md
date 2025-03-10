# HATEOAS Package for Go

This package provides a clean and reusable implementation of the HATEOAS (Hypertext as the Engine of Application State) pattern for RESTful APIs in Go.

## Features

- Automatic generation of resource links
- Support for various link types (self, collection, create, update, delete)
- Flexible resource configuration
- Pagination support
- Builder pattern for constructing links
- Ability to override specific links
- Integration with Gin (extendable to other frameworks)

## Installation

```bash
go get github.com/yourusername/yourproject/pkg/hateoas
```

## Basic Usage

### Initialization

```go
// In init() or at the start of your application
hateoas.GlobalInstance.Setup("/api/v1")

// Register resources
hateoas.GlobalInstance.RegisterResource("transaction", hateoas.ResourceConfig{
    ResourceName: "transactions",
    DefaultLinkTypes: []string{"self", "collection", "create", "update", "delete"},
    CustomLinks: map[string]string{
        "report": "{baseURL}/{resourceName}/report",
    },
})
```

### In a Controller (Gin)

For a single resource:

```go
func GetTransaction(ctx *gin.Context) {
    // Fetch transaction from the service
    transaction, err := service.FindById(id)
    if err != nil {
        // handle error
    }
    
    // Convert to response
    transactionResponse := FromEntity(transaction)
    
    // Create HATEOAS response
    response := hateoas.Single("transaction", transactionResponse, ctx, http.StatusOK)
    
    ctx.JSON(http.StatusOK, response)
}
```

For a collection:

```go
func ListTransactions(ctx *gin.Context) {
    page := 1 // get from query
    pageSize := 10 // get from query
    
    // Fetch transactions from the service
    transactions, err := service.FindAll()
    if err != nil {
        // handle error
    }
    
    // Convert to response
    transactionsResponse := FromEntities(transactions)
    
    // Create HATEOAS response
    response := hateoas.Collection(
        "transaction", 
        transactionsResponse, 
        ctx, 
        page, 
        pageSize, 
        len(transactions), 
        http.StatusOK,
    )
    
    ctx.JSON(http.StatusOK, response)
}
```

## Customization

### Overriding Links

```go
links := hateoas.GlobalInstance.GetLinksForResource("transaction", transaction, ctx)

// Add a custom link
customURL := fmt.Sprintf("/api/v1/transactions/%s/print", transaction.ID)
builder := hateoas.GlobalInstance.generator.For("transaction", transaction, ctx)
builder.Override("print", customURL)
links = builder.Build()

// Convert to response format
hateoasLinks := hateoas.ToLinks(links)
```

### Adding a New Resource Type

```go
hateoas.GlobalInstance.RegisterResource("category", hateoas.ResourceConfig{
    ResourceName: "categories",
    DefaultLinkTypes: []string{"self", "collection"},
    CustomLinks: map[string]string{
        "products": "{baseURL}/products?category={id}",
    },
})
```

## Extension for Other Frameworks

To use with frameworks other than Gin, implement the `RequestContext` interface for your specific framework.

## Response Format

```json
{
  "data": {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "categoryId": "456e4567-e89b-12d3-a456-426614174000",
    "amount": 100.50,
    "description": "Payment for services"
  },
  "_links": {
    "self": {
      "href": "http://api.example.com/api/v1/transactions/123e4567-e89b-12d3-a456-426614174000"
    },
    "collection": {
      "href": "http://api.example.com/api/v1/transactions"
    },
    "update": {
      "href": "http://api.example.com/api/v1/transactions/123e4567-e89b-12d3-a456-426614174000"
    },
    "delete": {
      "href": "http://api.example.com/api/v1/transactions/123e4567-e89b-12d3-a456-426614174000"
    }
  },
  "meta": {
    "timestamp": "2023-07-16T14:05:10Z",
    "statusCode": 200
  }
}
```

## License

MIT 