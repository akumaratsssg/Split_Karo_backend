package models

// Group represents a group in the EM_group table.
type Group struct {
	GroupID    uint   `gorm:"primaryKey;autoIncrement;unsigned;not null" json:"group_id"`
	GroupName  string `gorm:"type:varchar(45);not null;unique" json:"group_name"`
	GroupDesc  string `gorm:"type:varchar(45);default:null" json:"group_desc"`
	GroupAdmin uint   `gorm:"foreignKey:GroupAdmin; references:UserID; not null;unsigned" json:"group_admin"`

	// Define the relationship with User
	Admin User `gorm:"foreignKey:GroupAdmin;references:UserID" json:"group_admin_name"`
}

// TableName sets the table name for the Group struct
func (Group) TableName() string {
	return "EM_group"
}

type GroupInput struct {
	GroupName string `json:"group_name" binding:"required"`
	GroupDesc string `json:"group_desc" binding:"required"`
	// input the group_admin id into group details from logged in user details.
}
