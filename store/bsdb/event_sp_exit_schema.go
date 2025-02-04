package bsdb

import "github.com/forbole/juno/v4/common"

type EventStorageProviderExit struct {
	ID                uint64         `gorm:"column:id;primaryKey"`
	StorageProviderId uint32         `gorm:"column:storage_provider_id;index:idx_sp_id"`
	OperatorAddress   common.Address `gorm:"column:operator_address;type:BINARY(20)"`

	CreateAt     int64       `gorm:"column:create_at"`
	CreateTxHash common.Hash `gorm:"column:create_tx_hash;type:BINARY(32);not null"`
	CreateTime   int64       `gorm:"column:create_time"` // seconds
}

func (*EventStorageProviderExit) TableName() string {
	return EventStorageProviderExitTableName
}
