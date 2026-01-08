import { Router } from "express";
import { getBalance, getStatement } from "./controllers/accountController.js";

const router = Router();

router.get("/accounts/:id/balance", getBalance);
router.get("/accounts/:id/statement", getStatement);

export default router;
