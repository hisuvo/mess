package dto

import "time"

type ResponseMile struct {
	ID        uint      `json:"id"`
	TenantID  uint      `json:"tenant_id"`
	UserID    uint      `json:"user_id"`
	Date      time.Time `json:"date"`
	Breakfast float64   `json:"breakfast"`
	Lunch     float64   `json:"lunch"`
	Dinner    float64   `json:"dinner"`
	Total     float64   `json:"total"`
	Note      string    `json:"note"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ResponseMonthSummary struct {
	Year         int     `json:"year"`
	Month        int     `json:"month"`
	TotalMeals   int     `json:"total_meals"`
	TotalBreakfast float64   `json:"total_breakfast"`
	TotalLunch     float64   `json:"total_lunch"`
	TotalDinner    float64   `json:"total_dinner"`
	GrandTotal     float64   `json:"grand_total"`
	AvgPerDay      float64   `json:"avg_per_day"`
	ByDate         []DateDetail `json:"by_date"`
}

type DateDetail struct {
	Date      time.Time `json:"date"`
	Breakfast float64   `json:"breakfast"`
	Lunch     float64   `json:"lunch"`
	Dinner    float64   `json:"dinner"`
	Total     float64   `json:"total"`
}