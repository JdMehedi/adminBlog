

package category

import (
	"context"

	"blog/blog/storage"
	bgvc "blog/gunk/v1/category"
)

type CategoryCoreLink interface{
	CreateCat(context.Context,storage.Category)(int64,error)
	ListCat(context.Context)([]storage.Category, error)
	GetCat(context.Context, int64)(storage.Category, error)
	UpdateCat(context.Context, storage.Category) error
	DeleteCat(context.Context,int64)error
	SearchCategory(context.Context,string)([]storage.Category,error) 
}

type CategorySvc struct {
	bgvc.UnimplementedCategoryServiceServer
	store CategoryCoreLink
}

func NewCategorySvc(c CategoryCoreLink) *CategorySvc{
	return &CategorySvc{
		store: c,
	}
}


// type Po struct {
// 	idd int64
// }

// type co struct{
// s Po
// }