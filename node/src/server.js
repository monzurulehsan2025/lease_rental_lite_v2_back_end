import app from './app.js';

const PORT = process.env.PORT || 8080;

app.listen(PORT, () => {
    console.log(`Risk Assessment Service (Node.js) starting on port ${PORT}...`);
});
