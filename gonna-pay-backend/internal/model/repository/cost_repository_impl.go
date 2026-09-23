package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/domains"
)

type CostRepository interface {
	Create(ctx context.Context, cost *domains.Cost) (*domains.Cost, error)
	Update(ctx context.Context, id string, cost *domains.Cost) (*domains.Cost, error)
	FindAll(ctx context.Context, userID string, limit, offset int, filters domains.CostFilters) ([]*domains.Cost, int64, error)
	FindByID(ctx context.Context, id string) (*domains.Cost, error)
	Delete(ctx context.Context, id string) error
	FindStats(ctx context.Context, userID string, filters domains.CostFilters) (*domains.CostStats, error)
}

type costRepository struct {
	db *sql.DB
}

func NewCostRepository(db *sql.DB) CostRepository {
	return &costRepository{db: db}
}

type costRow struct {
	id              uuid.UUID
	userID          uuid.UUID
	groupID         sql.NullString
	groupName       sql.NullString
	costName        string
	totalValue      float64
	ownerPercentage float64
	category        string
	createdAt       time.Time
	updatedAt       time.Time
	splitCount      int
	splitsJSON      []byte
}

func costRowToDomain(r costRow) (*domains.Cost, error) {
	groupID := ""
	if r.groupID.Valid {
		groupID = r.groupID.String
	}

	groupName := ""
	if r.groupName.Valid {
		groupName = r.groupName.String
	}

	var splits []domains.Split
	if r.splitsJSON != nil {
		if err := json.Unmarshal(r.splitsJSON, &splits); err != nil {
			return nil, err
		}
	}

	splitCount := r.splitCount
	if len(splits) > 0 {
		splitCount = len(splits)
	}

	return &domains.Cost{
		ID:              r.id.String(),
		UserID:          r.userID.String(),
		GroupID:         groupID,
		GroupName:       groupName,
		CostName:        r.costName,
		TotalValue:      r.totalValue,
		OwnerPercentage: r.ownerPercentage,
		Category:        r.category,
		CreatedAt:       r.createdAt,
		UpdatedAt:       r.updatedAt,
		SplitCount:      splitCount,
		Splits:          splits,
	}, nil
}

func buildCostConditions(col, userID string, filters domains.CostFilters) ([]string, []any, int) {
	conditions := []string{col + "user_id = $1"}
	args := []any{userID}
	idx := 2

	if filters.Category != "" {
		conditions = append(conditions, fmt.Sprintf(col+"category = $%d", idx))
		args = append(args, filters.Category)
		idx++
	}

	switch filters.Type {
	case domains.CostTypeSolo:
		conditions = append(conditions, col+"group_id IS NULL")
	case domains.CostTypeGroup:
		conditions = append(conditions, col+"group_id IS NOT NULL")
	}

	switch filters.Period {
	case domains.CostPeriodMonth:
		conditions = append(conditions, "DATE_TRUNC('month', "+col+"created_at) = DATE_TRUNC('month', NOW())")
	case domains.CostPeriodWeek:
		conditions = append(conditions, col+"created_at >= NOW() - INTERVAL '7 days'")
	}

	if filters.MinValue != nil {
		conditions = append(conditions, fmt.Sprintf(col+"total_value >= $%d", idx))
		args = append(args, *filters.MinValue)
		idx++
	}
	if filters.MaxValue != nil {
		conditions = append(conditions, fmt.Sprintf(col+"total_value <= $%d", idx))
		args = append(args, *filters.MaxValue)
		idx++
	}

	return conditions, args, idx
}

func (r *costRepository) Create(ctx context.Context, cost *domains.Cost) (*domains.Cost, error) {
	costID := uuid.New()
	now := time.Now()

	var groupID interface{}
	if cost.GroupID != "" {
		groupID = cost.GroupID
	}

	for i := range cost.Splits {
		cost.Splits[i].ID = uuid.New().String()
	}

	var splitsJSON []byte
	if len(cost.Splits) > 0 {
		var err error
		splitsJSON, err = json.Marshal(cost.Splits)
		if err != nil {
			return nil, err
		}
	}

	_, err := r.db.ExecContext(ctx,
		`INSERT INTO cost_entities (id, user_id, group_id, cost_name, total_value, owner_percentage, category, splits, created_at, updated_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8::jsonb, $9, $10)`,
		costID, cost.UserID, groupID, cost.CostName,
		cost.TotalValue, cost.OwnerPercentage, cost.Category, splitsJSON, now, now,
	)
	if err != nil {
		return nil, err
	}

	return domains.NewCostWithID(
		costID.String(),
		cost.UserID,
		cost.GroupID,
		"",
		cost.CostName,
		cost.Category,
		cost.TotalValue,
		cost.OwnerPercentage,
		now,
		now,
		cost.Splits,
	), nil
}

func (r *costRepository) Update(ctx context.Context, id string, cost *domains.Cost) (*domains.Cost, error) {
	now := time.Now()

	var splitsJSON []byte
	if len(cost.Splits) > 0 {
		var err error
		splitsJSON, err = json.Marshal(cost.Splits)
		if err != nil {
			return nil, err
		}
	}

	_, err := r.db.ExecContext(ctx,
		`UPDATE cost_entities SET cost_name=$1, total_value=$2, owner_percentage=$3, category=$4, splits=$5::jsonb, updated_at=$6 WHERE id=$7`,
		cost.CostName, cost.TotalValue, cost.OwnerPercentage, cost.Category, splitsJSON, now, id,
	)
	if err != nil {
		return nil, err
	}

	return &domains.Cost{
		ID:              cost.ID,
		UserID:          cost.UserID,
		GroupID:         cost.GroupID,
		GroupName:       cost.GroupName,
		CostName:        cost.CostName,
		TotalValue:      cost.TotalValue,
		OwnerPercentage: cost.OwnerPercentage,
		Category:        cost.Category,
		CreatedAt:       cost.CreatedAt,
		UpdatedAt:       now,
		SplitCount:      len(cost.Splits),
		Splits:          cost.Splits,
	}, nil
}

func (r *costRepository) FindAll(ctx context.Context, userID string, limit, offset int, filters domains.CostFilters) ([]*domains.Cost, int64, error) {
	conditions, args, idx := buildCostConditions("c.", userID, filters)
	where := strings.Join(conditions, " AND ")

	var total int64
	err := r.db.QueryRowContext(ctx,
		fmt.Sprintf(`SELECT COUNT(*) FROM cost_entities c WHERE %s`, where),
		args...,
	).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	queryArgs := append(args, limit, offset)
	rows, err := r.db.QueryContext(ctx,
		fmt.Sprintf(`
			SELECT
				c.id, c.user_id, c.group_id, c.cost_name, c.total_value, c.owner_percentage, c.category, c.created_at, c.updated_at,
				g.name,
				COALESCE(jsonb_array_length(c.splits), 0) AS split_count
			FROM cost_entities c
			LEFT JOIN group_entities g ON g.id = c.group_id
			WHERE %s
			ORDER BY c.created_at DESC
			LIMIT $%d OFFSET $%d
		`, where, idx, idx+1),
		queryArgs...,
	)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var costs []*domains.Cost
	for rows.Next() {
		var row costRow
		if err := rows.Scan(
			&row.id, &row.userID, &row.groupID, &row.costName, &row.totalValue,
			&row.ownerPercentage, &row.category, &row.createdAt, &row.updatedAt,
			&row.groupName, &row.splitCount,
		); err != nil {
			return nil, 0, err
		}

		cost, err := costRowToDomain(row)
		if err != nil {
			return nil, 0, err
		}
		costs = append(costs, cost)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return costs, total, nil
}

func (r *costRepository) FindByID(ctx context.Context, id string) (*domains.Cost, error) {
	var row costRow
	err := r.db.QueryRowContext(ctx, `
		SELECT
			c.id, c.user_id, c.group_id, c.cost_name, c.total_value, c.owner_percentage, c.category, c.created_at, c.updated_at,
			g.name, c.splits
		FROM cost_entities c
		LEFT JOIN group_entities g ON g.id = c.group_id
		WHERE c.id = $1
	`, id).Scan(
		&row.id, &row.userID, &row.groupID, &row.costName, &row.totalValue,
		&row.ownerPercentage, &row.category, &row.createdAt, &row.updatedAt,
		&row.groupName, &row.splitsJSON,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domains.NewNotFoundError("cost not found")
		}
		return nil, err
	}

	return costRowToDomain(row)
}

func (r *costRepository) FindStats(ctx context.Context, userID string, filters domains.CostFilters) (*domains.CostStats, error) {
	conditions, args, _ := buildCostConditions("", userID, filters)
	where := strings.Join(conditions, " AND ")

	var thisMonth, inSplits, solo float64

	err := r.db.QueryRowContext(ctx, fmt.Sprintf(`
		SELECT
			COALESCE(SUM(total_value), 0),
			COALESCE(SUM(CASE WHEN group_id IS NOT NULL THEN total_value ELSE 0 END), 0),
			COALESCE(SUM(CASE WHEN group_id IS NULL THEN total_value ELSE 0 END), 0)
		FROM cost_entities
		WHERE %s
	`, where), args...).Scan(&thisMonth, &inSplits, &solo)
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, fmt.Sprintf(`
		SELECT category, COALESCE(SUM(total_value), 0)
		FROM cost_entities
		WHERE %s
		GROUP BY category
	`, where), args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rawBreakdown []domains.CostCategoryBreakdown
	var grandTotal float64

	for rows.Next() {
		var b domains.CostCategoryBreakdown
		if err := rows.Scan(&b.Category, &b.Total); err != nil {
			return nil, err
		}
		grandTotal += b.Total
		rawBreakdown = append(rawBreakdown, b)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	for i := range rawBreakdown {
		if grandTotal > 0 {
			rawBreakdown[i].Percentage = math.Round((rawBreakdown[i].Total/grandTotal)*100*100) / 100
		}
	}

	return &domains.CostStats{
		ThisMonth:  thisMonth,
		InSplits:   inSplits,
		Solo:       solo,
		ByCategory: rawBreakdown,
	}, nil
}

func (r *costRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM cost_entities WHERE id = $1`, id)
	return err
}
