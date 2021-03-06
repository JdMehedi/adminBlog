package postgres

import (
	"context"
	"fmt"

	"blog/blog/storage"
)

const insertPost = `
	INSERT INTO posts(
		title,
		description,
		category_id,
		image
	) VALUES(
		:title,
		:description,
		:category_id,
		:image
	)RETURNING id;
`

func (s Storage) Create(ctx context.Context, t storage.Post) (int64, error) {
	stmt, err := s.db.PrepareNamed(insertPost)
	if err != nil {
		return 0, err
	}
	var id int64
	if err := stmt.Get(&id, t); err != nil {
		return 0, err
	}

	return id, nil
}

func (s Storage) List(ctx context.Context) ([]storage.Post, error) {
	var list []storage.Post

	err := s.db.Select(&list, "SELECT posts.*, categories.title as name FROM posts LEFT JOIN categories ON posts.category_id = categories.id ")
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s Storage) Get(ctx context.Context, id int64) (storage.Post, error) {

	var t storage.Post
	err := s.db.Get(&t, "SELECT posts.*, categories.title as name FROM posts LEFT JOIN categories ON posts.category_id = categories.id WHERE posts.id=$1", id)
	if err != nil {
		return t, err
	}
	return t, nil
}

const updatePost = `

UPDATE posts
	SET
		title = :title,
		description= :description,
		category_id= :category_id,
		image= :image

	WHERE 
	id = :id
	RETURNING *;
`

func (s *Storage) Update(ctx context.Context, t storage.Post) error {

	stmt, err := s.db.PrepareNamed(updatePost)

	if err != nil {
		return err
	}
	var ut storage.Post
	if err := stmt.Get(&ut, t); err != nil {
		return err
	}

	return err
}

func (s Storage) Delete(ctx context.Context, id int64) error {
	var data storage.Post
	return s.db.Get(&data, "DELETE FROM posts WHERE id=$1 RETURNING *", id)

}


func (s Storage) SearchPost(ctx context.Context,title string) ([]storage.Post, error ){
	// fmt.Println("done")
		var data []storage.Post
		// err := s.db.Select(&list, "SELECT posts.*, categories.title as name FROM posts LEFT JOIN categories ON posts.category_id = categories.id ")

		err:= s.db.Select(&data, "SELECT posts.*, categories.title as name FROM posts LEFT JOIN categories ON posts.category_id = categories.id  WHERE posts.title ILIKE '%%' || $1 || '%%' OR categories.title ILIKE '%%' || $1 || '%%'",title )
		fmt.Println("2222222222222222")
		fmt.Println(data)
		fmt.Println("222222222222222222222222")
		if err!=nil{
			return nil,err
		}
		
		return data,nil
	}