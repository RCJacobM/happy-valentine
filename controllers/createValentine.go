package controllers

import (
	"hash/fnv"
	"math/rand"
	"net/http"
	"os"
	"valentoins/initializers"
	"valentoins/models"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
)

func generateId(seed string) string {
	h := fnv.New64()
	h.Write([]byte(seed))

	seedInt := int64(h.Sum64())

	for _, b := range []byte(seed) {
		seedInt += int64(b)
	}

	// Seed the random generator
	r := rand.New(rand.NewSource(seedInt))

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, 6)
	for i := range result {
		result[i] = charset[r.Intn(len(charset))]
	}
	return string(result)
}

func CreateValentine(ctx *gin.Context) {
	host := ctx.Request.Host
	// get both names
	var body struct {
		Receipient string `form:"receipient" binding:"required"`
		Sender     string `form:"sender" binding:"required"`
	}

	if ctx.ShouldBind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "retard",
		})
		return
	}

	//create id

	generated_id := generateId(body.Receipient + body.Sender)

	//push to db

	valentine := models.Valentines{Sender: body.Sender, Receipient: body.Receipient, CreateId: generated_id}
	res := initializers.DB.Create(&valentine)

	if res.Error != nil {

		if pgErr, ok := res.Error.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" { // send link here
				ctx.HTML(http.StatusOK, "card_ready.html", gin.H{"base": os.Getenv("base"), "id": generated_id, "host": host})
			} else {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message": "nop2",
				})
			}

		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "nop",
			})
		}

		return
	}

	//send link

	ctx.HTML(http.StatusOK, "card_ready.html", gin.H{"id": generated_id, "base": os.Getenv("base"), "host": host, "receipient": body.Receipient, "sender": body.Sender})

}
