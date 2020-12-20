package model

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
)

type AccountProfile struct {
	Id         primitive.ObjectID `json:"id" bson:"id"`
	Name       string             `json:"name"`
	Role       string             `json:"role"`
	Mobile     string             `json:"mobile"`
	Department string             `json:"department"`
	Position   string             `json:"position"`
	Extra      string             `json:"extra"`
	Email      string             `json:"email"`
	Status     string             `json:"status"`
	Subscribed bool               `json:"subscribed"`
}

func (m *AccountProfile) IsAdmin() bool {
	return strings.Contains(m.Role, "admin")
}

type ResAccountProfile struct {
	Key     string          `json:"key"`
	Data    *AccountProfile `json:"data"`
	Message string          `json:"message"`
}

//获得错误信息
func (m *ResAccountProfile) GetErrMsg() string {
	return fmt.Sprintf("key=%s msg=%s", m.Key, m.Message)
}
