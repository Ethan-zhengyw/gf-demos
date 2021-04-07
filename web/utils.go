package web

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"strings"
)

const (
	filterPrefix = "f_"
	searchPrefix = "s_"
)

// ApiParameterParser help to parse request parameters
type ApiParameterParser struct {}

func (parser *ApiParameterParser) GetParamsWithPrefix(r *ghttp.Request, prefix string) g.Map {
	paramsWithPrefix := make(g.Map)
	for key, value := range r.GetQueryMap() {
		if strings.HasPrefix(key, prefix) {
			paramsWithPrefix[strings.TrimPrefix(key, prefix)] = value
		t }
	}
	return paramsWithPrefix
}

func (parser *ApiParameterParser) GetFilterParams(r *ghttp.Request) g.Map {
	return parser.GetParamsWithPrefix(r, filterPrefix)
}

func (parser *ApiParameterParser) GetSearchParams(r *ghttp.Request) g.Map {
	return parser.GetParamsWithPrefix(r, searchPrefix)
}

func (parser *ApiParameterParser) GetSearchConditions(r *ghttp.Request) g.Map {
	searchConditions := make(g.Map)
	for key, value := range parser.GetSearchParams(r) {
		searchConditions[fmt.Sprintf("`%v` like ?", key)] = fmt.Sprintf("%%%v%%", value)
	}
	return searchConditions
}
