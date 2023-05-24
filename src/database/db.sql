create TABLE IF NOT EXISTS User (
    ID_USER INTEGER PRIMARY KEY,
    USERNAME TEXT,
    USER_PASSWORD TEXT
);

create TABLE IF NOT EXISTS Post (
    ID_POST INTEGER PRIMARY KEY,
    TIMESTAMP TEXT,
    CONTENT TEXT,
    ID_USER INTEGER,
    FOREIGN KEY(ID_USER) REFERENCES User(ID_USER)
);

-- Development
insert into User VALUES(1, "Tom", "Kitten");
insert into User VALUES(2, "Tyv", "dog");
insert into User VALUES(3, "Gurvan", "123");
insert into User VALUES(4, "Banos", "azerty");

insert into Post VALUES(1, "2023-05-24-9-37-10", "I love hotdogs", 2);
