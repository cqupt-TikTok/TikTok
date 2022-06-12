# TikTok

技术框架：Gin+mysql+七牛云

组织：https://github.com/cqupt-TikTok

项目仓库：https://github.com/cqupt-TikTok/TikTok

接口文档地址：https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18901444

极简抖音APP地址：https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7

注意事项：克隆仓库运行前请先修改config配置文件，否则可能运行失败

### 一、任务分工

| 分工   | 任务                                                      |                |
| ------ | --------------------------------------------------------- | -------------- |
| 板块一 | 数据库支持：设计，变更，维护，支持                        | 袁钰钿，赵彬文 |
| 板块三 | 对象储存(七牛/本地)：视频投稿接口+发布列表接口+视频流接口 | 袁钰钿，赵彬文 |
| 板块二 | 用户：用户注册，登录，信息接口                            | 冼文杰         |
| 板块五 | 扩展接口二                                                | 汪伯伦         |
| 板块四 | 扩展接口一                                                | 张政余，范国慷 |
| 板块六 | 资料整理，最后项目答辩                                    | 张政余，范国慷 |

### 二、项目结构

| 目录                | 功能                                   |
| ------------------- | -------------------------------------- |
| api                 | 放各类路由                             |
| apifubc             | 放各类路由需要调用的函数               |
| config              | 配置文件                               |
| dbfunc              | 放数据库操作相关函数                   |
| log                 | 日志文件目录有相关日志记录文件和logger |
| model               | 模板文件，各类结构体                   |
| router              | 存放路由api                            |
| storage             | 存放储存相关函数，包括七牛和gorm       |
| util                | 工具类                                 |
| main.go             | 主函数，入口函数                       |
| apiLatest_log.log   | 路由相关的最新日志                     |
| mysqlLatest_log.log | MySQL相关的最新日志                    |

![image](https://user-images.githubusercontent.com/93390152/173222370-90ed32c0-976b-4f37-b2f4-6f8dfa0eb699.png)



### 三、项目详解

##### 3.1、model目录

此目录用于存放模板文件，模板结构体

由于响应结构体和数据库结构体存在一些差异，特将两种分开建立，使结构更加清晰明了

每个文件都包含了一个gorm结构体和一个json响应结构体，以及一个gorm结构体转换为json响应结构的方法：ToResp（），调用方便

![image](https://user-images.githubusercontent.com/93390152/173222379-f614bcc4-a762-4a45-b579-754c95481fa1.png)




comment.go：评论model文件

relation.go：关系model文件，包括FollowRelation用户关注和FavoriteVideoRelation视频点赞

response.go：响应model文件各类json响应结构体

user.go：用户model文件

video.go：视频model文件

##### 3.2、storage目录

此目录用于存放储存相关文件

![image](https://user-images.githubusercontent.com/93390152/173222388-74e63213-d52f-4481-ba4b-b18c42a0d46a.png)


gorm.go：musql初始化文件

qiniu.go：七牛云储存相关文件

##### 3.3、util目录

此目录用于存放各类工具文件

![image](https://user-images.githubusercontent.com/93390152/173222393-3b8a18bd-0273-47f5-912c-0e3dc5f4dea9.png)


encrypt.go：加密文件，采用哈希加盐的加密方式

token.go：由于本项目的token直接存放于url中比较简单，就没有另起目录存放jwt中间件，直接存放在了工具类文件里

##### 3.4、log目录

此目录用于存放日志文件，包括日志原文件和logger函数

![image](https://user-images.githubusercontent.com/93390152/173222396-46a6b806-e741-4b52-8378-09a5d5653eb1.png)

apiLog：存放api相关的日志源文件

mysqlLog：存放mysql相关的日志源文件

logger.go：日志记录相关函数

##### 3.5、config目录

此目录用于存放相关配置文件

![image](https://user-images.githubusercontent.com/93390152/173222398-8c8c19f5-0478-455f-99c5-fe13a75f3507.png)


config.go：包含MySQL和七牛云的配置文件

##### 3.6、dbfunc目录

此目录用于存放操作数据库相关的函数

![image](https://user-images.githubusercontent.com/93390152/173222403-df1fae81-b803-46d0-9940-6a6f9f494df6.png)


##### 3.7、apifunc目录

此目录用于存放api相关的逻辑处理函数

![image](https://user-images.githubusercontent.com/93390152/173222408-528f7f4d-8ccf-4a38-afd7-0ca5bede260d.png)


##### 3.8、api目录

此目录用于存放api路由

![image](https://user-images.githubusercontent.com/93390152/173222417-d0f54f6b-c113-41ee-9d7a-f3e6d2d37450.png)


##### 3.9、router目录

此目录用于存放各路由router

![image](https://user-images.githubusercontent.com/93390152/173222423-9f97d1e8-cfaf-4862-926b-8081727d631a.png)


### 四、设计思路

一下是我自认为比较好的设计思路

##### 4.1、结构体设计

我们将结构体分为了gorm结构体和json响应结构体，使两者更加清晰明了，在建立数据表时可以直接使用gorm.AutoMigrate()自动迁移表，避免了手动建表带来的一些麻烦。并且为每个gorm结构体实现了一个ToResp()方法，调用简单方便

例如：

~~~go
// User 用户
type User struct {
	gorm.Model
	Name          string `gorm:"column:name;type:varchar(20);not null"`     // 用户名称
	Password      string `gorm:"column:password;type:varchar(20);not null"` //用户密码
	FollowCount   int64  `gorm:"column:follow_count;type:int;default:0"`    // 关注总数
	FollowerCount int64  `gorm:"column:follower_count;type:int;default:0"`  // 粉丝总数
}

// UserResp 响应结构体
type UserResp struct {
	Id            uint   `json:"id"`             // 用户id
	Name          string `json:"name"`           // 用户名称
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
}

// ToResp 转化为响应结构体，默认关注
func (U User) ToResp() (UR UserResp) {
	UR.Id = U.ID
	UR.Name = U.Name
	UR.FollowCount = U.FollowerCount
	UR.FollowerCount = U.FollowerCount
	UR.IsFollow = true
	return UR
}

// IsFollowJudge 关注校验，视情况调用
func (UR *UserResp) IsFollowJudge(UserId uint) {
	var FR FollowRelation
	storage.DB.Where("follower_id = ? AND user_id = ?", UR.Id, UserId).First(&FR)
	if FR.Id <= 0 {
		(*UR).IsFollow = false
	}
}
~~~

##### 4.2、密码加密

使用哈希加盐的加密方式给用户的密码进行加密储存，避免了密码的明文储存，以防出现盗库用户信息暴露的风险

~~~go
//ScryptPw 密码加密
func ScryptPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 6, 66, 22, 222, 11} //可以自定义，不一定是这几个数字
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}
~~~

##### 4.3、数据一致性

在对某些具有关联性的数据操作（如点赞，关注）时，采用mysql事务，一旦发生错误，立即回滚，保证了数据的高一致性

例如：

~~~go
// AddFavoriteVideo 点赞
func AddFavoriteVideo(videoId, userId uint) error {
	var favoriteVideoRelation = model.FavoriteVideoRelation{
		Id:           0,
		VideoId:      videoId,
		UserId:       userId,
		FavoriteDate: time.Now(),
	}
	//开始事务
	tx := storage.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	//查询点赞视频是否存在
	var v model.Video
	if err := tx.Model(&model.Video{}).Where("id = ?", videoId).First(&v).Error; err != nil {
		tx.Rollback()
		return err
	}
	//查询是否已经点赞
	var FVR model.FavoriteVideoRelation
	tx.Model(&model.FavoriteVideoRelation{}).Where("video_id = ? and user_id = ?", videoId, userId).First(&FVR)
	if FVR.Id > 0 {
		tx.Rollback()
		return errors.New("重复点赞")
	}
	//视频点赞总数favorite_count+1
	if err := tx.Model(&model.Video{}).Where("id = ? ", videoId).Update("favorite_count", gorm.Expr("favorite_count+ ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}
	//点赞表中写入数据
	if err := tx.Create(&favoriteVideoRelation).Error; err != nil {
		tx.Rollback()
		return err
	}
	//提交事务
	return tx.Commit().Error

}
~~~

##### 4.4、apifunc目录的建立

将api的一些处理逻辑操作放入其中，把api和操作函数分离，使两者更加简明，方便错误排查。

例如：

api 中user：用户注册

~~~go
// Register 用户注册
func Register(c *gin.Context) {
	var resp model.UserResponse
	var err error
	resp, err = apifunc.Register(c)   //此调用，避免了函数冗杂
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = "注册失败:" + err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.StatusCode = 0
	resp.StatusMsg = "注册成功"
	c.JSON(http.StatusOK, resp)
	return
}
~~~

apifunc中user：用户注册：

~~~go
// Register 用户注册
func Register(c *gin.Context) (model.UserResponse, error) {
	var userResponse model.UserResponse
	var token string
	username := c.Query("username")
	password := c.Query("password")
	userId, err := dbfunc.Register(username, password)
	if err != nil {
		return userResponse, err
	}
	token, err = util.SetToken(username, userId, time.Now().Add(time.Hour*240))
	if err != nil {
		return userResponse, err
	}
	userResponse.UserId = userId
	userResponse.Token = token
	return userResponse, nil
}
~~~

##### 4.5、对象储存

对象储存采用七牛云，可直接储存视频。

封面地址直接以视频播放地址和"?vframe/jpg/offset/1"拼接而成，以第一帧作为封面，也可以更改url的相关参数来设置封面。

文件命名：以时间戳和用户id拼接而成

文件上传采用分片上传的方式。但不知道是带宽问题还是什么原因，上传平均耗时在10~15秒。

##### 4.6、日志记录

通过引用一下包：

~~~go
"github.com/lestrrat-go/file-rotatelogs"
"github.com/rifflock/lfshook"
"github.com/sirupsen/logrus"
"gorm.io/gorm/logger"
~~~

实现了两类日志文件的格式化记录，每个日志文件最多保存一周，到期自动清除。

##### 4.7、防SQL注入

本项目的所有查询都是通过结构体查询，没有sql语句的拼接，所以不存在sql注入的问题
