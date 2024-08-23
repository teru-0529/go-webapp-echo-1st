-- is_master_table=false

-- 1.受注(receivings)

-- Create Table
DROP TABLE IF EXISTS orders.receivings CASCADE;
CREATE TABLE orders.receivings (
  order_no varchar(10) NOT NULL check (order_no ~* '^RO-[0-9]{7}$'),
  order_date date NOT NULL DEFAULT current_timestamp,
  operator_name varchar(30) NOT NULL check (LENGTH(operator_name) >= 3),
  customer_name varchar(50) NOT NULL,
  created_at timestamp NOT NULL DEFAULT current_timestamp,
  updated_at timestamp NOT NULL DEFAULT current_timestamp,
  created_by varchar(58),
  updated_by varchar(58)
);

-- Set Table Comment
COMMENT ON TABLE orders.receivings IS '受注';

-- Set Column Comment
COMMENT ON COLUMN orders.receivings.order_no IS '受注番号';
COMMENT ON COLUMN orders.receivings.order_date IS '受注日';
COMMENT ON COLUMN orders.receivings.operator_name IS '処理担当者名';
COMMENT ON COLUMN orders.receivings.customer_name IS '得意先名称';
COMMENT ON COLUMN orders.receivings.created_at IS '作成日時';
COMMENT ON COLUMN orders.receivings.updated_at IS '更新日時';
COMMENT ON COLUMN orders.receivings.created_by IS '作成者';
COMMENT ON COLUMN orders.receivings.updated_by IS '更新者';

-- Set PK Constraint
ALTER TABLE orders.receivings ADD PRIMARY KEY (
  order_no
);

-- Create 'set_update_at' Trigger
CREATE TRIGGER set_updated_at
  BEFORE UPDATE
  ON orders.receivings
  FOR EACH ROW
EXECUTE PROCEDURE set_updated_at();

-- Create 'append_history' Function
DROP FUNCTION IF EXISTS orders.receivings_audit();
CREATE OR REPLACE FUNCTION orders.receivings_audit() RETURNS TRIGGER AS $$
BEGIN
  IF (TG_OP = 'DELETE') THEN
    INSERT INTO operation_histories(schema_name, table_name, operation_type, table_key)
    SELECT TG_TABLE_SCHEMA, TG_TABLE_NAME, 'DELETE', OLD.order_no;
  ELSIF (TG_OP = 'UPDATE') THEN
    INSERT INTO operation_histories(operated_by, schema_name, table_name, operation_type, table_key)
    SELECT NEW.updated_by, TG_TABLE_SCHEMA, TG_TABLE_NAME, 'UPDATE', NEW.order_no;
  ELSIF (TG_OP = 'INSERT') THEN
    INSERT INTO operation_histories(operated_by, schema_name, table_name, operation_type, table_key)
    SELECT NEW.updated_by, TG_TABLE_SCHEMA, TG_TABLE_NAME, 'INSERT', NEW.order_no;
  END IF;
  RETURN null;
END;
$$ LANGUAGE plpgsql;

-- Create 'audit' Trigger
CREATE TRIGGER audit
  AFTER INSERT OR UPDATE OR DELETE
  ON orders.receivings
  FOR EACH ROW
EXECUTE PROCEDURE orders.receivings_audit();
