package repository

import (
	"database/sql"
	"encoding/json"
	"errors"
	"math"
	"time"

	"github.com/google/uuid"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/domains"
)

type CostRepository interface {
	Create(cost *domains.Cost, members []domains.Member) (*domains.Cost, error)
	Update(id string, cost *domains.Cost) (*domains.Cost, error)
	FindAll(userID string, limit, offset int) ([]*domains.Cost, int64, error)
	FindByID(id string) (*domains.Cost, error)
	Delete(id string) error
	GetGroupByID(groupID string) (*domains.Group, error)
	GetGroupMembers(groupID string) ([]domains.Member, error)
	FindStats(userID string) (*domains.CostStats, error)
}

type costRepository struct {
	db *sql.DB
}

func NewCostRepository(db *sql.DB) CostRepository {
	return &costRepository{db: db}
}

func (r *costRepository) Create(cost *domains.Cost, members []domains.Member) (*domains.Cost, error) {
	costID := uuid.New()
	now := time.Now()

	var groupID interface{}
	if cost.GroupID != "" {
		groupID = cost.GroupID
	}

	var splits []domains.Split
	var splitsJSON []byte

	if len(members) > 0 {
		memberPercentage := math.Round(((100-cost.OwnerPercentage)/float64(len(members)))*100) / 100
		memberValue := math.Round((cost.TotalValue*memberPercentage/100)*100) / 100

		splits = make([]domains.Split, len(members))
		for i, m := range members {
			splits[i] = domains.Split{
				ID:          uuid.New().String(),
				ContactID:   m.ID,
				ContactName: m.Name,
				Value:       memberValue,
				Percentage:  memberPercentage,
			}
		}

		var err error
		splitsJSON, err = json.Marshal(splits)
		if err != nil {
			return nil, err
		}
	}

	_, err := r.db.Exec(
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
		splits,
	), nil
}

func (r *costRepository) Update(id string, cost *domains.Cost) (*domains.Cost, error) {
	now := time.Now()

	var splitsJSON []byte
	if len(cost.Splits) > 0 {
		memberPercentage := math.Round(((100-cost.OwnerPercentage)/float64(len(cost.Splits)))*100) / 100
		memberValue := math.Round((cost.TotalValue*memberPercentage/100)*100) / 100

		for i := range cost.Splits {
			cost.Splits[i].Percentage = memberPercentage
			cost.Splits[i].Value = memberValue
		}

		var err error
		splitsJSON, err = json.Marshal(cost.Splits)
		if err != nil {
			return nil, err
		}
	}

	_, err := r.db.Exec(
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

func (r *costRepository) FindAll(userID string, limit, offset int) ([]*domains.Cost, int64, error) {
	var total int64

	err := r.db.QueryRow(`SELECT COUNT(*) FROM cost_entities WHERE user_id = $1`, userID).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	rows, err := r.db.Query(`
		SELECT
			c.id, c.user_id, c.group_id, c.cost_name, c.total_value, c.owner_percentage, c.category, c.created_at, c.updated_at,
			g.name,
			COALESCE(jsonb_array_length(c.splits), 0) AS split_count
		FROM cost_entities c
		LEFT JOIN group_entities g ON g.id = c.group_id
		WHERE c.user_id = $1
		LIMIT $2 OFFSET $3
	`, userID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var costs []*domains.Cost
	for rows.Next() {
		var (
			id              uuid.UUID
			userIDField     uuid.UUID
			groupIDRaw      sql.NullString
			costName        string
			totalValue      float64
			ownerPercentage float64
			category        string
			createdAt       time.Time
			updatedAt       time.Time
			groupName       sql.NullString
			splitCount      int
		)

		if err := rows.Scan(
			&id, &userIDField, &groupIDRaw, &costName, &totalValue, &ownerPercentage, &category, &createdAt, &updatedAt,
			&groupName,
			&splitCount,
		); err != nil {
			return nil, 0, err
		}

		groupID := ""
		if groupIDRaw.Valid {
			groupID = groupIDRaw.String
		}

		groupNameStr := ""
		if groupName.Valid {
			groupNameStr = groupName.String
		}

		costs = append(costs, &domains.Cost{
			ID:              id.String(),
			UserID:          userIDField.String(),
			GroupID:         groupID,
			GroupName:       groupNameStr,
			CostName:        costName,
			TotalValue:      totalValue,
			OwnerPercentage: ownerPercentage,
			Category:        category,
			CreatedAt:       createdAt,
			UpdatedAt:       updatedAt,
			SplitCount:      splitCount,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return costs, total, nil
}

func (r *costRepository) FindByID(id string) (*domains.Cost, error) {
	var (
		costID          uuid.UUID
		userID          uuid.UUID
		groupIDRaw      sql.NullString
		costName        string
		totalValue      float64
		ownerPercentage float64
		category        string
		createdAt       time.Time
		updatedAt       time.Time
		groupName       sql.NullString
		splitsJSON      []byte
	)

	err := r.db.QueryRow(`
		SELECT
			c.id, c.user_id, c.group_id, c.cost_name, c.total_value, c.owner_percentage, c.category, c.created_at, c.updated_at,
			g.name, c.splits
		FROM cost_entities c
		LEFT JOIN group_entities g ON g.id = c.group_id
		WHERE c.id = $1
	`, id).Scan(
		&costID, &userID, &groupIDRaw, &costName, &totalValue, &ownerPercentage, &category, &createdAt, &updatedAt,
		&groupName, &splitsJSON,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domains.NewNotFoundError("cost not found")
		}
		return nil, err
	}

	groupID := ""
	if groupIDRaw.Valid {
		groupID = groupIDRaw.String
	}

	groupNameStr := ""
	if groupName.Valid {
		groupNameStr = groupName.String
	}

	var splits []domains.Split
	if splitsJSON != nil {
		if err := json.Unmarshal(splitsJSON, &splits); err != nil {
			return nil, err
		}
	}

	return &domains.Cost{
		ID:              costID.String(),
		UserID:          userID.String(),
		GroupID:         groupID,
		GroupName:       groupNameStr,
		CostName:        costName,
		TotalValue:      totalValue,
		OwnerPercentage: ownerPercentage,
		Category:        category,
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
		SplitCount:      len(splits),
		Splits:          splits,
	}, nil
}

func (r *costRepository) FindStats(userID string) (*domains.CostStats, error) {
	var thisMonth, inSplits, solo float64

	err := r.db.QueryRow(`
		SELECT
			COALESCE(SUM(total_value), 0),
			COALESCE(SUM(CASE WHEN group_id IS NOT NULL THEN total_value ELSE 0 END), 0),
			COALESCE(SUM(CASE WHEN group_id IS NULL THEN total_value ELSE 0 END), 0)
		FROM cost_entities
		WHERE user_id = $1
		  AND DATE_TRUNC('month', created_at) = DATE_TRUNC('month', NOW())
	`, userID).Scan(&thisMonth, &inSplits, &solo)
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(`
		SELECT category, COALESCE(SUM(total_value), 0)
		FROM cost_entities
		WHERE user_id = $1
		  AND DATE_TRUNC('month', created_at) = DATE_TRUNC('month', NOW())
		GROUP BY category
	`, userID)
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

func (r *costRepository) Delete(id string) error {
	_, err := r.db.Exec(`DELETE FROM cost_entities WHERE id = $1`, id)
	return err
}

func (r *costRepository) GetGroupByID(groupID string) (*domains.Group, error) {
	var id, ownerID uuid.UUID
	err := r.db.QueryRow(
		`SELECT id, owner_id FROM group_entities WHERE id = $1`,
		groupID,
	).Scan(&id, &ownerID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domains.NewNotFoundError("group not found")
		}
		return nil, err
	}
	return &domains.Group{ID: id.String(), OwnerID: ownerID.String()}, nil
}

func (r *costRepository) GetGroupMembers(groupID string) ([]domains.Member, error) {
	rows, err := r.db.Query(`
		SELECT gm.contact_id, c.name
		FROM group_member_entities gm
		JOIN contact_entities c ON c.id = gm.contact_id AND c.deleted_at IS NULL
		WHERE gm.group_id = $1
	`, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []domains.Member
	for rows.Next() {
		var m domains.Member
		if err := rows.Scan(&m.ID, &m.Name); err != nil {
			return nil, err
		}
		members = append(members, m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return members, nil
}
