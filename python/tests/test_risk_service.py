from app.services import RiskService
from app.models import RenterProfile

def test_low_risk_renter():
    service = RiskService()
    profile = RenterProfile(
        id="r1",
        first_name="John",
        last_name="Doe",
        email="john@example.com",
        annual_income=100000,
        credit_score=780,
        monthly_rent=2000,
        employment_status="Employed"
    )
    result = service.assess_renter(profile)
    assert result.qualified is True
    assert result.risk_score <= 30
    assert result.recommendation == "Approve"

def test_high_risk_renter():
    service = RiskService()
    profile = RenterProfile(
        id="r2",
        first_name="Jane",
        last_name="Doe",
        email="jane@example.com",
        annual_income=50000,
        credit_score=550,
        monthly_rent=2500,
        employment_status="Self-Employed"
    )
    result = service.assess_renter(profile)
    assert result.qualified is False
    assert result.risk_score >= 70
    assert result.recommendation == "Reject - High Risk"
