import { Router } from "express";
import {
  transfer,
  deposit,
  withdraw
} from "./controllers/transactionController.js";

const router = Router();

router.post("/transactions/transfer", transfer);
router.post("/transactions/deposit", deposit);
router.post("/transactions/withdraw", withdraw);

export default router;
