CREATE VIRTUAL TABLE IF NOT EXISTS pages USING fts5(title, article, version);

