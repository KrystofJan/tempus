drop trigger if exists delete_entries_on_task_delete;
drop trigger if exists clear_current_entry_on_finish;
drop trigger if exists set_current_entry_on_insert;
drop trigger if exists update_task_on_end;
drop trigger if exists entry_set_finished;

drop table if exists current_entry;
drop table if exists entry;
drop table if exists task;
