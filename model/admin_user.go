package model

type AdminUsers struct {
	Id       int    `gorm:"type:int;primaryKey" json:"id"`
	UserName string `gorm:"type:string;size:30" json:"userName"`
	Password string `gorm:"type:string;size:50" json:"password"`
}
