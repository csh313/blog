package models

import "server/models/ctype"

// AuthModel 用户表
type UserModel struct {
	MODEL
	NickName       string           `gorm:"size:36" json:"nick_name"`                                                              //昵称
	UserName       string           `gorm:"size:36" json:"user_name"`                                                              //用户名
	Password       string           `gorm:"size:128" json:"-"`                                                                     //查询用户密码不显示                                                       //密码
	Avatar         string           `gorm:"size:256" json:"avatar_id"`                                                             //头像id
	Email          string           `gorm:"size:128" json:"email"`                                                                 //邮箱
	Tel            string           `gorm:"size:18" json:"tel"`                                                                    //手机号
	Addr           string           `gorm:"size:64" json:"addr"`                                                                   //地址
	Token          string           `gorm:"size:648" json:"token"`                                                                 //token其他平台的唯一验证id
	IP             string           `gorm:"size:20" json:"ip"`                                                                     //ip
	Role           ctype.Role       `gorm:"size:4;default:1" json:"role"`                                                          //权限 1管理员 2普通用户 3游客 4封禁
	SignStatus     ctype.SignStatus `gorm:"type=smallint(6)" json:"sign_status"`                                                   //注册来源
	ArticleModels  []ArticleModel   `gorm:"foreignKey:UserID" json:"-"`                                                            //发布的文章
	CollectsModels []ArticleModel   `gorm:"many2many:user_collect_models;joinForeignKey:UserID;JoinReferences:ArticleID" json:"-"` //收藏的文章
}

/*
在 GORM 中，你提供的代码片段展示了两种不同的数据库关系：

ArticleModels 字段是一个 ArticleModel 结构体的切片，
通过 gorm:"size:36" 标签设置了该字段在数据库中的大小为36字节。
这个标签通常用于 varchar 或类似类型的字段，但在这里它对于切片类型似乎没有直接意义，
因为切片的大小是由其内部元素的数量决定的，而不是固定的大小。json:"-" 表示在 JSON 序列化时忽略该字段。

CollectsModels 字段是一个 CollectsModel 结构体的切片，
它表示的是多对多关系。这里的 GORM 标签说明如下：

Go
gorm:"many2many:auth2_collects;joinForeignKey:AuthID;joinReferences:ArticleID;"
many2many:auth2_collects 表明这是一个多对多关系，并且使用名为 auth2_collects 的中间关联表。
joinForeignKey:AuthID 指定当前结构体（可能是某个用户或其他身份认证实体）的外键字段是 AuthID。
joinReferences:ArticleID 指定被关联的 CollectsModel 所对应的 ArticleModel 的外键字段是 ArticleID。
这样设置后，GORM 会在数据库中自动维护一个 auth2_collects 表，
其中包含 AuthID 和 ArticleID 两列，分别记录了用户身份和文章 ID，从而实现了用户收藏文章的多对多关系。

总结一下，这段代码描述了一个模型，其中包含了发布的文章列表（ArticleModels），以及通过多对多关系实现的收藏列表（CollectsModels）。
*/
/*
在数据库设计中，使用多对多（Many-to-Many）映射关系的原因通常是当一个实体可以与多个其他实体相关联，
而每个其他实体也可以与多个这个实体相关联时。具体到您给出的例子：

ArticleModels 可能代表文章模型，一篇文章可以被多个用户收藏。
CollectsModels 可能代表收藏模型，每个用户可以收藏多篇文章。
在这种情况下，如果我们只使用一对多关系，那么就需要在文章或者用户模型中添加冗余信息来记录所有的收藏关系，
显然这是不合适的。因此，引入一个中间表（在这个例子中是 auth2_collects 表）来作为多对多关系的桥梁，这样就能简洁且高效地表达这种“一篇文章可以被多个用户收藏，一个用户也可以收藏多篇文章”的关系。

中间表 auth2_collects 存储了用户的标识（AuthID）和文章的标识（ArticleID），
通过这两个外键就可以轻松查询出某个用户收藏的所有文章，或者查询出某篇文章被哪些用户收藏了。
*/
