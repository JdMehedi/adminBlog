package post

import (
	"blog/blog/storage"
	bgv "blog/gunk/v1/post"
	"context"
)

type PostCoreLink interface{
	Create(context.Context,storage.Post)(int64,error)
	List(context.Context)([]storage.Post, error)
	Get(context.Context, int64)(storage.Post, error)
	Update(context.Context, storage.Post) error
	Delete(context.Context,int64)error
	SearchPost(context.Context,string)([]storage.Post,error) 

}

type PostSvc struct {
	bgv.UnimplementedPostServiceServer
	store PostCoreLink
}

func NewPostSvc(s PostCoreLink) *PostSvc{
	return &PostSvc{
		store: s,
	}
}