import router from './controllers/healthchecks';
import express, { Express } from 'express';

const app: Express = express();
const port = 3000;

app.use(router);

app.listen(port, () => {
  console.log(`Listening on port ${port}`);
});
