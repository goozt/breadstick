package menu

import (
	"bytes"
	"encoding/gob"
	"errors"
	"math"
	"strings"

	"github.com/google/uuid"
	"goozt.org/breakstick/api/db"
)

type Item struct {
	ID       string  `json:"id"`
	Name     string  `validate:"required" json:"name"`
	Category string  `validate:"required" json:"category"`
	Price    float64 `validate:"required,numeric" json:"price"`
}

type Menu struct {
	Items []Item `json:"items"`
}

// Get menu store
var menuDB = db.Datastore.GetStore("menu")

// New Item instance
func NewItem() *Item {
	item := new(Item)

	// Create the uuid
	id, _ := uuid.NewUUID()
	item.ID = id.String()

	return item
}

// Copy menu item with new item
func (item *Item) Copy(newItem *Item) {
	if newItem.Name != "" {
		(*item).Name = strings.TrimSpace(newItem.Name)
	}
	if newItem.Price != 0.0 {
		(*item).Price = math.Round(newItem.Price*100) / 100
	}
	if newItem.Category != "" {
		(*item).Category = strings.TrimSpace(newItem.Category)
	}
}

// Encode item
func (item *Item) Encode() ([]byte, error) {
	if item == nil {
		return nil, errors.New("item not found")
	}
	var buff bytes.Buffer
	err := gob.NewEncoder(&buff).Encode(item)
	return buff.Bytes(), err
}

// Decode item
func Decode(data []byte, item *Item) error {
	buff := bytes.NewBuffer(data)
	return gob.NewDecoder(buff).Decode(item)
}
