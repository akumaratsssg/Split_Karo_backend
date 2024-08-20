package services

import (
	"Expense_Management/backend/models"
	"Expense_Management/backend/repositories"
)

type GroupService struct {
	groupRepo *repositories.GroupRepository
}

func NewGroupService(groupRepo *repositories.GroupRepository) *GroupService {
	return &GroupService{groupRepo: groupRepo}
}

// GetGroupsByUser fetches the groups created by a user.
func (service *GroupService) GetGroupsByUser(userID uint) ([]models.Group, error) {
	return service.groupRepo.GetGroupsByUserID(userID)
}
