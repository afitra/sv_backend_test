-- +migrate Up

CREATE TABLE IF NOT EXISTS posts (
                                     id INT AUTO_INCREMENT PRIMARY KEY,
                                     title VARCHAR(200) NOT NULL,
    content TEXT,
    category VARCHAR(100),
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    status VARCHAR(100)
    );

-- +migrate StatementEnd