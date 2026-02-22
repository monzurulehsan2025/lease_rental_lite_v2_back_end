from datetime import datetime
from .models import RenterProfile, RiskAssessmentResult

class RiskService:
    def assess_renter(self, profile: RenterProfile) -> RiskAssessmentResult:
        rent_to_income_ratio = (profile.monthly_rent * 12) / profile.annual_income
        
        risk_score = 50  # Base score

        # Credit score impact
        if profile.credit_score > 750:
            risk_score -= 30
        elif profile.credit_score > 650:
            risk_score -= 15
        elif profile.credit_score < 600:
            risk_score += 25

        # Income ratio impact
        if rent_to_income_ratio > 0.4:
            risk_score += 20
        elif rent_to_income_ratio < 0.25:
            risk_score -= 10

        # Qualification criteria
        qualified = risk_score < 70 and profile.credit_score >= 600
        
        recommendation = "Approve"
        if not qualified:
            recommendation = "Reject - High Risk"
        elif risk_score > 40:
            recommendation = "Approve with Co-signer"

        max_coverage = profile.monthly_rent * 12 if qualified else 0.0

        return RiskAssessmentResult(
            renter_id=profile.id,
            risk_score=risk_score,
            qualified=qualified,
            max_coverage=max_coverage,
            recommendation=recommendation,
            evaluated_at=datetime.now()
        )
