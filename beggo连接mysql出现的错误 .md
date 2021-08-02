# 一.beego连接mysql

## 错误：1.操作数据库的util.go的init()方法没有执行，导致数据库初始化失败！原因：init()命名为了initmysql().

## 心得：分析beego执行流程，观察终端运行。还有方法名成灰色代表没有使用。

## 错误2：执行models.InsertUser(user)方法失败。原因：user的id属性有非空限制，没有给默认值。解决方法：navicat将id属性设为自增。
