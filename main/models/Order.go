package models
type Order struct {
	Order_id     int   `json:"order_id"`
	Table_id     int   `json:"table_id"`
	Waiter_id    int   `json:"waiter_id"`
	Items        []int `json:"items"`
	Priority     int   `json:"priority"`
	Max_wait     int   `json:"max_wait"`
	Pick_up_time int   `json:"pick_up_time"`
}
