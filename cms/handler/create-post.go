package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gorilla/mux"

	bgvc "blog/gunk/v1/category"
	bgv "blog/gunk/v1/post"
)

type formData struct {
	Post     Post
	Category []Category
	Errors   map[string]string
}

type Post struct {
	ID           int64
	Title        string
	Description  string
	CategoryId   int64
	Image        string
	IsCompleted  bool
	CategoryName string
}

func (c *Post) validate() error {

	return validation.ValidateStruct(c,
		validation.Field(&c.Title,
			validation.Required.Error("This filed cannot be null"),
		),
		
	
	)
}

func (h *Handler) createPost(rw http.ResponseWriter, r *http.Request) {
	Post := Post{}
	Errors := map[string]string{}

	res, err := h.cc.ListCategory(r.Context(), &bgvc.ListCategoryRequest{})
	if err != nil {
		log.Fatal(err)
	}
	category := []Category{}
	for _, value := range res.Category {
		category = append(category, Category{
			ID:    value.ID,
			Title: value.Title,
		})
	}
	h.loadCreatedPostForm(rw, category, Post, Errors)
}
func (h *Handler) storePost(rw http.ResponseWriter, r *http.Request) {

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	var post Post

	if err := h.decoder.Decode(&post, r.PostForm); err != nil {

		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := h.cc.ListCategory(r.Context(), &bgvc.ListCategoryRequest{})
	if err != nil {
		log.Fatal(err)
	}
	category := []Category{}
	for _, value := range res.Category {
		category = append(category, Category{
			ID:    value.ID,
			Title: value.Title,
		})
	}

	file, _, err := r.FormFile("Image")
	if err != nil {
		Errors:= map[string]string{
			"Image":"This field is required",
		}
		h.loadCreatedPostForm(rw, category, post, Errors)
		return
	}
	defer file.Close()

	var img = "image-*.png"
	tempFile, err := ioutil.TempFile("cms/assets/images", img)
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	tempFile.Write(fileBytes)
	image := tempFile.Name()

	

	if err := post.validate(); err != nil {
		vErrors, ok := err.(validation.Errors)
		if ok {
			vErrs := make(map[string]string)
			for key, value := range vErrors {
				vErrs[strings.Title(key)] = value.Error()

			}
			h.loadCreatedPostForm(rw, category, post, vErrs)
			return
		}

		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return

	}
	_, err = h.tc.CreatePost(r.Context(), &bgv.CreatePostRequest{
		Post: &bgv.Post{
			Title:        post.Title,
			Description:  post.Description,
			Image:        image,
			CategoryId:   post.CategoryId,
			CategoryName: post.CategoryName,
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	http.Redirect(rw, r, "/", http.StatusTemporaryRedirect)

}

func (h *Handler) loadCreatedPostForm(rw http.ResponseWriter, categories []Category, posts Post, errs map[string]string) {

	form := formData{
		Post:     posts,
		Category: categories,
		Errors:   errs,
	}
	if err := h.templates.ExecuteTemplate(rw, "create-post.html", form); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h *Handler) editPost(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		http.Error(rw, "invalid ", http.StatusTemporaryRedirect)
		return
	}

	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}


	ress, err := h.cc.ListCategory(r.Context(), &bgvc.ListCategoryRequest{})
	if err != nil {
		log.Fatal(err)
	}
	
	
	category := []Category{}
	for _, value := range ress.Category {
		category = append(category, Category{
			ID:    value.ID,
			Title: value.Title,
		})
	}

	Errors := map[string]string{}

	res, err := h.tc.GetPost(r.Context(), &bgv.GetPostRequest{
		ID: Id,
	})
	
	if err != nil {
		log.Fatal(err)
	}
	Post := Post{
		ID:           res.Post.ID,
		Title:        res.Post.Title,
		Description:  res.Post.Description,
		Image:        res.Post.Image,
		CategoryId:   res.Post.CategoryId,
		CategoryName: res.Post.CategoryName,
		IsCompleted:  res.Post.IsCompleted,
	}

	

	h.loadUpdatedPostForm(rw, Post, category, Errors)
}

func (h Handler) updatePost(rw http.ResponseWriter, r *http.Request) {
	
	cat, err := h.cc.ListCategory(r.Context(), &bgvc.ListCategoryRequest{})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	var category []Category
	for _, val := range cat.Category {
		category = append(category, Category{
			ID:     val.ID,
			Title: val.Title,
		})
	}

	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(rw, "invalid URL", http.StatusInternalServerError)
		return
	}
	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := h.tc.GetPost(r.Context(), &bgv.GetPostRequest{
		ID: Id,
	})

	if err != nil {
		http.Error(rw, "invalid URL", http.StatusInternalServerError)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(rw, "invalid URL", http.StatusInternalServerError)
		return
	}

	var post Post
	

	if err := h.decoder.Decode(&post, r.PostForm); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	file, _, err := r.FormFile("Image")
    
	var imageName string
	
    if err == nil {
		defer file.Close()
		tempFile, err := ioutil.TempFile("cms/assets/images", "image-*.png")
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		defer tempFile.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		
		tempFile.Write(fileBytes)
		
		imageName = tempFile.Name()

		if err := os.Remove(res.Post.Image); err != nil {
				http.Error(rw, "Unable to delete image", http.StatusInternalServerError)
				return
			}
	} else {
		imageName = res.Post.Image
	}

	if err := post.validate(); err != nil {
		vErrors, ok := err.(validation.Errors)
		if ok {
			vErrs := make(map[string]string)
			for key, value := range vErrors {
				vErrs[key] = value.Error()
			}
			h.loadUpdatedPostForm(rw, post, category, vErrs)
			return
		}
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = h.tc.UpdatePost(r.Context(), &bgv.UpdatePostRequest{
		Post: &bgv.Post{
			ID:           Id,
			Title:        post.Title,
			Description:  post.Description,
			CategoryId:   post.CategoryId,
			Image:        imageName,
		},
	})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(rw, r, "/", http.StatusTemporaryRedirect)
}



func (h *Handler) viewPost(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	Errors := map[string]string{}
	if id == "" {
		http.Error(rw, "invalid request", http.StatusTemporaryRedirect)
		return
	}
	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := h.cc.ListCategory(r.Context(), &bgvc.ListCategoryRequest{})
	if err != nil {
		log.Fatal(err)
	}
	category := []Category{}
	for _, value := range res.Category {
		category = append(category, Category{
			ID:    value.ID,
			Title: value.Title,
		})
	}

	ress, err := h.tc.GetPost(r.Context(), &bgv.GetPostRequest{
		ID: Id,
	})
	if err != nil {
		log.Fatal(err)
	}
	Post := Post{
		ID:           ress.Post.ID,
		Title:        ress.Post.Title,
		Description:  ress.Post.Description,
		Image:        ress.Post.Image,
		CategoryId:   ress.Post.CategoryId,
		CategoryName: ress.Post.CategoryName,
		IsCompleted:  ress.Post.IsCompleted,
	}

	h.loadShowForm(rw, Post, category, Errors)
}

func (h *Handler) loadUpdatedPostForm(rw http.ResponseWriter, posts Post, categories []Category, errs map[string]string) {

	form := formData{
		Post:     posts,
		Category: categories,
		Errors:   errs,
	}
	if err := h.templates.ExecuteTemplate(rw, "edit-post.html", form); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

}
func (h *Handler) loadShowForm(rw http.ResponseWriter, posts Post, categories []Category, errs map[string]string) {

	form := formData{
		Post:     posts,
		Category: categories,
		Errors:   errs,
	}
	if err := h.templates.ExecuteTemplate(rw, "show-post.html", form); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

}
