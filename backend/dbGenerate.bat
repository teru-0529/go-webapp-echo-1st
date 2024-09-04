@echo off

sqlboiler psql -c .config/orders.yaml
sqlboiler psql -c .config/public.yaml
