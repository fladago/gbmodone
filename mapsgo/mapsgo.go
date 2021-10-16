package mapsgo

// nolint
// Функция удаления из мапы по ключу
func f(m map[string]string, key string) {
	delete(m, key)
}

// m := make(map[string]string)
// 	m["hello"] = "World"
// 	m["q"] = "a"
// 	if v, ok := m["hello"]; ok {
// 		fmt.Println(v)
// 	}

// 	mNew := make(map[string]string, len(m))
// 	for k, v := range m {
// 		fmt.Println(k, v)
// 		mNew[k] = v
// 	}
// 	fmt.Println(mNew)

// 	f(m, "q")

// 	s := []int{120, 3, 0, 44}
// 	sa := []int{1, 2, 3, 4}
// 	s = append(s, sa...)
// 	fmt.Println(s)
