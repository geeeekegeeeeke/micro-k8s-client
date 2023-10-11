package service

type NodeService struct{}

func NewNodeService() *NodeService {
	return &NodeService{}
}

/*func (this *NodeService) User(param map[string]string) *dnc.User {
	user := &dnc.User{}
	dnc.UserPvder.GetUser(context.TODO(), []interface{}{param["id"]}, user)
	return user
}*/
