-- operation_afert_create_tables

-- 4.キャンセル指示(cancel_instructions)

-- Set FK Constraint
ALTER TABLE orders.cancel_instructions DROP CONSTRAINT IF EXISTS cancel_instructions_foreignKey_1;
ALTER TABLE orders.cancel_instructions ADD CONSTRAINT cancel_instructions_foreignKey_1 FOREIGN KEY (
  order_no,
  product_id
) REFERENCES orders.receiving_details (
  order_no,
  product_id
) ON DELETE CASCADE ON UPDATE CASCADE;
