package service

import "gin-dubbogo-consumer/repo"

var (
	commonRepos    = repo.NewCommonRepo()
	k8sClusterRepo = repo.NewIK8sClusterRepo()
)
