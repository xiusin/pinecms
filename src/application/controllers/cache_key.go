package controllers

// NOTE: `pinecms.` 开始的会加入到清理数据缓存逻辑

const CacheTheme  = "theme"
const CacheMemCollect  = "pinecms.mem.collect"
const CacheAdminMenuByRoleIdAndMenuId = "pinecms.admin_menu_%d_%d"
const CacheAdminPriv = "pinecms.admin_priv_%d"
const CacheSetting = "pinecms.setting"
const CacheDocumentModelPrefix  = "pinecms.document_model.%d"
const CacheModels  = "pinecms.models"
const CacheCategories  = "pinecms.categories"
const CacheCategoryPosPrefix  = "pinecms.category.pos.%d"
const CacheCategoryInfoPrefix  = "pinecms.category.%d"
const CacheCategoryContentPrefix  = "pinecms.content.%d_%d"
