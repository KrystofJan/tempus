ALTER TABLE current_entry
RENAME COLUMN current_task_id to current_entry_id; 

ALTER TABLE current_task
RENAME COLUMN current_entry_id to current_task_id; 
