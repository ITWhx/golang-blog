/**
*
*@author 吴昊轩
*@create 2020-04-1615:43
 */
package main

import "golang-blog/db"

type UserProcessor struct {
	Uid string
}

func (u *UserProcessor) GetFansIds() (fansIds []string, err error) {
	fansIds, err = db.G_redisClient.SMembers("fans:" + u.Uid).Result()
	return
}
