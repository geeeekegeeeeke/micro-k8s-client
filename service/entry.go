package service

import "micro-k8s-client/repo"

var (
	commonRepos    = repo.NewCommonRepo()
	k8sClusterRepo = repo.NewIK8sClusterRepo()
)
