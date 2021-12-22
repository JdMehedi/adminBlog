package postgres

import (
	"blog/blog/storage"
	"context"
	"log"
	"testing"

)

func TestCreatePost(t *testing.T) {
	s := newTestStorage(t)

	tests := []struct {
		name    string
		in      storage.Post
		want    int64
		wantErr bool
	}{
		{
			name: "CREATE_BLOG_SUCCESS",
			in: storage.Post{
				Title:       "This is title",
				Description: "This is description",
				CategoryId: 1,
			},
			want: 1,
		},
		{
			name: "IF_NOT_UNIQUE",
			in: storage.Post{
				Title:       "This is title",
				Description: "This is description",
				CategoryId: 2,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Create(context.Background(), tt.in)
			log.Printf("%#v", got)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
