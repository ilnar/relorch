// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: workflow.sql

package sqlc

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

const createWorkflow = `-- name: CreateWorkflow :one
INSERT INTO workflows (
  id,
  current_node,
  status,
  graph
) VALUES (
  $1, $2, $3, $4
)
RETURNING id, current_node, status, graph, created_at, next_action_at
`

type CreateWorkflowParams struct {
	ID          uuid.UUID       `json:"id"`
	CurrentNode string          `json:"current_node"`
	Status      string          `json:"status"`
	Graph       json.RawMessage `json:"graph"`
}

func (q *Queries) CreateWorkflow(ctx context.Context, arg CreateWorkflowParams) (Workflow, error) {
	row := q.db.QueryRowContext(ctx, createWorkflow,
		arg.ID,
		arg.CurrentNode,
		arg.Status,
		arg.Graph,
	)
	var i Workflow
	err := row.Scan(
		&i.ID,
		&i.CurrentNode,
		&i.Status,
		&i.Graph,
		&i.CreatedAt,
		&i.NextActionAt,
	)
	return i, err
}

const getNextWorkflows = `-- name: GetNextWorkflows :many
SELECT id, current_node, status, graph, created_at, next_action_at FROM workflows
WHERE status = 'running'
  AND next_action_at <= now()
LIMIT 10
`

func (q *Queries) GetNextWorkflows(ctx context.Context) ([]Workflow, error) {
	rows, err := q.db.QueryContext(ctx, getNextWorkflows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Workflow
	for rows.Next() {
		var i Workflow
		if err := rows.Scan(
			&i.ID,
			&i.CurrentNode,
			&i.Status,
			&i.Graph,
			&i.CreatedAt,
			&i.NextActionAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getWorkflow = `-- name: GetWorkflow :one
SELECT id, current_node, status, graph, created_at, next_action_at FROM workflows
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetWorkflow(ctx context.Context, id uuid.UUID) (Workflow, error) {
	row := q.db.QueryRowContext(ctx, getWorkflow, id)
	var i Workflow
	err := row.Scan(
		&i.ID,
		&i.CurrentNode,
		&i.Status,
		&i.Graph,
		&i.CreatedAt,
		&i.NextActionAt,
	)
	return i, err
}

const updateWorkflowNextAction = `-- name: UpdateWorkflowNextAction :one
UPDATE workflows
SET current_node = $2, next_action_at = now() + interval '10 seconds'
WHERE id = $1
RETURNING id, current_node, status, graph, created_at, next_action_at
`

type UpdateWorkflowNextActionParams struct {
	ID          uuid.UUID `json:"id"`
	CurrentNode string    `json:"current_node"`
}

func (q *Queries) UpdateWorkflowNextAction(ctx context.Context, arg UpdateWorkflowNextActionParams) (Workflow, error) {
	row := q.db.QueryRowContext(ctx, updateWorkflowNextAction, arg.ID, arg.CurrentNode)
	var i Workflow
	err := row.Scan(
		&i.ID,
		&i.CurrentNode,
		&i.Status,
		&i.Graph,
		&i.CreatedAt,
		&i.NextActionAt,
	)
	return i, err
}

const updateWorkflowStatus = `-- name: UpdateWorkflowStatus :one
UPDATE workflows
SET status = $2
WHERE id = $1
RETURNING id, current_node, status, graph, created_at, next_action_at
`

type UpdateWorkflowStatusParams struct {
	ID     uuid.UUID `json:"id"`
	Status string    `json:"status"`
}

func (q *Queries) UpdateWorkflowStatus(ctx context.Context, arg UpdateWorkflowStatusParams) (Workflow, error) {
	row := q.db.QueryRowContext(ctx, updateWorkflowStatus, arg.ID, arg.Status)
	var i Workflow
	err := row.Scan(
		&i.ID,
		&i.CurrentNode,
		&i.Status,
		&i.Graph,
		&i.CreatedAt,
		&i.NextActionAt,
	)
	return i, err
}
