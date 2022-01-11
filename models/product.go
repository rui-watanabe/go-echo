package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	Price     uint      `json:"price" gorm:"not null;default:0"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *Product) GetById(id uint) *gorm.DB {
	return DB.Where("id=?", id).First(&p)
}

func (p *Product) Save() (tx *gorm.DB) {
	return DB.Save(&p)
}

func (p *Product) Create() *gorm.DB {
	return DB.Create(&p)
}

func (p *Product) Updates() *gorm.DB {
	return DB.Model(&p).Updates(p)
}

func (p *Product) Delete() (tx *gorm.DB) {
	return DB.Delete(&p)
}

func (p *Product) DeleteById(id uint) *gorm.DB {
	return DB.Where("id = ?", id).Delete(&p)
}
