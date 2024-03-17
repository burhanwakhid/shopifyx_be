CREATE TABLE bank (
  id SERIAL PRIMARY KEY,
  name VARCHAR NOT NULL,  -- Bank name with min length 5 and max length 15
  account_name VARCHAR NOT NULL,  -- Bank account name with min length 5 and max length 15
  account_number VARCHAR NOT NULL  -- Bank account number with min length 5 and max length 15
);