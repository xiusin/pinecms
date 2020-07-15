package frontend

func (c *IndexController) Click() {

	//aid, _ := c.Ctx().GetInt64("aid")
	//tid, _ := c.Ctx().GetInt64("tid")
	//if aid < 1 || tid < 1 {
	//	c.Ctx().Abort(http.StatusNotFound)
	//	return
	//}
	//clickCache := fmt.Sprintf("click_%d_%d", tid, aid)
	//info := c.Ctx().GetCookie(clickCache)
	//if len(info) == 0 {
	//	res, err := di.MustGet("orm").(*xorm.Engine).Table(models.NewCategoryModel().GetTable(tid)).ID(aid).Incr("visit_count").Exec()
	//	if err != nil {
	//		logger.Error("无法更新点击数据", err)
	//		return
	//	}
	//	if affe, _ := res.RowsAffected(); affe > 0 {
	//		c.Ctx().SetCookie(clickCache, "1", 0)
	//	}
	//}
}

func (c *IndexController) GetClick() {
	//aid, _ := c.Ctx().GetInt64("aid")
	//tid, _ := c.Ctx().GetInt64("tid")
	//if aid < 1 || tid < 1 {
	//	c.Ctx().Abort(http.StatusNotFound)
	//	return
	//}
	//fmt.Println(di.MustGet("orm").(*xorm.Engine).Table(models.NewCategoryModel().GetTable(tid)).ID(aid).Select("visit_count").QueryString())
}