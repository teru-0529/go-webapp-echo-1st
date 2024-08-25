-- 受注ステータス
CREATE TYPE orders.order_status AS enum (
  'WORK_IN_PROGRESS',
  'CANCELED',
  'COMPLETED',
  'PREPARING'
);
