package handlers

import (
	"context"

	iris "github.com/kataras/iris/v12"

	strconv "strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	madmin "github.com/minio/madmin-go/v2"
	resph "github.com/rzrbld/adminio-api/response"
)

var GrSetStatus = func(ctx iris.Context) {
	var group = ctx.FormValue("group")
	var status = ctx.FormValue("status")

	if resph.CheckAuthBeforeRequest(ctx) {
		var status = madmin.GroupStatus(status)
		err = madmClnt.SetGroupStatus(context.Background(), group, status)
		var res = resph.DefaultResHandler(ctx, err)
		ctx.JSON(res)
	} else {
		ctx.JSON(resph.DefaultAuthError())
	}
}

var GrSetDescription = func(ctx iris.Context) {
	var group = ctx.FormValue("group")

	if resph.CheckAuthBeforeRequest(ctx) {
		grp, err := madmClnt.GetGroupDescription(context.Background(), group)
		var res = resph.BodyResHandler(ctx, err, grp)
		ctx.JSON(res)
	} else {
		ctx.JSON(resph.DefaultAuthError())
	}
}

var GrUpdateMembers = func(ctx iris.Context) {
	gar := madmin.GroupAddRemove{}
	gar.Group = ctx.FormValue("group")
	if ctx.FormValue("members") != "" {
		gar.Members = strings.Split(ctx.FormValue("members"), ",")
	}

	gar.IsRemove, err = strconv.ParseBool(ctx.FormValue("IsRemove"))
	if err != nil {
		log.Errorln(err)
		ctx.JSON(iris.Map{"error": err.Error()})
	}

	if resph.CheckAuthBeforeRequest(ctx) {
		err = madmClnt.UpdateGroupMembers(context.Background(), gar)
		var res = resph.DefaultResHandler(ctx, err)
		ctx.JSON(res)
	} else {
		ctx.JSON(resph.DefaultAuthError())
	}

}

var GrList = func(ctx iris.Context) {
	lg, err := madmClnt.ListGroups(context.Background())
	var res = resph.BodyResHandler(ctx, err, lg)
	ctx.JSON(res)
}
