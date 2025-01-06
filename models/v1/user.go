// By Saif Hamdan, Team Lead
// Date: 2025/1/6
package v1

type User struct {
	Id           int    `json:"id" gorm:"column:id;primary_key"`
	Username     string `json:"username" gorm:"column:username"`
	UserPassword string `json:"user_password" gorm:"column:user_password"`
	FirstName    string `json:"first_name" gorm:"column:first_name"`
	FamilyName   string `json:"family_name" gorm:"column:family_name"`
	Email        string `json:"email" gorm:"column:email"`
	Phone        string `json:"phone" gorm:"column:phone"`
	CommonModel
}
