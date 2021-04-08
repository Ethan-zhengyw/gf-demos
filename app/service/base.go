package service

import (
	"github.com/gogf/gf-demos/web"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type ListResult struct {
	Data  gdb.Result `json:"data"`
	Total int        `json:"total"`
	Limit int        `json:"limit"`
	Page  int        `json:"page"`
	Count int        `json:"count"`
}

type tableNameGetter interface {
	getTableName() string
}

type SimpleService struct {}

type SimpleAddService struct {
	SimpleService
}

type SimpleListService struct {
	SimpleService
	apiParamsParser web.ApiParameterParser
}

type SimpleDetailService struct {
	SimpleService
}

type SimpleEditService struct {
	SimpleService
}

type SimpleDeleteService struct {
	SimpleService
}

// SimpleCurdService is for avoiding redundant code,
// and curd method provided by SimpleCurdService can not return accurate (list of)model,
// If result iof model type list needed, please use dao directly
type SimpleCurdService struct {
	SimpleAddService
	SimpleListService
	SimpleDetailService
	SimpleEditService
	SimpleDeleteService
}

func (s *SimpleService) getTableName() string {
	return "unknown table"
}

func (s *SimpleService) getGdbModel(tableName string) *gdb.Model {
	return g.DB("default").Model(tableName).Safe(false)
}

func (s *SimpleAddService) Add(tng tableNameGetter, r ...interface{}) error {
	gdbModel := s.getGdbModel(tng.getTableName())
	if _, err := gdbModel.Insert(r); err != nil {
		return err
	}
	return nil
}

func (s *SimpleListService) List(tng tableNameGetter, r *ghttp.Request) (ListResult, error) {
	gdbModel := s.getGdbModel(tng.getTableName())

	listParams := s.apiParamsParser.GetListParams(r)
	lr := ListResult{Page: listParams.Page, Limit: listParams.Limit}

	// query Total
	if total, err := gdbModel.WherePri(listParams.FilterParams).WherePri(listParams.SearchConditions).Order(listParams.OrderName, listParams.Order).Count(); err != nil {
		return lr, err
	} else {
		lr.Total = total
	}

	// query Count
	if count, err := gdbModel.Page(listParams.Page, listParams.Limit).Count(); err != nil {
		return lr, err
	} else {
		lr.Count = count
	}

	// query Data
	if data, err := gdbModel.FindAll(); err != nil {
		return lr, err
	} else {
		lr.Data = data
	}

	return lr, nil
}

func (s *SimpleEditService) Edit(tng tableNameGetter, id int, r interface{}) error {
	gdbModel := s.getGdbModel(tng.getTableName())
	if m, err := gdbModel.FindOne("id", id); err != nil || m == nil {
		return err
	} else {
		if _, err := gdbModel.Update(r, "id", id); err != nil {
			return err
		}
	}
	return nil
}

func (s *SimpleDeleteService) Delete(tng tableNameGetter, id int) error {
	gdbModel := s.getGdbModel(tng.getTableName())
	if m, err := gdbModel.Delete("id", id); err != nil || m == nil {
		return err
	}
	return nil
}

func (s *SimpleDetailService) Detail(tng tableNameGetter, id int) (gdb.Record, error) {
	gdbModel := s.getGdbModel(tng.getTableName())
	return gdbModel.FindOne("id", id)
}
