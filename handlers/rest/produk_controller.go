package rest

import (
	"WallE/domains"
	"WallE/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type productController struct {
	services domains.ProductService
}

func (cont *productController) AmbilKategori(c echo.Context) error {
	kategori := cont.services.AmbilKategori()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"kode":     http.StatusOK,
		"kategori": kategori,
	})
}

func (cont *productController) TambahSaldo(c echo.Context) error {
	saldobaru, _ := strconv.Atoi(c.FormValue("saldo_baru"))
	kategoriid, _ := strconv.Atoi(c.FormValue("kategori_id"))
	err := cont.services.TambahSaldo(saldobaru, uint(kategoriid))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"kode":  http.StatusInternalServerError,
			"pesan": err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"kode":    http.StatusOK,
		"message": "sukses",
	})
}
func (cont *productController) AmbilProdukBerdasarkanKategori(c echo.Context) error {
	kategoriid, _ := strconv.Atoi(c.FormValue("kategori_id"))
	fmt.Println(kategoriid)
	produk := cont.services.AmbilProdukBerdasarkanKategori(uint(kategoriid))
	return c.JSON(http.StatusOK, map[string]interface{}{
		"kode":   http.StatusOK,
		"produk": produk,
	})
}

func (cont *productController) AmbilProdukBerdasarkanProviderKategori(c echo.Context) error {
	kategoriid, _ := strconv.Atoi(c.FormValue("kategori_id"))
	providerid, _ := strconv.Atoi(c.FormValue("provider_id"))
	produk := cont.services.AmbilProdukBerdasarkanProviderKategori(uint(kategoriid), uint(providerid))
	return c.JSON(http.StatusOK, map[string]interface{}{
		"kode":   http.StatusOK,
		"produk": produk,
	})
}

func (cont *productController) TambahProduk(c echo.Context) error {
	produk := models.Produk{}
	c.Bind(&produk)
	err := cont.services.TambahProduk(produk)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"kode":  http.StatusInternalServerError,
			"pesan": err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"kode":    http.StatusOK,
		"message": "sukses",
	})
}