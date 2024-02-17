-- Coloque scripts iniciais aqui
CREATE TABLE clientes (
    id SMALLSERIAL,
    nome varchar(50) NOT NULL,
    limite integer NOT NULL,
    saldo integer DEFAULT 0
);

CREATE TABLE transacoes (
    user_id smallint,
    valor integer NOT NULL,
    tipo char(1) NOT NULL,
    descricao varchar(10) NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_cliente
        FOREIGN KEY(user_id)
            REFERENCES clientes(id)

);

DO $$
BEGIN
  INSERT INTO clientes (nome, limite)
  VALUES
    ('o barato sai caro', 1000 * 100),
    ('zan corp ltda', 800 * 100),
    ('les cruders', 10000 * 100),
    ('padaria joia de cocaia', 100000 * 100),
    ('kid mais', 5000 * 100);
END; $$