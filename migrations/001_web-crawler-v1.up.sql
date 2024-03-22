DROP TABLE IF EXISTS img;

CREATE TABLE img {
    id SERIAL PRIMARY KEY,
    img_urls TEXT[],
    img_descriptions TEXT[]
};

DROP TABLE IF EXISTS extracted_data;

CREATE TABLE extracted_data (
    id SERIAL PRIMARY KEY,
    crawled_url VARCHAR(255),
    title VARCHAR(255),
    content TEXT,
    related_urls TEXT[],
    img_id INT REFERENCES img(id),
    crawl_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

