-- DB version 2.0 schema
CREATE EXTENSION IF NOT EXISTS ulid;

CREATE OR REPLACE PROCEDURE drop_tables_constraint()
LANGUAGE plpgsql
AS $$
BEGIN
    EXECUTE 'ALTER TABLE tasks DROP CONSTRAINT IF EXISTS fk_workflow';
    EXECUTE 'ALTER TABLE actions DROP CONSTRAINT IF EXISTS fk_task';

    EXECUTE 'ALTER TABLE workflows
        ALTER COLUMN workflow_id TYPE VARCHAR(36)';

    EXECUTE 'ALTER TABLE tasks
        ALTER COLUMN task_id TYPE VARCHAR(36),
        ALTER COLUMN workflow_id TYPE VARCHAR(36)';

    EXECUTE 'ALTER TABLE actions
        ALTER COLUMN action_id TYPE VARCHAR(36),
        ALTER COLUMN task_id TYPE VARCHAR(36)';
END
$$;

CREATE OR REPLACE PROCEDURE update_workflow_task_tables_relationships()
LANGUAGE plpgsql
AS $$
DECLARE
    workflow_item workflows%ROWTYPE;
    new_workflow_name VARCHAR;
    new_workflow_id VARCHAR;
BEGIN
    FOR workflow_item IN SELECT * FROM workflows LOOP
        new_workflow_id = gen_ulid();
        new_workflow_name := concat('workflow_', new_workflow_id);
        EXECUTE
            'UPDATE workflows SET workflow_id = $1, name = $2 WHERE workflow_id = $3'
        USING new_workflow_id, new_workflow_name, workflow_item.workflow_id;

        EXECUTE
            'UPDATE tasks SET workflow_id = $1 WHERE workflow_id = $2'
        USING new_workflow_id, workflow_item.workflow_id;
    END LOOP;
END
$$;

CREATE OR REPLACE PROCEDURE update_task_action_tables_relationships()
    LANGUAGE plpgsql
AS $$
DECLARE
    action_item actions%ROWTYPE;
    new_task_name VARCHAR;
    new_action_name VARCHAR;
    new_task_id VARCHAR;
    new_action_id VARCHAR;
BEGIN
    FOR action_item IN SELECT * FROM actions LOOP
        new_action_id = gen_ulid();
        new_task_id = gen_ulid();

        new_task_name := concat('task_', new_task_id);
        new_action_name := concat('action_', new_action_id);

        EXECUTE
            'UPDATE actions SET action_id = $1, task_id = $2, name = $3 WHERE action_id = $4'
        USING new_action_id, new_task_id, new_action_name, action_item.action_id;

        EXECUTE
            'UPDATE tasks SET task_id = $1, name = $2 WHERE task_id = $3'
        USING new_task_id, new_task_name, action_item.task_id;
    END LOOP;
END
$$;

CREATE OR REPLACE PROCEDURE set_tables_constraint()
LANGUAGE plpgsql
AS $$
BEGIN
    EXECUTE 'ALTER TABLE actions
        ADD CONSTRAINT fk_task FOREIGN KEY (task_id) REFERENCES tasks(task_id)';

    EXECUTE 'ALTER TABLE tasks
        ADD CONSTRAINT fk_workflow FOREIGN KEY (workflow_id) REFERENCES workflows(workflow_id)';
END
$$;

DO $main$
    DECLARE
    BEGIN
        CALL drop_tables_constraint();
        CALL update_workflow_task_tables_relationships();
        CALL update_task_action_tables_relationships();
        CALL set_tables_constraint();
    END
$main$;