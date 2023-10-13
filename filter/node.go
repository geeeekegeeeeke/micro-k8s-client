package filter

import "github.com/gin-gonic/gin"

type NodeFilter struct {
	c *gin.Context
}

func NewNodeFilter(c *gin.Context) *NodeFilter {
	return &NodeFilter{c: c}
}
func (this *NodeFilter) NodeInfo() map[string]string {
	return map[string]string{
		"name": this.c.DefaultQuery("name", "master"),
	}
}
