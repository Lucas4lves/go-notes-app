package database

const CreateNotesTableQuery = `
create table if not exists notes
	(
		id integer primary key autoincrement,
		title text not null,
		content text not null,
		created_at text not null,
		updated_at text not null
	);
`
