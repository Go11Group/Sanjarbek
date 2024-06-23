CREATE TYPE transaction_type AS ENUM ('credit', 'debit');

CREATE TABLE transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    card_id UUID NOT NULL REFERENCES card(id),
    amount DECIMAL(10, 2) NOT NULL,
    terminal_id UUID DEFAULT NULL,
    transaction_type transaction_type NOT NULL
);