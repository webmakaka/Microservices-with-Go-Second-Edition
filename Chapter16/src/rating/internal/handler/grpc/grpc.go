package grpc

import (
	"context"
	"errors"

	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"movieexample.com/gen"
	"movieexample.com/rating/internal/controller/rating"
	"movieexample.com/rating/pkg/model"
)

var validate = validator.New()

// Handler defines a gRPC rating API handler.
type Handler struct {
	gen.UnimplementedRatingServiceServer
	ctrl *rating.Controller
}

// New creates a new rating gRPC handler.
func New(ctrl *rating.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

// GetAggregatedRating returns the aggregated rating for a record.
func (h *Handler) GetAggregatedRating(ctx context.Context, req *gen.GetAggregatedRatingRequest) (*gen.GetAggregatedRatingResponse, error) {
	if req == nil || req.RecordId == "" || req.RecordType == "" {
		return nil, status.Errorf(codes.InvalidArgument, "nil req or empty id/type")
	}
	v, err := h.ctrl.GetAggregatedRating(ctx, model.RecordID(req.RecordId), model.RecordType(req.RecordType))
	if err != nil && errors.Is(err, rating.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &gen.GetAggregatedRatingResponse{RatingValue: v}, nil
}

// PutRating writes a rating for a given record.
func (h *Handler) PutRating(ctx context.Context, req *gen.PutRatingRequest) (*gen.PutRatingResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "nil req")
	}
	rating := &model.Rating{UserID: model.UserID(req.UserId), Value: model.RatingValue(req.RatingValue)}
	if err := validate.Struct(rating); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation failed: %v", err)
	}
	if err := h.ctrl.PutRating(ctx, model.RecordID(req.RecordId), model.RecordType(req.RecordType), rating); err != nil {
		return nil, err
	}
	return &gen.PutRatingResponse{}, nil
}
