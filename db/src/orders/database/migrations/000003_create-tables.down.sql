DROP TABLE IF EXISTS orders.shipping_instructions CASCADE;
DROP TABLE IF EXISTS orders.cancel_instructions CASCADE;
DROP TABLE IF EXISTS orders.receiving_details CASCADE;
DROP TABLE IF EXISTS orders.receivings CASCADE;
DROP TABLE IF EXISTS orders.products CASCADE;

DROP FUNCTION IF EXISTS orders.shipping_instructions_audit();
DROP FUNCTION IF EXISTS orders.cancel_instructions_audit();
DROP FUNCTION IF EXISTS orders.receiving_details_audit();
DROP FUNCTION IF EXISTS orders.receivings_audit();
DROP FUNCTION IF EXISTS orders.products_audit();
