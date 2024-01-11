package model

type K8sCluster struct {
	BaseModel
	GroupID     uint   `gorm:"type:decimal" json:"groupId"`
	Name        string `gorm:"type:varchar(64);not null" json:"name"`
	UserId      string `gorm:"type:varchar(64);not null" json:"userId"`
	KubeConfig  string `gorm:"type:varchar(1024);not null" json:"kubeConfig"`
	Description string `gorm:"type:varchar(256)" json:"description"`
}
