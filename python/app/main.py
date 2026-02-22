from fastapi import FastAPI, HTTPException
from .models import RenterProfile, RiskAssessmentResult
from .services import RiskService

app = FastAPI(title="Risk Assessment Microservice")
risk_service = RiskService()

@app.post("/api/v1/assess", response_model=RiskAssessmentResult)
async def assess(profile: RenterProfile):
    try:
        return risk_service.assess_renter(profile)
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))

@app.get("/health")
async def health_check():
    return {"status": "up"}
