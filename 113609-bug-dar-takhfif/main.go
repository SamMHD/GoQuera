package main

import "fmt"

type (
	JalaliDate struct {
		Day   int
		Month int
		Year  int
	}
	Time struct {
		JalaliDate
		Hour   int
		Minute int
		Second int
	}
	TimePeriod struct {
		Start Time
		End   Time
	}

	Coordinate struct {
		X int
		Y int
	}
	Area struct {
		UpLeft    Coordinate
		DownRight Coordinate
	}

	Discount struct {
		Percent        uint64
		EffectiveArea  Area
		IsActive       bool
		ApplicableTime TimePeriod
	}
	Discounts map[string]*Discount

	SnappSystem struct {
		AllDiscounts        Discounts
		CurrentTime         Time
		BasePriceCalculatur func(TripRequest) Price
	}

	Price uint64

	TripRequest struct {
		From         Coordinate
		To           Coordinate
		basePrice    Price
		UserDiscount string
		RequestTime  Time
	}
)

func (system SnappSystem) findDiscountByName(userDiscount string) (*Discount, error) {
	discount, exists := system.AllDiscounts[userDiscount]
	if !exists {
		return nil, fmt.Errorf("discount [%s] doesn't exists", userDiscount)
	}
	return discount, nil
}

func (area Area) isPointInArea(point Coordinate) bool {
	return point.X <= area.DownRight.X &&
		area.UpLeft.X <= point.X &&
		point.Y <= area.DownRight.Y &&
		area.UpLeft.Y <= point.Y
}

func (t Time) StampValue() uint64 {
	var valueOrder = []int{t.Second, t.Minute, t.Hour, t.Day, t.Month, t.Year}
	result := uint64(0)
	for i, v := range valueOrder {
		result |= uint64(v) << (6 * i)
	}
	return result
}

func (t Time) isValid() bool {
	var DaysInMonth = []int{31, 31, 31, 31, 31, 31, 30, 30, 30, 30, 30, 30}

	result := (0 < t.Year && t.Year < 2000)
	result = result && (0 < t.Month && t.Month < 13)
	result = result && (0 < t.Day && t.Day <= DaysInMonth[t.Month])
	result = result && (0 <= t.Hour && t.Hour < 24)
	result = result && (0 <= t.Minute && t.Minute < 60)
	result = result && (0 <= t.Second && t.Second < 60)
	return result
}

func (base Time) IsAfter(time Time) bool {

	return base.StampValue() > time.StampValue()

	// faster response for common cases
	// absoloutelyAfter := base.Year >= time.Year &&
	// 	base.Month >= time.Month &&
	// 	base.Day > time.Day
	// if absoloutelyAfter {
	// 	return true
	// }

	// yEq := base.Year == time.Year
	// mEq := base.Month == time.Month
	// dEq := base.Day == time.Day

	// yAfter := base.Year > time.Year
	// mAfter := base.Month > time.Month
	// dAfter := base.Day > time.Day

	// if yAfter {
	// 	return true
	// } else if yEq && mAfter {
	// 	return true
	// } else if yEq && mEq && dAfter {
	// 	return true
	// } else if yEq && mEq && dEq {
	// 	if base.Hour > time.Hour {
	// 		return true
	// 	} else if base.Hour == time.Hour && base.Minute > time.Minute {
	// 		return true
	// 	} else if base.Hour == time.Hour &&
	// 		base.Minute == time.Minute &&
	// 		base.Second > time.Second {
	// 		return true
	// 	} else {
	// 		return false
	// 	}
	// } else {
	// 	return false
	// }
}

func (period TimePeriod) isTimeInPeriod(t Time) bool {
	if !t.isValid() {
		return false
	}
	return period.Start.StampValue() <= t.StampValue() && t.StampValue() <= period.End.StampValue()
}

func (disc Discount) Apply(basePrice Price) Price {
	return Price(((100 - disc.Percent) * uint64(basePrice)) / 100)
}

func (system SnappSystem) CalculateEffectivePrice(trip TripRequest) (Price, error) {

	discount, err := system.findDiscountByName(trip.UserDiscount)
	if err != nil {
		return 0, err
	}

	if !discount.EffectiveArea.isPointInArea(trip.From) {
		return 0, fmt.Errorf("user is not in discount area")
	}

	if !discount.IsActive {
		return 0, fmt.Errorf("discount is not currently active")
	}

	if !discount.ApplicableTime.isTimeInPeriod(trip.RequestTime) {
		return 0, fmt.Errorf("discount is not applicable for this time")
	}

	if !system.CurrentTime.IsAfter(trip.RequestTime) {
		return 0, fmt.Errorf("trip time can't be in the future")
	}

	trip.basePrice = system.BasePriceCalculatur(trip)
	finalPrice := discount.Apply(trip.basePrice)

	return finalPrice, nil
}

// func main() {
// 	t1 := Time{JalaliDate{1400, 5, 2}, 12, 23, 10}
// 	t2 := Time{JalaliDate{1400, 5, 3}, 12, 23, 10}
// 	p := TimePeriod{t1, t2}

// 	ti := Time{JalaliDate{2003, 5, 2}, 13, 23, 10}
// 	fmt.Println(ti.isValid())
// 	fmt.Println(p.isTimeInPeriod(ti))
// }
