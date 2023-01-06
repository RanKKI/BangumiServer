// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/bangumi/server/dal/dao"
)

func newPrivateMessage(db *gorm.DB, opts ...gen.DOOption) privateMessage {
	_privateMessage := privateMessage{}

	_privateMessage.privateMessageDo.UseDB(db, opts...)
	_privateMessage.privateMessageDo.UseModel(&dao.PrivateMessage{})

	tableName := _privateMessage.privateMessageDo.TableName()
	_privateMessage.ALL = field.NewAsterisk(tableName)
	_privateMessage.ID = field.NewUint32(tableName, "msg_id")
	_privateMessage.SenderID = field.NewUint32(tableName, "msg_sid")
	_privateMessage.ReceiverID = field.NewUint32(tableName, "msg_rid")
	_privateMessage.Folder = field.NewString(tableName, "msg_folder")
	_privateMessage.New = field.NewBool(tableName, "msg_new")
	_privateMessage.Title = field.NewString(tableName, "msg_title")
	_privateMessage.CreatedTime = field.NewUint32(tableName, "msg_dateline")
	_privateMessage.Content = field.NewString(tableName, "msg_message")
	_privateMessage.MainMessageID = field.NewUint32(tableName, "msg_related_main")
	_privateMessage.RelatedMessageID = field.NewUint32(tableName, "msg_related")
	_privateMessage.DeletedBySender = field.NewBool(tableName, "msg_sdeleted")
	_privateMessage.DeletedByReceiver = field.NewBool(tableName, "msg_rdeleted")

	_privateMessage.fillFieldMap()

	return _privateMessage
}

type privateMessage struct {
	privateMessageDo privateMessageDo

	ALL               field.Asterisk
	ID                field.Uint32
	SenderID          field.Uint32
	ReceiverID        field.Uint32
	Folder            field.String
	New               field.Bool
	Title             field.String
	CreatedTime       field.Uint32
	Content           field.String
	MainMessageID     field.Uint32
	RelatedMessageID  field.Uint32
	DeletedBySender   field.Bool
	DeletedByReceiver field.Bool

	fieldMap map[string]field.Expr
}

func (p privateMessage) Table(newTableName string) *privateMessage {
	p.privateMessageDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p privateMessage) As(alias string) *privateMessage {
	p.privateMessageDo.DO = *(p.privateMessageDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *privateMessage) updateTableName(table string) *privateMessage {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewUint32(table, "msg_id")
	p.SenderID = field.NewUint32(table, "msg_sid")
	p.ReceiverID = field.NewUint32(table, "msg_rid")
	p.Folder = field.NewString(table, "msg_folder")
	p.New = field.NewBool(table, "msg_new")
	p.Title = field.NewString(table, "msg_title")
	p.CreatedTime = field.NewUint32(table, "msg_dateline")
	p.Content = field.NewString(table, "msg_message")
	p.MainMessageID = field.NewUint32(table, "msg_related_main")
	p.RelatedMessageID = field.NewUint32(table, "msg_related")
	p.DeletedBySender = field.NewBool(table, "msg_sdeleted")
	p.DeletedByReceiver = field.NewBool(table, "msg_rdeleted")

	p.fillFieldMap()

	return p
}

func (p *privateMessage) WithContext(ctx context.Context) *privateMessageDo {
	return p.privateMessageDo.WithContext(ctx)
}

func (p privateMessage) TableName() string { return p.privateMessageDo.TableName() }

func (p privateMessage) Alias() string { return p.privateMessageDo.Alias() }

func (p *privateMessage) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *privateMessage) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 12)
	p.fieldMap["msg_id"] = p.ID
	p.fieldMap["msg_sid"] = p.SenderID
	p.fieldMap["msg_rid"] = p.ReceiverID
	p.fieldMap["msg_folder"] = p.Folder
	p.fieldMap["msg_new"] = p.New
	p.fieldMap["msg_title"] = p.Title
	p.fieldMap["msg_dateline"] = p.CreatedTime
	p.fieldMap["msg_message"] = p.Content
	p.fieldMap["msg_related_main"] = p.MainMessageID
	p.fieldMap["msg_related"] = p.RelatedMessageID
	p.fieldMap["msg_sdeleted"] = p.DeletedBySender
	p.fieldMap["msg_rdeleted"] = p.DeletedByReceiver
}

func (p privateMessage) clone(db *gorm.DB) privateMessage {
	p.privateMessageDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p privateMessage) replaceDB(db *gorm.DB) privateMessage {
	p.privateMessageDo.ReplaceDB(db)
	return p
}

type privateMessageDo struct{ gen.DO }

func (p privateMessageDo) Debug() *privateMessageDo {
	return p.withDO(p.DO.Debug())
}

func (p privateMessageDo) WithContext(ctx context.Context) *privateMessageDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p privateMessageDo) ReadDB() *privateMessageDo {
	return p.Clauses(dbresolver.Read)
}

func (p privateMessageDo) WriteDB() *privateMessageDo {
	return p.Clauses(dbresolver.Write)
}

func (p privateMessageDo) Session(config *gorm.Session) *privateMessageDo {
	return p.withDO(p.DO.Session(config))
}

func (p privateMessageDo) Clauses(conds ...clause.Expression) *privateMessageDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p privateMessageDo) Returning(value interface{}, columns ...string) *privateMessageDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p privateMessageDo) Not(conds ...gen.Condition) *privateMessageDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p privateMessageDo) Or(conds ...gen.Condition) *privateMessageDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p privateMessageDo) Select(conds ...field.Expr) *privateMessageDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p privateMessageDo) Where(conds ...gen.Condition) *privateMessageDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p privateMessageDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *privateMessageDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p privateMessageDo) Order(conds ...field.Expr) *privateMessageDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p privateMessageDo) Distinct(cols ...field.Expr) *privateMessageDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p privateMessageDo) Omit(cols ...field.Expr) *privateMessageDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p privateMessageDo) Join(table schema.Tabler, on ...field.Expr) *privateMessageDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p privateMessageDo) LeftJoin(table schema.Tabler, on ...field.Expr) *privateMessageDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p privateMessageDo) RightJoin(table schema.Tabler, on ...field.Expr) *privateMessageDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p privateMessageDo) Group(cols ...field.Expr) *privateMessageDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p privateMessageDo) Having(conds ...gen.Condition) *privateMessageDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p privateMessageDo) Limit(limit int) *privateMessageDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p privateMessageDo) Offset(offset int) *privateMessageDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p privateMessageDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *privateMessageDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p privateMessageDo) Unscoped() *privateMessageDo {
	return p.withDO(p.DO.Unscoped())
}

func (p privateMessageDo) Create(values ...*dao.PrivateMessage) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p privateMessageDo) CreateInBatches(values []*dao.PrivateMessage, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p privateMessageDo) Save(values ...*dao.PrivateMessage) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p privateMessageDo) First() (*dao.PrivateMessage, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*dao.PrivateMessage), nil
	}
}

func (p privateMessageDo) Take() (*dao.PrivateMessage, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*dao.PrivateMessage), nil
	}
}

func (p privateMessageDo) Last() (*dao.PrivateMessage, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*dao.PrivateMessage), nil
	}
}

func (p privateMessageDo) Find() ([]*dao.PrivateMessage, error) {
	result, err := p.DO.Find()
	return result.([]*dao.PrivateMessage), err
}

func (p privateMessageDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dao.PrivateMessage, err error) {
	buf := make([]*dao.PrivateMessage, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p privateMessageDo) FindInBatches(result *[]*dao.PrivateMessage, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p privateMessageDo) Attrs(attrs ...field.AssignExpr) *privateMessageDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p privateMessageDo) Assign(attrs ...field.AssignExpr) *privateMessageDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p privateMessageDo) Joins(fields ...field.RelationField) *privateMessageDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p privateMessageDo) Preload(fields ...field.RelationField) *privateMessageDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p privateMessageDo) FirstOrInit() (*dao.PrivateMessage, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*dao.PrivateMessage), nil
	}
}

func (p privateMessageDo) FirstOrCreate() (*dao.PrivateMessage, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*dao.PrivateMessage), nil
	}
}

func (p privateMessageDo) FindByPage(offset int, limit int) (result []*dao.PrivateMessage, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p privateMessageDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p privateMessageDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p privateMessageDo) Delete(models ...*dao.PrivateMessage) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *privateMessageDo) withDO(do gen.Dao) *privateMessageDo {
	p.DO = *do.(*gen.DO)
	return p
}
