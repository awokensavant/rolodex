// Package postgres contains PostgreSQL persistence models.
package postgres

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/awokensavant/rolodex/domain"
)

// VCard maps a domain vCard to a row in the vcards table.
// Data is stored as JSONB so repeated properties and extensions are preserved.
type VCard struct {
	ID   int64
	Data json.RawMessage
}

// NewVCard maps a domain vCard to its PostgreSQL representation.
func NewVCard(card domain.VCard) (VCard, error) {
	data, err := json.Marshal(card)
	if err != nil {
		return VCard{}, fmt.Errorf("marshal vCard: %w", err)
	}

	return VCard{Data: data}, nil
}

// Domain maps the PostgreSQL representation back to a domain vCard.
func (v VCard) Domain() (domain.VCard, error) {
	if len(v.Data) == 0 {
		return domain.VCard{}, errors.New("vCard data is empty")
	}

	var card domain.VCard
	if err := json.Unmarshal(v.Data, &card); err != nil {
		return domain.VCard{}, fmt.Errorf("unmarshal vCard: %w", err)
	}
	return card, nil
}
