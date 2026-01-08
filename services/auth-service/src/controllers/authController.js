import jwt from "jsonwebtoken";
import bcrypt from "bcryptjs";
import { db } from "../database.js";

export async function login(req, res) {
  const { email, password } = req.body;

  const result = await db.query("SELECT * FROM users WHERE email = $1", [email]);
  const user = result.rows[0];

  if (!user) return res.status(401).json({ error: "Invalid credentials" });

  const match = bcrypt.compareSync(password, user.password_hash);
  if (!match) return res.status(401).json({ error: "Invalid credentials" });

  const token = jwt.sign({ userId: user.id }, process.env.JWT_SECRET, { expiresIn: "1h" });

  return res.json({ token, "UserId": user.id });
}

export async function refresh(req, res) {
  const { token } = req.body;

  try {
    const decoded = jwt.verify(token, process.env.JWT_SECRET);
    const newToken = jwt.sign({ userId: decoded.userId }, process.env.JWT_SECRET, { expiresIn: "1h" });

    return res.json({ token: newToken });
  } catch (err) {
    return res.status(401).json({ error: "Invalid token" });
  }
}
