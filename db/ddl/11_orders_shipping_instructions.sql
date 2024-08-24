-- operation_afert_create_tables

-- 5.出荷指示(shipping_instructions)

-- Set FK Constraint
ALTER TABLE orders.shipping_instructions DROP CONSTRAINT IF EXISTS shipping_instructions_foreignKey_1;
ALTER TABLE orders.shipping_instructions ADD CONSTRAINT shipping_instructions_foreignKey_1 FOREIGN KEY (
  order_no,
  product_id
) REFERENCES orders.receiving_details (
  order_no,
  product_id
) ON DELETE CASCADE ON UPDATE CASCADE;
