package main

import (
	"fmt"
	"strings"
)

func main()  {
	var str string
	var step byte
	fmt.Println("请输入需要加密的字符和key, 用空格隔开！（不区分大小写）")
	_, _ = fmt.Scanf("%s %d",&str,&step)
	caesarEn(str, step)
}


func caesarEn(strRaw string, step byte) string {
	//1.将明文转成小写
	strRaw = strings.ToLower(strRaw)
	//2.将 明文字符串 转成 明文切片(( 内部 存放的 是 ACSII码 ))
	strSliceSrc := []byte(strRaw)
	//3.创建密文切片对象
	strSliceDst := make([]byte, len(strSliceSrc), len(strSliceSrc))
	//4.循环明文切片，将 ASCII码 + step位移值后 存入 密文切片
	for i := 0; i < len(strSliceSrc); i++ {
		//5.判断 明文字符的ASCII码 位移后 是否有超过 小写字母的范围，如果没有，则直接使用，如果有超过，则需要 -26
		if strSliceSrc[i] < 123-step {
			//直接加上 位移步长
			strSliceDst[i] = strSliceSrc[i] + step
		} else {
			strSliceDst[i] = strSliceSrc[i] + step - 26
		}
	}
	fmt.Println("明文：", strRaw, strSliceSrc)
	fmt.Println("密文：", string(strSliceDst), strSliceDst)


	return string(strSliceDst)
}


func caesarDe(strCipher string, step_move byte) string {
	//1.密文 转成 小写
	strCiphers := strings.ToLower(strCipher)
	//2.将字符串 转为 密文字符切片
	strSliceSrc := []byte(strCiphers)

	//3. 创建 明文字符切片
	strSliceDst := make([]byte, len(strSliceSrc), len(strSliceSrc))

	//4.循环密文切片
	for i := 0; i < len(strSliceSrc); i++ {
		//5.如果当前循环的 密文字符 在位移 范围内，则直接 减去 位移步长 存入 明文字符切片
		if strSliceSrc[i] >= 97+step_move {
			strSliceDst[i] = strSliceSrc[i] - step_move
		} else { //6.如果 密文字符 超出 范围，则 加上 26 后，再向左位移
			strSliceDst[i] = strSliceSrc[i] + 26 - step_move
		}
	}
	//7.输出结果
	fmt.Println("密文：", strCipher, strSliceSrc)
	fmt.Println("明文：", string(strSliceDst), strSliceDst)


	return string(strSliceDst)
}