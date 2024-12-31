package db_driver

import (
	"context"
	"database/sql"
	"feed_collector/domain/rss_feed"
	"feed_collector/slogger"
)

func Register(feed rss_feed.Feed, db *sql.DB, ctx context.Context) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		slogger.Logger.Error("Failed to begin a transaction.", "Caused by", err)
	}

	result, err := tx.Exec("INSERT INTO feeds (url, title, description, author, published_date) VALUES ($1, $2, $3, $4, $5)",
		feed.URL, feed.Title, feed.Description, feed.Author, feed.PublishedDate)
	if err != nil {
		slogger.Logger.Error("Failed to execute the query.", "Caused by", err)
		tx.Rollback()
	}

	err = tx.Commit()
	if err != nil {
		slogger.Logger.Error("Failed to commit the transaction.", "Caused by", err)
		tx.Rollback()
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		slogger.Logger.Error("Failed to get the number of rows affected.", "Caused by", err)
		tx.Rollback()
	}

	slogger.Logger.Info("Successfully registered the feed.", "Rows affected", rowsAffected)

	return nil
}
