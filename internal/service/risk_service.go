package service

import (
	"time"

	"github.com/securerent/risk-assessment-service/internal/model"
)

type RiskService interface {
	AssessRenter(profile model.RenterProfile) (model.RiskAssessmentResult, error)
}

type riskService struct{}

func NewRiskService() RiskService {
	return &riskService{}
}

// AssessRenter implements a basic risk assessment algorithm for rent coverage
func (s *riskService) AssessRenter(profile model.RenterProfile) (model.RiskAssessmentResult, error) {
	// Simple logic: Rent-to-income ratio and Credit Score
	rentToIncomeRatio := (profile.MonthlyRent * 12) / profile.AnnualIncome
	
	riskScore := 50 // Base score

	// Credit score impact
	if profile.CreditScore > 750 {
		riskScore -= 30
	} else if profile.CreditScore > 650 {
		riskScore -= 15
	} else if profile.CreditScore < 600 {
		riskScore += 25
	}

	// Income ratio impact
	if rentToIncomeRatio > 0.4 {
		riskScore += 20
	} else if rentToIncomeRatio < 0.25 {
		riskScore -= 10
	}

	// Qualification criteria
	qualified := riskScore < 70 && profile.CreditScore >= 600
	
	recommendation := "Approve"
	if !qualified {
		recommendation = "Reject - High Risk"
	} else if riskScore > 40 {
		recommendation = "Approve with Co-signer"
	}

	maxCoverage := 0.0
	if qualified {
		maxCoverage = profile.MonthlyRent * 12 // Coverage up to one year of rent
	}

	return model.RiskAssessmentResult{
		RenterID:       profile.ID,
		RiskScore:      riskScore,
		Qualified:      qualified,
		MaxCoverage:    maxCoverage,
		Recommendation: recommendation,
		EvaluatedAt:    time.Now(),
	}, nil
}
