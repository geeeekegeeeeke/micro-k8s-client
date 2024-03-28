package service

import (
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"micro-k8s-client/constant"
	"micro-k8s-client/dto"
	"micro-k8s-client/model"
)

type IK8sClusterService interface {
	List(req dto.ClusterSearch) ([]dto.ClusterInfo, error)
	Create(req dto.ClusterCreate) error
	Update(req dto.ClusterUpdate) error
	Delete(id uint) error
	Get(id uint) (dto.ClusterInfo, error)
}

// type K8sClusterService struct{}
//
//	func NewIK8sClusterService() IK8sClusterService {
//		return &K8sClusterService{}
//	}
type K8sClusterService struct{}

func NewIK8sClusterService() IK8sClusterService {
	return &K8sClusterService{}
}

func (u *K8sClusterService) List(req dto.ClusterSearch) ([]dto.ClusterInfo, error) {
	groups, err := k8sClusterRepo.GetList(commonRepo.WithOrderBy("created_at desc"))
	if err != nil {
		return nil, constant.ErrRecordNotFound
	}
	var dtoClusterInfo []dto.ClusterInfo
	for _, group := range groups {
		var item dto.ClusterInfo
		if err := copier.Copy(&item, &group); err != nil {
			return nil, errors.WithMessage(constant.ErrStructTransform, err.Error())
		}
		dtoClusterInfo = append(dtoClusterInfo, item)
	}
	return dtoClusterInfo, err
}

func (u *K8sClusterService) Create(req dto.ClusterCreate) error {
	group, _ := k8sClusterRepo.Get(commonRepo.WithByName(req.Name))
	if group.ID != 0 {
		return constant.ErrRecordExist
	}
	if err := copier.Copy(&group, &req); err != nil {
		return errors.WithMessage(constant.ErrStructTransform, err.Error())
	}
	if err := k8sClusterRepo.Create(&group); err != nil {
		return err
	}
	return nil
}

func (u *K8sClusterService) Update(req dto.ClusterUpdate) error {
	upMap := make(map[string]interface{})
	upMap["name"] = req.Name
	upMap["KubeConfig"] = req.KubeConfig
	upMap["description"] = req.Description
	return k8sClusterRepo.Update(req.ID, upMap)
}
func (u *K8sClusterService) Delete(id uint) error {
	group, _ := k8sClusterRepo.Get(commonRepo.WithByID(id))
	fmt.Println("group")
	fmt.Println(group)
	if group.ID == 0 {
		return constant.ErrRecordNotFound
	}
	return k8sClusterRepo.Delete(commonRepo.WithByID(id))

}
func (u *K8sClusterService) Get(id uint) (dto.ClusterInfo, error) {
	group, _ := k8sClusterRepo.Get(commonRepo.WithByID(id))
	if group.ID == 0 {
		return dto.ClusterInfo{}, constant.ErrRecordNotFound
	}
	baseModel := model.BaseModel{CreatedAt: group.CreatedAt}
	return dto.ClusterInfo{ID: group.ID, Name: group.Name, Description: group.Description, KubeConfig: group.KubeConfig, BaseModel: baseModel}, nil
}
