CREATE TABLE book
				(
				id Integer PRIMARY KEY,
				name varchar(255),
				format varchar(64),
				content text,
				tags varchar(1024),
				categories varchar(1024),
				pan_url varchar(255),
				view_count int,
				dl_count int,
				created_at timestamp,
				updated_at timestamp
				)
