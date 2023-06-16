### 一、为什么要使用`gorm-gen`来生成实体类和查询

* 1、根据`gorm`官网地址，正常的写法是先写数据模型,然后由数据模型自动同步生成到数据库中,但是这样的工作量会比较大,对于写后端的人来说都熟悉`sql`语句,正常来说都是先自己手动创建表,利用工具将表字段同步到项目实体类中，[之前我就写过一个这样的工具](https://github.com/kuangshp/generate-model)，对于小项目来说都能满足

* 2、使用上面的方式在项目中写出的代码大概是如此，这样如果数据库字段改了,程序是无法识别出来，不能很好的规避错误

  ![image-20230616164124139](E:\git\2023\gorm_gen\assets\image-20230616164124139.png)

* 3、[`gorm-gen`](https://gorm.io/gen/)刚好可以满足这个需求，写出的代码更好的维护，比如下面的代码

  ![image-20230616164318982](E:\git\2023\gorm_gen\assets\image-20230616164318982.png)

### 二、在项目中配置`gorm-gen`来生成`sql`操作

* 1、手动新创建表

  ```sql
  DROP TABLE IF EXISTS `user`;
  CREATE TABLE `user`  (
    `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `username` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '用户名',
    `age` int UNSIGNED NOT NULL COMMENT '年龄',
    `password` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '密码',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT '软删除时间',
    `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '手机号码',
    PRIMARY KEY (`id`) USING BTREE
  ) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;
  ```

* 2、根据官网来安装依赖包

  ```properties
  go get -u gorm.io/gen
  go get -u gorm.io/driver/mysql
  ```

* 3、在项目的根目录下创建一个`generate.go`的文件

  ```go
  package main
  
  import (
  	"fmt"
  	"gorm.io/driver/mysql"
  	"gorm.io/gen"
  	"gorm.io/gen/field"
  	"gorm.io/gorm"
  	"strings"
  )
  
  const dsn = "root:123456@(localhost:3306)/test001?charset=utf8mb4&parseTime=True&loc=Local"
  
  // Case2Camel 下划线转驼峰(大驼峰)
  func Case2Camel(name string) string {
  	name = strings.Replace(name, "_", " ", -1) // 根据_来替换成
  	name = strings.Title(name)                 // 全部大写
  	return strings.Replace(name, " ", "", -1)  // 删除空格
  }
  
  // LowerCamelCase 转换为小驼峰
  func LowerCamelCase(name string) string {
  	name = Case2Camel(name)
  	return strings.ToLower(name[:1]) + name[1:]
  }
  func main() {
  	var tableName string
  	fmt.Print("请输入表名,如果不输入将全部同步数据库:")
  	if _, err := fmt.Scanln(&tableName); err != nil {
  		fmt.Println("没有输入表名将全部同步数据库", err.Error())
  	}
  	var fileName = Case2Camel(tableName) // 转为首字目大写
  	fmt.Println(fileName)
  	// 连接数据库
  	db, err := gorm.Open(mysql.Open(dsn))
  	if err != nil {
  		panic(fmt.Errorf("cannot establish db connection: %w", err))
  	}
  
  	// 构造生成器实例
  	g := gen.NewGenerator(gen.Config{
  		// 相对执行`go run`时的路径, 会自动创建目录
  
  		OutPath:      "./dao",   //curd代码的输出路径
  		ModelPkgPath: "./model", //model代码的输出路径
  
  		// WithDefaultQuery 生成默认查询结构体(作为全局变量使用), 即`Q`结构体和其字段(各表模型)
  		// WithoutContext 生成没有context调用限制的代码供查询
  		// WithQueryInterface 生成interface形式的查询代码(可导出), 如`Where()`方法返回的就是一个可导出的接口类型
  		Mode: gen.WithDefaultQuery | gen.WithoutContext | gen.WithQueryInterface,
  
  		// 表字段可为 null 值时, 对应结体字段使用指针类型
  		FieldNullable: false, // generate pointer when field is nullable
  
  		// 表字段默认值与模型结构体字段零值不一致的字段, 在插入数据时需要赋值该字段值为零值的, 结构体字段须是指针类型才能成功, 即`FieldCoverable:true`配置下生成的结构体字段.
  		// 因为在插入时遇到字段为零值的会被GORM赋予默认值. 如字段`age`表默认值为10, 即使你显式设置为0最后也会被GORM设为10提交.
  		// 如果该字段没有上面提到的插入时赋零值的特殊需要, 则字段为非指针类型使用起来会比较方便.
  		FieldCoverable: false, // generate pointer when field has default value, to fix problem zero value cannot be assign: https://gorm.io/docs/create.html#Default-Values
  
  		// 模型结构体字段的数字类型的符号表示是否与表字段的一致, `false`指示都用有符号类型
  		FieldSignable: false, // detect integer field's unsigned type, adjust generated data type
  		// 生成 gorm 标签的字段索引属性
  		FieldWithIndexTag: false, // generate with gorm index tag
  		// 生成 gorm 标签的字段类型属性
  		FieldWithTypeTag: true, // generate with gorm column type tag
  	})
  	// 设置目标 db
  	g.UseDB(db)
  
  	// 自定义字段的数据类型
  	// 统一数字类型为int64,兼容protobuf和thrift
  	dataMap := map[string]func(detailType gorm.ColumnType) (dataType string){
  		"tinyint":   func(detailType gorm.ColumnType) (dataType string) { return "int64" },
  		"smallint":  func(detailType gorm.ColumnType) (dataType string) { return "int64" },
  		"mediumint": func(detailType gorm.ColumnType) (dataType string) { return "int64" },
  		"bigint":    func(detailType gorm.ColumnType) (dataType string) { return "int64" },
  		"int":       func(detailType gorm.ColumnType) (dataType string) { return "int64" },
  		"timestamp": func(detailType gorm.ColumnType) (dataType string) { return "LocalTime" }, // 自定义时间
  		"decimal":   func(detailType gorm.ColumnType) (dataType string) { return "Decimal" },   // 金额类型全部转换为第三方库,github.com/shopspring/decimal
  	}
  	// 要先于`ApplyBasic`执行
  	g.WithDataTypeMap(dataMap)
  
  	// 自定义模型结体字段的标签
  	// 将特定字段名的 json 标签加上`string`属性,即 MarshalJSON 时该字段由数字类型转成字符串类型
  	jsonField := gen.FieldJSONTagWithNS(func(columnName string) (tagContent string) {
  		toStringField := `id, `
  		if strings.Contains(toStringField, columnName) {
  			return columnName + ",string"
  		} else if strings.Contains(`deleted_at`, columnName) {
  			return "-"
  		}
  		return LowerCamelCase(columnName) // 下划线转小驼峰
  	})
  	// 将非默认字段名的字段定义为自动时间戳和软删除字段;
  	// 自动时间戳默认字段名为:`updated_at`、`created_at, 表字段数据类型为: INT 或 DATETIME
  	// 软删除默认字段名为:`deleted_at`, 表字段数据类型为: DATETIME
  	autoUpdateTimeField := gen.FieldGORMTag("updated_at", func(tag field.GormTag) field.GormTag {
  		return map[string]string{
  			"column":  "updated_at",
  			"comment": "更新时间",
  		}
  	})
  	autoCreateTimeField := gen.FieldGORMTag("created_at", func(tag field.GormTag) field.GormTag {
  		return map[string]string{
  			"column":  "created_at",
  			"comment": "创建时间",
  		}
  	})
  	softDeleteField := gen.FieldType("deleted_at", "gorm.DeletedAt")
  	// 模型自定义选项组
  	fieldOpts := []gen.ModelOpt{jsonField, autoCreateTimeField, autoUpdateTimeField, softDeleteField}
  	//fieldOpts := []gen.ModelOpt{jsonField, softDeleteField}
  
  	// 创建模型的结构体,生成文件在 model 目录; 先创建的结果会被后面创建的覆盖
  	// 这里创建个别模型仅仅是为了拿到`*generate.QueryStructMeta`类型对象用于后面的模型关联操作中
  	//User := g.GenerateModel("user")
  	// 如果传递了表名的时候就单独创建单独的表
  	if tableName != "" {
  		//allModel := g.GenerateAllTable(fieldOpts...)
  		allModel := g.GenerateModelAs(tableName, fileName, fieldOpts...)
  		// 创建有关联关系的模型文件
  		// 可以用于指定外键
  		//Score := g.GenerateModel("score",
  		// append(
  		//    fieldOpts,
  		//    // user 一对多 address 关联, 外键`uid`在 address 表中
  		//    gen.FieldRelate(field.HasMany, "user", User, &field.RelateConfig{GORMTag: "foreignKey:UID"}),
  		// )...,
  		//)
  
  		// 创建模型的方法,生成文件在 query 目录; 先创建结果不会被后创建的覆盖
  		//g.ApplyBasic(User)
  		g.ApplyBasic(allModel)
  	} else {
  		allModel := g.GenerateAllTable(fieldOpts...)
  		g.ApplyBasic(allModel...)
  	}
  
  	g.Execute()
  }
  ```

* 4、在`utils`文件夹下创建一个`db.go`的文件,这个地方是连接数据库的，先运行上面的文件才会生成`dao`文件

  ```go
  const dsn = "root:123456@(localhost:3306)/test001?charset=utf8mb4&parseTime=True&loc=Local"
  
  var DB = dao.Query{}
  
  func InitDB() {
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
  		Logger:                                   logger.Default.LogMode(logger.Info),
  		DisableForeignKeyConstraintWhenMigrating: true, // 自动创建表的时候不创建外键
  		NamingStrategy: schema.NamingStrategy{ // 自动创建表时候表名的配置
  			SingularTable: true,
  		},
  	})
  	if err != nil {
  		panic(err)
  	}
      // 进行关联
  	dao.SetDefault(db)
  	q := dao.Q
  	DB = *q
  }
  
  func init() {
  	InitDB()
  }
  ```

* 5、可以正常使用内置封装好的方法了