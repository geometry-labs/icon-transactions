// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: transaction_count_index.proto

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

type TransactionCountIndexORM struct {
	TransactionHash string `gorm:"primary_key"`
}

// TableName overrides the default tablename generated by GORM
func (TransactionCountIndexORM) TableName() string {
	return "transaction_count_indices"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *TransactionCountIndex) ToORM(ctx context.Context) (TransactionCountIndexORM, error) {
	to := TransactionCountIndexORM{}
	var err error
	if prehook, ok := interface{}(m).(TransactionCountIndexWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.TransactionHash = m.TransactionHash
	if posthook, ok := interface{}(m).(TransactionCountIndexWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *TransactionCountIndexORM) ToPB(ctx context.Context) (TransactionCountIndex, error) {
	to := TransactionCountIndex{}
	var err error
	if prehook, ok := interface{}(m).(TransactionCountIndexWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.TransactionHash = m.TransactionHash
	if posthook, ok := interface{}(m).(TransactionCountIndexWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type TransactionCountIndex the arg will be the target, the caller the one being converted from

// TransactionCountIndexBeforeToORM called before default ToORM code
type TransactionCountIndexWithBeforeToORM interface {
	BeforeToORM(context.Context, *TransactionCountIndexORM) error
}

// TransactionCountIndexAfterToORM called after default ToORM code
type TransactionCountIndexWithAfterToORM interface {
	AfterToORM(context.Context, *TransactionCountIndexORM) error
}

// TransactionCountIndexBeforeToPB called before default ToPB code
type TransactionCountIndexWithBeforeToPB interface {
	BeforeToPB(context.Context, *TransactionCountIndex) error
}

// TransactionCountIndexAfterToPB called after default ToPB code
type TransactionCountIndexWithAfterToPB interface {
	AfterToPB(context.Context, *TransactionCountIndex) error
}

// DefaultCreateTransactionCountIndex executes a basic gorm create call
func DefaultCreateTransactionCountIndex(ctx context.Context, in *TransactionCountIndex, db *gorm1.DB) (*TransactionCountIndex, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionCountIndexORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionCountIndexORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type TransactionCountIndexORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type TransactionCountIndexORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm1.DB) error
}

// DefaultApplyFieldMaskTransactionCountIndex patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskTransactionCountIndex(ctx context.Context, patchee *TransactionCountIndex, patcher *TransactionCountIndex, updateMask *field_mask1.FieldMask, prefix string, db *gorm1.DB) (*TransactionCountIndex, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors1.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"TransactionHash" {
			patchee.TransactionHash = patcher.TransactionHash
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListTransactionCountIndex executes a gorm list call
func DefaultListTransactionCountIndex(ctx context.Context, db *gorm1.DB) ([]*TransactionCountIndex, error) {
	in := TransactionCountIndex{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionCountIndexORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm2.ApplyCollectionOperators(ctx, db, &TransactionCountIndexORM{}, &TransactionCountIndex{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionCountIndexORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("transaction_hash")
	ormResponse := []TransactionCountIndexORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionCountIndexORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*TransactionCountIndex{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type TransactionCountIndexORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type TransactionCountIndexORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type TransactionCountIndexORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm1.DB, *[]TransactionCountIndexORM) error
}
