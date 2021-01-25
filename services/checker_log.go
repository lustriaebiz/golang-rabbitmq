package services

import (
	"context"
	"fmt"
	"log"

	"adb/config"
	"adb/models"
)

const (
	table          = "checker_log"
	layoutDateTime = "2021-01-25 15:04:05"
)

// GetAll
func CheckerLogAll(ctx context.Context) ([]models.CheckerLogModel, error) {

	var checker_logs []models.CheckerLogModel

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id DESC", table)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var checker_log models.CheckerLogModel

		if err = rowQuery.Scan(
			&checker_log.ID,
			&checker_log.TopicName,
			&checker_log.QueueName,
			&checker_log.Message,
			&checker_log.CreatedAt,
			&checker_log.UpdatedAt,
		); err != nil {
			return nil, err
		}

		checker_logs = append(checker_logs, checker_log)
	}

	return checker_logs, nil
}
