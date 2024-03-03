// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/bezhai/multi-bot-task/biz/mysql/model"
)

func newUserModuleAccess(db *gorm.DB, opts ...gen.DOOption) userModuleAccess {
	_userModuleAccess := userModuleAccess{}

	_userModuleAccess.userModuleAccessDo.UseDB(db, opts...)
	_userModuleAccess.userModuleAccessDo.UseModel(&model.UserModuleAccess{})

	tableName := _userModuleAccess.userModuleAccessDo.TableName()
	_userModuleAccess.ALL = field.NewAsterisk(tableName)
	_userModuleAccess.UserID = field.NewInt32(tableName, "user_id")
	_userModuleAccess.ModuleAccessID = field.NewInt32(tableName, "module_access_id")

	_userModuleAccess.fillFieldMap()

	return _userModuleAccess
}

// userModuleAccess 关联用户和模块准入权限
type userModuleAccess struct {
	userModuleAccessDo userModuleAccessDo

	ALL            field.Asterisk
	UserID         field.Int32 // 用户ID
	ModuleAccessID field.Int32 // 模块准入权限ID

	fieldMap map[string]field.Expr
}

func (u userModuleAccess) Table(newTableName string) *userModuleAccess {
	u.userModuleAccessDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userModuleAccess) As(alias string) *userModuleAccess {
	u.userModuleAccessDo.DO = *(u.userModuleAccessDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userModuleAccess) updateTableName(table string) *userModuleAccess {
	u.ALL = field.NewAsterisk(table)
	u.UserID = field.NewInt32(table, "user_id")
	u.ModuleAccessID = field.NewInt32(table, "module_access_id")

	u.fillFieldMap()

	return u
}

func (u *userModuleAccess) WithContext(ctx context.Context) *userModuleAccessDo {
	return u.userModuleAccessDo.WithContext(ctx)
}

func (u userModuleAccess) TableName() string { return u.userModuleAccessDo.TableName() }

func (u userModuleAccess) Alias() string { return u.userModuleAccessDo.Alias() }

func (u userModuleAccess) Columns(cols ...field.Expr) gen.Columns {
	return u.userModuleAccessDo.Columns(cols...)
}

func (u *userModuleAccess) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userModuleAccess) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 2)
	u.fieldMap["user_id"] = u.UserID
	u.fieldMap["module_access_id"] = u.ModuleAccessID
}

func (u userModuleAccess) clone(db *gorm.DB) userModuleAccess {
	u.userModuleAccessDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userModuleAccess) replaceDB(db *gorm.DB) userModuleAccess {
	u.userModuleAccessDo.ReplaceDB(db)
	return u
}

type userModuleAccessDo struct{ gen.DO }

func (u userModuleAccessDo) Debug() *userModuleAccessDo {
	return u.withDO(u.DO.Debug())
}

func (u userModuleAccessDo) WithContext(ctx context.Context) *userModuleAccessDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userModuleAccessDo) ReadDB() *userModuleAccessDo {
	return u.Clauses(dbresolver.Read)
}

func (u userModuleAccessDo) WriteDB() *userModuleAccessDo {
	return u.Clauses(dbresolver.Write)
}

func (u userModuleAccessDo) Session(config *gorm.Session) *userModuleAccessDo {
	return u.withDO(u.DO.Session(config))
}

func (u userModuleAccessDo) Clauses(conds ...clause.Expression) *userModuleAccessDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userModuleAccessDo) Returning(value interface{}, columns ...string) *userModuleAccessDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userModuleAccessDo) Not(conds ...gen.Condition) *userModuleAccessDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userModuleAccessDo) Or(conds ...gen.Condition) *userModuleAccessDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userModuleAccessDo) Select(conds ...field.Expr) *userModuleAccessDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userModuleAccessDo) Where(conds ...gen.Condition) *userModuleAccessDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userModuleAccessDo) Order(conds ...field.Expr) *userModuleAccessDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userModuleAccessDo) Distinct(cols ...field.Expr) *userModuleAccessDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userModuleAccessDo) Omit(cols ...field.Expr) *userModuleAccessDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userModuleAccessDo) Join(table schema.Tabler, on ...field.Expr) *userModuleAccessDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userModuleAccessDo) LeftJoin(table schema.Tabler, on ...field.Expr) *userModuleAccessDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userModuleAccessDo) RightJoin(table schema.Tabler, on ...field.Expr) *userModuleAccessDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userModuleAccessDo) Group(cols ...field.Expr) *userModuleAccessDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userModuleAccessDo) Having(conds ...gen.Condition) *userModuleAccessDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userModuleAccessDo) Limit(limit int) *userModuleAccessDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userModuleAccessDo) Offset(offset int) *userModuleAccessDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userModuleAccessDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *userModuleAccessDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userModuleAccessDo) Unscoped() *userModuleAccessDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userModuleAccessDo) Create(values ...*model.UserModuleAccess) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userModuleAccessDo) CreateInBatches(values []*model.UserModuleAccess, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userModuleAccessDo) Save(values ...*model.UserModuleAccess) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userModuleAccessDo) First() (*model.UserModuleAccess, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserModuleAccess), nil
	}
}

func (u userModuleAccessDo) Take() (*model.UserModuleAccess, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserModuleAccess), nil
	}
}

func (u userModuleAccessDo) Last() (*model.UserModuleAccess, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserModuleAccess), nil
	}
}

func (u userModuleAccessDo) Find() ([]*model.UserModuleAccess, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserModuleAccess), err
}

func (u userModuleAccessDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserModuleAccess, err error) {
	buf := make([]*model.UserModuleAccess, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userModuleAccessDo) FindInBatches(result *[]*model.UserModuleAccess, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userModuleAccessDo) Attrs(attrs ...field.AssignExpr) *userModuleAccessDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userModuleAccessDo) Assign(attrs ...field.AssignExpr) *userModuleAccessDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userModuleAccessDo) Joins(fields ...field.RelationField) *userModuleAccessDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userModuleAccessDo) Preload(fields ...field.RelationField) *userModuleAccessDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userModuleAccessDo) FirstOrInit() (*model.UserModuleAccess, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserModuleAccess), nil
	}
}

func (u userModuleAccessDo) FirstOrCreate() (*model.UserModuleAccess, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserModuleAccess), nil
	}
}

func (u userModuleAccessDo) FindByPage(offset int, limit int) (result []*model.UserModuleAccess, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userModuleAccessDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userModuleAccessDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userModuleAccessDo) Delete(models ...*model.UserModuleAccess) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userModuleAccessDo) withDO(do gen.Dao) *userModuleAccessDo {
	u.DO = *do.(*gen.DO)
	return u
}
