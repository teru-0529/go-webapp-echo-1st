ALTER TABLE orders.receiving_details DROP CONSTRAINT IF EXISTS receiving_details_foreignKey_1;
ALTER TABLE orders.receiving_details DROP CONSTRAINT IF EXISTS receiving_details_foreignKey_2;

ALTER TABLE orders.cancel_instructions DROP CONSTRAINT IF EXISTS cancel_instructions_foreignKey_1;

ALTER TABLE orders.shipping_instructions DROP CONSTRAINT IF EXISTS shipping_instructions_foreignKey_1;
