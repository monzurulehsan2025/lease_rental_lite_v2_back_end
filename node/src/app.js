import express from 'express';
import { assess, healthCheck } from './handlers/riskHandler.js';

const app = express();

app.use(express.json());

// Routes
app.post('/api/v1/assess', assess);
app.get('/health', healthCheck);

export default app;
