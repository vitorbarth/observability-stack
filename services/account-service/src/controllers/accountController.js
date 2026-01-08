import { db } from "../database.js";

// GET /accounts/:id/balance
export async function getBalance(req, res) {
  const { id } = req.params;

  try {
    const result = await db.query(
      "SELECT balance FROM accounts WHERE id = $1",
      [id]
    );

    if (result.rowCount === 0) {
      return res.status(404).json({ error: "Account not found" });
    }

    return res.json({ id, balance: result.rows[0].balance });
  } catch (err) {
    console.error("Error fetching balance", err);
    return res.status(500).json({ error: "Internal server error" });
  }
}

// GET /accounts/:id/statement
export async function getStatement(req, res) {
  const { id } = req.params;

  try {
    const result = await db.query(
      `
      SELECT 
        id,
        type,
        amount,
        created_at
      FROM transactions
      WHERE account_id = $1
      ORDER BY created_at DESC
      `,
      [id]
    );

    return res.json({ accountId: id, statement: result.rows });
  } catch (err) {
    console.error("Error fetching statement", err);
    return res.status(500).json({ error: "Internal server error" });
  }
}
