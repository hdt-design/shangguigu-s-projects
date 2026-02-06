package model

// 定义一个结构体
type student struct {
	Name  string
	Score float64
}

//因为student结构体首字母小写，因此只能在model包使用
//通过工厂模式解决问题

func Newstudent(name string, score float64) *student {
	return &student{
		Name:  name,
		Score: score,
	}
}
