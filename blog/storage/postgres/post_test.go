package postgres

import (
	"blog/blog/storage"
	"context"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
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
				CategoryId:  1,
			},
			want: 1,
		},
		{
			name: "CREATE_BLOG_SUCCESS",
			in: storage.Post{
				Title:       "This is title 2",
				Description: "This is description 2",
				CategoryId:  1,
			},
			want: 2,
		},
		{
			name: "IF_NOT_UNIQUE",
			in: storage.Post{
				Title:       "This is title",
				Description: "This is description",
				CategoryId:  2,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Create(context.Background(), tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Storage.Create() = got -, want + %v", cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestListPost(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		want    []storage.Post
		wantErr bool
	}{
		{
			name: "GET_LIST_POST_SUCCESS",
			want: []storage.Post{
				{
					ID:           1,
					Title:        "This is title",
					Description:  "This is description",
					CategoryName: "This is category 4",
					CategoryId:   1,
				},
				{
					ID:           2,
					Title:        "This is title 2",
					Description:  "This is description 2",
					CategoryName: "This is category 4",
					CategoryId:   1,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			gotList, err := s.List(context.Background())

			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			wantList := tt.want

			sort.Slice(wantList, func(i, j int) bool {
				return wantList[i].ID < wantList[j].ID
			})

			sort.Slice(gotList, func(i, j int) bool {
				return gotList[i].ID < gotList[j].ID
			})

			for i, got := range gotList {

				if !cmp.Equal(got, wantList[i]) {
					t.Errorf("Diff: got -, want += %v", cmp.Diff(got, wantList[i]))
				}

			}
		})
	}
}

func TestGetPost(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		in      int64
		want    storage.Post
		wantErr bool
	}{
		{
			name: "GET_POST_SUCCESS",
			in:   1,
			want: storage.Post{
				ID:           1,
				Title:        "This is title",
				Description:  "This is description",
				CategoryName: "This is category 4",
				CategoryId:   1,
			},
			
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Get(context.Background(), tt.in)
			// log.Printf("%#v", got)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.CreateCat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got,tt.want) {
				t.Errorf("Diff: got - want + %v", cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestUpdatePost(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		in      storage.Post
		want    int64
		wantErr bool
	}{
		{
			name: "GET_POST_SUCCESS",
			in: storage.Post{
				ID:           1,
				Title:        "This is post",
				Description:  "This is description post",
				CategoryName: "This is category 4",
				CategoryId:   1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			 err := s.Update(context.Background(), tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.CreateCat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			
		})
	}
}

func TestDeletePost(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		in      int64
		want    int64
		wantErr bool
	}{
		{
			name: "DELETE_POST_SUCCESS",
			in:   1,
			want: 1,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			 err := s.Delete(context.Background(), tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.CreateCat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			
		})
	}
}
