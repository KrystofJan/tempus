DELETE FROM task;
DELETE FROM entry;

INSERT INTO task (name, start_timestamp, end_timestamp, finished, recorded_time)
VALUES
  ('Write Documentation', strftime('%s', 'now') - 7200, strftime('%s', 'now') - 3600, 1, 3600),
  ('Implement Authentication', strftime('%s', 'now') - 10800, NULL, 0, NULL),
  ('Fix Bug #123', strftime('%s', 'now') - 5400, strftime('%s', 'now') - 1800, 1, 3600),
  ('Create Database Schema', strftime('%s', 'now') - 20000, strftime('%s', 'now') - 10000, 1, 10000),
  ('Refactor Code', strftime('%s', 'now') - 5000, NULL, 0, NULL);

INSERT INTO entry (task_id, start_timestamp, end_timestamp, recorded_time, finished)
VALUES
  (1, strftime('%s', 'now') - 7500, strftime('%s', 'now') - 7200, 300, 1),
  (1, strftime('%s', 'now') - 4000, strftime('%s', 'now') - 3600, 400, 1);

INSERT INTO entry (task_id, start_timestamp, end_timestamp, recorded_time, finished)
VALUES
  (2, strftime('%s', 'now') - 8000, NULL, NULL, 0);

INSERT INTO entry (task_id, start_timestamp, end_timestamp, recorded_time, finished)
VALUES
  (3, strftime('%s', 'now') - 5000, strftime('%s', 'now') - 1800, 3200, 1);

INSERT INTO entry (task_id, start_timestamp, end_timestamp, recorded_time, finished)
VALUES
  (4, strftime('%s', 'now') - 15000, strftime('%s', 'now') - 10000, 5000, 1),
  (4, strftime('%s', 'now') - 12000, strftime('%s', 'now') - 10000, 2000, 1);

INSERT INTO entry (task_id, start_timestamp, end_timestamp, recorded_time, finished)
VALUES
  (5, strftime('%s', 'now') - 4000, NULL, NULL, 0);

DELETE FROM current_entry;
DELETE FROM current_task;

INSERT INTO current_entry (id, current_entry_id)
VALUES (1, (SELECT id FROM entry WHERE finished = 0 LIMIT 1));

INSERT INTO current_task (id, current_task_id)
VALUES (1, (SELECT id FROM task WHERE finished = 0 LIMIT 1));
