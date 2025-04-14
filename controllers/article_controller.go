package controllers

import (
	"MyBlog/global"
	"MyBlog/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateArticle(ctx *gin.Context) {
	var article models.Article
	if err := ctx.ShouldBind(&article); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if err := global.Db.AutoMigrate(&article); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if err := global.Db.Create(&article).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusCreated, article)
}

func GetArticle(ctx *gin.Context) {
	var articles models.Article
	if err := global.Db.First(&articles).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, articles)
}
