package controllers

import (
	"fmt"
	"hummingbird/databases"
	"hummingbird/models"

	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
)

func AddUpExls(c *gin.Context) {
	// 单文件
	name := c.PostForm("Name")
	DB := databases.GetDB() //到数据库
	var studenInfo = models.StudenTables{}
	file, _ := c.FormFile("file")
	path := "./studentFile/" + file.Filename
	c.SaveUploadedFile(file, path)

	excelFileName := path
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Printf("open failed: %s\n", err)
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for i := 0; i < 6; i++ {
				studenInfo.ClassName = name
				switch {
				case i == 1:
					studenInfo.Name = row.Cells[i].String()
				case i == 2:
					studenInfo.StudenID = row.Cells[i].String()
				case i == 3:
					studenInfo.QQNumble = row.Cells[i].String()
				case i == 4:
					studenInfo.Email = row.Cells[i].String()
				case i == 5:
					studenInfo.Address = row.Cells[i].String()
				}
			}
			result := DB.Create(&studenInfo)
			if result != nil {
				fmt.Println("存储成功")
			} else {
				fmt.Println("导入失败")
			}
		}
	}
}
