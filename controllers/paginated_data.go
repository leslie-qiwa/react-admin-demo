package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
)

type paginateParam struct {
	offset    int64
	limit     int64
	order     string
	sort      string
	startDate time.Time
}

func parseQueryPagination(ctx *gin.Context) (*paginateParam, error) {
	start, err := strconv.Atoi(ctx.Query("_start"))
	if err != nil {
		return nil, err
	}
	end, err := strconv.Atoi(ctx.Query("_end"))
	if err != nil {
		return nil, err
	}
	order := ""
	switch strings.ToLower(ctx.Query("_order")) {
	case "":
		// do nothing
	case "asc":
		order = "asc"
	case "desc":
		order = "desc"
	default:
		return nil, errors.New("not supported _order param")
	}
	if ctx.Query("date_gte") != "" {
		fmt.Println(ctx.Query("date_gte"))
	}
	return &paginateParam{
		offset: int64(start), limit: int64(end - start),
		order: order, sort: ctx.Query("_sort")}, nil
}
