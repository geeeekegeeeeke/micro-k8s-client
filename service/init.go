package service

import "gin-dubbogo-consumer/repo"

func init() {
	commonRepo = repo.NewCommonRepo()

}
