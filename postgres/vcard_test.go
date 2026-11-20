package postgres

import (
	"reflect"
	"testing"

	"github.com/awokensavant/rolodex/domain"
)

func TestVCardMappingRoundTrip(t *testing.T) {
	card := domain.VCard{
		Version: "4.0",
		FormattedName: []domain.Property{
			{Name: "FN", Value: "Ada Lovelace"},
			{Name: "FN", Value: "Countess of Lovelace"},
		},
		Email: []domain.Property{{
			Name:       "EMAIL",
			Value:      "ada@example.com",
			Parameters: map[string][]string{"TYPE": {"work", "internet"}},
		}},
		Extensions: []domain.Property{{Name: "X-EXAMPLE", Value: "kept"}},
	}

	row, err := NewVCard(card)
	if err != nil {
		t.Fatalf("NewVCard() error = %v", err)
	}
	got, err := row.Domain()
	if err != nil {
		t.Fatalf("Domain() error = %v", err)
	}
	if !reflect.DeepEqual(got, card) {
		t.Errorf("Domain() = %#v, want %#v", got, card)
	}
}

func TestVCardDomainRejectsInvalidData(t *testing.T) {
	tests := []struct {
		name string
		data []byte
	}{
		{name: "empty", data: nil},
		{name: "invalid JSON", data: []byte("{")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := (VCard{Data: tt.data}).Domain()
			if err == nil {
				t.Fatal("Domain() error = nil, want error")
			}
		})
	}
}
