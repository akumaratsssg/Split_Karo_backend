// /internal/models/group_participant.go

package models

type GroupParticipant struct {
	PartUserID  uint  `gorm:"primaryKey"`
	PartGroupID uint  `gorm:"primaryKey"`
	User        User  `gorm:"foreignKey:PartUserID;constraint:OnDelete:CASCADE;"`
	Group       Group `gorm:"foreignKey:PartGroupID;constraint:OnDelete:CASCADE;"`
}

// TableName sets the table name for the Group struct
func (GroupParticipant) TableName() string {
	return "EM_group_participants"
}
