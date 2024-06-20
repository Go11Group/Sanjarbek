CREATE TABLE made_in (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    country VARCHAR(255)
);

INSERT INTO made_in (id, country)
VALUES ('1b51c4b1-28af-4b0c-b8ab-94f892265c3c', 'Japan'),
    ('298cd9fd-2491-4da1-902e-3c7be16f7f5e', 'USA');