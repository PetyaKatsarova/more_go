package model

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {            // adding json tags, struct tags, key
	OrderID		uint64			`json:"order_id"`
	CustomerID	uuid.UUID		`json:"customer_id"`
	LineItems	[]LineItem		`json:"line_items"`
	CreatedAt	*time.Time		`json:"created_at"`
	ShippedAt	*time.Time      `json:"shipped_at"`
	CompletedAt	*time.Time		`json:"completed_at"`
}

type LineItem struct {
	ItemID		uuid.UUID
	Quantity	uint
	Price		uint   
}

/*
UUIDv4, short for "Universally Unique Identifier version 4," is a type of universally unique identifier defined by RFC 4122, a specification created by the Internet
Engineering Task Force (IETF). UUIDs are standardized 128-bit identifiers that are used to uniquely identify resources or entities in a distributed computing environment.
Here are some key characteristics of UUIDv4:
Randomness: UUIDv4 is generated using random or pseudo-random numbers, which makes it highly unlikely that two UUIDs generated at different times and places will be the same.
128-Bit Length: A UUIDv4 is 128 bits long, represented as a 32-character hexadecimal string with five groups separated by hyphens, such as
 "550e8400-e29b-41d4-a716-446655440000."
Uniqueness: While the use of randomness ensures a high degree of uniqueness, it's important to note that UUIDv4 uniqueness is not guaranteed across all time and space.
The uniqueness depends on the quality of the random number generator used to create the UUID.
Version 4: UUIDv4 is one of several UUID versions defined by RFC 4122. It is specifically identified by having a version number of 4 in its hexadecimal representation.
Variants: UUIDv4 follows the "Leach-Salz" variant defined in RFC 4122, which is one of the variants allowed for UUIDs.
UUIDv4 is commonly used in various applications and systems, including database records, distributed systems, and as identifiers in various programming languages.
 It is particularly useful in scenarios where multiple systems need to generate unique identifiers without centralized coordination. However, because UUIDs are relatively
  long and contain randomness, they are not suitable for use as human-readable identifiers. Instead, they are primarily used as machine-readable unique identifiers.
  --- UUID package
   go get github.com/google/uuid
*/