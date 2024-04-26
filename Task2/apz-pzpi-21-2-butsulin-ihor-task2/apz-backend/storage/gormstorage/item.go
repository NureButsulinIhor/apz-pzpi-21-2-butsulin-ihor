package gormstorage

import (
	"apz-backend/types/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (g GormStorage) AddItem(item models.Item) error {
	db := g.db

	result := db.Create(&item)

	return result.Error
}

func (g GormStorage) GetItem(id uint) (*models.Item, error) {
	db := g.db

	var item models.Item
	result := db.Preload(clause.Associations).First(&item, id)

	return &item, result.Error
}

func (g GormStorage) UpdateItem(itemID uint, name, description string) error {
	db := g.db

	result := db.Model(&models.Item{}).Updates(models.Item{
		Model:       gorm.Model{ID: itemID},
		Name:        name,
		Description: description,
	})

	return result.Error
}
