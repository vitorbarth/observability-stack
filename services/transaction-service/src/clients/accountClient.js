import axios from "axios" 
import dotenv from "dotenv";

dotenv.config();

export default class accountClient {
    constructor() {
        this.baseUrl = process.env.ACCOUNT_SERVICE_URL;
    }

    async getBalance(accountId) {
        try {
            const response = await axios.get(`${this.baseUrl}/accounts/${accountId}/balance`);
            return response.data.balance;
        } catch (err) {
            console.error('Error calling account API:', err.response?.data || err.message);
            throw new Error('ACCOUNT_SERVICE_UNAVAILABLE');
        }
    }
}