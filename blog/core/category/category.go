package category

import (
	"context"
	"blog/blog/storage"
	"blog/blog/storage/postgres"
)
type CategoryCoreSvc struct {
	core *postgres.Storage
}

func NewCategoryCoreSvc(s *postgres.Storage) *CategoryCoreSvc {
	return &CategoryCoreSvc{
		core: s,
	}
}

func (cs CategoryCoreSvc) CreateCat(ctx context.Context, t storage.Category) (int64, error) {

	return cs.core.CreateCat(ctx,t)
}
func (cs CategoryCoreSvc) ListCat(ctx context.Context) ([]storage.Category, error) {

	return cs.core.ListCat(ctx)
}
func (cs CategoryCoreSvc) GetCat(ctx context.Context, id int64) (storage.Category, error) {

	return cs.core.GetCat(ctx, id)
}
func (cs CategoryCoreSvc) UpdateCat(ctx context.Context, t storage.Category) error {

	return cs.core.UpdateCat(ctx, t)
}
func (cs CategoryCoreSvc) DeleteCat(ctx context.Context, id int64) error {

	return cs.core.DeleteCat(ctx, id)
}
func (cs CategoryCoreSvc) SearchCategory(ctx context.Context, t string) ([]storage.Category,error ){

	return cs.core.SearchCategory(ctx,t)
}
