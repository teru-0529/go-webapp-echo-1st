-- is_master_table=false

-- 5.出荷指示(shipping_instructions)

-- Create Table
DROP TABLE IF EXISTS orders.shipping_instructions CASCADE;
CREATE TABLE orders.shipping_instructions (
  sipping_no serial NOT NULL,
  order_no varchar(10) NOT NULL check (order_no ~* '^RO-[0-9]{7}$'),
  product_id varchar(5) NOT NULL check (product_id ~* '^P[0-9]{4}$'),
  sipping_date date NOT NULL DEFAULT current_timestamp,
  operator_name varchar(30) NOT NULL check (LENGTH(operator_name) >= 3),
  shipping_quantity integer NOT NULL DEFAULT 0 check (0 <= shipping_quantity AND shipping_quantity <= 1000),
  created_at timestamp NOT NULL DEFAULT current_timestamp,
  updated_at timestamp NOT NULL DEFAULT current_timestamp,
  created_by varchar(58),
  updated_by varchar(58)
);

-- Set Table Comment
COMMENT ON TABLE orders.shipping_instructions IS '出荷指示';

-- Set Column Comment
COMMENT ON COLUMN orders.shipping_instructions.sipping_no IS '出荷NO';
COMMENT ON COLUMN orders.shipping_instructions.order_no IS '受注番号';
COMMENT ON COLUMN orders.shipping_instructions.product_id IS '商品ID';
COMMENT ON COLUMN orders.shipping_instructions.sipping_date IS '出荷日';
COMMENT ON COLUMN orders.shipping_instructions.operator_name IS '処理担当者名';
COMMENT ON COLUMN orders.shipping_instructions.shipping_quantity IS '出荷数';
COMMENT ON COLUMN orders.shipping_instructions.created_at IS '作成日時';
COMMENT ON COLUMN orders.shipping_instructions.updated_at IS '更新日時';
COMMENT ON COLUMN orders.shipping_instructions.created_by IS '作成者';
COMMENT ON COLUMN orders.shipping_instructions.updated_by IS '更新者';

-- Set PK Constraint
ALTER TABLE orders.shipping_instructions ADD PRIMARY KEY (
  sipping_no
);

-- Create 'set_update_at' Trigger
CREATE TRIGGER set_updated_at
  BEFORE UPDATE
  ON orders.shipping_instructions
  FOR EACH ROW
EXECUTE PROCEDURE set_updated_at();

-- Create 'append_history' Function
DROP FUNCTION IF EXISTS orders.shipping_instructions_audit();
CREATE OR REPLACE FUNCTION orders.shipping_instructions_audit() RETURNS TRIGGER AS $$
BEGIN
  IF (TG_OP = 'DELETE') THEN
    INSERT INTO operation_histories(schema_name, table_name, operation_type, table_key)
    SELECT TG_TABLE_SCHEMA, TG_TABLE_NAME, 'DELETE', OLD.sipping_no;
  ELSIF (TG_OP = 'UPDATE') THEN
    INSERT INTO operation_histories(operated_by, schema_name, table_name, operation_type, table_key)
    SELECT NEW.updated_by, TG_TABLE_SCHEMA, TG_TABLE_NAME, 'UPDATE', NEW.sipping_no;
  ELSIF (TG_OP = 'INSERT') THEN
    INSERT INTO operation_histories(operated_by, schema_name, table_name, operation_type, table_key)
    SELECT NEW.updated_by, TG_TABLE_SCHEMA, TG_TABLE_NAME, 'INSERT', NEW.sipping_no;
  END IF;
  RETURN null;
END;
$$ LANGUAGE plpgsql;

-- Create 'audit' Trigger
CREATE TRIGGER audit
  AFTER INSERT OR UPDATE OR DELETE
  ON orders.shipping_instructions
  FOR EACH ROW
EXECUTE PROCEDURE orders.shipping_instructions_audit();
