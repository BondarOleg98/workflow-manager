-- DB version 2.0 schema
CREATE TABLE IF NOT EXISTS workflows (
    workflow_id VARCHAR(36) NOT NULL,
    name VARCHAR (50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    PRIMARY KEY(workflow_id)
);

CREATE TABLE IF NOT EXISTS tasks (
    task_id VARCHAR(36) NOT NULL,
    workflow_id VARCHAR(36) NOT NULL,
    name VARCHAR (50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    PRIMARY KEY(task_id),
    CONSTRAINT fk_workflow
        FOREIGN KEY(workflow_id)
            REFERENCES workflows(workflow_id)
);

CREATE TABLE IF NOT EXISTS actions (
    action_id VARCHAR(36) NOT NULL,
    task_id VARCHAR(36) NOT NULL,
    name VARCHAR (50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    PRIMARY KEY(action_id),
    CONSTRAINT fk_task
        FOREIGN KEY(task_id)
            REFERENCES tasks(task_id)
);
