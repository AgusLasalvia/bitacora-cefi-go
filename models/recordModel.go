package models

import (
	"bitacora/config"
	"bitacora/core"
)

func AddRecord(record core.Record) error {
	_, err := config.Database.Exec(`
		UPDATE records
		SET name=?, lab=?, equipment=?, startDateTime=?, endDateTime=?, 
		    received=?, returned=?, comments=?, timestamp=?
		WHERE id=?`,
		record.Name,
		record.Lab,
		record.Equipment,
		record.StartDateTime,
		record.EndDateTime,
		record.Received,
		record.Returned,
		record.Comments,
		record.Timestamp,
		record.ID,
	)
	return err
}
