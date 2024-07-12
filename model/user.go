package model

type Users struct {
	Id       int    `gorm:"type:int;primaryKey" json:"id"`
	Avatar   string `gorm:"type:string;size:255" json:"avatar"`
	Username string `gorm:"type:string;size:30" json:"userName"`
}
