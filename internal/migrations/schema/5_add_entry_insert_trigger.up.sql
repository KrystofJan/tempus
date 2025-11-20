CREATE TRIGGER set_current_task_on_insert
AFTER INSERT ON task
FOR EACH ROW
BEGIN
    UPDATE current_task
    SET current_task_id = NEW.id
    WHERE id = 1;
END;

CREATE TRIGGER clear_current_task_on_finish
AFTER UPDATE OF finished ON task
FOR EACH ROW
WHEN NEW.finished = 1 AND OLD.finished = 0
BEGIN
    UPDATE current_task
    SET current_task_id = NULL
    WHERE id = 1 AND current_task_id = NEW.id;
END;
