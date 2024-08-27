-- is_master_table=false

-- 2.受注明細(receiving_details)

-- Create Table
DROP TABLE IF EXISTS orders.receiving_details CASCADE;
CREATE TABLE orders.receiving_details (
  order_no varchar(10) NOT NULL check (order_no ~* '^RO-[0-9]{7}$'),
  product_id varchar(5) NOT NULL check (product_id ~* '^P[0-9]{4}$'),
  receiving_quantity integer NOT NULL DEFAULT 0 check (0 <= receiving_quantity AND receiving_quantity <= 1000),
  shipping_quantity integer NOT NULL DEFAULT 0 check (0 <= shipping_quantity AND shipping_quantity <= 1000),
  cancel_quantity integer NOT NULL DEFAULT 0 check (0 <= cancel_quantity AND cancel_quantity <= 1000),
  remaining_quantity integer NOT NULL DEFAULT 0 check (0 <= remaining_quantity AND remaining_quantity <= 1000),
  sellling_price integer NOT NULL DEFAULT 0 check (0 <= sellling_price AND sellling_price <= 9999999),
  cost_price integer NOT NULL DEFAULT 0 check (0 <= cost_price AND cost_price <= 9999999),
  profit_rate numeric NOT NULL DEFAULT 0.00 check (0 <= profit_rate AND profit_rate <= 1),
  order_status orders.order_status NOT NULL DEFAULT 'PREPARING',
  created_at timestamp NOT NULL DEFAULT current_timestamp,
  updated_at timestamp NOT NULL DEFAULT current_timestamp,
  created_by varchar(58),
  updated_by varchar(58)
);

-- Set Table Comment
COMMENT ON TABLE orders.receiving_details IS '受注明細';

-- Set Column Comment
COMMENT ON COLUMN orders.receiving_details.order_no IS '受注番号';
COMMENT ON COLUMN orders.receiving_details.product_id IS '商品ID';
COMMENT ON COLUMN orders.receiving_details.receiving_quantity IS '受注数';
COMMENT ON COLUMN orders.receiving_details.shipping_quantity IS '出荷数';
COMMENT ON COLUMN orders.receiving_details.cancel_quantity IS 'キャンセル数';
COMMENT ON COLUMN orders.receiving_details.remaining_quantity IS '受注残数';
COMMENT ON COLUMN orders.receiving_details.sellling_price IS '販売単価';
COMMENT ON COLUMN orders.receiving_details.cost_price IS '商品原価';
COMMENT ON COLUMN orders.receiving_details.profit_rate IS '利益率';
COMMENT ON COLUMN orders.receiving_details.order_status IS '受注ステータス';
COMMENT ON COLUMN orders.receiving_details.created_at IS '作成日時';
COMMENT ON COLUMN orders.receiving_details.updated_at IS '更新日時';
COMMENT ON COLUMN orders.receiving_details.created_by IS '作成者';
COMMENT ON COLUMN orders.receiving_details.updated_by IS '更新者';

-- Set PK Constraint
ALTER TABLE orders.receiving_details ADD PRIMARY KEY (
  order_no,
  product_id
);

-- Create 'set_update_at' Trigger
CREATE TRIGGER set_updated_at
  BEFORE UPDATE
  ON orders.receiving_details
  FOR EACH ROW
EXECUTE PROCEDURE set_updated_at();

-- Create 'append_history' Function
DROP FUNCTION IF EXISTS orders.receiving_details_audit();
CREATE OR REPLACE FUNCTION orders.receiving_details_audit() RETURNS TRIGGER AS $$
BEGIN
  IF (TG_OP = 'DELETE') THEN
    INSERT INTO operation_histories(schema_name, table_name, operation_type, table_key)
    SELECT TG_TABLE_SCHEMA, TG_TABLE_NAME, 'DELETE', OLD.order_no || '-' || OLD.product_id;
  ELSIF (TG_OP = 'UPDATE') THEN
    INSERT INTO operation_histories(operated_by, schema_name, table_name, operation_type, table_key)
    SELECT NEW.updated_by, TG_TABLE_SCHEMA, TG_TABLE_NAME, 'UPDATE', NEW.order_no || '-' || NEW.product_id;
  ELSIF (TG_OP = 'INSERT') THEN
    INSERT INTO operation_histories(operated_by, schema_name, table_name, operation_type, table_key)
    SELECT NEW.updated_by, TG_TABLE_SCHEMA, TG_TABLE_NAME, 'INSERT', NEW.order_no || '-' || NEW.product_id;
  END IF;
  RETURN null;
END;
$$ LANGUAGE plpgsql;

-- Create 'audit' Trigger
CREATE TRIGGER audit
  AFTER INSERT OR UPDATE OR DELETE
  ON orders.receiving_details
  FOR EACH ROW
EXECUTE PROCEDURE orders.receiving_details_audit();
