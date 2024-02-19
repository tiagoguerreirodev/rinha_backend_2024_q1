-- Coloque scripts iniciais aqui
CREATE TABLE IF NOT EXISTS clientes (
    id SMALLSERIAL UNIQUE PRIMARY KEY,
    limite integer NOT NULL,
    saldo integer DEFAULT 0
);

CREATE TABLE IF NOT EXISTS transacoes (
    user_id smallint,
    valor integer NOT NULL,
    tipo char(1) NOT NULL,
    descricao varchar(10) NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_user_id ON transacoes (user_id);
CREATE INDEX idx_id ON clientes (id);
CREATE INDEX desc_index ON transacoes (created_at DESC);

DO $$
BEGIN
  INSERT INTO clientes (limite)
  VALUES
    (1000 * 100),
    (800 * 100),
    (10000 * 100),
    (100000 * 100),
    (5000 * 100);
END; $$