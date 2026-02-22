package com.securerent.riskassessment.service;

import com.securerent.riskassessment.model.RenterProfile;
import com.securerent.riskassessment.model.RiskAssessmentResult;
import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.*;

public class RiskServiceTest {

    private final RiskService service = new RiskService();

    @Test
    void testLowRiskRenter() {
        RenterProfile profile = new RenterProfile();
        profile.setAnnualIncome(100000);
        profile.setCreditScore(780);
        profile.setMonthlyRent(2000);

        RiskAssessmentResult result = service.assessRenter(profile);
        assertTrue(result.isQualified());
        assertTrue(result.getRiskScore() <= 30);
        assertEquals("Approve", result.getRecommendation());
    }

    @Test
    void testHighRiskRenter() {
        RenterProfile profile = new RenterProfile();
        profile.setAnnualIncome(50000);
        profile.setCreditScore(550);
        profile.setMonthlyRent(2500);

        RiskAssessmentResult result = service.assessRenter(profile);
        assertFalse(result.isQualified());
        assertTrue(result.getRiskScore() >= 70);
        assertEquals("Reject - High Risk", result.getRecommendation());
    }
}
