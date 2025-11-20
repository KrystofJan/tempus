CREATE TABLE task (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    start_timestamp INTEGER NOT NULL default (strftime('%s', 'now')),
    end_timestamp INTEGER,
    finished INTEGER NOT NULL CHECK( finished >= 0 AND finished <= 1 ) default 0,
    recorded_time INTEGER 
);

CREATE TABLE entry (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    task_id INTEGER NOT NULL,
    start_timestamp INTEGER NOT NULL default (strftime('%s', 'now')),
    end_timestamp INTEGER,
    recorded_time INTEGER,
    finished INTEGER NOT NULL CHECK( finished >= 0 AND finished <= 1 ) default 0,
    FOREIGN KEY(task_id) REFERENCES task(id)
);

CREATE TABLE current_entry (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    current_entry_id INTEGER,
    FOREIGN KEY(current_entry_id) REFERENCES entry(id)
);

CREATE TRIGGER entry_set_finished
AFTER UPDATE OF finished ON entry
FOR EACH ROW
WHEN NEW.finished = 1 AND OLD.finished = 0
BEGIN
    -- Set end_timestamp to now
    UPDATE entry
    SET
        end_timestamp = strftime('%s', 'now'),
        recorded_time = strftime('%s', 'now') - start_timestamp
    WHERE id = NEW.id;
END;

CREATE TRIGGER update_task_on_end
AFTER UPDATE OF end_timestamp ON task
FOR EACH ROW
WHEN NEW.end_timestamp IS NOT NULL
BEGIN
    UPDATE task
    SET
        recorded_time = NEW.end_timestamp - NEW.start_timestamp,
        finished = 1
    WHERE id = NEW.id;
END;


CREATE TRIGGER set_current_entry_on_insert
AFTER INSERT ON entry
FOR EACH ROW
BEGIN
    UPDATE current_entry
    SET current_entry_id = NEW.id
    WHERE id = 1;
END;

CREATE TRIGGER clear_current_entry_on_finish
AFTER UPDATE OF finished ON entry
FOR EACH ROW
WHEN NEW.finished = 1 AND OLD.finished = 0
BEGIN
    UPDATE current_entry
    SET current_entry_id = NULL
    WHERE id = 1 AND current_entry_id = NEW.id;
END;

CREATE TRIGGER delete_entries_on_task_delete
AFTER DELETE ON task
FOR EACH ROW
BEGIN
    DELETE FROM entry
    WHERE task_id = OLD.id;
END;
