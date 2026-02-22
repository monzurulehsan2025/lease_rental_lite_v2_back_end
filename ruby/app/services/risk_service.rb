require 'time'

class RiskService
  def assess_renter(profile)
    annual_income = profile['annual_income'].to_f
    credit_score = profile['credit_score'].to_i
    monthly_rent = profile['monthly_rent'].to_f

    rent_to_income_ratio = (monthly_rent * 12) / annual_income
    
    risk_score = 50 # Base score

    # Credit score impact
    if credit_score > 750
      risk_score -= 30
    elsif credit_score > 650
      risk_score -= 15
    elsif credit_score < 600
      risk_score += 25
    end

    # Income ratio impact
    if rent_to_income_ratio > 0.4
      risk_score += 20
    elsif rent_to_income_ratio < 0.25
      risk_score -= 10
    end

    # Qualification criteria
    qualified = risk_score < 70 && credit_score >= 600
    
    recommendation = "Approve"
    if !qualified
      recommendation = "Reject - High Risk"
    elsif risk_score > 40
      recommendation = "Approve with Co-signer"
    end

    max_coverage = qualified ? monthly_rent * 12 : 0.0

    {
      renter_id: profile['id'],
      risk_score: risk_score,
      qualified: qualified,
      max_coverage: max_coverage,
      recommendation: recommendation,
      evaluated_at: Time.now.iso8601
    }
  end
end
