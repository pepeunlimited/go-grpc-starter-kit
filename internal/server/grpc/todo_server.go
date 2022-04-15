package grpc

import (
	"context"
	"math"
	"math/rand"
	"sort"
	"sync"
	"time"

	"github.com/pepeunlimited/go-grpc-starter-kit/pkg/api/v1/resources"
	"github.com/pepeunlimited/go-grpc-starter-kit/pkg/api/v1/services"
	"github.com/pepeunlimited/go-grpc-starter-kit/pkg/fm"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/mennanov/fmutils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TodoServer auto generate:
// => CMD + SHIFT + A
// => Go: Generate Interaface Stubs
// => server TodoServer services.TodoServiceServer
type TodoServer struct {
	todos     map[int64]*resources.Todo
	keysOrder []int64
	// @see https://stackoverflow.com/a/69302744/3913343
	services.UnimplementedTodoServiceServer
	// @see https://kylewbanks.com/blog/when-to-use-defer-in-go
	mu *sync.RWMutex
}

func (server *TodoServer) CreateTodo(ctx context.Context, request *services.CreateTodoRequest) (*resources.Todo, error) {
	server.mu.Lock()
	defer server.mu.Unlock()
	generatedId := server.generateId()
	createdAt := time.Now().UTC()
	resourcesTodo := &resources.Todo{
		Id:        generatedId,
		Title:     request.Title,
		CreatedAt: createdAt.Format(time.RFC3339),
		UpdatedAt: createdAt.Format(time.RFC3339),
		IsDone:    false,
	}
	server.todos[generatedId] = resourcesTodo
	server.keysOrder = append(server.keysOrder, generatedId)
	return resourcesTodo, nil
}

func (server *TodoServer) generateId() int64 {
	rand.Seed(time.Now().UnixNano())
	for {
		id := rand.Int63n(math.MaxInt64)
		if _, ok := server.todos[id]; !ok {
			return id
		}
	}
}

func (server *TodoServer) GetTodo(ctx context.Context, request *services.GetTodoRequest) (*resources.Todo, error) {
	server.mu.Lock()
	defer server.mu.Unlock()
	if fm.IsNilOrPathEmpty(request.FieldMask) {
		// set all possible fields to paths
		request.FieldMask = &fieldmaskpb.FieldMask{
			Paths: []string{
				"id",
				"title",
				"created_at",
				"updated_at",
				"is_done",
			},
		}
	}
	// => m proto.Message
	todoWithMaskedFieldMask := &resources.Todo{}
	// normalize and validate the field mask before using it.
	request.FieldMask.Normalize()
	if !request.FieldMask.IsValid(todoWithMaskedFieldMask) {
		return nil, status.Error(codes.InvalidArgument, "Invalid field_mask")
	}
	// IMPORTANT: sort the fieldMask paths
	sort.StringSlice(request.FieldMask.Paths).Sort()
	// check if should execute external database join queries
	//withParents := fm.PathContains(params.FieldMask.Paths, "foo")

	// next step should be query database with request ID.
	// now we memcache as 'db'
	resourcesTodo, ok := server.todos[request.Id]
	if !ok {
		return nil, status.Error(codes.NotFound, "Todo entity not exist")
	}
	// redact the request according to the provided field mask.
	fmutils.Filter(resourcesTodo, request.FieldMask.Paths)
	// now that the request is vetted we can merge it with the user entity.
	proto.Merge(todoWithMaskedFieldMask, resourcesTodo)
	return resourcesTodo, nil
}

func (server *TodoServer) ListTodo(ctx context.Context, request *services.ListTodoRequest) (*services.ListTodoResponse, error) {
	server.mu.Lock()
	defer server.mu.Unlock()
	if fm.IsNilOrPathEmpty(request.FieldMask) {
		// set all possible fields to paths
		request.FieldMask = &fieldmaskpb.FieldMask{
			Paths: []string{
				"id",
				"title",
				"created_at",
				"updated_at",
				"is_done",
			},
		}
	}
	// normalize and validate the field mask before using it.
	request.FieldMask.Normalize()
	if !request.FieldMask.IsValid(&resources.Todo{}) {
		return nil, status.Error(codes.InvalidArgument, "Invalid field_mask")
	}
	// IMPORTANT: sort the fieldMask paths
	sort.StringSlice(request.FieldMask.Paths).Sort()
	// check if should execute external database join queries
	//withParents := fm.PathContains(params.FieldMask.Paths, "foo")

	// next step should be query database with request pageToken and nextPageToken.
	// now we memcache as 'db'
	todos := make([]*resources.Todo, 0)
	//nextPageToken := int64

	isNextPageToken := false

	for _, x := range server.keysOrder {
		if len(todos) > int(request.PageSize) {
			break
		}
		if request.PageToken == 0 {
			if entity, ok := server.todos[x]; ok {
				todos = append(todos, entity)
			}
		} else {
			if request.PageToken == x {
				isNextPageToken = true
			} else {
				if isNextPageToken {
					if entity, ok := server.todos[x]; ok {
						todos = append(todos, entity)
					}
				}
			}
		}
	}
	nextPageToken := todos[len(todos)-1].Id
	// create a mask only once and reuse it.
	mask := fmutils.NestedMaskFromPaths(request.FieldMask.Paths)
	for _, todo := range todos {
		mask.Filter(todo)
	}
	return &services.ListTodoResponse{
		Todos:         todos,
		NextPageToken: nextPageToken,
	}, nil
}

func (server *TodoServer) DeleteTodo(ctx context.Context, request *services.DeleteTodoRequest) (*emptypb.Empty, error) {
	if _, err := server.GetTodo(ctx, &services.GetTodoRequest{
		Id: request.TodoId,
		FieldMask: &fieldmaskpb.FieldMask{
			Paths: []string{
				"id",
			},
		},
	}); err != nil {
		return nil, err
	}
	delete(server.todos, request.TodoId)
	var firstIndex int64
	for i, x := range server.keysOrder {
		if x == request.TodoId {
			firstIndex = int64(i)
		}
	}
	copy(server.keysOrder[firstIndex:], server.keysOrder[firstIndex+1:])
	server.keysOrder = server.keysOrder[:len(server.keysOrder)-1]
	return &emptypb.Empty{}, nil
}

func (server *TodoServer) UpdateTodo(ctx context.Context, request *services.UpdateTodoRequest) (*resources.Todo, error) {
	if request.Operation.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "operation.id is required")
	}
	if fm.IsNilOrPathEmpty(request.FieldMask) {
		return nil, status.Error(codes.InvalidArgument, "field_mask is required")
	}

	// check bok existence and => m proto.Message
	updateWithMaskedFieldMask, err := server.GetTodo(ctx, &services.GetTodoRequest{
		Id: request.Operation.Id,
	})
	if err != nil {
		return nil, err
	}
	// normalize and validate the field mask before using it.
	request.FieldMask.Normalize()
	if !request.FieldMask.IsValid(updateWithMaskedFieldMask) {
		return nil, status.Error(codes.InvalidArgument, "Invalid field_mask")
	}
	// IMPORTANT: sort the fieldMask paths
	sort.StringSlice(request.FieldMask.Paths).Sort()
	// Redact the request according to the provided field mask.
	fmutils.Filter(request.Operation, request.FieldMask.Paths)
	// Now that the request is vetted we can merge it with the bookmark entity.
	proto.Merge(updateWithMaskedFieldMask, request.Operation)
	// next step should be update database with request ID.
	// now we memcache as 'db'
	updateWithMaskedFieldMask.UpdatedAt = time.Now().UTC().Format(time.RFC3339)
	server.todos[request.Operation.Id] = updateWithMaskedFieldMask
	return updateWithMaskedFieldMask, nil
}

func NewTodoServer() *TodoServer {
	return &TodoServer{
		todos:     make(map[int64]*resources.Todo),
		keysOrder: make([]int64, 0),
		mu:        &sync.RWMutex{},
	}
}
