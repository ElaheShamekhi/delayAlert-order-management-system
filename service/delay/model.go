package delay

import "time"

type CreateOrderDelay struct {
	OrderId int `json:"order_id"`
}

type VendorDelayWeeklyReport struct {
	WeekStart    time.Time `json:"week_start"`
	VendorId     uint      `json:"vendor_id"`
	TotalMinutes float64   `json:"total_minutes"`
}
