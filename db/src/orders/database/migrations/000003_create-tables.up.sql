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


-- ----+----+----+----+----+----+----+


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


-- ----+----+----+----+----+----+----+


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
  order_status orders.order_status NOT NULL DEFAULT 'WORK_IN_PROGRESS',
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


-- ----+----+----+----+----+----+----+


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


-- ----+----+----+----+----+----+----+


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


-- ----+----+----+----+----+----+----+
