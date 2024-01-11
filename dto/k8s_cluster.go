package dto

import "gin-dubbogo-consumer/model"

type ClusterCreate struct {
	ID   uint   `json:"id"`
	Name string `json:"name" validate:"required"`
	//Name      string `json:"name"`
	GroupID     string `json:"groupId"`
	UserId      string `json:"userId"`
	Description string `json:"description"`
	KubeConfig  string `json:"kubeConfig"  validate:"required"`
}

type ClusterSearch struct {
	Name string `json:"name"`
}

type ClusterUpdate struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	//Name      string `json:"name"`
	GroupID     string `json:"groupId"`
	UserId      string `json:"userId"`
	Description string `json:"description"`
	KubeConfig  string `json:"kubeConfig"`
}

type ClusterInfo struct {
	model.BaseModel
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	KubeConfig  string `json:"kubeConfig"`
	GroupID     string `json:"groupId"`
	UserId      string `json:"userId"`

	//IsDefault bool   `json:"isDefault"`
	//GroupID     uint   `gorm:"type:decimal" json:"group_id"`
	//User        string `gorm:"type:varchar(64);not null" json:"user"`
}
