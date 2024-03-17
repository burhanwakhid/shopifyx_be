CREATE TABLE IF NOT EXISTS product (
  id SERIAL PRIMARY KEY,
  name VARCHAR(60) NOT NULL ,  -- Name with min length 5 and max length 60
  price DECIMAL(10,2) NOT NULL,        -- Price (decimal) with minimum 0
  imageUrl VARCHAR(255) NOT NULL,
  stock INTEGER NOT NULL,               -- Stock (integer) with minimum 0
  condition VARCHAR(10) CHECK (condition IN ('new', 'second')),  -- Condition with enum validation ("new" or "second")
  tags TEXT[] NOT NULL DEFAULT '{}',                          -- Tags as text array (allows multiple tags)
  owner_id INTEGER NOT NULL REFERENCES users(id),          -- Foreign key referencing users.id
  CONSTRAINT fk_product_owner FOREIGN KEY (owner_id) REFERENCES users(id) ON DELETE CASCADE
);