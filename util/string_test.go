package util

import (
	"fmt"
	"testing"
)

func TestTrimSpace(t *testing.T) {
	t4s := make([]*Test4, 0)
	t5s := make([]Test5, 0)
	t6s := make([]*Test6, 0)
	t4s = append(t4s, &Test4{Name: " 关 羽4 "})
	t4s = append(t4s, &Test4{Name: " 关 羽4 "})
	t4s = append(t4s, &Test4{Name: " 赵 云4 "})

	t5s = append(t5s, Test5{Name: " 关 羽5 "})
	t5s = append(t5s, Test5{Name: " 关 羽5 "})
	t5s = append(t5s, Test5{Name: " 赵 云5 "})

	t6s = append(t6s, &Test6{Name: " 关 羽6 "})
	t6s = append(t6s, &Test6{Name: " 关 羽6 "})
	t6s = append(t6s, &Test6{Name: " 赵 云6 "})
	w := &TestWrap{
		Name:      "张三 ",
		nameSpace: " Test ",
		Age:       20,
		T2: &Test1{
			T2: &Test2{Name: " 李 四 "},
			T3: Test3{
				Name: " 王 五 ",
			},
		},
		T4: t4s,
		T5: t5s,
		t6: t6s,
	}
	t.Log(Struct2String(w), fmt.Sprintf(" [%s] ", w.nameSpace), Struct2String(w.t6))
	TrimSpace(w)
	t.Log(Struct2String(w), fmt.Sprintf(" [%s] ", w.nameSpace), Struct2String(w.t6))
}

type TestWrap struct {
	Name      string
	nameSpace string `json:"name_space"` // 非导出，不修改
	Age       uint32
	T2        *Test1
	T4        []*Test4
	T5        []Test5
	t6        []*Test6 `json:"t_6"` // 非导出，不修改
}

type Test1 struct {
	T2 *Test2
	T3 Test3
}

type Test2 struct {
	Name string
}

type Test3 struct {
	Name string
}

type Test4 struct {
	Name string
}

type Test5 struct {
	Name string
}

type Test6 struct {
	Name string
}
