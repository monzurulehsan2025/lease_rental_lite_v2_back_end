import { RiskService } from '../services/riskService.js';

const riskService = new RiskService();

export const assess = (req, res) => {
    try {
        const profile = req.body;

        // Basic validation
        if (!profile.annual_income || !profile.credit_score || !profile.monthly_rent) {
            return res.status(400).json({ error: "Missing required fields" });
        }

        const result = riskService.assessRenter(profile);
        res.json(result);
    } catch (error) {
        res.status(500).json({ error: "Internal server error during assessment" });
    }
};

export const healthCheck = (req, res) => {
    res.json({ status: "up" });
};
