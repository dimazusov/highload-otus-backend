package auth_token

const TableName = "auth_token"

const timeExpiredSeconds = 2592000

type AuthToken struct {
	Token     string `json:"token" db:"token"`
	UserId    uint   `json:"user_id" db:"user_id"`
	ExpiredAt uint   `json:"expired_at" db:"expired_at"`
}

func (m AuthToken) TableName() string {
	return TableName
}
