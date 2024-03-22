DROP TABLE IF EXISTS extracted_data;

CREATE TABLE extracted_data (
    id SERIAL PRIMARY KEY,
    crawled_url VARCHAR(255),
    title VARCHAR(255),
    content TEXT,
    related_urls TEXT[],
    line_count INT,
    word_count INT,
    char_count BIGINT,
    average_word_length FLOAT,
    crawl_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS img;

CREATE TABLE img (
    id SERIAL PRIMARY KEY,
    img_urls TEXT[],
    img_descriptions TEXT[],
    extracted_data_id INT REFERENCES extracted_data(id)
);

DROP TABLE IF EXISTS word_frequency;

CREATE TABLE word_frequency (
    id SERIAL PRIMARY KEY,
    word VARCHAR(255) NOT NULL,
    frequency INT NOT NULL,
    extracted_data_id INT REFERENCES extracted_data(id)
);