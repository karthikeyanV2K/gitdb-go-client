# GitDB Go Client

A Go client library for interacting with GitDB - GitHub-backed NoSQL database.

## Installation

```bash
go get github.com/karthikeyanV2K/gitdb-client
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"
    "github.com/karthikeyanV2K/gitdb-client/gitdb"
)

func main() {
    // Create a new GitDB client
    client := gitdb.NewClient("your_github_token", "your_username", "your_repo")

    // Health check
    if err := client.Health(); err != nil {
        log.Fatalf("Health check failed: %v", err)
    }

    // Create a collection
    if err := client.CreateCollection("users"); err != nil {
        log.Fatalf("Failed to create collection: %v", err)
    }

    // Insert a document
    user := gitdb.Document{
        "name":  "John Doe",
        "email": "john@example.com",
        "age":   30,
        "active": true,
    }

    docID, err := client.Insert("users", user)
    if err != nil {
        log.Fatalf("Failed to insert document: %v", err)
    }
    fmt.Printf("Inserted document with ID: %s\n", docID)
}
```

## API Reference

### Client Creation

```go
// Create client with default URL (http://localhost:7896)
client := gitdb.NewClient(token, owner, repo)

// Create client with custom URL
client := gitdb.NewClient(token, owner, repo)
client.SetBaseURL("http://your-server:7896")
```

### Collections

```go
// Create a collection
err := client.CreateCollection("users")

// List all collections
collections, err := client.ListCollections()

// Delete a collection
err := client.DeleteCollection("users")
```

### Documents

```go
// Insert a document
doc := gitdb.Document{
    "name": "John Doe",
    "email": "john@example.com",
    "age": 30,
}
docID, err := client.Insert("users", doc)

// Find documents
query := gitdb.Query{"age": gitdb.Query{"$gte": 25}}
documents, err := client.Find("users", query)

// Find one document
doc, err := client.FindOne("users", gitdb.Query{"email": "john@example.com"})

// Find by ID
doc, err := client.FindById("users", docID)

// Update a document
update := gitdb.Update{
    "$set": gitdb.Document{"age": 31},
}
err := client.Update("users", docID, update)

// Update many documents
modifiedCount, err := client.UpdateMany("users", query, update)

// Delete a document
err := client.Delete("users", docID)

// Delete many documents
deletedCount, err := client.DeleteMany("users", query)

// Count documents
count, err := client.Count("users", query)
```

### GraphQL

```go
query := `
    query {
        collections
        documents(collection: "users") {
            _id
            name
            email
        }
    }
`

response, err := client.GraphQL(query, nil)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("GraphQL response: %+v\n", response.Data)
```

## Error Handling

The client uses Go's standard error handling pattern:

```go
docID, err := client.Insert("users", document)
if err != nil {
    // Handle error
    log.Printf("Insert failed: %v", err)
    return
}
```

## Query Operators

The client supports MongoDB-style query operators:

```go
// Greater than or equal
query := gitdb.Query{"age": gitdb.Query{"$gte": 25}}

// Less than
query := gitdb.Query{"age": gitdb.Query{"$lt": 50}}

// In array
query := gitdb.Query{"status": gitdb.Query{"$in": []string{"active", "pending"}}}

// And operator
query := gitdb.Query{
    "age": gitdb.Query{"$gte": 25},
    "active": true,
}
```

## Update Operators

```go
// Set fields
update := gitdb.Update{
    "$set": gitdb.Document{
        "age": 31,
        "lastUpdated": "2024-01-15",
    },
}

// Increment
update := gitdb.Update{
    "$inc": gitdb.Document{"views": 1},
}

// Push to array
update := gitdb.Update{
    "$push": gitdb.Document{"tags": "new-tag"},
}
```

## Examples

See the `examples/` directory for complete working examples:

- `basic_usage.go` - Basic CRUD operations
- `advanced_queries.go` - Complex queries and updates
- `graphql_example.go` - GraphQL usage

## License

MIT License - see LICENSE file for details. 