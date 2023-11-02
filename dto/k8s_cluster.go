package dto

type ClusterCreate struct {
	ID   uint   `json:"id"`
	Name string `json:"name" validate:"required"`
	//Name      string `json:"name"`
	GroupID     string `json:"groupID"`
	User        string `json:"user"`
	Description string `json:"description"`
	KubeConfig  string `json:"kubeConfig"`
}

type ClusterSearch struct {
	Name string `json:"name"`
}

type ClusterUpdate struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	KubeConfig  string `json:"kubeConfig"`
}

type ClusterInfo struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	KubeConfig  string `json:"kubeConfig"`
	//IsDefault bool   `json:"isDefault"`
	//GroupID     uint   `gorm:"type:decimal" json:"group_id"`
	//User        string `gorm:"type:varchar(64);not null" json:"user"`
}
