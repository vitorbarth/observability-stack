import { db } from "../database.js";

// Helper – registra histórico
async function insertTransaction(accountId, type, amount) {
  await db.query(
    `
      INSERT INTO transactions (account_id, type, amount)
      VALUES ($1, $2, $3)
    `,
    [accountId, type, amount]
  );
}

// -------------------------------------------
// POST /transactions/deposit
// -------------------------------------------
export async function deposit(req, res) {
  const { accountId, amount } = req.body;

  if (!accountId || !amount || amount <= 0) {
    return res.status(400).json({ error: "Invalid input" });
  }

  try {
    await db.query("BEGIN");

    await db.query(
      "UPDATE accounts SET balance = balance + $1 WHERE id = $2",
      [amount, accountId]
    );

    await insertTransaction(accountId, "DEPOSIT", amount);

    await db.query("COMMIT");

    return res.json({ message: "Deposit successful" });
  } catch (err) {
    await db.query("ROLLBACK");
    console.error("Deposit error", err);
    return res.status(500).json({ error: "Internal server error" });
  }
}

// -------------------------------------------
// POST /transactions/withdraw
// -------------------------------------------
export async function withdraw(req, res) {
  const { accountId, amount } = req.body;

  if (!accountId || !amount || amount <= 0) {
    return res.status(400).json({ error: "Invalid input" });
  }

  try {
    await db.query("BEGIN");

    const result = await db.query(
      "SELECT balance FROM accounts WHERE id = $1",
      [accountId]
    );

    if (result.rows[0].balance < amount) {
      await db.query("ROLLBACK");
      return res.status(400).json({ error: "Insufficient funds" });
    }

    await db.query(
      "UPDATE accounts SET balance = balance - $1 WHERE id = $2",
      [amount, accountId]
    );

    await insertTransaction(accountId, "WITHDRAW", amount);

    await db.query("COMMIT");

    return res.json({ message: "Withdraw successful" });
  } catch (err) {
    await db.query("ROLLBACK");
    console.error("Withdraw error", err);
    return res.status(500).json({ error: "Internal server error" });
  }
}

// -------------------------------------------
// POST /transactions/transfer
// -------------------------------------------
export async function transfer(req, res) {
  const { fromAccount, toAccount, amount } = req.body;

  if (!fromAccount || !toAccount || !amount || amount <= 0) {
    return res.status(400).json({ error: "Invalid input" });
  }

  try {
    await db.query("BEGIN");

    // saldo origem
    const result = await db.query(
      "SELECT balance FROM accounts WHERE id = $1",
      [fromAccount]
    );

    if (result.rows[0].balance < amount) {
      await db.query("ROLLBACK");
      return res.status(400).json({ error: "Insufficient funds" });
    }

    // debita origem
    await db.query(
      "UPDATE accounts SET balance = balance - $1 WHERE id = $2",
      [amount, fromAccount]
    );

    // credita destino
    await db.query(
      "UPDATE accounts SET balance = balance + $1 WHERE id = $2",
      [amount, toAccount]
    );

    // histórico
    await insertTransaction(fromAccount, "TRANSFER_OUT", amount);
    await insertTransaction(toAccount, "TRANSFER_IN", amount);

    await db.query("COMMIT");

    return res.json({ message: "Transfer successful" });
  } catch (err) {
    await db.query("ROLLBACK");
    console.error("Transfer error", err);
    return res.status(500).json({ error: "Internal server error" });
  }
}
