export class RiskService {
  /**
   * Assess a renter's risk based on their profile.
   * @param {Object} profile - The renter's profile data
   * @returns {Object} The result of the risk assessment
   */
  assessRenter(profile) {
    const { annual_income, credit_score, monthly_rent } = profile;
    const rentToIncomeRatio = (monthly_rent * 12) / annual_income;

    let riskScore = 50; // Base score

    // Credit score impact
    if (credit_score > 750) {
      riskScore -= 30;
    } else if (credit_score > 650) {
      riskScore -= 15;
    } else if (credit_score < 600) {
      riskScore += 25;
    }

    // Income ratio impact
    if (rentToIncomeRatio > 0.4) {
      riskScore += 20;
    } else if (rentToIncomeRatio < 0.25) {
      riskScore -= 10;
    }

    // Qualification criteria
    const qualified = riskScore < 70 && credit_score >= 600;

    let recommendation = "Approve";
    if (!qualified) {
      recommendation = "Reject - High Risk";
    } else if (riskScore > 40) {
      recommendation = "Approve with Co-signer";
    }

    const maxCoverage = qualified ? monthly_rent * 12 : 0.0;

    return {
      renter_id: profile.id,
      risk_score: riskScore,
      qualified: qualified,
      max_coverage: maxCoverage,
      recommendation: recommendation,
      evaluated_at: new Date().toISOString()
    };
  }
}
