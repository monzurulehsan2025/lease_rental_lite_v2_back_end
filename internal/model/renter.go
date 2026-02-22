package model

import "time"

// RenterProfile represents the data provided by a renter for qualification
type RenterProfile struct {
	ID            string    `json:"id"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Email         string    `json:"email"`
	AnnualIncome  float64   `json:"annual_income"`
	CreditScore   int       `json:"credit_score"`
	MonthlyRent   float64   `json:"monthly_rent"`
	EmploymentStatus string `json:"employment_status"`
	CreatedAt     time.Time `json:"created_at"`
}

// RiskAssessmentResult represents the outcome of a risk evaluation
type RiskAssessmentResult struct {
	RenterID       string    `json:"renter_id"`
	RiskScore      int       `json:"risk_score"`       // Scale 0-100 (higher is riskier)
	Qualified      bool      `json:"qualified"`
	MaxCoverage    float64   `json:"max_coverage"`
	Recommendation string    `json:"recommendation"`
	EvaluatedAt    time.Time `json:"evaluated_at"`
}
