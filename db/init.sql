-- ===========================================
-- Criar Tabelas
-- ===========================================

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE accounts (
    id SERIAL PRIMARY KEY,
    owner_id INTEGER REFERENCES users(id),
    account_number VARCHAR(20) UNIQUE NOT NULL,
    balance NUMERIC(12,2) DEFAULT 0,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    account_id INTEGER REFERENCES accounts(id),
    type VARCHAR(20),
    amount NUMERIC(12,2),
    created_at TIMESTAMP DEFAULT NOW()
);

-- ===========================================
-- Índices para performance
-- ===========================================

CREATE INDEX idx_transactions_account_id ON transactions(account_id);
CREATE INDEX idx_accounts_owner_id ON accounts(owner_id);

-- ===========================================
-- Usuários de exemplo
-- ===========================================

INSERT INTO users (email, password_hash)
VALUES
  ('alice@example.com', '$2a$10$HvMw7qVuVggtSRPKTUZti.upl0F6P9C6Z9ngQZgps/kgDqmcGzDGG'), -- senha: 123456
  ('bob@example.com',   '$2a$10$HvMw7qVuVggtSRPKTUZti.upl0F6P9C6Z9ngQZgps/kgDqmcGzDGG'); -- senha: 123456 (mesma hash)

-- ===========================================
-- Contas de exemplo
-- ===========================================

INSERT INTO accounts (owner_id, account_number, balance)
VALUES
  (1, 'ACC1001', 1500.00),
  (2, 'ACC2001', 800.00);

-- ===========================================
-- Transações de exemplo
-- ===========================================

INSERT INTO transactions (account_id, type, amount)
VALUES
  (1, 'DEPOSIT', 1500.00),
  (2, 'DEPOSIT', 800.00);

-- ===========================================
-- Transferência inicial entre contas
-- ===========================================

-- Debita 100 da conta 1
INSERT INTO transactions (account_id, type, amount) VALUES (1, 'TRANSFER_OUT', 100.00);
UPDATE accounts SET balance = balance - 100 WHERE id = 1;

-- Credita 100 na conta 2
INSERT INTO transactions (account_id, type, amount) VALUES (2, 'TRANSFER_IN', 100.00);
UPDATE accounts SET balance = balance + 100 WHERE id = 2;

-- Final
