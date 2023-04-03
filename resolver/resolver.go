package resolver

import (
	"app/gen"
	"app/model"

	"github.com/99designs/gqlgen/graphql"
)

/*
This channel-based implementation is just one approach to handling subscriptions.
Remember that there are other tools available for handling subscriptions beyond channels.
For example, you can use Redis or a similar message broker to manage subscriptions across
multiple instances of your application.
These tools can provide features like scalability, message persistence, and automatic subscription management.
However, using these tools may require additional setup and configuration,
so be sure to consider your specific needs and constraints before deciding which approach to take.
*/
type Resolver struct {
	// ModelTodos is the list of todos.
	ModelTodos []*model.Todo
	// OnTodoCreated is the channel that is used to notify the client that a todo was created.
	OnTodoCreated chan []*model.Todo
	// OnTodoUpdated is the channel that is used to notify the client that a todo was updated.
	OnTodoUpdated chan []*model.Todo
}

// NewSchema creates a graphql executable schema.
func NewSchema() graphql.ExecutableSchema {
	return gen.NewExecutableSchema(gen.Config{
		Resolvers: &Resolver{
			// Initialize the list of todos.
			ModelTodos: []*model.Todo{},
			// Initialize the channels that are used to notify the client that a todo was created.
			OnTodoCreated: make(chan []*model.Todo),
			// Initialize the channels that are used to notify the client that a todo was updated.
			OnTodoUpdated: make(chan []*model.Todo),
		},
	})
}
