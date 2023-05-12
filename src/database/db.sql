create TABLE IF NOT EXISTS user (
    firstname VARCHAR(50),
    lastname VARCHAR(50)
);

create TABLE IF NOT EXISTS post (
    title VARCHAR(255)
);

INSERT INTO user VALUES ("tom", "jegou");