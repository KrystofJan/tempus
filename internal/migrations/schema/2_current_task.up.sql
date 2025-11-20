CREATE TABLE current_task (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    current_entry_id INTEGER,
    FOREIGN KEY(current_entry_id) REFERENCES entry(id)
);
