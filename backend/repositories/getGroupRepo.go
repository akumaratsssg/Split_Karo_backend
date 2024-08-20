package repositories

import (
	"Expense_Management/backend/models"

	"gorm.io/gorm"
)

type GroupRepository struct {
	db *gorm.DB
}

func NewGroupRepository(db *gorm.DB) *GroupRepository {
	return &GroupRepository{db: db}
}

// GetGroupsByUserID fetches groups where the user is a participant (including as an admin).
func (repo *GroupRepository) GetGroupsByUserID(userID uint) ([]models.Group, error) {
	var groups []models.Group
	err := repo.db.Joins("JOIN EM_group_participants ON EM_group_participants.part_group_id = EM_group.group_id").
		Where("EM_group_participants.part_user_id = ?", userID).
		Preload("Admin", func(db *gorm.DB) *gorm.DB {
			return db.Select("user_id, user_name")
		}).
		Find(&groups).Error
	if err != nil {
		return nil, err
	}
	return groups, nil
}
