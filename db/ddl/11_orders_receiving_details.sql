-- operation_afert_create_tables

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
