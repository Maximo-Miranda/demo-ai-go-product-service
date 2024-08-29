SELECT 'CREATE DATABASE product_db'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'product_db')\gexec