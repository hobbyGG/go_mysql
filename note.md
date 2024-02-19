1->准备连接到数据库

首先加载目标数据库的驱动，sql.Open(数据库驱动名称, 数据源名称) 返回sql.DB(struct)

sql.DB用于操作数据库，代表连接池。多个goroutine对DB的使用是安全的。

Open()并不会连接，只是把链接所需的struct设置好。真正的连接是在被需要时才进行的

sql.DB不需要关闭，只是用来处理数据库的，而不是实际的连接

使用sql.DB时可以作为全局变量，也可以将它作为函数参数进行传递

正常获取驱动是要用sql.Register()和一个实现driver.Driver()接口的struct来注册数据库驱动,sql sever会自我注册

