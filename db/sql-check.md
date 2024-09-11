# チェック用登録SQL

``` SQL
-- 商品
INSERT INTO orders.products VALUES (default,'日本刀',20000,default,default,'100001-P0673822','100001-P0673822');
INSERT INTO orders.products VALUES (default,'火縄銃',40000,default,default,'100001-P0673822','100001-P0673822');
INSERT INTO orders.products VALUES (default,'弓',15000,default,default,'100001-P0673822','100001-P0673822');

-- 受注
INSERT INTO orders.receivings VALUES (default,default,'徳川家康','織田物産',default,default,default,default,default,'100002-P0673822','100002-P0673822');

-- 受注明細
INSERT INTO orders.receiving_details VALUES ('RO-0000001','P0001',5,default,default,default,34800,default,default,default,default,default,'100002-P0673822','100002-P0673822');
INSERT INTO orders.receiving_details VALUES ('RO-0000001','P0002',2,default,default,default,106400,default,default,default,default,default,'100002-P0673822','100002-P0673822');

-- キャンセル指示
INSERT INTO orders.cancel_instructions VALUES (default,'RO-0000001','P0001',default,'徳川綱吉',2,'受注数登録ミス',default,default,'100003-P0673822','100003-P0673822');
INSERT INTO orders.cancel_instructions VALUES (default,'RO-0000001','P0002',default,'徳川吉宗',1,'顧客都合',default,default,'100004-P0673822','100004-P0673822');
INSERT INTO orders.cancel_instructions VALUES (default,'RO-0000001','P0002',default,'徳川吉宗',1,'顧客都合',default,default,'100005-P0673822','100005-P0673822');

-- 出荷指示
INSERT INTO orders.shipping_instructions VALUES (default,'RO-0000001','P0001',default,'徳川慶喜',3,default,default,'100005-P0673822','100005-P0673822');
```
