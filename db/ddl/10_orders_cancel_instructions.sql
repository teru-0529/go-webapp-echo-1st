-- is_master_table=false

-- 4.キャンセル指示(cancel_instructions)

-- Create Table
DROP TABLE IF EXISTS orders.cancel_instructions CASCADE;
CREATE TABLE orders.cancel_instructions (
  cancel_no serial NOT NULL,
  order_no varchar(10) NOT NULL check (order_no ~* '^RO-[0-9]{7}$'),
  product_id varchar(5) NOT NULL check (product_id ~* '^P[0-9]{4}$'),
  cancel_date date NOT NULL DEFAULT current_timestamp,
  operator_name varchar(30) NOT NULL check (LENGTH(operator_name) >= 3),
  cancel_quantity integer NOT NULL DEFAULT 0 check (0 <= cancel_quantity AND cancel_quantity <= 1000),
  cancel_reason text,
  created_at timestamp NOT NULL DEFAULT current_timestamp,
  updated_at timestamp NOT NULL DEFAULT current_timestamp,
  created_by varchar(58),
  updated_by varchar(58)
);

-- Set Table Comment
COMMENT ON TABLE orders.cancel_instructions IS 'キャンセル指示';

-- Set Column Comment
COMMENT ON COLUMN orders.cancel_instructions.cancel_no IS 'キャンセルNO';
COMMENT ON COLUMN orders.cancel_instructions.order_no IS '受注番号';
COMMENT ON COLUMN orders.cancel_instructions.product_id IS '商品ID';
COMMENT ON COLUMN orders.cancel_instructions.cancel_date IS 'キャンセル日';
COMMENT ON COLUMN orders.cancel_instructions.operator_name IS '処理担当者名';
COMMENT ON COLUMN orders.cancel_instructions.cancel_quantity IS 'キャンセル数';
COMMENT ON COLUMN orders.cancel_instructions.cancel_reason IS 'キャンセル理由';
COMMENT ON COLUMN orders.cancel_instructions.created_at IS '作成日時';
COMMENT ON COLUMN orders.cancel_instructions.updated_at IS '更新日時';
COMMENT ON COLUMN orders.cancel_instructions.created_by IS '作成者';
COMMENT ON COLUMN orders.cancel_instructions.updated_by IS '更新者';

-- Set PK Constraint
ALTER TABLE orders.cancel_instructions ADD PRIMARY KEY (
  cancel_no
);

-- Create 'set_update_at' Trigger
CREATE TRIGGER set_updated_at
  BEFORE UPDATE
  ON orders.cancel_instructions
  FOR EACH ROW
EXECUTE PROCEDURE set_updated_at();

-- Create 'append_history' Function
DROP FUNCTION IF EXISTS orders.cancel_instructions_audit();
CREATE OR REPLACE FUNCTION orders.cancel_instructions_audit() RETURNS TRIGGER AS $$
BEGIN
  IF (TG_OP = 'DELETE') THEN
    INSERT INTO operation_histories(schema_name, table_name, operation_type, table_key)
    SELECT TG_TABLE_SCHEMA, TG_TABLE_NAME, 'DELETE', OLD.cancel_no;
  ELSIF (TG_OP = 'UPDATE') THEN
    INSERT INTO operation_histories(operated_by, schema_name, table_name, operation_type, table_key)
    SELECT NEW.updated_by, TG_TABLE_SCHEMA, TG_TABLE_NAME, 'UPDATE', NEW.cancel_no;
  ELSIF (TG_OP = 'INSERT') THEN
    INSERT INTO operation_histories(operated_by, schema_name, table_name, operation_type, table_key)
    SELECT NEW.updated_by, TG_TABLE_SCHEMA, TG_TABLE_NAME, 'INSERT', NEW.cancel_no;
  END IF;
  RETURN null;
END;
$$ LANGUAGE plpgsql;

-- Create 'audit' Trigger
CREATE TRIGGER audit
  AFTER INSERT OR UPDATE OR DELETE
  ON orders.cancel_instructions
  FOR EACH ROW
EXECUTE PROCEDURE orders.cancel_instructions_audit();
