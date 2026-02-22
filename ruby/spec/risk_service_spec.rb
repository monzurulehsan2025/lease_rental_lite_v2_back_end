require_relative '../app/services/risk_service'

RSpec.describe RiskService do
  let(:service) { RiskService.new }

  it 'qualifies a low risk renter' do
    profile = {
      'id' => 'r1',
      'annual_income' => 100000,
      'credit_score' => 780,
      'monthly_rent' => 2000
    }
    result = service.assess_renter(profile)
    expect(result[:qualified]).to be true
    expect(result[:risk_score]).to be <= 30
    expect(result[:recommendation]).to eq 'Approve'
  end

  it 'rejects a high risk renter with low credit' do
    profile = {
      'id' => 'r2',
      'annual_income' => 50000,
      'credit_score' => 550,
      'monthly_rent' => 2500
    }
    result = service.assess_renter(profile)
    expect(result[:qualified]).to be false
    expect(result[:risk_score]).to be >= 70
    expect(result[:recommendation]).to eq 'Reject - High Risk'
  end
end
