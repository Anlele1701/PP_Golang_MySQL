package common

type Paging struct {
	Page  int   `json:"page" form:"page"`
	Size  int   `json:"size" form:"size"`
	Total int64 `json:"total" form:"-"`
}

func (p *Paging) Offset() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Size <= 0 || p.Size >= 100 {
		p.Size = 10
	}
}
