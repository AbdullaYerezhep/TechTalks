PRAGMA foreign_keys = ON;
-- DROP TABLE IF EXISTS `users`;

-- DROP TABLE IF EXISTS `session`;

-- DROP TABLE IF EXISTS `category`; 

-- DROP TABLE IF EXISTS `post`;

-- DROP TABLE IF EXISTS `post_category`;

DROP TABLE IF EXISTS `comment`;

DROP TABLE IF EXISTS `post_rating`;

CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY,
    email VARCHAR UNIQUE NOT NULL,
    username VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS session (
    user_id INTEGER UNIQUE NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token VARCHAR UNIQUE NOT NULL,
    expiration_date DATETIME
);

CREATE TABLE IF NOT EXISTS category (
    id INTEGER PRIMARY KEY,
    name VARCHAR UNIQUE NOT NULL
);

INSERT OR IGNORE INTO category (name) 
SELECT 'Movies' UNION ALL
SELECT 'Games' UNION ALL
SELECT 'Anime' UNION ALL
SELECT 'Cartoons' UNION ALL
SELECT 'Books' UNION ALL
SELECT 'Comix' UNION ALL
SELECT 'Manga';

CREATE TABLE IF NOT EXISTS post (
    id INTEGER PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    author VARCHAR NOT NULL,
    title VARCHAR NOT NULL,
    content TEXT NOT NULL,
    created DATETIME,
    updated DATETIME
);

CREATE TABLE IF NOT EXISTS post_category (
    post_id INTEGER NOT NULL REFERENCES post(id) ON DELETE CASCADE,
    category_name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS comment (
    id INTEGER PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    post_id INTEGER NOT NULL REFERENCES post(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    created DATETIME
    updated DATETIME
);

CREATE TABLE IF NOT EXISTS post_rating (
    user_id INTEGER NOT NULL REFERENCES users(id),
    post_id INTEGER NOT NULL REFERENCES post(id) ON DELETE CASCADE,
    islike INTEGER CHECK (islike IN (-1, 1)),
    PRIMARY KEY (user_id, post_id)
);


CREATE TABLE IF NOT EXISTS comment_rating (
    comment_id INTEGER PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    islike INTEGER CHECK (islike IN (-1, 1))
);