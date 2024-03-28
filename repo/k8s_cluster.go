package repo

import (
	"fmt"
	"micro-k8s-client/global"
	"micro-k8s-client/model"
)

type K8sClusterRepo struct{}

type IK8sClusterRepo interface {
	Get(opts ...DBOption) (model.K8sCluster, error)
	GetList(opts ...DBOption) ([]model.K8sCluster, error)
	Create(group *model.K8sCluster) error
	Update(id uint, vars map[string]interface{}) error
	Delete(opts ...DBOption) error
	CancelDefault(groupType string) error
	WithByIsDefault(isDefault bool) DBOption
}

func NewIK8sClusterRepo() IK8sClusterRepo {
	return &K8sClusterRepo{}
}

func (k K8sClusterRepo) Get(opts ...DBOption) (model.K8sCluster, error) {
	var k8sCluster model.K8sCluster
	db := global.DB
	for _, opt := range opts {
		db = opt(db)
	}
	err := db.First(&k8sCluster).Error
	return k8sCluster, err
}

func (k K8sClusterRepo) GetList(opts ...DBOption) ([]model.K8sCluster, error) {
	var k8sCluster []model.K8sCluster
	db := global.DB.Model(&model.K8sCluster{})
	for _, opt := range opts {
		db = opt(db)
	}
	err := db.Find(&k8sCluster).Error
	return k8sCluster, err
}

func (k K8sClusterRepo) Create(group *model.K8sCluster) error {
	return global.DB.Create(group).Error
}

func (k K8sClusterRepo) Update(id uint, vars map[string]interface{}) error {
	return global.DB.Model(&model.K8sCluster{}).Where("id = ?", id).Updates(vars).Error
}

func (k K8sClusterRepo) Delete(opts ...DBOption) error {
	db := global.DB
	for _, opt := range opts {
		db = opt(db)
	}
	fmt.Println("db")
	fmt.Println(db)
	return db.Delete(&model.K8sCluster{}).Error
}

func (k K8sClusterRepo) CancelDefault(groupType string) error {
	//TODO implement me
	panic("implement me")
}

func (k K8sClusterRepo) WithByIsDefault(isDefault bool) DBOption {
	//TODO implement me
	panic("implement me")
}
