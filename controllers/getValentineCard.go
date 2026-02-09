package controllers

import (
	"net/http"
	"os"
	"valentoins/initializers"
	"valentoins/models"

	"github.com/gin-gonic/gin"
)

func GetValentineCard(ctx *gin.Context) {
	var param struct {
		CreateId string `form:"id"`
	}
	if err := ctx.ShouldBindQuery(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var valentine models.Valentines
	initializers.DB.First(&valentine, "create_id = ?", param.CreateId)
	if valentine.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no"})
		return
	}

	ctx.HTML(200, "card.html", gin.H{"receipient": valentine.Receipient, "sender": valentine.Sender, "id": valentine.CreateId, "base": os.Getenv("base")})

}

func Result(ctx *gin.Context) {
	var body struct {
		CreateId        string `form:"id"`
		ClickedYesFirst string `form:"clickyesfirst"`
	}

	if ctx.ShouldBind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "nope",
		})
		return
	}

	var valentine models.Valentines
	initializers.DB.First(&valentine, "create_id = ?", body.CreateId)
	if valentine.ClickYesFirst != "" {
		return
	}

	initializers.DB.Model(&valentine).Where("create_id = ?", body.CreateId).Update("click_yes_first", body.ClickedYesFirst)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
