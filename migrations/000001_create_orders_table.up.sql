CREATE TABLE orders (
                        id SERIAL PRIMARY KEY,
                        customer_name VARCHAR(255) NOT NULL,
                        total_amount DECIMAL(10,2) NOT NULL,
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
