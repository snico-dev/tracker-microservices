package acls

//IMapACL contract
type IMapACL interface {
	GetAddressName(latitude float64, longitude float64) (string, error)
}
