package controllers

const CacheTheme = "theme"
const CacheStatistics = "pinecms_persist.statistics"
const CacheRefer = "statistics_refer"
const CacheMemCollect = "pinecms.mem.collect"
const CacheAdminMenuByRoleIdAndMenuId = "pinecms.admin_menu_%d_%d"
const CacheAdminPriv = "pinecms.admin_priv_%d"
const CacheSetting = "pinecms.setting"
const CacheFeTplList = "pinecms.fe_tpl_list"
const CacheDocumentModelPrefix = "pinecms.document_model.%d"
const CacheModels = "pinecms.models"
const CacheCategories = "pinecms.categories"
const CacheCategoryPosPrefix = "pinecms.category.pos.%d"
const CacheCategoryInfoPrefix = "pinecms.category.%d"
const CacheCategoryContentPrefix = "pinecms.content.%d_%d"

// CacheKeyAll 缓存管理迭代
var CacheKeyAll = []map[string]interface{}{
	{CacheTheme: "主题"},
}
