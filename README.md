# GitDB Go Client

Official Go client for GitDB - GitHub-backed NoSQL database.

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
    // Create a new client
    client := gitdb.NewClient("your-github-token", "owner", "repo")
    
    // Check server health
    err := client.Health()
    if err != nil {
        log.Fatal("Server health check failed:", err)
    }
    
    // Create a collection
    err = client.CreateCollection("users")
    if err != nil {
        log.Fatal("Failed to create collection:", err)
    }
    
    // Insert a document
    document := map[string]interface{}{
        "name":  "John Doe",
        "email": "john@example.com",
        "age":   30,
    }
    
    id, err := client.Insert("users", document)
    if err != nil {
        log.Fatal("Failed to insert document:", err)
    }
    fmt.Printf("Inserted document with ID: %s\n", id)
    
    // Find documents
    query := map[string]interface{}{
        "age": map[string]interface{}{
            "$gte": 25,
        },
    }
    
    documents, err := client.Find("users", query)
    if err != nil {
        log.Fatal("Failed to find documents:", err)
    }
    fmt.Printf("Found %d documents\n", len(documents))
    
    // Update a document
    update := map[string]interface{}{
        "age": 31,
    }
    
    err = client.Update("users", id, update)
    if err != nil {
        log.Fatal("Failed to update document:", err)
    }
    
    // Delete a document
    err = client.Delete("users", id)
    if err != nil {
        log.Fatal("Failed to delete document:", err)
    }
}
```

## Features

- ✅ **Simple API** - Easy to use Go interface
- ✅ **Full CRUD operations** - Create, Read, Update, Delete documents
- ✅ **Query support** - MongoDB-style query operators
- ✅ **Collection management** - Create, list, delete collections
- ✅ **Error handling** - Comprehensive error handling
- ✅ **Type safety** - Strong typing throughout
- ✅ **HTTP client** - Built-in HTTP client with retry logic
- ✅ **JSON handling** - Native JSON serialization/deserialization

## Configuration

### GitHub Token

You'll need a GitHub Personal Access Token with the following permissions:
- `repo` - Full control of private repositories
- `workflow` - Update GitHub Action workflows

Create a token at: https://github.com/settings/tokens

### Client Initialization

```go
import "github.com/karthikeyanV2K/gitdb-client/gitdb"

// Basic initialization
client := gitdb.NewClient("token", "owner", "repo")

// With custom base URL (for self-hosted instances)
client := gitdb.NewClientWithURL("token", "owner", "repo", "http://localhost:7896")
```

## API Reference

### Client Creation

```go
// Create a new client
client := gitdb.NewClient(token, owner, repo)

// Create client with custom URL
client := gitdb.NewClientWithURL(token, owner, repo, "http://localhost:7896")
```

### Health Check

```go
// Check if server is healthy
err := client.Health()
if err != nil {
    log.Fatal("Server is not healthy:", err)
}
```

### Collection Management

```go
// Create a collection
err := client.CreateCollection("users")

// List all collections
collections, err := client.ListCollections()
for _, collection := range collections {
    fmt.Printf("Collection: %s (%d documents)\n", collection.Name, collection.Count)
}

// Delete a collection
err := client.DeleteCollection("users")
```

### Document Operations

#### Insert

```go
document := map[string]interface{}{
    "name":  "Alice",
    "email": "alice@example.com",
    "age":   25,
}

id, err := client.Insert("users", document)
if err != nil {
    log.Fatal("Insert failed:", err)
}
fmt.Printf("Inserted with ID: %s\n", id)
```

#### Find

```go
// Find all documents
documents, err := client.Find("users", nil)

// Find with query
query := map[string]interface{}{
    "age": map[string]interface{}{
        "$gte": 30,
    },
}
documents, err := client.Find("users", query)

// Find one document
document, err := client.FindOne("users", query)

// Find by ID
document, err := client.FindByID("users", "document-id")
```

#### Update

```go
update := map[string]interface{}{
    "age": 26,
    "last_updated": "2024-01-01",
}

err := client.Update("users", "document-id", update)
```

#### Delete

```go
// Delete by ID
err := client.Delete("users", "document-id")

// Delete multiple documents
query := map[string]interface{}{
    "age": map[string]interface{}{
        "$lt": 18,
    },
}
deletedCount, err := client.DeleteMany("users", query)
```

### Batch Operations

```go
// Insert multiple documents
documents := []map[string]interface{}{
    {"name": "Alice", "age": 25},
    {"name": "Bob", "age": 30},
    {"name": "Charlie", "age": 35},
}

for _, doc := range documents {
    id, err := client.Insert("users", doc)
    if err != nil {
        log.Printf("Failed to insert document: %v", err)
    }
}

// Update multiple documents
query := map[string]interface{}{
    "age": map[string]interface{}{
        "$gte": 25,
    },
}

update := map[string]interface{}{
    "category": "senior",
}

modifiedCount, err := client.UpdateMany("users", query, update)
```

### Query Operators

The Go client supports MongoDB-style query operators:

```go
query := map[string]interface{}{}

// Equal
query["age"] = 30

// Greater than
query["age"] = map[string]interface{}{
    "$gt": 25,
}

// Greater than or equal
query["age"] = map[string]interface{}{
    "$gte": 25,
}

// Less than
query["age"] = map[string]interface{}{
    "$lt": 50,
}

// Less than or equal
query["age"] = map[string]interface{}{
    "$lte": 50,
}

// In array
query["status"] = map[string]interface{}{
    "$in": []string{"active", "pending"},
}

// Not in array
query["status"] = map[string]interface{}{
    "$nin": []string{"inactive", "deleted"},
}

// Logical AND
query["$and"] = []map[string]interface{}{
    {"age": map[string]interface{}{"$gte": 18}},
    {"status": "active"},
}

// Logical OR
query["$or"] = []map[string]interface{}{
    {"status": "active"},
    {"status": "pending"},
}
```

## Error Handling

The SDK provides comprehensive error handling:

```go
document, err := client.FindByID("users", "non-existent-id")
if err != nil {
    if strings.Contains(err.Error(), "not found") {
        fmt.Println("Document not found")
    } else {
        log.Printf("Unexpected error: %v", err)
    }
    return
}
```

## Advanced Usage

### Custom HTTP Client

```go
import (
    "net/http"
    "time"
    "github.com/karthikeyanV2K/gitdb-client/gitdb"
)

// Create custom HTTP client
httpClient := &http.Client{
    Timeout: 30 * time.Second,
}

// Create GitDB client with custom HTTP client
client := gitdb.NewClientWithHTTPClient("token", "owner", "repo", httpClient)
```

### Context Support

```go
import (
    "context"
    "time"
)

ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

// Use context for operations
documents, err := client.FindWithContext(ctx, "users", query)
```

### Retry Logic

```go
// The client includes built-in retry logic for transient errors
// You can configure retry behavior if needed
client.SetMaxRetries(3)
client.SetRetryDelay(1 * time.Second)
```

## Examples

### User Management System

```go
package main

import (
    "fmt"
    "log"
    "time"
    "github.com/karthikeyanV2K/gitdb-client/gitdb"
)

type User struct {
    ID        string    `json:"id,omitempty"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    Age       int       `json:"age"`
    Status    string    `json:"status"`
    CreatedAt time.Time `json:"created_at"`
}

type UserManager struct {
    client *gitdb.Client
}

func NewUserManager(token, owner, repo string) *UserManager {
    return &UserManager{
        client: gitdb.NewClient(token, owner, repo),
    }
}

func (um *UserManager) CreateUser(name, email string, age int) (string, error) {
    user := User{
        Name:      name,
        Email:     email,
        Age:       age,
        Status:    "active",
        CreatedAt: time.Now(),
    }
    
    return um.client.Insert("users", user)
}

func (um *UserManager) FindUserByEmail(email string) (map[string]interface{}, error) {
    query := map[string]interface{}{
        "email": email,
    }
    
    return um.client.FindOne("users", query)
}

func (um *UserManager) UpdateUserStatus(userID, status string) error {
    update := map[string]interface{}{
        "status": status,
    }
    
    return um.client.Update("users", userID, update)
}

func (um *UserManager) GetActiveUsers() ([]map[string]interface{}, error) {
    query := map[string]interface{}{
        "status": "active",
    }
    
    return um.client.Find("users", query)
}

func (um *UserManager) DeleteInactiveUsers() (int, error) {
    query := map[string]interface{}{
        "status": "inactive",
    }
    
    return um.client.DeleteMany("users", query)
}

func main() {
    userManager := NewUserManager("your-token", "owner", "repo")
    
    // Create user
    userID, err := userManager.CreateUser("John Doe", "john@example.com", 30)
    if err != nil {
        log.Fatal("Failed to create user:", err)
    }
    
    // Find user
    user, err := userManager.FindUserByEmail("john@example.com")
    if err != nil {
        log.Fatal("Failed to find user:", err)
    }
    
    // Update status
    err = userManager.UpdateUserStatus(userID, "inactive")
    if err != nil {
        log.Fatal("Failed to update user status:", err)
    }
    
    // Get active users
    activeUsers, err := userManager.GetActiveUsers()
    if err != nil {
        log.Fatal("Failed to get active users:", err)
    }
    
    fmt.Printf("Active users: %d\n", len(activeUsers))
}
```

## Testing

```go
package main

import (
    "testing"
    "github.com/karthikeyanV2K/gitdb-client/gitdb"
)

func TestClientCreation(t *testing.T) {
    client := gitdb.NewClient("token", "owner", "repo")
    
    if client == nil {
        t.Error("Client should not be nil")
    }
}

func TestInsertAndFind(t *testing.T) {
    client := gitdb.NewClient("token", "owner", "repo")
    
    // Test document
    document := map[string]interface{}{
        "name": "Test User",
        "age":  25,
    }
    
    // Insert
    id, err := client.Insert("test", document)
    if err != nil {
        t.Errorf("Insert failed: %v", err)
    }
    
    // Find by ID
    found, err := client.FindByID("test", id)
    if err != nil {
        t.Errorf("FindByID failed: %v", err)
    }
    
    if found["name"] != "Test User" {
        t.Error("Document not found correctly")
    }
    
    // Cleanup
    client.Delete("test", id)
}
```

## Troubleshooting

### Common Issues

1. **Authentication Error**
   - Verify your GitHub token is valid
   - Ensure token has required permissions
   - Check token hasn't expired

2. **Repository Access**
   - Verify repository exists
   - Check you have access to the repository
   - Ensure repository is not private (unless using private GitDB)

3. **Network Issues**
   - Check internet connection
   - Verify GitHub API is accessible
   - Check firewall settings

4. **Rate Limiting**
   - GitHub API has rate limits
   - Implement exponential backoff for retries
   - Consider using authenticated requests

### Debug Mode

Enable debug mode to see detailed request/response information:

```go
// Set debug mode (if supported by the client)
client.SetDebug(true)
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## License

This project is licensed under the MIT License - see the [LICENSE](../../LICENSE) file for details.

## Support

- GitHub Issues: https://github.com/karthikeyanV2K/GitDB/issues
- Documentation: https://github.com/karthikeyanV2K/GitDB
- Email: Support@afot.in

## Changelog

### v1.0.0
- Initial release
- Full CRUD operations
- Query support with MongoDB-style operators
- Error handling
- Comprehensive documentation 