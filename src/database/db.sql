CREATE TABLE IF NOT EXISTS User (
    id_user INTEGER NOT NULL PRIMARY KEY DEFAULT 0,
    email TEXT,
    username TEXT,
    date_of_birth TEXT,
    password_hash TEXT
);

CREATE TABLE IF NOT EXISTS Post (
    id_post INTEGER NOT NULL PRIMARY KEY DEFAULT 0,
    title TEXT,
    date_post TEXT,
    content TEXT,
    id_user INTEGER,
    FOREIGN KEY(id_user) REFERENCES User(id_user)
);

CREATE TABLE IF NOT EXISTS Comment (
    id_comment INTEGER NOT NULL PRIMARY KEY DEFAULT 0,
    content TEXT,
    date_content TEXT,
    id_post INTEGER,
    id_user INTEGER,
    FOREIGN KEY(id_post) REFERENCES Post(id_post),
    FOREIGN KEY(id_user) REFERENCES User(id_user)
);

CREATE TABLE IF NOT EXISTS "Like" (
    id_post INTEGER,
    id_user INTEGER,
    CONSTRAINT PK_Like PRIMARY KEY(id_post, id_user),
    FOREIGN KEY(id_post) REFERENCES Post(id_post),
    FOREIGN KEY(id_user) REFERENCES User(id_user)
)