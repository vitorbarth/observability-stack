import express from "express";
import dotenv from "dotenv";
import router from "./routes.js";
import * as api from '@opentelemetry/api';
import { AsyncHooksContextManager } from '@opentelemetry/context-async-hooks';

dotenv.config();

const contextManager = new AsyncHooksContextManager();
contextManager.enable();
api.context.setGlobalContextManager(contextManager);

const app = express();
app.use(express.json());
app.use(router);

const port = process.env.PORT || 3002;
app.listen(port, () => console.log(`account-service running on port ${port}`));
