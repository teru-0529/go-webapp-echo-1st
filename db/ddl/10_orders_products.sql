-- is_master_table=false

-- 3.商品(products)

-- Create Table
DROP TABLE IF EXISTS orders.products CASCADE;
CREATE TABLE orders.products (
  product_id varchar(5) NOT NULL check (product_id ~* '^P[0-9]{4}$'),
  product_name varchar(30) NOT NULL,
  cost_price integer NOT NULL DEFAULT 0 check (0 <= cost_price AND cost_price <= 9999999),
  created_at timestamp NOT NULL DEFAULT current_timestamp,
  updated_at timestamp NOT NULL DEFAULT current_timestamp,
  created_by varchar(58),
  updated_by varchar(58)
);

-- Set Table Comment
COMMENT ON TABLE orders.products IS '商品';

-- Set Column Comment
COMMENT ON COLUMN orders.products.product_id IS '商品ID';
COMMENT ON COLUMN orders.products.product_name IS '商品名';
COMMENT ON COLUMN orders.products.cost_price IS '商品原価';
COMMENT ON COLUMN orders.products.created_at IS '作成日時';
COMMENT ON COLUMN orders.products.updated_at IS '更新日時';
COMMENT ON COLUMN orders.products.created_by IS '作成者';
COMMENT ON COLUMN orders.products.updated_by IS '更新者';

-- Set PK Constraint
ALTER TABLE orders.products ADD PRIMARY KEY (
  product_id
);

-- Create 'set_update_at' Trigger
CREATE TRIGGER set_updated_at
  BEFORE UPDATE
  ON orders.products
  FOR EACH ROW
EXECUTE PROCEDURE set_updated_at();

-- Create 'append_history' Function
DROP FUNCTION IF EXISTS orders.products_audit();
CREATE OR REPLACE FUNCTION orders.products_audit() RETURNS TRIGGER AS $$
BEGIN
  IF (TG_OP = 'DELETE') THEN
    INSERT INTO operation_histories(schema_name, table_name, operation_type, table_key)
    SELECT TG_TABLE_SCHEMA, TG_TABLE_NAME, 'DELETE', OLD.product_id;
  ELSIF (TG_OP = 'UPDATE') THEN
    INSERT INTO operation_histories(operated_by, schema_name, table_name, operation_type, table_key)
    SELECT NEW.updated_by, TG_TABLE_SCHEMA, TG_TABLE_NAME, 'UPDATE', NEW.product_id;
  ELSIF (TG_OP = 'INSERT') THEN
    INSERT INTO operation_histories(operated_by, schema_name, table_name, operation_type, table_key)
    SELECT NEW.updated_by, TG_TABLE_SCHEMA, TG_TABLE_NAME, 'INSERT', NEW.product_id;
  END IF;
  RETURN null;
END;
$$ LANGUAGE plpgsql;

-- Create 'audit' Trigger
CREATE TRIGGER audit
  AFTER INSERT OR UPDATE OR DELETE
  ON orders.products
  FOR EACH ROW
EXECUTE PROCEDURE orders.products_audit();
