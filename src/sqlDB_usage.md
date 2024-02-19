db *sql.DB

Query   //将select传给DB

QueryContext    //带上下文

上面两个返回类型是type Rows struct{},这货的方法如下

func (rs *Rows) Close() error   //查询完要关闭
func (rs *Rows) ColumnTypes() ([]*ColumnType, error)    //返回列类型
ColumnTypes() ([]string, err)   //返回列名
Err() error     //若查询过程中有错误，会返回错误err
Next() bool     //按行读取结果集

QueryRow  //当你确定你的DB最多返回一条记录

这个返回的是Row struct{}

只有2种方法，

Err()与Scan(dest interface{}) error