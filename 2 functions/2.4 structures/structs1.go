type TestStruct struct {
	On bool
	Ammo, Power int
}

func (c *TestStruct) Shoot() bool {
	if c.On == false {
		return false
	}
	if c.Ammo <= 0 {
		return false
	}
	c.Ammo--
	return true
}

func (c *TestStruct) RideBike() bool {
	if c.On == false {
		return false
	}
	if c.Power <= 0 {
		return false
	}
	c.Power--
	return true
}

func main() {
    testStruct := &TestStruct{}

// Вам необходимо реализовать структуру со свойствами-полями On, Ammo и Power, с типами bool, int, int соответственно.
// У этой структуры должны быть методы: Shoot и RideBike, которые не принимают аргументов, но возвращают значение bool.
// Если значение On == false, то оба метода вернут false.
