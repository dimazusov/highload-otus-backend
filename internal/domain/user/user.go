package user

const TableName = "user"

const Female = 0
const Male = 1

type User struct {
	ID       uint   `json:"id" db:"id" gorm:"primary_key"`
	Name     string `json:"name" db:"name"`
	Age      uint   `json:"age" db:"age"`
	Sex      bool   `json:"sex" db:"sex"`
	City     string `json:"city" db:"city"`
	Interest string `json:"interest" db:"interest"`
}

func (m User) TableName() string {
	return TableName
}
