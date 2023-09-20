CREATE TABLE IF NOT EXISTS todos (
    id SERIAL PRIMARY KEY,
    text varchar(200) NOT NULL,
    done BOOLEAN DEFAULT false
);