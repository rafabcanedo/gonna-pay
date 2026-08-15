package repository

import (
	"database/sql"
	"math"
	"time"

	"github.com/google/uuid"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/domains"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/repository/entity"
)

type CostRepository interface {
	Create(domain domains.CostDomainInterface, memberIDs []string) (domains.CostDomainInterface, error)
	Update(id string, domain domains.CostDomainInterface) (domains.CostDomainInterface, error)
	FindAll(userID string) ([]domains.CostDomainInterface, error)
	FindByID(id string) (domains.CostDomainInterface, error)
	Delete(id string) error
	GetGroupMemberIDs(groupID string) ([]string, error)
	GetGroupName(groupID string) (string, error)
}

type costRepository struct {
	db *sql.DB
}

func NewCostRepository(db *sql.DB) CostRepository {
	return &costRepository{db: db}
}

func (r *costRepository) Create(domain domains.CostDomainInterface, memberIDs []string) (domains.CostDomainInterface, error) {
	costID := uuid.New()
	now := time.Now()

	var groupID interface{}
	if domain.GetGroupID() != "" {
		groupID = domain.GetGroupID()
	}

	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	_, err = tx.Exec(
		`INSERT INTO cost_entities (id, user_id, group_id, cost_name, total_value, owner_percentage, category, created_at, updated_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		costID, domain.GetUserID(), groupID, domain.GetCostName(),
		domain.GetTotalValue(), domain.GetOwnerPercentage(), domain.GetCategory(), now, now,
	)
	if err != nil {
		return nil, err
	}

	memberPercentage := 0.0
	memberValue := 0.0
	if len(memberIDs) > 0 {
		memberPercentage = math.Round(((100-domain.GetOwnerPercentage())/float64(len(memberIDs)))*100) / 100
		memberValue = math.Round((domain.GetTotalValue()*memberPercentage/100)*100) / 100
	}

	for _, contactID := range memberIDs {
		_, err = tx.Exec(
			`INSERT INTO cost_split_entities (id, cost_id, contact_id, value, percentage, created_at, updated_at)
			 VALUES ($1, $2, $3, $4, $5, $6, $7)`,
			uuid.New(), costID, contactID, memberValue, memberPercentage, now, now,
		)
		if err != nil {
			return nil, err
		}
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return domains.NewCostDomainWithID(
		costID.String(),
		domain.GetUserID(),
		domain.GetGroupID(),
		domain.GetCostName(),
		domain.GetCategory(),
		domain.GetTotalValue(),
		domain.GetOwnerPercentage(),
		now,
		now,
		nil,
	), nil
}

func (r *costRepository) Update(id string, domain domains.CostDomainInterface) (domains.CostDomainInterface, error) {
	now := time.Now()

	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	_, err = tx.Exec(
		`UPDATE cost_entities SET cost_name = $1, total_value = $2, owner_percentage = $3, category = $4, updated_at = $5 WHERE id = $6`,
		domain.GetCostName(), domain.GetTotalValue(), domain.GetOwnerPercentage(), domain.GetCategory(), now, id,
	)
	if err != nil {
		return nil, err
	}

	var splitCount int
	if err := tx.QueryRow(`SELECT COUNT(*) FROM cost_split_entities WHERE cost_id = $1`, id).Scan(&splitCount); err != nil {
		return nil, err
	}

	if splitCount > 0 {
		memberPercentage := math.Round(((100-domain.GetOwnerPercentage())/float64(splitCount))*100) / 100
		memberValue := math.Round((domain.GetTotalValue()*memberPercentage/100)*100) / 100

		_, err = tx.Exec(
			`UPDATE cost_split_entities SET value = $1, percentage = $2, updated_at = $3 WHERE cost_id = $4`,
			memberValue, memberPercentage, now, id,
		)
		if err != nil {
			return nil, err
		}
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return r.FindByID(id)
}

func (r *costRepository) FindAll(userID string) ([]domains.CostDomainInterface, error) {
	rows, err := r.db.Query(`
		SELECT
			c.id, c.user_id, c.group_id, c.cost_name, c.total_value, c.owner_percentage, c.category, c.created_at, c.updated_at,
			g.name,
			COUNT(cs.id) AS split_count
		FROM cost_entities c
		LEFT JOIN group_entities g ON g.id = c.group_id
		LEFT JOIN cost_split_entities cs ON cs.cost_id = c.id
		WHERE c.user_id = $1
		GROUP BY c.id, g.name
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var costs []domains.CostDomainInterface
	for rows.Next() {
		var (
			e          entity.CostEntity
			groupIDRaw sql.NullString
			groupName  sql.NullString
			splitCount int
		)

		if err := rows.Scan(
			&e.ID, &e.UserID, &groupIDRaw, &e.CostName, &e.TotalValue, &e.OwnerPercentage, &e.Category, &e.CreatedAt, &e.UpdatedAt,
			&groupName,
			&splitCount,
		); err != nil {
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

		domain := domains.NewCostDomainWithID(
			e.ID.String(),
			e.UserID.String(),
			groupID,
			e.CostName,
			string(e.Category),
			e.TotalValue,
			e.OwnerPercentage,
			e.CreatedAt,
			e.UpdatedAt,
			nil,
		)
		domain.SetGroupName(groupNameStr)
		domain.SetSplitCount(splitCount)
		costs = append(costs, domain)
	}

	return costs, nil
}

func (r *costRepository) FindByID(id string) (domains.CostDomainInterface, error) {
	rows, err := r.db.Query(`
		SELECT
			c.id, c.user_id, c.group_id, c.cost_name, c.total_value, c.owner_percentage, c.category, c.created_at, c.updated_at,
			g.name,
			cs.id, cs.contact_id, ct.name, cs.value, cs.percentage
		FROM cost_entities c
		LEFT JOIN group_entities g ON g.id = c.group_id
		LEFT JOIN cost_split_entities cs ON cs.cost_id = c.id
		LEFT JOIN contact_entities ct ON ct.id = cs.contact_id
		WHERE c.id = $1
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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
		splits          []domains.SplitDomain
		found           bool
	)

	for rows.Next() {
		var (
			splitID      sql.NullString
			contactID    sql.NullString
			contactName  sql.NullString
			splitValue   sql.NullFloat64
			splitPercent sql.NullFloat64
		)

		if err := rows.Scan(
			&costID, &userID, &groupIDRaw, &costName, &totalValue, &ownerPercentage, &category, &createdAt, &updatedAt,
			&groupName,
			&splitID, &contactID, &contactName, &splitValue, &splitPercent,
		); err != nil {
			return nil, err
		}

		found = true

		if splitID.Valid {
			splits = append(splits, domains.SplitDomain{
				ID:          splitID.String,
				ContactID:   contactID.String,
				ContactName: contactName.String,
				Value:       splitValue.Float64,
				Percentage:  splitPercent.Float64,
			})
		}
	}

	if !found {
		return nil, domains.NewNotFoundError("cost not found")
	}

	groupID := ""
	if groupIDRaw.Valid {
		groupID = groupIDRaw.String
	}

	groupNameStr := ""
	if groupName.Valid {
		groupNameStr = groupName.String
	}

	domain := domains.NewCostDomainWithID(
		costID.String(), userID.String(), groupID, costName, category,
		totalValue, ownerPercentage, createdAt, updatedAt, splits,
	)
	domain.SetGroupName(groupNameStr)
	domain.SetSplitCount(len(splits))

	return domain, nil
}

func (r *costRepository) Delete(id string) error {
	_, err := r.db.Exec(`DELETE FROM cost_entities WHERE id = $1`, id)
	return err
}

func (r *costRepository) GetGroupMemberIDs(groupID string) ([]string, error) {
	rows, err := r.db.Query(
		`SELECT contact_id FROM group_member_entities WHERE group_id = $1`,
		groupID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var memberIDs []string
	for rows.Next() {
		var contactID string
		if err := rows.Scan(&contactID); err != nil {
			return nil, err
		}
		memberIDs = append(memberIDs, contactID)
	}

	return memberIDs, nil
}

func (r *costRepository) GetGroupName(groupID string) (string, error) {
	var name string
	err := r.db.QueryRow(
		`SELECT name FROM group_entities WHERE id = $1`,
		groupID,
	).Scan(&name)
	if err != nil {
		return "", err
	}
	return name, nil
}
