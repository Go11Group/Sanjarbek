ALTER TABLE cars
ADD COLUMN made_in_id UUID REFERENCES made_in(id);
