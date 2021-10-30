package controllers

const ServiceConfig = "pinecms.config"

const ServiceSiteConfig = "pinecms.site.config"

const ServiceICache = "cache.AbstractCache"

const ServiceTablePrefix = "pinecms.table_prefix"

const ServiceJetEngine = "pinecms.jet"

const ServiceXorm = "*xorm.Engine"

const ServiceCasbinEnforcer = "pinecms.casbin.enforcer"

const ServiceCasbinClearPolicy = "pinecms.casbin.enforcer.policy.clear"

const ServiceCasbinAddPolicy = "pinecms.casbin.enforcer.policy.add"

const ServiceUploader = "pinecms.uploader"

const ServiceUploaderEngine = "pinecms.uploader.%s"

const ServiceApplication = "pinecms.application"

const ServiceBackendRouter = "pinecms.router.backend"

// 允许插件自动注册上传驱动 并注册服务进DI
// DI内自动获取选中 (根据驱动名称) 驱动
