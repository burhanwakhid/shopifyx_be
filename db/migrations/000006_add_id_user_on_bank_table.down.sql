-- Down migration to remove id_user column

ALTER TABLE bank
DROP CONSTRAINT fk_bank_user;

ALTER TABLE bank
DROP COLUMN id_user;