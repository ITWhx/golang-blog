/**
*
*@author 吴昊轩
*@create 2020-04-1519:32
 */
package service

import (
	"golang-blog/model"
	"time"
)

type BlogService struct {
}

var bid int64 = 1

func NewBlogService() (blogservice *BlogService) {
	blogservice = &BlogService{}
	return
}
func (b *BlogService) CreateBlog(uid int64, content string) *model.Blog {

	blog := &model.Blog{bid, content, time.Now().Unix(), uid}
	bid++
	return blog
}
