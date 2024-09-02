package models

// ParamList :
type ParamList struct {
	Page       int    `json:"page" query:"page" valid:"Required"`
	PerPage    int    `json:"perpage" query:"perpage" valid:"Required"`
	Search     string `json:"search,omitempty" query:"search" `
	InitSearch string `json:"initsearch,omitempty" query:"initsearch"`
	SortField  string `json:"sortfield,omitempty" query:"sortfield"`
}

type ParamDynamicList struct {
	ParamList
	MenuUrl   string `json:"menu_url" valid:"Required"`
	LineNo    int    `json:"line_no,omitempty"`
	ParamView string `json:"param_view,omitempty"`
}

type PostMulti struct {
	MenuUrl string      `json:"menu_url" valid:"Required"`
	LineNo  int         `json:"line_no,omitempty"`
	InData  interface{} `json:"in_data,omitempty"`
}
