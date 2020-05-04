package acls

type mapBoxACL struct {
}

//NewMapBoxACL contructor
func NewMapBoxACL() IMapACL {
	return &mapBoxACL{}
}

func (mp *mapBoxACL) GetAddressName(latitude float64, longitude float64) (string, error) {
	return "Teste de endereço", nil
}
