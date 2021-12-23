package postgres

import (
	"blog/blog/storage"
	"context"
	"log"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCreateCat(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		in      storage.Category
		want    int64
		wantErr bool
	}{
		{
			name: "CREATE_CATEGORY_SUCCESS",
			in: storage.Category{
				ID:    1,
				Title: "This is Category 1",
			},
			want: 1,
		},
		{
			name: "CREATE_CATEGORY_SUCCESS",
			in: storage.Category{
				ID:    2,
				Title: "This is Category 2",
			},
			want: 2,
		},
		{
			name: "IF_NOT_UNIQUE",
			in: storage.Category{
				Title: "This is Category 1",
			},
			wantErr: true,
		},

		// {
		// 	name: "IF_EMPTY",
		// 	in: storage.Category{
		// 		Title: "",
		// 	},
		// 	wantErr:true,
		// },
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.CreateCat(context.Background(), tt.in)
			log.Printf("%#v", got)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.CreateCat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.CreateCat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListCat(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		want    []storage.Category
		wantErr bool
	}{
		{
			name: "GET_ALL_CATEGORY_SUCCESS",
			want: []storage.Category{
				{
					ID:    1,
					Title: "This is Category 1",
				},
				{
					ID:    2,
					Title: "This is Category 2",
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			gotList, err := s.ListCat(context.Background())

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

func TestGetCat(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		in      int64
		want    storage.Category
		wantErr bool
	}{
		{
			name: "GET_CATEGORY_SUCCESS",
			in:   1,
			want: storage.Category{ID: 1,
				Title: "This is Category 1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetCat(context.Background(), tt.in)
			// log.Printf("%#v", got)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.CreateCat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.CreateCat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateCat(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		in      storage.Category
		wantErr bool
	}{
		{
			name: "GET_CATEGORY_SUCCESS",
			in: storage.Category{
				ID:    1,
				Title: "This is category 4",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			 err := s.UpdateCat(context.Background(), tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.CreateCat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			
		})
	}
}

func TestDeleteCat(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		in      int64
		wantErr bool
	}{
		{
			name: "DELETE_CATEGORY_SUCCESS",
			in:   1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetCat(context.Background(), tt.in)
			log.Printf("%#v", got)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.CreateCat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			
		})
	}
}