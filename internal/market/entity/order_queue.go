package entity

// a maioria dessas funcoes são de manipulação da nossa fila. Podemos até ter ferramentas dessas funções salvas.

type OrderQueue struct {
	Orders []*Order
}

func (oq *OrderQueue) Less(i, j int) bool {
	return oq.Orders[i].Price < oq.Orders[j].Price
}

func (oq *OrderQueue) Swap(i, j int) {
	oq.Orders[i], oq.Orders[j] = oq.Orders[j], oq.Orders[i]
}

func (oq *OrderQueue) Len() int {
	return len(oq.Orders)
}

func (oq *OrderQueue) Push(x interface{}) {
	oq.Orders = append(oq.Orders, x.(*Order))
}

func (oq *OrderQueue) Pop() interface{} {
	oldOrders := oq.Orders                 // copia nossas orders
	ordersQty := len(oldOrders)            // pega quantas orders temos
	lastOrder := oldOrders[ordersQty-1]    // pega a ultima order
	oq.Orders = oldOrders[0 : ordersQty-1] // remove a ultima order do slice/array
	return lastOrder                       // retorna a ultima order
}

func NewOrderQueue() *OrderQueue {
	return &OrderQueue{}
}
