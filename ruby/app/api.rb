require 'sinatra'
require 'json'
require_relative 'services/risk_service'

set :port, 8080

risk_service = RiskService.new

post '/api/v1/assess' do
  content_type :json
  begin
    profile = JSON.parse(request.body.read)
    result = risk_service.assess_renter(profile)
    result.to_json
  catch => e
    status 500
    { error: e.message }.to_json
  end
end

get '/health' do
  content_type :json
  { status: 'up' }.to_json
end
