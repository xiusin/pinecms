package tables

type Content struct {
	Id          int64  `form:"id" json:"id"`
	Catid       int64  `form:"catid" json:"catid"`
	Title       string `form:"title" json:"title"`
	Thumb       string `form:"thumb" json:"thumb"`
	Keywords    string `form:"keywords" json:"keywords"`
	Description string `form:"description" json:"description"`
	Content     string `form:"content" json:"content"`
	Listorder   int64  `form:"listorder" json:"listorder"`
	Status      int64  `form:"status" json:"status"`
	Recommend   int64  `form:"recommend" json:"recommend"`
	PwdType     int64  `form:"pwd_type" json:"pwd_type"`
	Money       int64  `form:"money" json:"money"`
	Userid      int64  `form:"userid" json:"userid"`
	CreatedAt   int64  `form:"created_at" json:"created_at"`
	UpdatedAt   int64  `form:"updated_at" json:"updated_at"`
	DeletedAt   int64  `form:"deleted_at" json:"deleted_at"`
	SourceUrl   string `form:"source_url" json:"source_url"`
	SourcePwd   string `form:"source_pwd" json:"source_pwd"`
	Catids      string `form:"catids" json:"catids"`
	Tags        string `form:"tags" json:"tags"`
}
