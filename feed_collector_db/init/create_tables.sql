CREATE DATABASE IF NOT EXISTS feed_collector_db;

\c feed_collector_db

CREATE TABLE IF NOT EXISTS rss_feed_links (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    url VARCHAR(128) NOT NULL UNIQUE,
);

CREATE INDEX IF NOT EXISTS idx_rss_feed_links_url ON rss_feed_links(url);

CREATE TABLE IF NOT EXISTS feeds (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    rss_feed_link_id INTEGER NOT NULL,
    author TEXT NOT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    url TEXT NOT NULL,
    published_date TIMESTAMP NOT NULL,
    last_updated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(rss_feed_link_id, url),
    FOREIGN KEY (rss_feed_link_id) REFERENCES rss_feed_links(id) ON DELETE CASCADE
);

CREATE OR REPLACE FUNCTION update_search_vector() RETURNS trigger AS $$
BEGIN
    NEW.search_vector :=
        to_tsvector('simple', coalesce(NEW.title, '')) ||
        to_tsvector('simple', coalesce(NEW.description, ''));
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_update_search_vector
BEFORE INSERT OR UPDATE ON feeds
FOR EACH ROW
EXECUTE FUNCTION update_search_vector();

CREATE INDEX IF NOT EXISTS idx_feeds_search_vector ON feeds USING GIN(search_vector);
CREATE INDEX IF NOT EXISTS idx_feeds_rss_feed_link_id ON feeds(rss_feed_link_id);
CREATE INDEX IF NOT EXISTS idx_feeds_last_updated ON feeds(last_updated);


