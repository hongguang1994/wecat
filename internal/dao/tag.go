package dao

import (
	"wecat/internal/models"
	"wecat/pkg/app"
)

func (d *Dao) CountTag(name string, state uint8) (int64, error) {
	tag := models.Tag{Name: name, State: state}
	return tag.Count(d.engine)
}

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*models.Tag, error) {
	tag := models.Tag{Name: name, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateTag(name string, state uint8, createBy string) error {
	tag := models.Tag{
		Name:  name,
		State: state,
		Model: &models.Model{CreatedBy: createBy},
	}

	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(id uint32, name string, state uint8, moddifiedBy string) error {
	tag := models.Tag{
		Model: &models.Model{ID: id},
	}
	values := map[string]interface{}{
		"state":       state,
		"modified_by": moddifiedBy,
	}
	if name != "" {
		values["name"] = name
	}
	// tag := models.Tag{
	// 	Name:  name,
	// 	State: state,
	// 	Model: &models.Model{ID: id, ModifiedBy: moddifiedBy},
	// }
	return tag.Update(d.engine, values)
}

func (d *Dao) DeleteTag(id uint32) error {
	tag := models.Tag{Model: &models.Model{ID: id}}
	return tag.Delete(d.engine)
}
