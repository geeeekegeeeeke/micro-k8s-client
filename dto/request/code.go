package request

type CodeCLoneReq struct {
	Url      string `form:"url" json:"url"`
	Directoy string `form:"directoy" json:"directoy"`
	Rewrite  bool   `form:"rewrite" json:"rewrite"`
}
