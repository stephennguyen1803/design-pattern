package main

type ServiceLogistics struct {
	name  string
	doors int
	price float32
}

func (sL *ServiceLogistics) fillDefault() {
	if sL.name == "" {
		sL.name = "truct 10"
	}

	if sL.doors == 0 {
		sL.doors = 10
	}

	if sL.price == 0 {
		sL.price = 100000
	}
}

func newServiceLogistics(name string, doors int, price float32) *ServiceLogistics {
	if name == "" && doors == 0 && price == 0 {
		sL := ServiceLogistics{}
		sL.fillDefault()
		return &sL
	}

	sL := ServiceLogistics{
		name:  name,
		doors: doors,
		price: price,
	}

	return &sL
}

func GetServiceLogicstics(cargoVolumn int) *ServiceLogistics {
	switch cargoVolumn {
	default:
		return newServiceLogistics("", 0, 0)
	case 20:
		return newServiceLogistics("truck 20", 16, 1000000)
	}
}
