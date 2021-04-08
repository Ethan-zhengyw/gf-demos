package web

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"strconv"
	"strings"
)

const (
	filterPrefix = "f_"
	searchPrefix = "s_"
)

// ApiParameterParser help to parse request parameters
type ApiParameterParser struct { }

// List parameter
type ListParams struct {
	FilterParams     g.Map
	SearchConditions g.Map
	OrderName        string
	Order            string
	Limit            int
	Page             int
}

func (parser *ApiParameterParser) GetParamsWithPrefix(r *ghttp.Request, prefix string) g.Map {
	paramsWithPrefix := make(g.Map)
	for key, value := range r.GetQueryMap() {
		if strings.HasPrefix(key, prefix) {
			paramsWithPrefix[strings.TrimPrefix(key, prefix)] = value
		}
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

func (parser *ApiParameterParser) GetOrderParams(r *ghttp.Request) (orderName string, order string) {
	const (
		keyOrderName = "order_name"
		keyOrder = "order"
		defaultOrderName = "id"
		defaultOrder = "desc"
	)

	// default order option
	orderName, order = defaultOrderName, defaultOrder

	for key, value := range r.GetQueryMap() {
		if key == keyOrderName {
			orderName = fmt.Sprintf("%v", value)
		}
		if key == keyOrder {
			order = fmt.Sprintf("%v", value)
		}
	}

	return orderName, order
}

func (parser *ApiParameterParser) GetPageParams(r *ghttp.Request) (limit int, page int) {
	const (
		keyLimit = "limit"
		keyPage = "page"
		defaultLimit = 10
		defaultPage = 1
	)

	limit, page = defaultLimit, defaultPage

	for key, value := range r.GetQueryMap() {
		if key == keyLimit {
			limit, _ = strconv.Atoi(value.(string))
		}
		if key == keyPage {
			page, _ = strconv.Atoi(value.(string))
		}
	}

	return limit, page
}

func (parser *ApiParameterParser) GetListParams(r *ghttp.Request) *ListParams  {
	filterParams := parser.GetFilterParams(r)
	searchConditions := parser.GetSearchConditions(r)
	orderName, order := parser.GetOrderParams(r)
	limit, page := parser.GetPageParams(r)

	return &ListParams{
		FilterParams:     filterParams,
		SearchConditions: searchConditions,
		OrderName:        orderName, Order:order,
		Limit:limit, Page:page,
	}
}