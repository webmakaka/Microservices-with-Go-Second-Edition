package grpc

import (
	"context"
	"errors"

	"github.com/uber-go/tally"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"movieexample.com/gen"
	"movieexample.com/metadata/internal/controller/metadata"
	"movieexample.com/metadata/pkg/model"
)

// Handler defines a movie metadata gRPC handler.
type Handler struct {
	gen.UnimplementedMetadataServiceServer
	ctrl  *metadata.Controller
	scope tally.Scope
}

// New creates a new movie metadata gRPC handler.
func New(ctrl *metadata.Controller, scope tally.Scope) *Handler {
	return &Handler{ctrl: ctrl, scope: scope.Tagged(map[string]string{"component": "handler"})}
}

// GetMetadata returns movie metadata.
func (h *Handler) GetMetadata(ctx context.Context, req *gen.GetMetadataRequest) (*gen.GetMetadataResponse, error) {
	scope := h.scope.Tagged(map[string]string{"endpoint": "GetMetadata"})
	scope.Counter("call").Inc(1)
	if req == nil || req.MovieId == "" {
		scope.Tagged(map[string]string{"error": "invalid_argument"}).Counter("error").Inc(1)
		return nil, status.Errorf(codes.InvalidArgument, "nil req or empty id")
	}
	m, err := h.ctrl.Get(ctx, req.MovieId)
	if err != nil && errors.Is(err, metadata.ErrNotFound) {
		scope.Tagged(map[string]string{"error": "not_found"}).Counter("error").Inc(1)
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil {
		scope.Tagged(map[string]string{"error": "internal"}).Counter("error").Inc(1)
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	scope.Counter("success").Inc(1)
	return &gen.GetMetadataResponse{Metadata: model.MetadataToProto(m)}, nil
}

// PutMetadata puts movie metadata to repository.
func (h *Handler) PutMetadata(ctx context.Context, req *gen.PutMetadataRequest) (*gen.PutMetadataResponse, error) {
	scope := h.scope.Tagged(map[string]string{"endpoint": "PutMetadata"})
	scope.Counter("call").Inc(1)
	if req == nil || req.Metadata == nil {
		scope.Tagged(map[string]string{"error": "invalid_argument"}).Counter("error").Inc(1)
		return nil, status.Errorf(codes.InvalidArgument, "nil req or metadata")
	}
	if err := h.ctrl.Put(ctx, model.MetadataFromProto(req.Metadata)); err != nil {
		scope.Tagged(map[string]string{"error": "internal"}).Counter("error").Inc(1)
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	scope.Counter("success").Inc(1)
	return &gen.PutMetadataResponse{}, nil
}
