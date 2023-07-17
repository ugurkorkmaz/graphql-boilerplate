package resolver

import (
	"context"

	"app/gen"
	"app/model"

	"github.com/google/uuid"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.CreateTodoInput) (*model.Todo, error) {
	// Create a new todo.
	todo := &model.Todo{
		ID:   uuid.New(), // Generate a new UUID.
		Text: input.Text, // Set the text from the input.
		Done: false,      // Set the done status to false.
	}
	// Add the todo to the list of todos.
	r.ModelTodos = append(r.ModelTodos, todo)
	// Notify the client that the todo was created.
	r.OnTodoCreated <- r.ModelTodos
	return todo, nil
}

// UpdateTodoDone is the resolver for the updateTodoDone field.
func (r *mutationResolver) UpdateTodoDone(ctx context.Context, id uuid.UUID, done bool) (*model.Todo, error) {
	// Find the todo with the given ID.
	var todo *model.Todo
	// Iterate over the list of todos.
	for _, t := range r.ModelTodos {
		if t.ID == id {
			todo = t
			break
		}
	}
	// Return an error if the todo was not found.
	if todo == nil {
		return nil, nil
	}
	// Update the todo.
	todo.Done = done
	// Notify the client that the todo was updated.
	r.OnTodoUpdated <- r.ModelTodos
	// Return the updated todo.
	return todo, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	// Return the list of todos.
	return r.ModelTodos, nil
}

// TodoCreated is the resolver for the todoCreated field.
func (r *subscriptionResolver) TodoCreated(ctx context.Context) (<-chan []*model.Todo, error) {
	// Return the channel that is used to notify the client that a todo was created.
	return r.OnTodoCreated, nil
}

// TodoUpdated is the resolver for the todoUpdated field.
func (r *subscriptionResolver) TodoUpdated(ctx context.Context) (<-chan []*model.Todo, error) {
	// Return the channel that is used to notify the client that a todo was updated.
	return r.OnTodoUpdated, nil
}

// Mutation returns gen.MutationResolver implementation.
func (r *Resolver) Mutation() gen.MutationResolver { return &mutationResolver{r} }

// Query returns gen.QueryResolver implementation.
func (r *Resolver) Query() gen.QueryResolver { return &queryResolver{r} }

// Subscription returns gen.SubscriptionResolver implementation.
func (r *Resolver) Subscription() gen.SubscriptionResolver { return &subscriptionResolver{r} }

type (
	mutationResolver     struct{ *Resolver }
	queryResolver        struct{ *Resolver }
	subscriptionResolver struct{ *Resolver }
)
