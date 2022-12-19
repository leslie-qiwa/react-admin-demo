package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/leslie-qiwa/react-admin-demo/helpers"
	"github.com/leslie-qiwa/react-admin-demo/infra/database"
	"strconv"
	"strings"
	"time"
)

type RAController struct{}

type queryParam struct {
	categoryID int
	offset     int64
	limit      int64
	order      string
	sort       string
	status     string
	startDate  time.Time
}

func parseQueryPagination(ctx *gin.Context) (*queryParam, error) {
	var (
		start, end int
		err        error
	)
	if ctx.Query("_start") != "" {
		start, err = strconv.Atoi(ctx.Query("_start"))
		if err != nil {
			return nil, err
		}
	}
	if ctx.Query("_end") != "" {
		end, err = strconv.Atoi(ctx.Query("_end"))
		if err != nil {
			return nil, err
		}
	}
	param := &queryParam{
		offset: int64(start), limit: int64(end - start),
		sort:   ctx.Query("_sort"),
		status: ctx.Query("status"),
	}
	switch strings.ToLower(ctx.Query("_order")) {
	case "":
		// do nothing
	case "asc":
		param.order = "asc"
	case "desc":
		param.order = "desc"
	default:
		return nil, errors.New("not supported _order param")
	}
	if ctx.Query("date_gte") != "" {
		fmt.Println(ctx.Query("date_gte"))
	}

	if ctx.Query("category_id") != "" {
		param.categoryID, err = strconv.Atoi(ctx.Query("category_id"))
		if err != nil {
			return nil, err
		}
	}

	return param, nil
}

func mkPaginateParam(query *queryParam) *helpers.Param {
	param := &helpers.Param{
		DB:           database.DB,
		Offset:       query.offset,
		Limit:        query.limit,
		ForeignTable: "Baskets",
	}
	if query.order != "" {
		if query.sort == "" {
			param.OrderBy = "id " + query.order
		} else {
			param.OrderBy = query.sort + " " + query.order
		}
	}
	if query.status != "" {
		param.Query = "status = '" + query.status + "'"
	} else if query.categoryID != 0 {
		param.Query = "category_id = " + strconv.Itoa(query.categoryID) + ""
	}
	return param
}
