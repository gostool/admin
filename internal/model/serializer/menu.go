package serializer

type Menu struct {
	OrmCommon
	Title       string `json:"title"     dc:"名称"  `      // 名称
	Icon        string `json:"icon"       dc:"图标" `      // 图标
	KeepAlive   int    `json:"keepAlive"  dc:"是否缓存" `    // 是否缓存
	DefaultMenu int    `json:"defaultMenu" dc:"是否基础路由" ` // 是否基础路由
	Hidden      int    `json:"hidden"      dc:"是否隐藏" `   //
	Sort        int    `json:"sort"       dc:"排序" `      //
	ParentId    int    `json:"parentId"   dc:"父ID" `     //
	Name        string `json:"name"       dc:"名字"  `     //
	Path        string `json:"path"      dc:"路径"   `     //
	Component   string `json:"component"   dc:"组件路径" `   //
}

type MenuDetail struct {
	Menu
	Children []*MenuDetail `json:"children"`
}
