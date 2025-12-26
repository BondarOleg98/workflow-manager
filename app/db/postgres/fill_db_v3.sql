-- DB version 3.0 schema
CREATE OR REPLACE PROCEDURE fill_workflow_table(entities_counts jsonb)
LANGUAGE plpgsql
AS $$
DECLARE
    workflows_count INT;
    tasks_under_workflow_count INT;
    time_current TIMESTAMP;
    workflow_id UUID;
    workflow_name VARCHAR;
    workflow_state VARCHAR;
BEGIN
    workflows_count := (entities_counts->>'workflows_count')::INT;
    tasks_under_workflow_count := (entities_counts->>'tasks_under_workflow_count')::INT;
    FOR workflow_counter IN 1..workflows_count LOOP
        time_current := now();
        workflow_name := concat('workflow_', time_current);
        workflow_state := 'SUCCESS';
        RAISE NOTICE '% - Creating % with state %',
            workflow_counter, workflow_name, workflow_state;
        EXECUTE
            'INSERT INTO workflows(name, created_at, state, updated_at) ' ||
            'VALUES ($1, $2, $3, $4)' ||
            'RETURNING workflow_id'
        INTO workflow_id
        USING workflow_name, time_current, workflow_state, time_current;
        RAISE NOTICE '% - test', workflow_id;
        CALL fill_task_table(workflow_id, workflow_name, tasks_under_workflow_count);
    END LOOP;
END
$$;

CREATE OR REPLACE PROCEDURE fill_task_table(workflow_id UUID, workflow_name VARCHAR, tasks_under_workflow_count INT)
LANGUAGE plpgsql
AS $$
DECLARE
    time_current TIMESTAMP;
    task_id UUID;
    task_name VARCHAR;
    task_state VARCHAR;
BEGIN
    FOR task_counter IN 1..tasks_under_workflow_count LOOP
        time_current := now();
        task_name := concat('task_', time_current);
        task_state := 'SUCCESS';
        RAISE NOTICE '% - Creating % with state % under %',
            task_counter, task_name, task_state, workflow_name;
        EXECUTE
            'INSERT INTO tasks(name, created_at, state, workflow_id, updated_at) ' ||
            'VALUES ($1, $2, $3, $4, $5) ' ||
            'RETURNING task_id'
        INTO task_id
        USING task_name, time_current, task_state, workflow_id, time_current;
        CALL fill_action_table(task_id, task_name);
    END LOOP;
END
$$;


CREATE OR REPLACE PROCEDURE fill_action_table(task_id UUID, task_name VARCHAR)
LANGUAGE plpgsql
AS $$
DECLARE
    time_current TIMESTAMP;
    action_name VARCHAR;
    action_state VARCHAR;
BEGIN
    time_current := now();
    action_name := concat('action_', time_current);
    action_state := 'SUCCESS';
    RAISE NOTICE 'Creating % with state % under %',
        action_name, action_state, task_name;
    EXECUTE
        'INSERT INTO actions(name, created_at, state, task_id, updated_at)' ||
        'VALUES ($1, $2, $3, $4, $5)'
    USING action_name, time_current, action_state, task_id, time_current;
END
$$;

\set workflows_count :workflows_count
\set tasks_under_workflow_count :tasks_under_workflow_count
\set entities_counts '{"workflows_count": ':workflows_count', "tasks_under_workflow_count": ' :tasks_under_workflow_count '}'

CALL fill_workflow_table(:'entities_counts'::jsonb);