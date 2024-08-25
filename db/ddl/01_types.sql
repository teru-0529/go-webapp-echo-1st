-- Enum Type DDL

-- 受注ステータス
DROP TYPE IF EXISTS orders.order_status;
CREATE TYPE orders.order_status AS enum (
  'WORK_IN_PROGRESS',
  'CANCELED',
  'COMPLETED',
  'PREPARING'
);

