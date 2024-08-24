-- 2.受注明細(receiving_details)

-- Set FK Constraint
ALTER TABLE orders.receiving_details DROP CONSTRAINT IF EXISTS receiving_details_foreignKey_1;
ALTER TABLE orders.receiving_details ADD CONSTRAINT receiving_details_foreignKey_1 FOREIGN KEY (
  order_no
) REFERENCES orders.receivings (
  order_no
) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE orders.receiving_details DROP CONSTRAINT IF EXISTS receiving_details_foreignKey_2;
ALTER TABLE orders.receiving_details ADD CONSTRAINT receiving_details_foreignKey_2 FOREIGN KEY (
  product_id
) REFERENCES orders.products (
  product_id
);


-- ----+----+----+----+----+----+----+


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


-- ----+----+----+----+----+----+----+


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


-- ----+----+----+----+----+----+----+
