-- DB version 2.0 schema

CREATE EXTENSION IF NOT EXISTS ulid;

CREATE OR REPLACE PROCEDURE fill_workflow_table()
LANGUAGE plpgsql
AS $$
DECLARE
    workflows_count INT;
    time_current TIMESTAMP;
    workflow_id VARCHAR;
    workflow_name VARCHAR;
BEGIN
    workflows_count := 5;
    FOR workflow_counter IN 1..workflows_count LOOP
        time_current := now();
        workflow_id := gen_ulid();
        workflow_name := concat('workflow_', workflow_id);
        RAISE NOTICE '% - Creating %s with time %s', workflow_counter, workflow_name, time_current;
        EXECUTE
            'INSERT INTO workflows(workflow_id, name, created_at, updated_at) ' ||
            'VALUES ($1, $2, $3, $3)'
        USING workflow_id, workflow_name, time_current;
        CALL fill_task_table(workflow_id, workflow_name);
    END LOOP;
END
$$;

CREATE OR REPLACE PROCEDURE fill_task_table(workflow_id VARCHAR, workflow_name VARCHAR)
LANGUAGE plpgsql
AS $$
DECLARE
    tasks_count INT;
    time_current TIMESTAMP;
    task_id VARCHAR;
    task_name VARCHAR;
BEGIN
    tasks_count := 10;
    FOR task_counter IN 1..tasks_count LOOP
        time_current := now();
        task_id := gen_ulid();
        task_name := concat('task_', task_id);
        RAISE NOTICE '% - Creating %s with time %s under %s', task_counter, task_name, time_current, workflow_name;
        EXECUTE
            'INSERT INTO tasks(task_id, name, created_at, updated_at, workflow_id) ' ||
            'VALUES ($1, $2, $3, $3, $4)'
            USING task_id, task_name, time_current, workflow_id;
        CALL fill_action_table(task_id, task_name);
    END LOOP;
END
$$;


CREATE OR REPLACE PROCEDURE fill_action_table(task_id VARCHAR, task_name VARCHAR)
LANGUAGE plpgsql
AS $$
DECLARE
    time_current TIMESTAMP;
    action_id VARCHAR;
    action_name VARCHAR;
BEGIN
    time_current := now();
    action_id := gen_ulid();
    action_name := concat('action_', action_id);
    RAISE NOTICE 'Creating %s with time %s under %s', action_name, time_current, task_name;
    EXECUTE
        'INSERT INTO actions(action_id, name, created_at, updated_at, task_id)' ||
        'VALUES ($1, $2, $3, $3, $4)'
    USING action_id, action_name, time_current, task_id;
END
$$;

DO $main$
BEGIN
    CALL fill_workflow_table();
END
$main$;