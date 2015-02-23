package models

import (
    "github.com/revel/revel"
    "github.com/coopernurse/gorp"
    "time"
)

type Car struct {
    Id              int64    `db:"id" json:"id"`
    Name            string   `db:"name" json:"name"`
    CreatedAt       time.Time  `db:"created_at" json:"created_at"`
    ModifiedAt      time.Time  `db:"modified_at" json:"modified_at"`
}


func (c *Car) Validate(v *revel.Validation) {
    v.Check(c.Name, revel.ValidRequired(), revel.ValidMaxSize(500))
}

func (c *Car) PreInsert(_ gorp.SqlExecutor) error {
  c.CreatedAt = time.Now()
  c.ModifiedAt = time.Now()

  return nil
}