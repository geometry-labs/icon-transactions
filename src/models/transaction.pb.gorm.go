// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: transaction.proto

package models

import (
	context "context"
	fmt "fmt"
	
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	math "math"

	gorm2 "github.com/infobloxopen/atlas-app-toolkit/gorm"
	errors1 "github.com/infobloxopen/protoc-gen-gorm/errors"
	gorm1 "github.com/jinzhu/gorm"
	field_mask1 "google.golang.org/genproto/protobuf/field_mask"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = math.Inf

type TransactionORM struct {
	BlockHash                 string
	BlockNumber               uint64 `gorm:"index:transaction_idx_block_number"`
	BlockTimestamp            uint64
	Data                      string
	DataType                  string
	Fee                       uint64
	FromAddress               string `gorm:"index:transaction_idx_from_address"`
	Hash                      string `gorm:"index:transaction_idx_hash"`
	Id                        uint64
	ItemId                    string
	ItemTimestamp             string
	LogIndex                  int32
	Nid                       uint32
	Nonce                     string
	ReceiptCumulativeStepUsed uint64
	ReceiptLogs               string
	ReceiptScoreAddress       string
	ReceiptStatus             uint32
	ReceiptStepPrice          uint64
	ReceiptStepUsed           uint64
	Signature                 string
	StepLimit                 uint64
	Timestamp                 string
	ToAddress                 string `gorm:"index:transaction_idx_to_address"`
	TransactionIndex          uint32
	Type                      string
	Value                     string
	Version                   string
}

// TableName overrides the default tablename generated by GORM
func (TransactionORM) TableName() string {
	return "transactions"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Transaction) ToORM(ctx context.Context) (TransactionORM, error) {
	to := TransactionORM{}
	var err error
	if prehook, ok := interface{}(m).(TransactionWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Type = m.Type
	to.Version = m.Version
	to.FromAddress = m.FromAddress
	to.ToAddress = m.ToAddress
	to.Value = m.Value
	to.StepLimit = m.StepLimit
	to.Timestamp = m.Timestamp
	to.BlockTimestamp = m.BlockTimestamp
	to.Nid = m.Nid
	to.Nonce = m.Nonce
	to.Hash = m.Hash
	to.TransactionIndex = m.TransactionIndex
	to.BlockHash = m.BlockHash
	to.BlockNumber = m.BlockNumber
	to.Fee = m.Fee
	to.Signature = m.Signature
	to.DataType = m.DataType
	to.Data = m.Data
	to.ReceiptCumulativeStepUsed = m.ReceiptCumulativeStepUsed
	to.ReceiptStepUsed = m.ReceiptStepUsed
	to.ReceiptStepPrice = m.ReceiptStepPrice
	to.ReceiptScoreAddress = m.ReceiptScoreAddress
	to.ReceiptLogs = m.ReceiptLogs
	to.ReceiptStatus = m.ReceiptStatus
	to.ItemId = m.ItemId
	to.ItemTimestamp = m.ItemTimestamp
	to.LogIndex = m.LogIndex
	to.Id = m.Id
	if posthook, ok := interface{}(m).(TransactionWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *TransactionORM) ToPB(ctx context.Context) (Transaction, error) {
	to := Transaction{}
	var err error
	if prehook, ok := interface{}(m).(TransactionWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Type = m.Type
	to.Version = m.Version
	to.FromAddress = m.FromAddress
	to.ToAddress = m.ToAddress
	to.Value = m.Value
	to.StepLimit = m.StepLimit
	to.Timestamp = m.Timestamp
	to.BlockTimestamp = m.BlockTimestamp
	to.Nid = m.Nid
	to.Nonce = m.Nonce
	to.Hash = m.Hash
	to.TransactionIndex = m.TransactionIndex
	to.BlockHash = m.BlockHash
	to.BlockNumber = m.BlockNumber
	to.Fee = m.Fee
	to.Signature = m.Signature
	to.DataType = m.DataType
	to.Data = m.Data
	to.ReceiptCumulativeStepUsed = m.ReceiptCumulativeStepUsed
	to.ReceiptStepUsed = m.ReceiptStepUsed
	to.ReceiptStepPrice = m.ReceiptStepPrice
	to.ReceiptScoreAddress = m.ReceiptScoreAddress
	to.ReceiptLogs = m.ReceiptLogs
	to.ReceiptStatus = m.ReceiptStatus
	to.ItemId = m.ItemId
	to.ItemTimestamp = m.ItemTimestamp
	to.LogIndex = m.LogIndex
	to.Id = m.Id
	if posthook, ok := interface{}(m).(TransactionWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Transaction the arg will be the target, the caller the one being converted from

// TransactionBeforeToORM called before default ToORM code
type TransactionWithBeforeToORM interface {
	BeforeToORM(context.Context, *TransactionORM) error
}

// TransactionAfterToORM called after default ToORM code
type TransactionWithAfterToORM interface {
	AfterToORM(context.Context, *TransactionORM) error
}

// TransactionBeforeToPB called before default ToPB code
type TransactionWithBeforeToPB interface {
	BeforeToPB(context.Context, *Transaction) error
}

// TransactionAfterToPB called after default ToPB code
type TransactionWithAfterToPB interface {
	AfterToPB(context.Context, *Transaction) error
}

// DefaultCreateTransaction executes a basic gorm create call
func DefaultCreateTransaction(ctx context.Context, in *Transaction, db *gorm1.DB) (*Transaction, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type TransactionORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type TransactionORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm1.DB) error
}

// DefaultReadTransaction executes a basic gorm read call
func DefaultReadTransaction(ctx context.Context, in *Transaction, db *gorm1.DB) (*Transaction, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Id == 0 {
		return nil, errors1.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(TransactionORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm2.ApplyFieldSelection(ctx, db, nil, &TransactionORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := TransactionORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(TransactionORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type TransactionORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type TransactionORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type TransactionORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm1.DB) error
}

func DefaultDeleteTransaction(ctx context.Context, in *Transaction, db *gorm1.DB) error {
	if in == nil {
		return errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == 0 {
		return errors1.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(TransactionORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&TransactionORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(TransactionORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type TransactionORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type TransactionORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm1.DB) error
}

func DefaultDeleteTransactionSet(ctx context.Context, in []*Transaction, db *gorm1.DB) error {
	if in == nil {
		return errors1.NilArgumentError
	}
	var err error
	keys := []uint64{}
	for _, obj := range in {
		ormObj, err := obj.ToORM(ctx)
		if err != nil {
			return err
		}
		if ormObj.Id == 0 {
			return errors1.EmptyIdError
		}
		keys = append(keys, ormObj.Id)
	}
	if hook, ok := (interface{}(&TransactionORM{})).(TransactionORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&TransactionORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&TransactionORM{})).(TransactionORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type TransactionORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*Transaction, *gorm1.DB) (*gorm1.DB, error)
}
type TransactionORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*Transaction, *gorm1.DB) error
}

// DefaultStrictUpdateTransaction clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateTransaction(ctx context.Context, in *Transaction, db *gorm1.DB) (*Transaction, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateTransaction")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &TransactionORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(TransactionORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(TransactionORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, err
}

type TransactionORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type TransactionORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type TransactionORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm1.DB) error
}

// DefaultPatchTransaction executes a basic gorm update call with patch behavior
func DefaultPatchTransaction(ctx context.Context, in *Transaction, updateMask *field_mask1.FieldMask, db *gorm1.DB) (*Transaction, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	var pbObj Transaction
	var err error
	if hook, ok := interface{}(&pbObj).(TransactionWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadTransaction(ctx, &Transaction{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(TransactionWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskTransaction(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(TransactionWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateTransaction(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(TransactionWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type TransactionWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *Transaction, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type TransactionWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *Transaction, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type TransactionWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *Transaction, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type TransactionWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *Transaction, *field_mask1.FieldMask, *gorm1.DB) error
}

// DefaultPatchSetTransaction executes a bulk gorm update call with patch behavior
func DefaultPatchSetTransaction(ctx context.Context, objects []*Transaction, updateMasks []*field_mask1.FieldMask, db *gorm1.DB) ([]*Transaction, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors1.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*Transaction, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchTransaction(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskTransaction patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskTransaction(ctx context.Context, patchee *Transaction, patcher *Transaction, updateMask *field_mask1.FieldMask, prefix string, db *gorm1.DB) (*Transaction, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors1.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"Type" {
			patchee.Type = patcher.Type
			continue
		}
		if f == prefix+"Version" {
			patchee.Version = patcher.Version
			continue
		}
		if f == prefix+"FromAddress" {
			patchee.FromAddress = patcher.FromAddress
			continue
		}
		if f == prefix+"ToAddress" {
			patchee.ToAddress = patcher.ToAddress
			continue
		}
		if f == prefix+"Value" {
			patchee.Value = patcher.Value
			continue
		}
		if f == prefix+"StepLimit" {
			patchee.StepLimit = patcher.StepLimit
			continue
		}
		if f == prefix+"Timestamp" {
			patchee.Timestamp = patcher.Timestamp
			continue
		}
		if f == prefix+"BlockTimestamp" {
			patchee.BlockTimestamp = patcher.BlockTimestamp
			continue
		}
		if f == prefix+"Nid" {
			patchee.Nid = patcher.Nid
			continue
		}
		if f == prefix+"Nonce" {
			patchee.Nonce = patcher.Nonce
			continue
		}
		if f == prefix+"Hash" {
			patchee.Hash = patcher.Hash
			continue
		}
		if f == prefix+"TransactionIndex" {
			patchee.TransactionIndex = patcher.TransactionIndex
			continue
		}
		if f == prefix+"BlockHash" {
			patchee.BlockHash = patcher.BlockHash
			continue
		}
		if f == prefix+"BlockNumber" {
			patchee.BlockNumber = patcher.BlockNumber
			continue
		}
		if f == prefix+"Fee" {
			patchee.Fee = patcher.Fee
			continue
		}
		if f == prefix+"Signature" {
			patchee.Signature = patcher.Signature
			continue
		}
		if f == prefix+"DataType" {
			patchee.DataType = patcher.DataType
			continue
		}
		if f == prefix+"Data" {
			patchee.Data = patcher.Data
			continue
		}
		if f == prefix+"ReceiptCumulativeStepUsed" {
			patchee.ReceiptCumulativeStepUsed = patcher.ReceiptCumulativeStepUsed
			continue
		}
		if f == prefix+"ReceiptStepUsed" {
			patchee.ReceiptStepUsed = patcher.ReceiptStepUsed
			continue
		}
		if f == prefix+"ReceiptStepPrice" {
			patchee.ReceiptStepPrice = patcher.ReceiptStepPrice
			continue
		}
		if f == prefix+"ReceiptScoreAddress" {
			patchee.ReceiptScoreAddress = patcher.ReceiptScoreAddress
			continue
		}
		if f == prefix+"ReceiptLogs" {
			patchee.ReceiptLogs = patcher.ReceiptLogs
			continue
		}
		if f == prefix+"ReceiptStatus" {
			patchee.ReceiptStatus = patcher.ReceiptStatus
			continue
		}
		if f == prefix+"ItemId" {
			patchee.ItemId = patcher.ItemId
			continue
		}
		if f == prefix+"ItemTimestamp" {
			patchee.ItemTimestamp = patcher.ItemTimestamp
			continue
		}
		if f == prefix+"LogIndex" {
			patchee.LogIndex = patcher.LogIndex
			continue
		}
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListTransaction executes a gorm list call
func DefaultListTransaction(ctx context.Context, db *gorm1.DB) ([]*Transaction, error) {
	in := Transaction{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm2.ApplyCollectionOperators(ctx, db, &TransactionORM{}, &Transaction{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []TransactionORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*Transaction{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type TransactionORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type TransactionORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type TransactionORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm1.DB, *[]TransactionORM) error
}
