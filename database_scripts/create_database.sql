CREATE TABLE IF NOT EXISTS pages(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    currentVersion INTEGER,
    FOREIGN KEY (currentVersion) REFERENCES versions(id)
);

CREATE TABLE IF NOT EXISTS versions(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    dateCreated INTEGER,
    pageID INTEGER,
    pageContentID INTEGER,
    FOREIGN KEY(pageID) REFERENCES pages(rowid)
);

CREATE VIRTUAL TABLE IF NOT EXISTS pageContent USING fts5(
    title,
    article
);

