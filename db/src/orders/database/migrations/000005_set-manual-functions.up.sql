-- シーケンス
DROP SEQUENCE IF EXISTS orders.product_id_seed;
CREATE SEQUENCE orders.product_id_seed START 1;

DROP SEQUENCE IF EXISTS orders.order_no_seed;
CREATE SEQUENCE orders.order_no_seed START 1;


-- ----+----+----+----+----+----+----+


-- 商品:登録「前」処理
-- Create Function
CREATE OR REPLACE FUNCTION orders.products_pre_process() RETURNS TRIGGER AS $$
BEGIN
  -- 導出属性の算出:商品ID
  NEW.product_id:='P'||to_char(nextval('orders.product_id_seed'),'FM0000');
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create Trigger
CREATE TRIGGER pre_process
  BEFORE INSERT
  ON orders.products
  FOR EACH ROW
EXECUTE PROCEDURE orders.products_pre_process();


-- ----+----+----+----+----+----+----+


-- 受注:登録「前」処理
-- Create Function
CREATE OR REPLACE FUNCTION orders.receivings_pre_process() RETURNS TRIGGER AS $$
BEGIN
  -- 導出属性の算出:受注番号
  NEW.order_no:='RO-'||to_char(nextval('orders.order_no_seed'),'FM0000000');
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create Trigger
CREATE TRIGGER pre_process
  BEFORE INSERT
  ON orders.receivings
  FOR EACH ROW
EXECUTE PROCEDURE orders.receivings_pre_process();


-- ----+----+----+----+----+----+----+


-- 受注明細:チェック制約(受注数が正の数)

-- Create Constraint
ALTER TABLE orders.receiving_details DROP CONSTRAINT IF EXISTS receiving_quantity_positive_check;
ALTER TABLE orders.receiving_details ADD CONSTRAINT receiving_quantity_positive_check CHECK (
  0 < receiving_quantity
);


-- ----+----+----+----+----+----+----+


-- 受注明細:チェック制約(販売単価が正の数)

-- Create Constraint
ALTER TABLE orders.receiving_details DROP CONSTRAINT IF EXISTS sellling_price_positive_check;
ALTER TABLE orders.receiving_details ADD CONSTRAINT sellling_price_positive_check CHECK (
  0 < sellling_price
);


-- ----+----+----+----+----+----+----+


-- 受注明細:登録「前」処理
-- Create Function
CREATE OR REPLACE FUNCTION orders.receiving_details_pre_process() RETURNS TRIGGER AS $$
BEGIN
  -- (INSERT時のみ)
  IF (TG_OP = 'INSERT') THEN
    -- 導出属性の算出:商品原価
    NEW.cost_price:= (SELECT cost_price FROM orders.products WHERE product_id = NEW.product_id);
    -- 導出属性の算出:利益率
    IF (NEW.sellling_price < NEW.cost_price) THEN
      NEW.profit_rate:= 0.00;
    ELSE
      NEW.profit_rate:= ROUND(1.0 * (NEW.sellling_price - NEW.cost_price) / NEW.sellling_price, 2);
    END IF;
  END IF;
  -- (INSERT時のみ)

  -- 導出属性の算出:受注残数
  NEW.remaining_quantity:= NEW.receiving_quantity - NEW.shipping_quantity - NEW.cancel_quantity;
  -- 導出属性の算出:受注ステータス
  IF NEW.remaining_quantity != 0 THEN
    NEW.order_status:= 'WORK_IN_PROGRESS';
  ELSIF NEW.shipping_quantity = 0 THEN
    NEW.order_status:= 'CANCELED';
  ELSE
    NEW.order_status:= 'COMPLETED';
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create Trigger
CREATE TRIGGER pre_process
  BEFORE INSERT OR UPDATE
  ON orders.receiving_details
  FOR EACH ROW
EXECUTE PROCEDURE orders.receiving_details_pre_process();


-- ----+----+----+----+----+----+----+


-- キャンセル指示:登録「後」処理
-- Create Function
CREATE OR REPLACE FUNCTION orders.cancel_instructions_post_process() RETURNS TRIGGER AS $$
BEGIN
  -- 「受注明細」:キャンセル数更新
  UPDATE orders.receiving_details
  SET cancel_quantity = cancel_quantity + NEW.cancel_quantity,
      updated_by = NEW.created_by
  WHERE order_no = NEW.order_no
    AND product_id = NEW.product_id;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create Trigger
CREATE TRIGGER post_process
  AFTER INSERT
  ON orders.cancel_instructions
  FOR EACH ROW
EXECUTE PROCEDURE orders.cancel_instructions_post_process();


-- ----+----+----+----+----+----+----+


-- 出荷指示:登録「後」処理
-- Create Function
CREATE OR REPLACE FUNCTION orders.shipping_instructions_post_process() RETURNS TRIGGER AS $$
BEGIN
  -- 「受注明細」:出荷数更新
  UPDATE orders.receiving_details
  SET shipping_quantity = shipping_quantity + NEW.shipping_quantity,
      updated_by = NEW.created_by
  WHERE order_no = NEW.order_no
    AND product_id = NEW.product_id;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create Trigger
CREATE TRIGGER post_process
  AFTER INSERT
  ON orders.shipping_instructions
  FOR EACH ROW
EXECUTE PROCEDURE orders.shipping_instructions_post_process();


-- ----+----+----+----+----+----+----+
