CREATE TABLE shops (
    id CHAR(36) PRIMARY KEY,
    name VARCHAR(255),
    description TEXT,
    latitud VARCHAR(100),
    longitude VARCHAR(100)
) ENGINE=InnoDB;
