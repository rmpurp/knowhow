CREATE TABLE IF NOT EXISTS pages(
    id INTEGER PRIMARY KEY AUTOINCREMENT
);


CREATE VIRTUAL TABLE IF NOT EXISTS pageContent USING fts5(
    title,
    article
);

CREATE TABLE IF NOT EXISTS versions(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    dateCreated INTEGER,
    pageID INTEGER,
    pageContentID INTEGER,
    isCurrentVersion BOOLEAN,
    FOREIGN KEY(pageID) REFERENCES pages(id),
    FOREIGN KEY(pageContentID) REFERENCES pageContent(rowid)
);

