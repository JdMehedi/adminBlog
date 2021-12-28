package handler

import (
	"fmt"
	"log"
	"strconv"

	bgv "blog/gunk/v1/post"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) Index (rw http.ResponseWriter, r *http.Request) {


	res,err:= h.tc.ListPost(r.Context(), &bgv.ListPostRequest{})
	fmt.Printf("%#v",res)
	if err!=nil{
		log.Fatal(err)
	}
	if err:= h.templates.ExecuteTemplate(rw,"index-post.html", res); err !=nil{
		http.Error(rw, err.Error(),http.StatusInternalServerError)
		return
	}
}


func (h *Handler) deletePost (rw http.ResponseWriter, r *http.Request) {
vars := mux.Vars(r)
	id := vars["id"]
	
	if id == "" {
		http.Error(rw, "invalid request", http.StatusTemporaryRedirect)
		return
	}
		Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	_,err = h.tc.DeletePost(r.Context(),&bgv.DeletePostRequest{
		ID: Id,
	})
	if err!=nil{
		log.Fatal(err)
	}
	http.Redirect(rw,r, "/", http.StatusTemporaryRedirect)
}

func (h *Handler)searchPost(rw http.ResponseWriter, r *http.Request){

	if err:=r.ParseForm(); err !=nil{
		http.Error(rw, err.Error(),http.StatusInternalServerError)
		return
	}

	res:=r.FormValue("search")
	fmt.Println("3333333333333333333")
	fmt.Println(res)
	fmt.Println("2222222222222222")

	getRes,err:= h.tc.SearchPost(r.Context(),&bgv.SearchPostRequest{
		Title: res,
	})
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("################")
	fmt.Println(getRes)
	fmt.Println("@@@@@@@@@@@@@@@@@@")
	if err:= h.templates.ExecuteTemplate(rw,"index-post.html", getRes); err !=nil{
		http.Error(rw, err.Error(),http.StatusInternalServerError)
		return
	}
}
