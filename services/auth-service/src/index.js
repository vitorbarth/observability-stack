import express from "express";
import dotenv from "dotenv";
import router from "./routes.js";
import { initTelemetry } from "./telemetry.js";

dotenv.config();
initTelemetry();

const app = express();
app.use(express.json());
app.use(router);

const port = process.env.PORT || 3001;
app.listen(port, () => console.log(`auth-service running on port ${port}`));
