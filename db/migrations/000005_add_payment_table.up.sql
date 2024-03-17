CREATE TABLE payment (
  id SERIAL PRIMARY KEY,
  product_id INTEGER NOT NULL REFERENCES product(id),  -- Foreign key referencing product.id
  CONSTRAINT fk_payment_product FOREIGN KEY (product_id) REFERENCES product(id) ON DELETE CASCADE,
  user_id INTEGER NOT NULL REFERENCES users(id),        -- Foreign key referencing user.id
  CONSTRAINT fk_payment_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  bank_id INTEGER REFERENCES bank(id),                 -- Foreign key referencing bank.id (optional)
  CONSTRAINT fk_payment_bank FOREIGN KEY (bank_id) REFERENCES bank(id) ON DELETE SET NULL,
  bank_account_id VARCHAR(255) NOT NULL,               -- Bank account ID
  payment_proof_imageUrl VARCHAR(255) NOT NULL,  -- Payment proof image URL with validation
  quantity INTEGER NOT NULL,      -- Quantity purchased (min 1)
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP -- Payment creation timestamp
);