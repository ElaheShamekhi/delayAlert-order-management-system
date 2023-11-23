package delay

import (
	"delayAlert-order-management-system/internal/validation"
	"delayAlert-order-management-system/storage"
	"time"
)

const (
	TripAssigned  = "ASSIGNED"
	TripAtVendor  = "AT_VENDOR"
	TripPicked    = "PICKED"
	TripDelivered = "Delivered"
)

func (s *Service) CreateOrderDelay(request CreateOrderDelay) (*NewEstimatedTime, error) {
	var result NewEstimatedTime
	order, err := s.store.GetOrderDetails(request.OrderId)
	if err != nil {
		return nil, err
	}
	v := validation.New().SetMethod("delay.CreateOrderDelay").Validate(
		validation.TimeValidToSubmitDelay(order.RegisterTime, order.DeliveryTime),
	)
	if v.Error() != nil {
		return nil, v.Error()
	}
	found, err := s.store.IsOrderInTripExists(request.OrderId)
	if err != nil {
		return nil, err
	}
	if found {
		trip, err := s.store.GetTripDetails(request.OrderId)
		if err != nil {
			return nil, err
		}
		if checkTripStateIsNotDelivered(trip.State) {
			/*
					!!!!!!This api url https://run.mocky.io/v3/122c2796-5df4-461c-ab75-87c1192b17f7 does not work generate
				    new time.Time that is order.RegisterTime.Add(time.Hour*2)!!!!!!
					newTime, err := s.client.GetNewEstimatedDelay()
					if err != nil {
						return err
					}
			*/
			result.NewTime = order.RegisterTime.Add(time.Hour * 2)
		}
	}
	err = s.store.CreateDelayReport(request.OrderId, order.RegisterTime.Add(time.Hour*2))
	if err != nil {
		return nil, err

	}
	return &result, nil
}

func (s *Service) AssignDelayToAgent(agentId int) error {
	report, err := s.store.GetFirstDelayNotChecked()
	if err != nil {
		return err
	}
	err = s.store.AssignDelayToAgent(report.ID, uint(agentId))
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) VendorsDelayWeeklyReport() ([]VendorDelayWeeklyReport, error) {
	reports, err := s.store.VendorsDelayWeeklyReport()
	if err != nil {
		return nil, err
	}
	results := make([]VendorDelayWeeklyReport, len(reports), len(reports))
	for i, report := range reports {
		results[i] = toVendorDelayWeeklyReport(report)
	}
	return results, nil
}

func toVendorDelayWeeklyReport(report storage.VendorsDelayWeeklyReport) VendorDelayWeeklyReport {
	return VendorDelayWeeklyReport{
		WeekStart:    report.WeekStart,
		VendorId:     report.VendorId,
		TotalMinutes: report.TotalMinutes,
	}
}

func checkTripStateIsNotDelivered(state string) bool {
	switch state {
	case TripAssigned, TripAtVendor, TripPicked:
		return true
	}
	return false
}
