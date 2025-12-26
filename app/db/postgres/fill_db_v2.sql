-- DB schema version 2.0
CREATE EXTENSION IF NOT EXISTS ulid;

CREATE OR REPLACE PROCEDURE fill_workflow_table(entities_counts jsonb)
LANGUAGE plpgsql
AS $$
DECLARE
    workflows_count INT;
    tasks_under_workflow_count INT;
    time_current TIMESTAMP;
    workflow_id VARCHAR;
    workflow_name VARCHAR;
    workflow_state VARCHAR;
BEGIN
    workflows_count := (entities_counts->>'workflows_count')::INT;
    tasks_under_workflow_count := (entities_counts->>'tasks_under_workflow_count')::INT;
    FOR workflow_counter IN 1..workflows_count LOOP
        time_current := now();
        workflow_id := gen_ulid();
        workflow_name := concat('workflow_', workflow_id);
        workflow_state := 'SUCCESS';
        RAISE NOTICE '% - Creating % with time % and state %',
            workflow_counter, workflow_name, time_current, workflow_state;
        EXECUTE
            'INSERT INTO workflows(workflow_id, name, created_at, state, updated_at) ' ||
            'VALUES ($1, $2, $3, $4, $5)'
        USING workflow_id, workflow_name, time_current, workflow_state, time_current;
        CALL fill_task_table(workflow_id, workflow_name, tasks_under_workflow_count);
    END LOOP;
END
$$;

CREATE OR REPLACE PROCEDURE fill_task_table(workflow_id VARCHAR, workflow_name VARCHAR, tasks_under_workflow_count INT)
LANGUAGE plpgsql
AS $$
DECLARE
    time_current TIMESTAMP;
    task_id VARCHAR;
    task_name VARCHAR;
    task_state VARCHAR;
BEGIN
    FOR task_counter IN 1..tasks_under_workflow_count LOOP
        time_current := now();
        task_id := gen_ulid();
        task_name := concat('task_', task_id);
        task_state := 'SUCCESS';
        RAISE NOTICE '% - Creating % with time % and state % under %',
            task_counter, task_name, time_current, task_state, workflow_name;
        EXECUTE
            'INSERT INTO tasks(task_id, name, created_at, state, workflow_id, updated_at) ' ||
            'VALUES ($1, $2, $3, $4, $5, $6)'
            USING task_id, task_name, time_current, task_state, workflow_id, time_current;
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
    action_state VARCHAR;
BEGIN
    time_current := now();
    action_id := gen_ulid();
    action_name := concat('action_', action_id);
    action_state := 'SUCCESS';
    RAISE NOTICE 'Creating % with time % and status % under %',
        action_name, time_current, action_state, task_name;
    EXECUTE
        'INSERT INTO actions(action_id, name, created_at, state, task_id, updated_at)' ||
        'VALUES ($1, $2, $3, $4, $5, $6)'
    USING action_id, action_name, time_current, action_state, task_id, time_current;
END
$$;

\set workflows_count :workflows_count
\set tasks_under_workflow_count :tasks_under_workflow_count
\set entities_counts '{"workflows_count": ':workflows_count', "tasks_under_workflow_count": ' :tasks_under_workflow_count '}'

CALL fill_workflow_table(:'entities_counts'::jsonb);