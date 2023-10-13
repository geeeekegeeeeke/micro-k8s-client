package filter

import "github.com/gin-gonic/gin"

type PodFilter struct {
	c *gin.Context
}

func NewPodFilter(c *gin.Context) *PodFilter {
	return &PodFilter{c: c}
}
func (this *PodFilter) PodInfo() map[string]string {
	return map[string]string{
		"name": this.c.DefaultQuery("name", "master"),
	}
}
