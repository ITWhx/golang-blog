package model

import "github.com/jinzhu/gorm"

/**
 * Created by 吴昊轩 on 2020/4/15.
 */

type User struct {
	gorm.Model
	Uid  string
	Name string
}
