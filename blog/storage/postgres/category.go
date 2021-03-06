package postgres

import (
	"blog/blog/storage"
	"context"
	"fmt"
	"log"
)

const insertCategory = `
	INSERT INTO categories(
		title
	) VALUES(
		:title
	)RETURNING id;
`

func (s Storage) CreateCat(ctx context.Context, t storage.Category) (int64, error) {
	stmt, err := s.db.PrepareNamed(insertCategory)
	if err != nil {
		return 0, err
	}
	var id int64
	if err := stmt.Get(&id, t); err != nil {
		return 0, err
	}
	log.Println("Category ID: ", id)
	return id, nil
}

func (s Storage) ListCat(ctx context.Context) ([]storage.Category, error) {
	var l []storage.Category

	err := s.db.Select(&l, "SELECT *from categories")
	if err != nil {
		return nil, err
	}
	return l,nil
}

func (s Storage) GetCat(ctx context.Context, id int64)(storage.Category, error) {

	var t storage.Category
	err := s.db.Get(&t,"SELECT * from categories WHERE id=$1",id)
	if err != nil {
		return t, err
	}
	fmt.Println(t)
	return t, nil
}


const updateCat = `

UPDATE categories
	SET
		title = :title
	WHERE 
	id = :id
	RETURNING *;
`

func (s *Storage) UpdateCat(ctx context.Context, t storage.Category) error{

	stmt, err := s.db.PrepareNamed(updateCat)
	log.Println(stmt)

	if err != nil {
		return  err
	}
	var ut storage.Category
	if err := stmt.Get(&ut,t); err != nil {
		return err
	}
	return err
}

func (s Storage) DeleteCat(ctx context.Context, id int64) error {
	// fmt.Println("done")
		var data storage.Category
		return s.db.Get(&data, "DELETE FROM categories WHERE id=$1 RETURNING *", id)
	
	}
func (s Storage) SearchCategory(ctx context.Context,title string) ([]storage.Category, error ){
	// fmt.Println("done")
		var data []storage.Category

		err:= s.db.Select(&data, "Select * FROM categories WHERE title ILIKE '%%' || $1 || '%%'",title )
		
		if err!=nil{
			return nil,err
		}
		return data,nil
	
	}
	