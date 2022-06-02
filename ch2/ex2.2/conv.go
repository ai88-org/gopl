package main

// 人民币转换成美元
func RToD(r Renminbi) Dollar {
	return Dollar(r * 0.01495)
}

func DToR(d Dollar) Renminbi {
	return Renminbi(d * 6.6912)
}
