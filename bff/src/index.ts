import healthchecksRouter from './controllers/healthchecks';
import graphqlRouter from './controllers/graphql';
import express, { Express } from 'express';

const app: Express = express();
const port = 3000;

app.use(healthchecksRouter);
app.use(graphqlRouter);

app.listen(port, () => {
  console.log(`Listening on port ${port}`);
});
