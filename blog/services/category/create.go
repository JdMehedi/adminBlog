package category

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"blog/blog/storage"
	bgvc "blog/gunk/v1/category"
)

func (s *CategorySvc) CreateCategory(ctx context.Context, req *bgvc.CreateCategoryRequest) (*bgvc.CreateCategoryResponse, error) {
	//Needs to validate Category
	categories := storage.Category{
		Title: req.GetCategory().Title,
	}
	id, err := s.store.CreateCat(context.Background(), categories)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create todo: %s", err)
	}

	return &bgvc.CreateCategoryResponse{
		ID: id,
	}, nil

}

func (s *CategorySvc) ListCategory(ctx context.Context, req *bgvc.ListCategoryRequest) (*bgvc.ListCategoryResponse, error) {
	//Needs to validate Category
	res, err := s.store.ListCat(context.Background())

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get Category: %s", err)
	}
	var ctl []*bgvc.Category

	for _, value := range res {
		ctl = append(ctl, &bgvc.Category{
			ID:    value.ID,
			Title: value.Title,
		})
	}
	return &bgvc.ListCategoryResponse{
		Category: ctl,
	}, nil
}
func (s *CategorySvc) GetCategory(ctx context.Context, req *bgvc.GetCategoryRequest) (*bgvc.GetCategoryResponse, error) {
	//Needs to validate Category
	res, err := s.store.GetCat(context.Background(), req.GetID())

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get Category: %s", err)
	}

	return &bgvc.GetCategoryResponse{
		Category: &bgvc.Category{
			ID:    res.ID,
			Title: res.Title,
		},
	}, nil
}

func (s *CategorySvc) UpdateCategory(ctx context.Context, req *bgvc.UpdateCategoryRequest) (*bgvc.UpdateCategoryResponse, error) {
	//Needs to validate Category
	categories := storage.Category{
		ID:    req.Category.ID,
		Title: req.Category.Title,
	}
	err := s.store.UpdateCat(context.Background(), categories)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get Category: %s", err)
	}

	return &bgvc.UpdateCategoryResponse{}, nil

}

func (s *CategorySvc) DeleteCategory(ctx context.Context, req *bgvc.DeleteCategoryRequest) (*bgvc.DeleteCategoryResponse, error) {
	//Needs to validate Category

	err := s.store.DeleteCat(context.Background(), req.GetID())

	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to Delete category.")
	}
	return &bgvc.DeleteCategoryResponse{}, nil
}
func (s *CategorySvc) SearchCategory(ctx context.Context, req *bgvc.SearchCategoryRequest) (*bgvc.SearchCategoryResponse, error) {
	//Needs to validate Category
	res,err := s.store.SearchCategory(context.Background(),req.GetTitle())

	var sec []*bgvc.Category

	for _, value := range res {
		sec = append(sec, &bgvc.Category{
			ID: value.ID,
			Title: value.Title,
		})
	}
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to Search Category.")
	}
	return &bgvc.SearchCategoryResponse{
		Category: sec,
	}, nil
}
