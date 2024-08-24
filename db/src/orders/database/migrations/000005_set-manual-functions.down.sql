-- 出荷指示:登録「後」処理
DROP FUNCTION IF EXISTS orders.shipping_instructions_post_process() CASCADE;

-- キャンセル指示:登録「後」処理
DROP FUNCTION IF EXISTS orders.cancel_instructions_post_process() CASCADE;

-- 受注明細:登録「前」処理
DROP FUNCTION IF EXISTS orders.receiving_details_pre_process() CASCADE;

-- 受注明細:チェック制約
ALTER TABLE orders.receiving_details DROP CONSTRAINT IF EXISTS sellling_price_positive_check;
ALTER TABLE orders.receiving_details DROP CONSTRAINT IF EXISTS receiving_quantity_positive_check;

-- 受注:登録「前」処理
DROP FUNCTION IF EXISTS orders.receivings_pre_process() CASCADE;

-- 商品:登録「前」処理
DROP FUNCTION IF EXISTS orders.products_pre_process() CASCADE;

-- シーケンス
DROP SEQUENCE IF EXISTS orders.product_id_seed;
DROP SEQUENCE IF EXISTS orders.order_no_seed;
