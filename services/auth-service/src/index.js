import express from "express";
import dotenv from "dotenv";
import router from "./routes.js";
import * as api from '@opentelemetry/api';
import { AsyncHooksContextManager } from '@opentelemetry/context-async-hooks';

const contextManager = new AsyncHooksContextManager();
contextManager.enable();
api.context.setGlobalContextManager(contextManager);

dotenv.config();

const app = express();
app.use(express.json());
app.use(router);

const port = process.env.PORT || 3001;
app.listen(port, () => console.log(`auth-service running on port ${port}`));
