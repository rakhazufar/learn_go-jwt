package productcontroller

import (
	"net/http"

	"github.com/rakhazufar/go-jwt/pkg/utils"
)

type Products struct {
	name        string
	price       float64
	description string
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	data := []map[string]interface{}{
		{
			"id": 1,
			"nameproduk": "anjay",
			"stok": 100,
		},
		{
			"id": 2,
			"nameproduk": "anjay",
			"stok": 200,
		},
		{
			"id": 3,
			"nameproduk": "anjay",
			"stok": 400,
		},
	}

	utils.SendJSONResponse(w, http.StatusOK, data)
}