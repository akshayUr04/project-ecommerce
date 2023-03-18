package response

type AdminData struct {
	Id           int
	Name         string
	Email        string
	IsSuperAdmin bool
}
type DashBoard struct {
	TotalRevenue        int
	TotalOrders         int
	TotalProductsSelled int
	TotalUsers          int
}
