/**
*
*@author 吴昊轩
*@create 2020-04-1519:27
 */
package dao

type User struct {
}

func (u *User) getFans() []int64 {
	fansId := []int64{2, 3, 4}
	return fansId
}
