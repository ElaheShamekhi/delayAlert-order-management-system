package storage

import (
	"time"
)

func (s *Storage) CreateDelayReport(orderId int, newTime time.Time) error {
	return s.db.Create(&DelayReport{
		NewEstimatedTime: newTime,
		OrderId:          uint(orderId),
	}).Error
}

func (s *Storage) GetOrderDetails(orderId int) (*Order, error) {
	var order Order
	query := s.db.Model(&Order{}).Where("id = ?", uint(orderId)).Find(&order)
	if query.Error != nil {
		return nil, query.Error
	}
	return &order, nil
}

func (s *Storage) IsOrderInTripExists(orderId int) (bool, error) {
	var IsExist bool
	query := s.db.Raw("SELECT EXISTS(SELECT id FROM trip WHERE order_id = ? and deleted_at IS NULL) ", orderId).
		First(&IsExist)
	if query.Error != nil {
		return false, query.Error
	}
	return IsExist, nil
}

func (s *Storage) GetTripDetails(orderId int) (*Trip, error) {
	var trip Trip
	query := s.db.Model(&Trip{}).Where("order_id = ?", orderId).Find(&trip)
	if query.Error != nil {
		return nil, query.Error
	}
	return &trip, nil
}

func (s *Storage) GetFirstDelayNotChecked() (*DelayReport, error) {
	var result DelayReport
	query := s.db.Model(&DelayReport{}).Where("order_id is null and is_checked = false").Order("created_at").First(&result)
	if query.Error != nil {
		return nil, query.Error
	}
	return &result, nil
}

func (s *Storage) AssignDelayToAgent(reportId, agentId uint) error {
	query := s.db.Model(&DelayReport{}).Where("id = ?", reportId).Updates(DelayReport{
		AgentId: agentId,
	})
	if query.Error != nil {
		return query.Error
	}
	return nil
}

type VendorsDelayWeeklyReport struct {
	WeekStart    time.Time
	VendorId     uint
	TotalMinutes float64
}

func (s *Storage) VendorsDelayWeeklyReport() ([]VendorsDelayWeeklyReport, error) {
	var results []VendorsDelayWeeklyReport
	query := s.db.Raw(`
    SELECT 
        DATE_TRUNC('week', "order"."register_time") AS "week_start",
        "order"."vendor_id",
        SUM(EXTRACT(EPOCH FROM ("order"."delivered_time" - "order"."register_time")) / 60) AS "total_minutes"
    FROM 
        "order"
    GROUP BY 
        "week_start", 
        "order"."vendor_id"`).Scan(&results)

	if query.Error != nil {
		return nil, query.Error
	}
	return results, nil
}
