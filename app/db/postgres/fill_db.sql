-- DB version 1.0 schema
DO $FN$
DECLARE
    workflow_id UUID;
    task_id UUID;
    action_id UUID;

    workflow_name VARCHAR;
    task_name VARCHAR;
    action_name VARCHAR;
    time_current TIMESTAMP;

    workflows_count INT;
    tasks_count INT;
BEGIN
    workflows_count := 5;
    tasks_count := 10;

    FOR workflow_counter IN 1..workflows_count LOOP
        time_current := now();
        workflow_id := gen_random_uuid();
        workflow_name := concat('workflow_', workflow_id);

        RAISE NOTICE '% - Creating %s with time %s', workflow_counter, workflow_name, time_current;

        EXECUTE $$
            INSERT INTO workflows(workflow_id, name, created_at, updated_at)
            VALUES ($1, $2, $3, $3)
        $$
            USING workflow_id, workflow_name, time_current;
        FOR task_counter IN 1..tasks_count LOOP
            time_current := now();
            task_id := gen_random_uuid();
            task_name := concat('task_', task_id);

            RAISE NOTICE '% - Creating %s with time %s under %s', task_counter, task_name, time_current, workflow_name;

            EXECUTE $$
                INSERT INTO tasks(task_id, name, created_at, updated_at, workflow_id)
                VALUES ($1, $2, $3, $3, $4)
            $$
                USING task_id, task_name, time_current, workflow_id;


            time_current := now();
            action_id := gen_random_uuid();
            action_name := concat('action_', action_id);

            RAISE NOTICE 'Creating %s with time %s under %s', action_name, time_current, task_name;

            EXECUTE $$
                INSERT INTO actions(action_id, name, created_at, updated_at, task_id)
                VALUES ($1, $2, $3, $3, $4)
            $$
                USING action_id, action_name, time_current, task_id;
        END LOOP;
    END LOOP;
END;
$FN$