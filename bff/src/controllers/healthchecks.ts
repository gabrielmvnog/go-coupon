import express, { Request, Response } from 'express';
const router = express.Router();

router.get('/health', function (req: Request, res: Response) {
  res.send('OK');
});

export default router;
