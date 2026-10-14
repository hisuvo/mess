package graphql

import "messify/internal/domain/users"

// Resolver provides dependencies to GraphQL resolvers.
// Business logic remains inside the service layer.
type Resolver struct {
	UserService users.Service
}
