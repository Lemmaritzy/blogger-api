package blog

import (
	"encoding/json"
	"fmt"
	pg_db "github.com/lemmaritzy/blogger/api/adapters/db/postgres"
	"github.com/lemmaritzy/blogger/internal/entities"
)

var (
	query string
	errDb error
)

func CreateBlogModel(userId string, d entities.Blog) (Record int64, status int64) {
	query = fmt.Sprintf("INSERT INTO blogs (title,description,content,creator) values ($1,$2,$3,$4)")
	errDb = pg_db.DB.QueryRow(pg_db.Ctx, query, d.Title, d.Description, d.Content, userId).Scan(&status)

	if errDb != nil {

		fmt.Println(errDb.Error())
	}
	return
}

func LoadBlogs() (Record int64, blogs json.RawMessage) {
	query = "select json_agg(blogs) from blogs"
	errDb = pg_db.DB.QueryRow(pg_db.Ctx, query).Scan(&blogs)

	if errDb != nil {
		fmt.Println(errDb.Error())
	}
	return
}

func LoadBlog(blogId int64) (Record int64, blog json.RawMessage) {
	query = fmt.Sprintf("select json_agg(blogs) from blogs where id=$1")
	errDb = pg_db.DB.QueryRow(pg_db.Ctx, query, blogId).Scan(&blog)

	if errDb != nil {
		fmt.Println(errDb.Error())
	}
	return
}

func UpdateBlog(blogId int64, d entities.Blog) (Record int64, Status int64) {
	query = fmt.Sprintf("update blogs set title=$1,description=$2,content=$3 where id=$4")
	errDb = pg_db.DB.QueryRow(pg_db.Ctx, query, d.Title, d.Description, d.Content, blogId).Scan(&Status)

	if errDb != nil {
		fmt.Println(errDb.Error())
	}

	return
}
