/**
*
*@author 吴昊轩
*@create 2020-04-1519:28
 */
package service

type UserService struct {
}

func (u *UserService) PublishBlog(uId int64, content string) error {
	blogservice := NewBlogService()
	blog := blogservice.CreateBlog(uId, content)

	return nil
}
