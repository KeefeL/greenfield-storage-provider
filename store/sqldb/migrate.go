package sqldb

import (
	"errors"
	"fmt"

	"github.com/bnb-chain/greenfield-storage-provider/core/spdb"
	"gorm.io/gorm"
)

const (
	SPExitProgressKey        = "sp_exit_progress"
	SwapOutProgressKey       = "swap_out_progress"
	BucketMigrateProgressKey = "bucket_migrate_progress"
)

// UpdateSPExitSubscribeProgress is used to update progress.
// insert a new one if it is not found in db.
func (s *SpDBImpl) UpdateSPExitSubscribeProgress(blockHeight uint64) error {
	var (
		result       *gorm.DB
		queryReturn  *MigrateSubscribeProgressTable
		needInsert   bool
		updateRecord *MigrateSubscribeProgressTable
	)
	queryReturn = &MigrateSubscribeProgressTable{}
	result = s.db.First(queryReturn, "event_name = ?", SPExitProgressKey)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	if result.Error != nil {
		needInsert = errors.Is(result.Error, gorm.ErrRecordNotFound)
	}
	updateRecord = &MigrateSubscribeProgressTable{
		EventName:                 SPExitProgressKey,
		LastSubscribedBlockHeight: blockHeight,
	}
	if needInsert {
		result = s.db.Create(updateRecord)
		if result.Error != nil || result.RowsAffected != 1 {
			return fmt.Errorf("failed to insert record in subscribe progress table: %s", result.Error)
		}

	} else { // update
		result = s.db.Model(&MigrateSubscribeProgressTable{}).
			Where("event_name = ?", SPExitProgressKey).Updates(updateRecord)
		if result.Error != nil {
			return fmt.Errorf("failed to update record in subscribe progress table: %s", result.Error)
		}
	}
	return nil
}

func (s *SpDBImpl) QuerySPExitSubscribeProgress() (uint64, error) {
	var (
		result      *gorm.DB
		queryReturn *MigrateSubscribeProgressTable
	)
	queryReturn = &MigrateSubscribeProgressTable{}
	result = s.db.First(queryReturn, "event_name = ?", SPExitProgressKey)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return 0, nil
	}
	if result.Error != nil {
		return 0, result.Error
	}
	return queryReturn.LastSubscribedBlockHeight, nil
}

func (s *SpDBImpl) UpdateSwapOutSubscribeProgress(blockHeight uint64) error {
	var (
		result       *gorm.DB
		queryReturn  *MigrateSubscribeProgressTable
		needInsert   bool
		updateRecord *MigrateSubscribeProgressTable
	)
	queryReturn = &MigrateSubscribeProgressTable{}
	result = s.db.First(queryReturn, "event_name = ?", SwapOutProgressKey)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	if result.Error != nil {
		needInsert = errors.Is(result.Error, gorm.ErrRecordNotFound)
	}
	updateRecord = &MigrateSubscribeProgressTable{
		EventName:                 SwapOutProgressKey,
		LastSubscribedBlockHeight: blockHeight,
	}
	if needInsert {
		result = s.db.Create(updateRecord)
		if result.Error != nil || result.RowsAffected != 1 {
			return fmt.Errorf("failed to insert record in subscribe progress table: %s", result.Error)
		}

	} else { // update
		result = s.db.Model(&MigrateSubscribeProgressTable{}).
			Where("event_name = ?", SwapOutProgressKey).Updates(updateRecord)
		if result.Error != nil {
			return fmt.Errorf("failed to update record in subscribe progress table: %s", result.Error)
		}
	}
	return nil
}

func (s *SpDBImpl) QuerySwapOutSubscribeProgress() (uint64, error) {
	var (
		result      *gorm.DB
		queryReturn *MigrateSubscribeProgressTable
	)
	queryReturn = &MigrateSubscribeProgressTable{}
	result = s.db.First(queryReturn, "event_name = ?", SwapOutProgressKey)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return 0, nil
	}
	if result.Error != nil {
		return 0, result.Error
	}
	return queryReturn.LastSubscribedBlockHeight, nil
}
func (s *SpDBImpl) UpdateBucketMigrateSubscribeProgress(blockHeight uint64) error {
	var (
		result       *gorm.DB
		queryReturn  *MigrateSubscribeProgressTable
		needInsert   bool
		updateRecord *MigrateSubscribeProgressTable
	)
	queryReturn = &MigrateSubscribeProgressTable{}
	result = s.db.First(queryReturn, "event_name = ?", BucketMigrateProgressKey)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	if result.Error != nil {
		needInsert = errors.Is(result.Error, gorm.ErrRecordNotFound)
	}
	updateRecord = &MigrateSubscribeProgressTable{
		EventName:                 BucketMigrateProgressKey,
		LastSubscribedBlockHeight: blockHeight,
	}
	if needInsert {
		result = s.db.Create(updateRecord)
		if result.Error != nil || result.RowsAffected != 1 {
			return fmt.Errorf("failed to insert record in subscribe progress table: %s", result.Error)
		}

	} else { // update
		result = s.db.Model(&MigrateSubscribeProgressTable{}).
			Where("event_name = ?", BucketMigrateProgressKey).Updates(updateRecord)
		if result.Error != nil {
			return fmt.Errorf("failed to update record in subscribe progress table: %s", result.Error)
		}
	}
	return nil
}

func (s *SpDBImpl) QueryBucketMigrateSubscribeProgress() (uint64, error) {
	var (
		result      *gorm.DB
		queryReturn *MigrateSubscribeProgressTable
	)
	queryReturn = &MigrateSubscribeProgressTable{}
	result = s.db.First(queryReturn, "event_name = ?", BucketMigrateProgressKey)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return 0, nil
	}
	if result.Error != nil {
		return 0, result.Error
	}
	return queryReturn.LastSubscribedBlockHeight, nil
}

func (s *SpDBImpl) InsertMigrateGVGUnit(meta *spdb.MigrateGVGUnitMeta) error {
	var (
		err              error
		result           *gorm.DB
		insertMigrateGVG *MigrateGVGTable
	)
	insertMigrateGVG = &MigrateGVGTable{
		MigrateKey:           meta.MigrateGVGKey,
		GlobalVirtualGroupID: meta.GlobalVirtualGroupID,
		VirtualGroupFamilyID: meta.VirtualGroupFamilyID,
		RedundancyIndex:      meta.RedundancyIndex,
		BucketID:             meta.BucketID,
		IsSecondary:          meta.IsSecondary,
		MigrateStatus:        meta.MigrateStatus,
	}
	result = s.db.Create(insertMigrateGVG)
	if result.Error != nil || result.RowsAffected != 1 {
		err = fmt.Errorf("failed to insert migrate gvg table: %s", result.Error)
		return err
	}
	return nil
}

func (s *SpDBImpl) DeleteMigrateGVGUnit(meta *spdb.MigrateGVGUnitMeta) error {
	// TODO:
	return nil
}

func (s *SpDBImpl) UpdateMigrateGVGUnitStatus(migrateKey string, migrateStatus int) error {
	if result := s.db.Model(&MigrateGVGTable{}).Where("migrate_key = ?", migrateKey).Updates(&MigrateGVGTable{
		MigrateStatus: migrateStatus,
	}); result.Error != nil {
		return fmt.Errorf("failed to update migrate gvg status: %s", result.Error)
	}
	return nil
}

func (s *SpDBImpl) UpdateMigrateGVGUnitLastMigrateObjectID(migrateKey string, lastMigratedObjectID uint64) error {
	if result := s.db.Model(&MigrateGVGTable{}).Where("migrate_key = ?", migrateKey).Updates(&MigrateGVGTable{
		LastMigratedObjectID: lastMigratedObjectID,
	}); result.Error != nil {
		return fmt.Errorf("failed to update migrate gvg progress: %s", result.Error)
	}
	return nil
}

func (s *SpDBImpl) QueryMigrateGVGUnit(migrateKey string) (*spdb.MigrateGVGUnitMeta, error) {
	var (
		result      *gorm.DB
		queryReturn *MigrateGVGTable
	)
	queryReturn = &MigrateGVGTable{}
	result = s.db.First(queryReturn, "migrate_key = ?", migrateKey)
	if result.Error != nil {
		return nil, result.Error
	}
	return &spdb.MigrateGVGUnitMeta{
		GlobalVirtualGroupID: queryReturn.GlobalVirtualGroupID,
		VirtualGroupFamilyID: queryReturn.VirtualGroupFamilyID,
		RedundancyIndex:      queryReturn.RedundancyIndex,
		BucketID:             queryReturn.BucketID,
		IsSecondary:          queryReturn.IsSecondary,
		IsConflicted:         queryReturn.IsConflicted,
		SrcSPID:              queryReturn.SrcSPID,
		DestSPID:             queryReturn.DestSPID,
		LastMigratedObjectID: queryReturn.LastMigratedObjectID,
		MigrateStatus:        queryReturn.MigrateStatus,
	}, nil
}

// ListMigrateGVGUnitsByFamilyID is used to src sp load to build execute plan.
func (s *SpDBImpl) ListMigrateGVGUnitsByFamilyID(familyID uint32, srcSP uint32) ([]*spdb.MigrateGVGUnitMeta, error) {
	var queryReturns []MigrateGVGTable
	result := s.db.Where("is_conflicted = false and is_secondary = false and is_remoted = false and redundancy_index = 0 and bucket_id = 0 and virtual_group_family_id = ? and src_sp_id = ?",
		familyID, srcSP).Find(&queryReturns)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to list migrate gvg table: %s", result.Error)
	}
	returns := make([]*spdb.MigrateGVGUnitMeta, 0)
	for _, queryReturn := range queryReturns {
		returns = append(returns, &spdb.MigrateGVGUnitMeta{
			GlobalVirtualGroupID: queryReturn.GlobalVirtualGroupID,
			VirtualGroupFamilyID: queryReturn.VirtualGroupFamilyID,
			RedundancyIndex:      queryReturn.RedundancyIndex,
			BucketID:             queryReturn.BucketID,
			IsSecondary:          queryReturn.IsSecondary,
			IsConflicted:         queryReturn.IsConflicted,
			SrcSPID:              queryReturn.SrcSPID,
			DestSPID:             queryReturn.DestSPID,
			LastMigratedObjectID: queryReturn.LastMigratedObjectID,
			MigrateStatus:        queryReturn.MigrateStatus,
		})
	}
	return returns, nil
}

// ListConflictedMigrateGVGUnitsByFamilyID is used to src sp load to build execute plan.
func (s *SpDBImpl) ListConflictedMigrateGVGUnitsByFamilyID(familyID uint32) ([]*spdb.MigrateGVGUnitMeta, error) {
	var queryReturns []MigrateGVGTable
	result := s.db.Where("is_conflicted = true and virtual_group_family_id = ?", familyID).Find(&queryReturns)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to list migrate gvg table: %s", result.Error)
	}
	returns := make([]*spdb.MigrateGVGUnitMeta, 0)
	for _, queryReturn := range queryReturns {
		returns = append(returns, &spdb.MigrateGVGUnitMeta{
			GlobalVirtualGroupID: queryReturn.GlobalVirtualGroupID,
			VirtualGroupFamilyID: queryReturn.VirtualGroupFamilyID,
			RedundancyIndex:      queryReturn.RedundancyIndex,
			BucketID:             queryReturn.BucketID,
			IsSecondary:          queryReturn.IsSecondary,
			IsConflicted:         queryReturn.IsConflicted,
			SrcSPID:              queryReturn.SrcSPID,
			DestSPID:             queryReturn.DestSPID,
			LastMigratedObjectID: queryReturn.LastMigratedObjectID,
			MigrateStatus:        queryReturn.MigrateStatus,
		})
	}
	return returns, nil
}

// ListRemotedMigrateGVGUnits is used to dest sp load to build migrate task runner.
func (s *SpDBImpl) ListRemotedMigrateGVGUnits() ([]*spdb.MigrateGVGUnitMeta, error) {
	var queryReturns []MigrateGVGTable
	result := s.db.Where("is_remoted = true").Find(&queryReturns)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to list migrate gvg table: %s", result.Error)
	}
	returns := make([]*spdb.MigrateGVGUnitMeta, 0)
	for _, queryReturn := range queryReturns {
		returns = append(returns, &spdb.MigrateGVGUnitMeta{
			GlobalVirtualGroupID: queryReturn.GlobalVirtualGroupID,
			VirtualGroupFamilyID: queryReturn.VirtualGroupFamilyID,
			RedundancyIndex:      queryReturn.RedundancyIndex,
			BucketID:             queryReturn.BucketID,
			IsSecondary:          queryReturn.IsSecondary,
			IsConflicted:         queryReturn.IsConflicted,
			SrcSPID:              queryReturn.SrcSPID,
			DestSPID:             queryReturn.DestSPID,
			LastMigratedObjectID: queryReturn.LastMigratedObjectID,
			MigrateStatus:        queryReturn.MigrateStatus,
		})
	}
	return returns, nil
}

func (s *SpDBImpl) ListMigrateGVGUnitsByBucketID(bucketID uint64) ([]*spdb.MigrateGVGUnitMeta, error) {
	var queryReturns []MigrateGVGTable
	result := s.db.Where("bucket_id = ?", bucketID).Find(&queryReturns)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to query migrate gvg table: %s", result.Error)
	}
	returns := make([]*spdb.MigrateGVGUnitMeta, 0)
	for _, queryReturn := range queryReturns {
		returns = append(returns, &spdb.MigrateGVGUnitMeta{
			GlobalVirtualGroupID: queryReturn.GlobalVirtualGroupID,
			VirtualGroupFamilyID: queryReturn.VirtualGroupFamilyID,
			RedundancyIndex:      queryReturn.RedundancyIndex,
			BucketID:             queryReturn.BucketID,
			IsSecondary:          queryReturn.IsSecondary,
			IsConflicted:         queryReturn.IsConflicted,
			SrcSPID:              queryReturn.SrcSPID,
			DestSPID:             queryReturn.DestSPID,
			LastMigratedObjectID: queryReturn.LastMigratedObjectID,
			MigrateStatus:        queryReturn.MigrateStatus,
		})
	}
	return returns, nil
}
