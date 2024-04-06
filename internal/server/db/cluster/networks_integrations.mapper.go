//go:build linux && cgo && !agent

package cluster

// The code below was generated by incus-generate - DO NOT EDIT!

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/lxc/incus/v6/internal/server/db/query"
	"github.com/lxc/incus/v6/shared/api"
)

var _ = api.ServerEnvironment{}

var networkIntegrationObjects = RegisterStmt(`
SELECT networks_integrations.id, networks_integrations.name, networks_integrations.description, networks_integrations.type
  FROM networks_integrations
  ORDER BY networks_integrations.name
`)

var networkIntegrationObjectsByName = RegisterStmt(`
SELECT networks_integrations.id, networks_integrations.name, networks_integrations.description, networks_integrations.type
  FROM networks_integrations
  WHERE ( networks_integrations.name = ? )
  ORDER BY networks_integrations.name
`)

var networkIntegrationObjectsByID = RegisterStmt(`
SELECT networks_integrations.id, networks_integrations.name, networks_integrations.description, networks_integrations.type
  FROM networks_integrations
  WHERE ( networks_integrations.id = ? )
  ORDER BY networks_integrations.name
`)

var networkIntegrationCreate = RegisterStmt(`
INSERT INTO networks_integrations (name, description, type)
  VALUES (?, ?, ?)
`)

var networkIntegrationID = RegisterStmt(`
SELECT networks_integrations.id FROM networks_integrations
  WHERE networks_integrations.name = ?
`)

var networkIntegrationRename = RegisterStmt(`
UPDATE networks_integrations SET name = ? WHERE name = ?
`)

var networkIntegrationUpdate = RegisterStmt(`
UPDATE networks_integrations
  SET name = ?, description = ?, type = ?
 WHERE id = ?
`)

var networkIntegrationDeleteByName = RegisterStmt(`
DELETE FROM networks_integrations WHERE name = ?
`)

// networkIntegrationColumns returns a string of column names to be used with a SELECT statement for the entity.
// Use this function when building statements to retrieve database entries matching the NetworkIntegration entity.
func networkIntegrationColumns() string {
	return "networks_integrations.id, networks_integrations.name, networks_integrations.description, networks_integrations.type"
}

// getNetworkIntegrations can be used to run handwritten sql.Stmts to return a slice of objects.
func getNetworkIntegrations(ctx context.Context, stmt *sql.Stmt, args ...any) ([]NetworkIntegration, error) {
	objects := make([]NetworkIntegration, 0)

	dest := func(scan func(dest ...any) error) error {
		n := NetworkIntegration{}
		err := scan(&n.ID, &n.Name, &n.Description, &n.Type)
		if err != nil {
			return err
		}

		objects = append(objects, n)

		return nil
	}

	err := query.SelectObjects(ctx, stmt, dest, args...)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch from \"networks_integrations\" table: %w", err)
	}

	return objects, nil
}

// getNetworkIntegrationsRaw can be used to run handwritten query strings to return a slice of objects.
func getNetworkIntegrationsRaw(ctx context.Context, tx *sql.Tx, sql string, args ...any) ([]NetworkIntegration, error) {
	objects := make([]NetworkIntegration, 0)

	dest := func(scan func(dest ...any) error) error {
		n := NetworkIntegration{}
		err := scan(&n.ID, &n.Name, &n.Description, &n.Type)
		if err != nil {
			return err
		}

		objects = append(objects, n)

		return nil
	}

	err := query.Scan(ctx, tx, sql, dest, args...)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch from \"networks_integrations\" table: %w", err)
	}

	return objects, nil
}

// GetNetworkIntegrations returns all available network_integrations.
// generator: network_integration GetMany
func GetNetworkIntegrations(ctx context.Context, tx *sql.Tx, filters ...NetworkIntegrationFilter) ([]NetworkIntegration, error) {
	var err error

	// Result slice.
	objects := make([]NetworkIntegration, 0)

	// Pick the prepared statement and arguments to use based on active criteria.
	var sqlStmt *sql.Stmt
	args := []any{}
	queryParts := [2]string{}

	if len(filters) == 0 {
		sqlStmt, err = Stmt(tx, networkIntegrationObjects)
		if err != nil {
			return nil, fmt.Errorf("Failed to get \"networkIntegrationObjects\" prepared statement: %w", err)
		}
	}

	for i, filter := range filters {
		if filter.Name != nil && filter.ID == nil {
			args = append(args, []any{filter.Name}...)
			if len(filters) == 1 {
				sqlStmt, err = Stmt(tx, networkIntegrationObjectsByName)
				if err != nil {
					return nil, fmt.Errorf("Failed to get \"networkIntegrationObjectsByName\" prepared statement: %w", err)
				}

				break
			}

			query, err := StmtString(networkIntegrationObjectsByName)
			if err != nil {
				return nil, fmt.Errorf("Failed to get \"networkIntegrationObjects\" prepared statement: %w", err)
			}

			parts := strings.SplitN(query, "ORDER BY", 2)
			if i == 0 {
				copy(queryParts[:], parts)
				continue
			}

			_, where, _ := strings.Cut(parts[0], "WHERE")
			queryParts[0] += "OR" + where
		} else if filter.ID != nil && filter.Name == nil {
			args = append(args, []any{filter.ID}...)
			if len(filters) == 1 {
				sqlStmt, err = Stmt(tx, networkIntegrationObjectsByID)
				if err != nil {
					return nil, fmt.Errorf("Failed to get \"networkIntegrationObjectsByID\" prepared statement: %w", err)
				}

				break
			}

			query, err := StmtString(networkIntegrationObjectsByID)
			if err != nil {
				return nil, fmt.Errorf("Failed to get \"networkIntegrationObjects\" prepared statement: %w", err)
			}

			parts := strings.SplitN(query, "ORDER BY", 2)
			if i == 0 {
				copy(queryParts[:], parts)
				continue
			}

			_, where, _ := strings.Cut(parts[0], "WHERE")
			queryParts[0] += "OR" + where
		} else if filter.ID == nil && filter.Name == nil {
			return nil, fmt.Errorf("Cannot filter on empty NetworkIntegrationFilter")
		} else {
			return nil, fmt.Errorf("No statement exists for the given Filter")
		}
	}

	// Select.
	if sqlStmt != nil {
		objects, err = getNetworkIntegrations(ctx, sqlStmt, args...)
	} else {
		queryStr := strings.Join(queryParts[:], "ORDER BY")
		objects, err = getNetworkIntegrationsRaw(ctx, tx, queryStr, args...)
	}

	if err != nil {
		return nil, fmt.Errorf("Failed to fetch from \"networks_integrations\" table: %w", err)
	}

	return objects, nil
}

// GetNetworkIntegrationConfig returns all available NetworkIntegration Config
// generator: network_integration GetMany
func GetNetworkIntegrationConfig(ctx context.Context, tx *sql.Tx, networkIntegrationID int, filters ...ConfigFilter) (map[string]string, error) {
	networkIntegrationConfig, err := GetConfig(ctx, tx, "network_integration", filters...)
	if err != nil {
		return nil, err
	}

	config, ok := networkIntegrationConfig[networkIntegrationID]
	if !ok {
		config = map[string]string{}
	}

	return config, nil
}

// GetNetworkIntegration returns the network_integration with the given key.
// generator: network_integration GetOne
func GetNetworkIntegration(ctx context.Context, tx *sql.Tx, name string) (*NetworkIntegration, error) {
	filter := NetworkIntegrationFilter{}
	filter.Name = &name

	objects, err := GetNetworkIntegrations(ctx, tx, filter)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch from \"networks_integrations\" table: %w", err)
	}

	switch len(objects) {
	case 0:
		return nil, api.StatusErrorf(http.StatusNotFound, "NetworkIntegration not found")
	case 1:
		return &objects[0], nil
	default:
		return nil, fmt.Errorf("More than one \"networks_integrations\" entry matches")
	}
}

// NetworkIntegrationExists checks if a network_integration with the given key exists.
// generator: network_integration Exists
func NetworkIntegrationExists(ctx context.Context, tx *sql.Tx, name string) (bool, error) {
	_, err := GetNetworkIntegrationID(ctx, tx, name)
	if err != nil {
		if api.StatusErrorCheck(err, http.StatusNotFound) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

// CreateNetworkIntegration adds a new network_integration to the database.
// generator: network_integration Create
func CreateNetworkIntegration(ctx context.Context, tx *sql.Tx, object NetworkIntegration) (int64, error) {
	// Check if a network_integration with the same key exists.
	exists, err := NetworkIntegrationExists(ctx, tx, object.Name)
	if err != nil {
		return -1, fmt.Errorf("Failed to check for duplicates: %w", err)
	}

	if exists {
		return -1, api.StatusErrorf(http.StatusConflict, "This \"networks_integrations\" entry already exists")
	}

	args := make([]any, 3)

	// Populate the statement arguments.
	args[0] = object.Name
	args[1] = object.Description
	args[2] = object.Type

	// Prepared statement to use.
	stmt, err := Stmt(tx, networkIntegrationCreate)
	if err != nil {
		return -1, fmt.Errorf("Failed to get \"networkIntegrationCreate\" prepared statement: %w", err)
	}

	// Execute the statement.
	result, err := stmt.Exec(args...)
	if err != nil {
		return -1, fmt.Errorf("Failed to create \"networks_integrations\" entry: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("Failed to fetch \"networks_integrations\" entry ID: %w", err)
	}

	return id, nil
}

// CreateNetworkIntegrationConfig adds new network_integration Config to the database.
// generator: network_integration Create
func CreateNetworkIntegrationConfig(ctx context.Context, tx *sql.Tx, networkIntegrationID int64, config map[string]string) error {
	referenceID := int(networkIntegrationID)
	for key, value := range config {
		insert := Config{
			ReferenceID: referenceID,
			Key:         key,
			Value:       value,
		}

		err := CreateConfig(ctx, tx, "network_integration", insert)
		if err != nil {
			return fmt.Errorf("Insert Config failed for NetworkIntegration: %w", err)
		}

	}

	return nil
}

// GetNetworkIntegrationID return the ID of the network_integration with the given key.
// generator: network_integration ID
func GetNetworkIntegrationID(ctx context.Context, tx *sql.Tx, name string) (int64, error) {
	stmt, err := Stmt(tx, networkIntegrationID)
	if err != nil {
		return -1, fmt.Errorf("Failed to get \"networkIntegrationID\" prepared statement: %w", err)
	}

	row := stmt.QueryRowContext(ctx, name)
	var id int64
	err = row.Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return -1, api.StatusErrorf(http.StatusNotFound, "NetworkIntegration not found")
	}

	if err != nil {
		return -1, fmt.Errorf("Failed to get \"networks_integrations\" ID: %w", err)
	}

	return id, nil
}

// RenameNetworkIntegration renames the network_integration matching the given key parameters.
// generator: network_integration Rename
func RenameNetworkIntegration(ctx context.Context, tx *sql.Tx, name string, to string) error {
	stmt, err := Stmt(tx, networkIntegrationRename)
	if err != nil {
		return fmt.Errorf("Failed to get \"networkIntegrationRename\" prepared statement: %w", err)
	}

	result, err := stmt.Exec(to, name)
	if err != nil {
		return fmt.Errorf("Rename NetworkIntegration failed: %w", err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Fetch affected rows failed: %w", err)
	}

	if n != 1 {
		return fmt.Errorf("Query affected %d rows instead of 1", n)
	}

	return nil
}

// DeleteNetworkIntegration deletes the network_integration matching the given key parameters.
// generator: network_integration DeleteOne-by-Name
func DeleteNetworkIntegration(ctx context.Context, tx *sql.Tx, name string) error {
	stmt, err := Stmt(tx, networkIntegrationDeleteByName)
	if err != nil {
		return fmt.Errorf("Failed to get \"networkIntegrationDeleteByName\" prepared statement: %w", err)
	}

	result, err := stmt.Exec(name)
	if err != nil {
		return fmt.Errorf("Delete \"networks_integrations\": %w", err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Fetch affected rows: %w", err)
	}

	if n == 0 {
		return api.StatusErrorf(http.StatusNotFound, "NetworkIntegration not found")
	} else if n > 1 {
		return fmt.Errorf("Query deleted %d NetworkIntegration rows instead of 1", n)
	}

	return nil
}

// UpdateNetworkIntegration updates the network_integration matching the given key parameters.
// generator: network_integration Update
func UpdateNetworkIntegration(ctx context.Context, tx *sql.Tx, name string, object NetworkIntegration) error {
	id, err := GetNetworkIntegrationID(ctx, tx, name)
	if err != nil {
		return err
	}

	stmt, err := Stmt(tx, networkIntegrationUpdate)
	if err != nil {
		return fmt.Errorf("Failed to get \"networkIntegrationUpdate\" prepared statement: %w", err)
	}

	result, err := stmt.Exec(object.Name, object.Description, object.Type, id)
	if err != nil {
		return fmt.Errorf("Update \"networks_integrations\" entry failed: %w", err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Fetch affected rows: %w", err)
	}

	if n != 1 {
		return fmt.Errorf("Query updated %d rows instead of 1", n)
	}

	return nil
}

// UpdateNetworkIntegrationConfig updates the network_integration Config matching the given key parameters.
// generator: network_integration Update
func UpdateNetworkIntegrationConfig(ctx context.Context, tx *sql.Tx, networkIntegrationID int64, config map[string]string) error {
	err := UpdateConfig(ctx, tx, "network_integration", int(networkIntegrationID), config)
	if err != nil {
		return fmt.Errorf("Replace Config for NetworkIntegration failed: %w", err)
	}

	return nil
}
