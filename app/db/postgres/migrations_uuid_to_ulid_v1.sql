-- DB version 2.0 schema
CREATE EXTENSION IF NOT EXISTS ulid;

DO $FN$
DECLARE
    workflow_item workflows%ROWTYPE;
    action_item actions%ROWTYPE;

    new_workflow_id VARCHAR;
    new_task_id VARCHAR;
    new_action_id VARCHAR;
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

    FOR workflow_item IN SELECT * FROM workflows LOOP
        new_workflow_id = gen_ulid();
        EXECUTE $$
            UPDATE workflows SET workflow_id = $1 WHERE workflow_id = $2;
        $$
            USING new_workflow_id, workflow_item.workflow_id;

        EXECUTE $$
            UPDATE tasks SET workflow_id = $1 WHERE workflow_id = $2
        $$
            USING new_workflow_id, workflow_item.workflow_id;
    END LOOP;

    FOR action_item IN SELECT * FROM actions LOOP
        new_action_id = gen_ulid();
        new_task_id = gen_ulid();

        EXECUTE $$
            UPDATE actions SET action_id = $1, task_id = $2 WHERE action_id = $3
        $$
            USING new_action_id, new_task_id, action_item.action_id;

        EXECUTE $$
                UPDATE tasks SET task_id = $1 WHERE task_id = $2
            $$
        USING new_task_id, action_item.task_id;
    END LOOP;

    EXECUTE 'ALTER TABLE actions
        ADD CONSTRAINT fk_task FOREIGN KEY (task_id) REFERENCES tasks(task_id)';

    EXECUTE 'ALTER TABLE tasks
        ADD CONSTRAINT fk_workflow FOREIGN KEY (workflow_id) REFERENCES workflows(workflow_id)';
END
$FN$
