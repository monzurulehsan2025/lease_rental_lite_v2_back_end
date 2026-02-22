package com.securerent.riskassessment.service;

import com.securerent.riskassessment.model.RenterProfile;
import com.securerent.riskassessment.model.RiskAssessmentResult;
import org.springframework.stereotype.Service;
import java.time.LocalDateTime;

@Service
public class RiskService {

    public RiskAssessmentResult assessRenter(RenterProfile profile) {
        double rentToIncomeRatio = (profile.getMonthlyRent() * 12) / profile.getAnnualIncome();

        int riskScore = 50; // Base score

        // Credit score impact
        if (profile.getCreditScore() > 750) {
            riskScore -= 30;
        } else if (profile.getCreditScore() > 650) {
            riskScore -= 15;
        } else if (profile.getCreditScore() < 600) {
            riskScore += 25;
        }

        // Income ratio impact
        if (rentToIncomeRatio > 0.4) {
            riskScore += 20;
        } else if (rentToIncomeRatio < 0.25) {
            riskScore -= 10;
        }

        // Qualification criteria
        boolean qualified = riskScore < 70 && profile.getCreditScore() >= 600;

        String recommendation = "Approve";
        if (!qualified) {
            recommendation = "Reject - High Risk";
        } else if (riskScore > 40) {
            recommendation = "Approve with Co-signer";
        }

        double maxCoverage = qualified ? profile.getMonthlyRent() * 12 : 0.0;

        return RiskAssessmentResult.builder()
                .renterId(profile.getId())
                .riskScore(riskScore)
                .qualified(qualified)
                .maxCoverage(maxCoverage)
                .recommendation(recommendation)
                .evaluatedAt(LocalDateTime.now())
                .build();
    }
}
