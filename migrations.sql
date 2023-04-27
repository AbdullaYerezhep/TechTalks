DROP TABLE IF EXISTS `users`;

DROP TABLE IF EXISTS `session`;

DROP TABLE IF EXISTS `category`;

DROP TABLE IF EXISTS `post`;

DROP TABLE IF EXISTS `post_category`;

DROP TABLE IF EXISTS `comment`;

DROP TABLE IF EXISTS `like_dislike`;

CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    email VARCHAR UNIQUE NOT NULL,
    username VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL
);

CREATE TABLE session (
    id INTEGER PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES user(id),
    token VARCHAR UNIQUE NOT NULL,
    expiration_date DATETIME
);

CREATE TABLE category (
    id INTEGER PRIMARY KEY,
    name VARCHAR NOT NULL
);

CREATE TABLE post (
    id INTEGER PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES user(id),
    author VARCHAR NOT NULL,
    title VARCHAR NOT NULL,
    content TEXT NOT NULL,
    created DATETIME,
    updated DATETIME
);

CREATE TABLE post_category (
    post_id INTEGER NOT NULL REFERENCES post(id),
    category_id INTEGER NOT NULL REFERENCES category(id),
    PRIMARY KEY (post_id, category_id)
);

CREATE TABLE comment (
    id INTEGER PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES user(id),
    post_id INTEGER NOT NULL REFERENCES post(id),
    content TEXT NOT NULL,
    created DATETIME
    updated DATETIME
);

CREATE TABLE like_dislike (
    id INTEGER PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES user(id),
    post_id INTEGER REFERENCES post(id),
    comment_id INTEGER REFERENCES comment(id),
    islike INTEGER CHECK (islike IN (-1, 0, 1))
);