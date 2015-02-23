package models

import (
    "github.com/revel/revel"
    "github.com/coopernurse/gorp"
    "time"
)

type Image struct {
    Id              int64    `db:"id" json:"id"`
    CarId           int64    `db:"car_id" json:"car_id"`
    Location        string   `db:"location" json:"location"`
    Name            string   `db:"name" json:"name"`
    Description     string   `db:"description" json:"description"`
    CreatedAt       time.Time  `db:"created_at" json:"created_at"`
    ModifiedAt      time.Time  `db:"modified_at" json:"modified_at"`
}


func (i *Image) Validate(v *revel.Validation) {
    v.Check(i.Name, revel.ValidRequired(), revel.ValidMaxSize(500))
    v.Check(i.Location, revel.ValidRequired(), revel.ValidMaxSize(1000))
    v.Check(i.Description, revel.ValidMaxSize(1000))
    v.Check(i.CarId, revel.ValidRequired())
}

func (i *Image) PreInsert(_ gorp.SqlExecutor) error {
  i.CreatedAt = time.Now()
  i.ModifiedAt = time.Now()

  return nil
}