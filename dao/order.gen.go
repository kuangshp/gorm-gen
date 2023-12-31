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

	"gorm_gen/model"
)

func newOrder(db *gorm.DB, opts ...gen.DOOption) order {
	_order := order{}

	_order.orderDo.UseDB(db, opts...)
	_order.orderDo.UseModel(&model.Order{})

	tableName := _order.orderDo.TableName()
	_order.ALL = field.NewAsterisk(tableName)
	_order.ID = field.NewInt64(tableName, "id")
	_order.IsDel = field.NewInt64(tableName, "is_del")
	_order.CreatedAt = field.NewField(tableName, "created_at")
	_order.UpdatedAt = field.NewField(tableName, "updated_at")
	_order.OrderSn = field.NewString(tableName, "order_sn")
	_order.TradeNo = field.NewString(tableName, "trade_no")
	_order.PayMethod = field.NewInt64(tableName, "pay_method")
	_order.UserID = field.NewInt64(tableName, "user_id")
	_order.ActivityID = field.NewInt64(tableName, "activity_id")
	_order.ActivityEndTime = field.NewField(tableName, "activity_end_time")
	_order.PaidAmount = field.NewField(tableName, "paid_amount")
	_order.StartPaidAt = field.NewField(tableName, "start_paid_at")
	_order.EndPaidAt = field.NewField(tableName, "end_paid_at")
	_order.Status = field.NewInt64(tableName, "status")
	_order.Remark = field.NewString(tableName, "remark")
	_order.CancelOrderReason = field.NewString(tableName, "cancel_order_reason")
	_order.SponsorID = field.NewInt64(tableName, "sponsor_id")
	_order.ActivityTypeID = field.NewInt64(tableName, "activity_type_id")
	_order.TicketTotal = field.NewInt64(tableName, "ticket_total")
	_order.DeletedAt = field.NewField(tableName, "deleted_at")
	_order.ActivityStartTime = field.NewField(tableName, "activity_start_time")
	_order.RefundStatus = field.NewInt64(tableName, "refund_status")
	_order.RefundTicketTotal = field.NewInt64(tableName, "refund_ticket_total")

	_order.fillFieldMap()

	return _order
}

type order struct {
	orderDo

	ALL               field.Asterisk
	ID                field.Int64 // 主键id
	IsDel             field.Int64 // 是否删除,1表示删除,0表示正常
	CreatedAt         field.Field // 创建时间
	UpdatedAt         field.Field // 更新时间
	OrderSn           field.String
	TradeNo           field.String // 支付交易号(微信支付的交易号)
	PayMethod         field.Int64  // 支付方式0是微信,1是支付宝
	UserID            field.Int64  // 交易的用户ID(小程序)
	ActivityID        field.Int64  // 订单关联活动的ID
	ActivityEndTime   field.Field  // 活动结束时间
	PaidAmount        field.Field  // 支付金额(总金额)
	StartPaidAt       field.Field  // 开始支付时间
	EndPaidAt         field.Field  // 支付成功时间
	Status            field.Int64  // 支付状态,0表示未支付,1表示已支付,2表示支付失败,3表示取消订单,4表示退款,5表示核销,6表示在支付等待支付回调
	Remark            field.String // 备注信息
	CancelOrderReason field.String // 取消订单原因
	SponsorID         field.Int64  // 订单关联主办方ID
	ActivityTypeID    field.Int64  // 活动类型id
	TicketTotal       field.Int64  // 该订单下票数量
	DeletedAt         field.Field  // 软删除时间
	ActivityStartTime field.Field  // 活动开始时间
	RefundStatus      field.Int64  // 退款状态
	RefundTicketTotal field.Int64  // 退票数量

	fieldMap map[string]field.Expr
}

func (o order) Table(newTableName string) *order {
	o.orderDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o order) As(alias string) *order {
	o.orderDo.DO = *(o.orderDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *order) updateTableName(table string) *order {
	o.ALL = field.NewAsterisk(table)
	o.ID = field.NewInt64(table, "id")
	o.IsDel = field.NewInt64(table, "is_del")
	o.CreatedAt = field.NewField(table, "created_at")
	o.UpdatedAt = field.NewField(table, "updated_at")
	o.OrderSn = field.NewString(table, "order_sn")
	o.TradeNo = field.NewString(table, "trade_no")
	o.PayMethod = field.NewInt64(table, "pay_method")
	o.UserID = field.NewInt64(table, "user_id")
	o.ActivityID = field.NewInt64(table, "activity_id")
	o.ActivityEndTime = field.NewField(table, "activity_end_time")
	o.PaidAmount = field.NewField(table, "paid_amount")
	o.StartPaidAt = field.NewField(table, "start_paid_at")
	o.EndPaidAt = field.NewField(table, "end_paid_at")
	o.Status = field.NewInt64(table, "status")
	o.Remark = field.NewString(table, "remark")
	o.CancelOrderReason = field.NewString(table, "cancel_order_reason")
	o.SponsorID = field.NewInt64(table, "sponsor_id")
	o.ActivityTypeID = field.NewInt64(table, "activity_type_id")
	o.TicketTotal = field.NewInt64(table, "ticket_total")
	o.DeletedAt = field.NewField(table, "deleted_at")
	o.ActivityStartTime = field.NewField(table, "activity_start_time")
	o.RefundStatus = field.NewInt64(table, "refund_status")
	o.RefundTicketTotal = field.NewInt64(table, "refund_ticket_total")

	o.fillFieldMap()

	return o
}

func (o *order) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *order) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 23)
	o.fieldMap["id"] = o.ID
	o.fieldMap["is_del"] = o.IsDel
	o.fieldMap["created_at"] = o.CreatedAt
	o.fieldMap["updated_at"] = o.UpdatedAt
	o.fieldMap["order_sn"] = o.OrderSn
	o.fieldMap["trade_no"] = o.TradeNo
	o.fieldMap["pay_method"] = o.PayMethod
	o.fieldMap["user_id"] = o.UserID
	o.fieldMap["activity_id"] = o.ActivityID
	o.fieldMap["activity_end_time"] = o.ActivityEndTime
	o.fieldMap["paid_amount"] = o.PaidAmount
	o.fieldMap["start_paid_at"] = o.StartPaidAt
	o.fieldMap["end_paid_at"] = o.EndPaidAt
	o.fieldMap["status"] = o.Status
	o.fieldMap["remark"] = o.Remark
	o.fieldMap["cancel_order_reason"] = o.CancelOrderReason
	o.fieldMap["sponsor_id"] = o.SponsorID
	o.fieldMap["activity_type_id"] = o.ActivityTypeID
	o.fieldMap["ticket_total"] = o.TicketTotal
	o.fieldMap["deleted_at"] = o.DeletedAt
	o.fieldMap["activity_start_time"] = o.ActivityStartTime
	o.fieldMap["refund_status"] = o.RefundStatus
	o.fieldMap["refund_ticket_total"] = o.RefundTicketTotal
}

func (o order) clone(db *gorm.DB) order {
	o.orderDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o order) replaceDB(db *gorm.DB) order {
	o.orderDo.ReplaceDB(db)
	return o
}

type orderDo struct{ gen.DO }

type IOrderDo interface {
	gen.SubQuery
	Debug() IOrderDo
	WithContext(ctx context.Context) IOrderDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IOrderDo
	WriteDB() IOrderDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IOrderDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IOrderDo
	Not(conds ...gen.Condition) IOrderDo
	Or(conds ...gen.Condition) IOrderDo
	Select(conds ...field.Expr) IOrderDo
	Where(conds ...gen.Condition) IOrderDo
	Order(conds ...field.Expr) IOrderDo
	Distinct(cols ...field.Expr) IOrderDo
	Omit(cols ...field.Expr) IOrderDo
	Join(table schema.Tabler, on ...field.Expr) IOrderDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IOrderDo
	RightJoin(table schema.Tabler, on ...field.Expr) IOrderDo
	Group(cols ...field.Expr) IOrderDo
	Having(conds ...gen.Condition) IOrderDo
	Limit(limit int) IOrderDo
	Offset(offset int) IOrderDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IOrderDo
	Unscoped() IOrderDo
	Create(values ...*model.Order) error
	CreateInBatches(values []*model.Order, batchSize int) error
	Save(values ...*model.Order) error
	First() (*model.Order, error)
	Take() (*model.Order, error)
	Last() (*model.Order, error)
	Find() ([]*model.Order, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Order, err error)
	FindInBatches(result *[]*model.Order, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Order) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IOrderDo
	Assign(attrs ...field.AssignExpr) IOrderDo
	Joins(fields ...field.RelationField) IOrderDo
	Preload(fields ...field.RelationField) IOrderDo
	FirstOrInit() (*model.Order, error)
	FirstOrCreate() (*model.Order, error)
	FindByPage(offset int, limit int) (result []*model.Order, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IOrderDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (o orderDo) Debug() IOrderDo {
	return o.withDO(o.DO.Debug())
}

func (o orderDo) WithContext(ctx context.Context) IOrderDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o orderDo) ReadDB() IOrderDo {
	return o.Clauses(dbresolver.Read)
}

func (o orderDo) WriteDB() IOrderDo {
	return o.Clauses(dbresolver.Write)
}

func (o orderDo) Session(config *gorm.Session) IOrderDo {
	return o.withDO(o.DO.Session(config))
}

func (o orderDo) Clauses(conds ...clause.Expression) IOrderDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o orderDo) Returning(value interface{}, columns ...string) IOrderDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o orderDo) Not(conds ...gen.Condition) IOrderDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o orderDo) Or(conds ...gen.Condition) IOrderDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o orderDo) Select(conds ...field.Expr) IOrderDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o orderDo) Where(conds ...gen.Condition) IOrderDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o orderDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IOrderDo {
	return o.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (o orderDo) Order(conds ...field.Expr) IOrderDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o orderDo) Distinct(cols ...field.Expr) IOrderDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o orderDo) Omit(cols ...field.Expr) IOrderDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o orderDo) Join(table schema.Tabler, on ...field.Expr) IOrderDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o orderDo) LeftJoin(table schema.Tabler, on ...field.Expr) IOrderDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o orderDo) RightJoin(table schema.Tabler, on ...field.Expr) IOrderDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o orderDo) Group(cols ...field.Expr) IOrderDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o orderDo) Having(conds ...gen.Condition) IOrderDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o orderDo) Limit(limit int) IOrderDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o orderDo) Offset(offset int) IOrderDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o orderDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IOrderDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o orderDo) Unscoped() IOrderDo {
	return o.withDO(o.DO.Unscoped())
}

func (o orderDo) Create(values ...*model.Order) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o orderDo) CreateInBatches(values []*model.Order, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o orderDo) Save(values ...*model.Order) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o orderDo) First() (*model.Order, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Order), nil
	}
}

func (o orderDo) Take() (*model.Order, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Order), nil
	}
}

func (o orderDo) Last() (*model.Order, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Order), nil
	}
}

func (o orderDo) Find() ([]*model.Order, error) {
	result, err := o.DO.Find()
	return result.([]*model.Order), err
}

func (o orderDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Order, err error) {
	buf := make([]*model.Order, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o orderDo) FindInBatches(result *[]*model.Order, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o orderDo) Attrs(attrs ...field.AssignExpr) IOrderDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o orderDo) Assign(attrs ...field.AssignExpr) IOrderDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o orderDo) Joins(fields ...field.RelationField) IOrderDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o orderDo) Preload(fields ...field.RelationField) IOrderDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o orderDo) FirstOrInit() (*model.Order, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Order), nil
	}
}

func (o orderDo) FirstOrCreate() (*model.Order, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Order), nil
	}
}

func (o orderDo) FindByPage(offset int, limit int) (result []*model.Order, count int64, err error) {
	result, err = o.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = o.Offset(-1).Limit(-1).Count()
	return
}

func (o orderDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o orderDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o orderDo) Delete(models ...*model.Order) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *orderDo) withDO(do gen.Dao) *orderDo {
	o.DO = *do.(*gen.DO)
	return o
}
