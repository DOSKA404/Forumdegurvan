create TABLE IF NOT EXISTS User (
    id_user INTEGER PRIMARY KEY DEFAULT 0,
    email TEXT,
    username TEXT,
    date_of_birth TEXT,
    password_hash TEXT
);

create TABLE IF NOT EXISTS Post (
    id_post INTEGER PRIMARY KEY DEFAULT 0,
    date_post TEXT,
    content TEXT,
    id_user INTEGER,
    FOREIGN KEY(id_user) REFERENCES User(id_user)
);
