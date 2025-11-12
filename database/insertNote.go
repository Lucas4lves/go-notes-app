package database

const InsertNoteStmt = `
insert into notes (title, content, created_at, updated_at)
values (?,?,?,?)
`
