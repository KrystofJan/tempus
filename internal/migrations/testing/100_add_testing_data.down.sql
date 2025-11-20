DELETE FROM entry
WHERE task_id IN (
    SELECT id FROM task
    WHERE name IN (
        'Write Documentation',
        'Implement Authentication',
        'Fix Bug #123',
        'Create Database Schema',
        'Refactor Code'
    )
);

DELETE FROM task
WHERE name IN (
    'Write Documentation',
    'Implement Authentication',
    'Fix Bug #123',
    'Create Database Schema',
    'Refactor Code'
);

DELETE FROM sqlite_sequence WHERE name IN ('task', 'entry');

UPDATE current_entry SET current_entry_id = NULL WHERE id = 1;
UPDATE current_task SET current_task_id = NULL WHERE id = 1;
