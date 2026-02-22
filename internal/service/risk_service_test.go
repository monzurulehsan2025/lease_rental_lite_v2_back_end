package service

import (
	"testing"

	"github.com/securerent/risk-assessment-service/internal/model"
)

func TestAssessRenter(t *testing.T) {
	s := NewRiskService()

	tests := []struct {
		name           string
		profile        model.RenterProfile
		wantQualified  bool
		minRiskScore   int
		maxRiskScore   int
	}{
		{
			name: "Low risk renter",
			profile: model.RenterProfile{
				AnnualIncome: 100000,
				CreditScore:  780,
				MonthlyRent:  2000,
			},
			wantQualified: true,
			minRiskScore:  0,
			maxRiskScore:  30,
		},
		{
			name: "High risk - low credit",
			profile: model.RenterProfile{
				AnnualIncome: 50000,
				CreditScore:  550,
				MonthlyRent:  2500,
			},
			wantQualified: false,
			minRiskScore:  70,
			maxRiskScore:  100,
		},
		{
			name: "Moderate risk - high rent ratio",
			profile: model.RenterProfile{
				AnnualIncome: 60000,
				CreditScore:  680,
				MonthlyRent:  2500,
			},
			wantQualified: true,
			minRiskScore:  40,
			maxRiskScore:  65,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.AssessRenter(tt.profile)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got.Qualified != tt.wantQualified {
				t.Errorf("Qualified = %v, want %v", got.Qualified, tt.wantQualified)
			}
			if got.RiskScore < tt.minRiskScore || got.RiskScore > tt.maxRiskScore {
				t.Errorf("RiskScore = %d, want between %d and %d", got.RiskScore, tt.minRiskScore, tt.maxRiskScore)
			}
		}
	}
}
