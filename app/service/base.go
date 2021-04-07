package service

import (
	"github.com/gogf/gf-demos/web"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)


type tableNameGetter interface {
	getTableName() string
}

// SimpleCurdService is for avoiding redundant code,
// and curd method provided by SimpleCurdService can not return accurate (list of)model,
// If result iof model type list needed, please use dao directly
type SimpleCurdService struct {
	apiParamsParser web.ApiParameterParser
}

func (s * SimpleCurdService) getTableName() string {
	return "unknown table"
}

func (s *SimpleCurdService) getGdbModel(tableName string) *gdb.Model {
	return g.DB("default").Model(tableName).Safe()
}

func (s *SimpleCurdService) Add(tng tableNameGetter, r ...interface{}) error {
	gdbModel := s.getGdbModel(tng.getTableName())
	if _, err := gdbModel.Insert(r); err != nil {
		return err
	}
	return nil
}

func (s *SimpleCurdService) List(tng tableNameGetter, r *ghttp.Request) (gdb.Result, error) {
	filterParams := s.apiParamsParser.GetFilterParams(r)
	searchConditions := s.apiParamsParser.GetSearchConditions(r)
	gdbModel := s.getGdbModel(tng.getTableName())
	return gdbModel.WherePri(searchConditions).FindAll(filterParams)
}

func (s *SimpleCurdService) Edit(tng tableNameGetter, id int, r interface{}) error {
	gdbModel := s.getGdbModel(tng.getTableName())
	if m, err := gdbModel.FindOne("id", id); err != nil || m == nil {
		return err
	}
	if res, err := gdbModel.Update(r, "id", id); err != nil {
		return err
	}
	return nil
}

func (s *SimpleCurdService) Delete(tng tableNameGetter, id int) error {
	gdbModel := s.getGdbModel(tng.getTableName())
	if m, err := gdbModel.Delete("id", id); err != nil || m == nil {
		return err
	}
	return nil
}

func (s *SimpleCurdService) Detail(tng tableNameGetter, id int) (gdb.Record, error) {
	gdbModel := s.getGdbModel(tng.getTableName())
	return gdbModel.FindOne("id", id)
}
