package repository

import "todo/models"

var Context, err = models.PostgresDataBase{}.Connect()