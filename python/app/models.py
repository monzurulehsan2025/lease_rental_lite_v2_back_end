from pydantic import BaseModel, EmailStr
from datetime import datetime
from typing import Optional

class RenterProfile(BaseModel):
    id: str
    first_name: str
    last_name: str
    email: EmailStr
    annual_income: float
    credit_score: int
    monthly_rent: float
    employment_status: str
    created_at: Optional[datetime] = None

class RiskAssessmentResult(BaseModel):
    renter_id: str
    risk_score: int
    qualified: bool
    max_coverage: float
    recommendation: str
    evaluated_at: datetime
