CREATE TABLE products (
    id CHAR(36) PRIMARY KEY,
    name VARCHAR(255),
    description TEXT,
    price_currency VARCHAR(5),
    price_amount BIGINT,
    shop_id CHAR(36)
) ENGINE=InnoDB;
ALTER TABLE products
ADD CONSTRAINT FOREIGN KEY (shop_id) REFERENCES shops (id);
