DROP TABLE IF EXISTS articles;

	CREATE TABLE articles( 
		id SERIAL NOT NULL UNIQUE PRIMARY KEY,
		author TEXT,
		title TEXT,
		description TEXT,
		url TEXT,
		urlToImage TEXT,
		publishedAt BIGINT,
		content TEXT
	);