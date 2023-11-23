package delay

import (
	"delayAlert-order-management-system/internal/validation"
	"delayAlert-order-management-system/storage"
)

const (
	TripAssigned  = "ASSIGNED"
	TripAtVendor  = "AT_VENDOR"
	TripPicked    = "PICKED"
	TripDelivered = "Delivered"
)

func (s *Service) CreateOrderDelay(request CreateOrderDelay) error {
	order, err := s.store.GetOrderDetails(request.OrderId)
	if err != nil {
		return err //TODO: change error format
	}
	v := validation.New().SetMethod("delay.CreateOrderDelay").Validate(
		validation.TimeValidToSubmitDelay(order.RegisterTime, order.DeliveryTime),
	)
	if v.Error() != nil {
		return v.Error()
	}
	found, err := s.store.IsOrderInTripExists(request.OrderId)
	if err != nil {
		return err
	}
	if found {
		trip, err := s.store.GetTripDetails(request.OrderId)
		if err != nil {
			return err
		}
		if checkTripStateIsNotDelivered(trip.State) {
			newTime, err := s.client.GetNewEstimatedDelay()
			if err != nil {
				return err
			}
			err = s.store.CreateDelayReport(request.OrderId, *newTime)
			if err != nil {
				return err
			}
		} else {
			// TODO : insert order to delay queue
		}
	} else {
		// TODO : insert order to delay queue
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
