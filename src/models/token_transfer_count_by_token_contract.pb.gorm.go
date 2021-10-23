// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: token_transfer_count_by_token_contract.proto

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

type TokenTransferCountByTokenContractORM struct {
	Count           uint64
	LogIndex        uint64
	TokenContract   string `gorm:"primary_key"`
	TransactionHash string
}

// TableName overrides the default tablename generated by GORM
func (TokenTransferCountByTokenContractORM) TableName() string {
	return "token_transfer_count_by_token_contracts"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *TokenTransferCountByTokenContract) ToORM(ctx context.Context) (TokenTransferCountByTokenContractORM, error) {
	to := TokenTransferCountByTokenContractORM{}
	var err error
	if prehook, ok := interface{}(m).(TokenTransferCountByTokenContractWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.TransactionHash = m.TransactionHash
	to.LogIndex = m.LogIndex
	to.TokenContract = m.TokenContract
	to.Count = m.Count
	if posthook, ok := interface{}(m).(TokenTransferCountByTokenContractWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *TokenTransferCountByTokenContractORM) ToPB(ctx context.Context) (TokenTransferCountByTokenContract, error) {
	to := TokenTransferCountByTokenContract{}
	var err error
	if prehook, ok := interface{}(m).(TokenTransferCountByTokenContractWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.TransactionHash = m.TransactionHash
	to.LogIndex = m.LogIndex
	to.TokenContract = m.TokenContract
	to.Count = m.Count
	if posthook, ok := interface{}(m).(TokenTransferCountByTokenContractWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type TokenTransferCountByTokenContract the arg will be the target, the caller the one being converted from

// TokenTransferCountByTokenContractBeforeToORM called before default ToORM code
type TokenTransferCountByTokenContractWithBeforeToORM interface {
	BeforeToORM(context.Context, *TokenTransferCountByTokenContractORM) error
}

// TokenTransferCountByTokenContractAfterToORM called after default ToORM code
type TokenTransferCountByTokenContractWithAfterToORM interface {
	AfterToORM(context.Context, *TokenTransferCountByTokenContractORM) error
}

// TokenTransferCountByTokenContractBeforeToPB called before default ToPB code
type TokenTransferCountByTokenContractWithBeforeToPB interface {
	BeforeToPB(context.Context, *TokenTransferCountByTokenContract) error
}

// TokenTransferCountByTokenContractAfterToPB called after default ToPB code
type TokenTransferCountByTokenContractWithAfterToPB interface {
	AfterToPB(context.Context, *TokenTransferCountByTokenContract) error
}

// DefaultCreateTokenTransferCountByTokenContract executes a basic gorm create call
func DefaultCreateTokenTransferCountByTokenContract(ctx context.Context, in *TokenTransferCountByTokenContract, db *gorm1.DB) (*TokenTransferCountByTokenContract, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TokenTransferCountByTokenContractORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TokenTransferCountByTokenContractORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type TokenTransferCountByTokenContractORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type TokenTransferCountByTokenContractORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm1.DB) error
}

// DefaultApplyFieldMaskTokenTransferCountByTokenContract patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskTokenTransferCountByTokenContract(ctx context.Context, patchee *TokenTransferCountByTokenContract, patcher *TokenTransferCountByTokenContract, updateMask *field_mask1.FieldMask, prefix string, db *gorm1.DB) (*TokenTransferCountByTokenContract, error) {
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
		if f == prefix+"LogIndex" {
			patchee.LogIndex = patcher.LogIndex
			continue
		}
		if f == prefix+"TokenContract" {
			patchee.TokenContract = patcher.TokenContract
			continue
		}
		if f == prefix+"Count" {
			patchee.Count = patcher.Count
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListTokenTransferCountByTokenContract executes a gorm list call
func DefaultListTokenTransferCountByTokenContract(ctx context.Context, db *gorm1.DB) ([]*TokenTransferCountByTokenContract, error) {
	in := TokenTransferCountByTokenContract{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TokenTransferCountByTokenContractORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm2.ApplyCollectionOperators(ctx, db, &TokenTransferCountByTokenContractORM{}, &TokenTransferCountByTokenContract{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TokenTransferCountByTokenContractORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("token_contract")
	ormResponse := []TokenTransferCountByTokenContractORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TokenTransferCountByTokenContractORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*TokenTransferCountByTokenContract{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type TokenTransferCountByTokenContractORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type TokenTransferCountByTokenContractORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type TokenTransferCountByTokenContractORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm1.DB, *[]TokenTransferCountByTokenContractORM) error
}
