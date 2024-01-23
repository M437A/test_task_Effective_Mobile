package api

import (
	rep "name_service/iternal/repository"
)

func Delete(id int64) error {
	return rep.DeleteByID(id)
}
