package domain

import (
	"bia-backend/structures"
	"gorm.io/gorm"
)

func FetchMonthly(db *gorm.DB, startDate string, endDate string) (*[]structures.DataGraph, error) {
	var consumptions []structures.DataGraph

	query := "SELECT EXTRACT(YEAR FROM TO_DATE(date,'YYYY-MM-DD')) AS year, EXTRACT(MONTH FROM TO_DATE(date,'YYYY-MM-DD')) AS month,  " +
		"SUM(active_energy) AS active_energy, SUM(reactive_energy) AS reactive_energy, SUM(capacitive_reactive) AS capacitive_reactive, SUM(solar) AS solar " +
		"FROM consumptions " +
		"EXTRACT(YEAR FROM TO_DATE(date,'YYYY-MM-DD')) BETWEEN EXTRACT(YEAR FROM TO_DATE(?,'YYYY-MM-DD')) AND EXTRACT(YEAR FROM TO_DATE(?,'YYYY-MM-DD'))" +
		"AND EXTRACT(MONTH FROM TO_DATE(date,'YYYY-MM-DD')) BETWEEN EXTRACT(MONTH FROM TO_DATE(?,'YYYY-MM-DD')) AND EXTRACT(MONTH FROM TO_DATE(?,'YYYY-MM-DD'))" +
		"GROUP BY GROUP BY year, month"

	err := db.Raw(query).Scan(&consumptions).Error
	if err != nil {
		return nil, err
	}

	return &consumptions, nil
}
