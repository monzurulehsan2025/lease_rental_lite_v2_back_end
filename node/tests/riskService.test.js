import { RiskService } from '../src/services/riskService.js';

describe('RiskService', () => {
    const service = new RiskService();

    test('Low risk renter qualification', () => {
        const profile = {
            id: 'r1',
            annual_income: 100000,
            credit_score: 780,
            monthly_rent: 2000
        };
        const result = service.assessRenter(profile);
        expect(result.qualified).toBe(true);
        expect(result.risk_score).toBeLessThanOrEqual(30);
        expect(result.recommendation).toBe('Approve');
    });

    test('High risk - low credit score rejection', () => {
        const profile = {
            id: 'r2',
            annual_income: 50000,
            credit_score: 550,
            monthly_rent: 2500
        };
        const result = service.assessRenter(profile);
        expect(result.qualified).toBe(false);
        expect(result.risk_score).toBeGreaterThanOrEqual(70);
        expect(result.recommendation).toBe('Reject - High Risk');
    });

    test('Moderate risk - high rent ratio qualification', () => {
        const profile = {
            id: 'r3',
            annual_income: 60000,
            credit_score: 680,
            monthly_rent: 2500
        };
        const result = service.assessRenter(profile);
        expect(result.qualified).toBe(true);
        expect(result.risk_score).toBeGreaterThanOrEqual(40);
        expect(result.risk_score).toBeLessThanOrEqual(65);
        expect(result.recommendation).toBe('Approve with Co-signer');
    });
});
