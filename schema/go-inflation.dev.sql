CREATE TABLE IF NOT EXISTS products (
	id SERIAL PRIMARY KEY,
	ean VARCHAR(13) UNIQUE NOT NULL, -- ean can vary from 8 to 13 characters
	external_id VARCHAR(36), -- to support large alphanumerical values or uuids
	name VARCHAR(255) NOT NULL,
	description TEXT,
	brand VARCHAR(255) NOT NULL,
	image VARCHAR(255), -- intended to be an url pointing to an image
	type VARCHAR(255) NOT NULL,
	creation_time TIMESTAMP NOT NULL,
	save_time TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS products_price_history (
	id SERIAL PRIMARY KEY,
	product_id INTEGER,
	price DOUBLE PRECISION NOT NULL,
	discounted_price DOUBLE PRECISION,
	registered_at TIMESTAMP NOT NULL,
	save_time TIMESTAMP NOT NULL,
	CONSTRAINT fk_product_price_history_product FOREIGN KEY(product_id)
		REFERENCES product(id)
		ON DELETE CASCADE
);