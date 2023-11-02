package model

type K8sCluster struct {
	BaseModel
	GroupID     uint   `gorm:"type:decimal" json:"group_id"`
	Name        string `gorm:"type:varchar(64);not null" json:"name"`
	User        string `gorm:"type:varchar(64);not null" json:"user"`
	KubeConfig  string `gorm:"type:varchar(1024);not null" json:"kubeConfig"`
	Description string `gorm:"type:varchar(256)" json:"description"`
}
