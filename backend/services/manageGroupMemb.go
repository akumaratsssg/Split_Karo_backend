// /backend/services/groupService.go

package services

import (
	"Expense_Management/backend/models"
	"Expense_Management/backend/repositories"
	"errors"
)

func AddMemberByGroupNameAndEmail(groupName, email string, requestingUserID uint) error {
	user, err := repositories.FindUserByEmail(email)
	if err != nil {
		return errors.New("user not found")
	}

	group, err := repositories.GetGroupByName(groupName)
	if err != nil {
		return errors.New("group not found")
	}

	// fmt.Println(group.GroupAdmin, " ", requestingUserID) // check for add member error

	if group.GroupAdmin != requestingUserID {
		return errors.New("only group admin can add members")
	}

	participant := models.GroupParticipant{
		PartUserID:  user.UserID,
		PartGroupID: group.GroupID,
	}

	return repositories.AddParticipant(&participant)
}

func RemoveMemberByGroupNameAndEmail(groupName, email string, requestingUserID uint) error {
	user, err := repositories.FindUserByEmail(email)
	if err != nil {
		return errors.New("user not found")
	}

	group, err := repositories.GetGroupByName(groupName)
	if err != nil {
		return errors.New("group not found")
	}

	if group.GroupAdmin != requestingUserID {
		return errors.New("only group admin can remove members")
	}

	return repositories.RemoveParticipant(user.UserID, group.GroupID)
}
