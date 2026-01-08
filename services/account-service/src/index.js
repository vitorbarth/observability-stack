import express from "express";
import dotenv from "dotenv";
import router from "./routes.js";
import { initTelemetry } from "./telemetry.js";

dotenv.config();
initTelemetry();

const app = express();
app.use(express.json());
app.use(router);

const port = process.env.PORT || 3002;
app.listen(port, () => console.log(`account-service running on port ${port}`));
