package main

import (
	"fmt"
	"log"

	"github.com/karthikeyanV2K/gitdb-client/gitdb"
)

func main() {
	// Create a new GitDB client
	client := gitdb.NewClient("your_github_token", "your_username", "your_repo")

	// Set custom base URL if needed
	// client.SetBaseURL("http://localhost:7896")

	// Health check
	if err := client.Health(); err != nil {
		log.Fatalf("Health check failed: %v", err)
	}
	fmt.Println("✅ GitDB server is healthy")

	// Create a collection
	collectionName := "users"
	if err := client.CreateCollection(collectionName); err != nil {
		log.Fatalf("Failed to create collection: %v", err)
	}
	fmt.Printf("✅ Created collection: %s\n", collectionName)

	// List collections
	collections, err := client.ListCollections()
	if err != nil {
		log.Fatalf("Failed to list collections: %v", err)
	}
	fmt.Printf("📋 Collections: %+v\n", collections)

	// Insert a document
	user := gitdb.Document{
		"name":  "John Doe",
		"email": "john@example.com",
		"age":   30,
		"active": true,
		"tags":  []string{"developer", "golang"},
		"address": gitdb.Document{
			"street":  "123 Main St",
			"city":    "New York",
			"country": "USA",
		},
	}

	docID, err := client.Insert(collectionName, user)
	if err != nil {
		log.Fatalf("Failed to insert document: %v", err)
	}
	fmt.Printf("✅ Inserted document with ID: %s\n", docID)

	// Insert another document
	user2 := gitdb.Document{
		"name":  "Jane Smith",
		"email": "jane@example.com",
		"age":   25,
		"active": true,
		"tags":  []string{"designer", "ui"},
		"address": gitdb.Document{
			"street":  "456 Oak Ave",
			"city":    "San Francisco",
			"country": "USA",
		},
	}

	docID2, err := client.Insert(collectionName, user2)
	if err != nil {
		log.Fatalf("Failed to insert document: %v", err)
	}
	fmt.Printf("✅ Inserted document with ID: %s\n", docID2)

	// Find document by ID
	doc, err := client.FindByID(collectionName, docID)
	if err != nil {
		log.Fatalf("Failed to find document: %v", err)
	}
	fmt.Printf("📄 Found document: %+v\n", doc)

	// Find documents with query
	query := gitdb.Query{
		"active": true,
		"age": gitdb.Query{
			"$gte": 25,
		},
	}

	documents, err := client.Find(collectionName, query)
	if err != nil {
		log.Fatalf("Failed to find documents: %v", err)
	}
	fmt.Printf("🔍 Found %d active users aged 25+: %+v\n", len(documents), documents)

	// Find one document
	oneDoc, err := client.FindOne(collectionName, gitdb.Query{"email": "john@example.com"})
	if err != nil {
		log.Fatalf("Failed to find one document: %v", err)
	}
	fmt.Printf("🔍 Found one document: %+v\n", oneDoc)

	// Count documents
	count, err := client.Count(collectionName, gitdb.Query{"active": true})
	if err != nil {
		log.Fatalf("Failed to count documents: %v", err)
	}
	fmt.Printf("📊 Active users count: %d\n", count)

	// Update a document
	update := gitdb.Update{
		"$set": gitdb.Document{
			"age": 31,
			"lastUpdated": "2024-01-15",
		},
	}

	if err := client.Update(collectionName, docID, update); err != nil {
		log.Fatalf("Failed to update document: %v", err)
	}
	fmt.Printf("✅ Updated document: %s\n", docID)

	// Update many documents
	updateMany := gitdb.Update{
		"$set": gitdb.Document{
			"lastUpdated": "2024-01-15",
		},
	}

	modifiedCount, err := client.UpdateMany(collectionName, gitdb.Query{"active": true}, updateMany)
	if err != nil {
		log.Fatalf("Failed to update many documents: %v", err)
	}
	fmt.Printf("✅ Updated %d documents\n", modifiedCount)

	// Delete a document
	if err := client.Delete(collectionName, docID2); err != nil {
		log.Fatalf("Failed to delete document: %v", err)
	}
	fmt.Printf("✅ Deleted document: %s\n", docID2)

	// Delete many documents
	deletedCount, err := client.DeleteMany(collectionName, gitdb.Query{"active": false})
	if err != nil {
		log.Fatalf("Failed to delete many documents: %v", err)
	}
	fmt.Printf("✅ Deleted %d inactive documents\n", deletedCount)

	// GraphQL query example
	graphqlQuery := `
		query {
			collections
			documents(collection: "users") {
				_id
				name
				email
				age
			}
		}
	`

	response, err := client.GraphQL(graphqlQuery, nil)
	if err != nil {
		log.Fatalf("Failed to execute GraphQL query: %v", err)
	}
	fmt.Printf("🔮 GraphQL response: %+v\n", response.Data)

	// Delete collection
	if err := client.DeleteCollection(collectionName); err != nil {
		log.Fatalf("Failed to delete collection: %v", err)
	}
	fmt.Printf("✅ Deleted collection: %s\n", collectionName)

	fmt.Println("🎉 All operations completed successfully!")
} 