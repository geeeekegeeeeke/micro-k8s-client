package service

import (
	"gin-dubbogo-consumer/constant"
	"gin-dubbogo-consumer/dto"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
)

type IK8sClusterService interface {
	List(req dto.ClusterSearch) ([]dto.ClusterInfo, error)
	Create(req dto.ClusterCreate) error
	Update(req dto.ClusterUpdate) error
	Delete(id uint) error
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
	groups, err := k8sClusterRepo.GetList(commonRepo.WithByName(req.Name), commonRepo.WithOrderBy("is_default desc"), commonRepo.WithOrderBy("created_at desc"))
	if err != nil {
		return nil, constant.ErrRecordNotFound
	}
	var dtoUsers []dto.ClusterInfo
	for _, group := range groups {
		var item dto.ClusterInfo
		if err := copier.Copy(&item, &group); err != nil {
			return nil, errors.WithMessage(constant.ErrStructTransform, err.Error())
		}
		dtoUsers = append(dtoUsers, item)
	}
	return dtoUsers, err
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
	//upMap["is_default"] = req.IsDefault

	return k8sClusterRepo.Update(req.ID, upMap)
}
func (u *K8sClusterService) Delete(id uint) error {
	group, _ := k8sClusterRepo.Get(commonRepo.WithByID(id))
	if group.ID == 0 {
		return constant.ErrRecordNotFound
	}
	return k8sClusterRepo.Delete(commonRepo.WithByID(id))
}
