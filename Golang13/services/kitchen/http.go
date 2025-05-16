package main

// import (
// 	"context"
// 	"html/template"
// 	"log"
// 	"net/http"
// 	"source-base-go/common/genproto/orders"
// 	"time"

// 	"github.com/gin-gonic/gin"
// )

// type httpServer struct {
// 	addr string
// }
// func NewHttpServer(addr string) *httpServer{
// 	return &httpServer{addr: addr}

// }

// type CreateOrderInput struct {
// 	CustomerID int32 `json:"customer_id"`
// 	ProductID  int32 `json:"product_id"`
// 	Quantity   int32 `json:"quantity"`
// }

// func (s *httpServer) Run() error {
// 	router := gin.Default()

// 	conn := NewGRPCClient(":9000")
// 	defer conn.Close()

// 	router.POST("/orders", func(c *gin.Context){
// 		var payload CreateOrderInput
// 		if err := c.ShouldBindJSON(&payload); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
// 			return
// 		}
// 		orderClient  := orders.NewOrderServiceClient(conn)

// 		ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
// 		defer cancel()

// 		_, err := orderClient.CreateOrder(ctx, &orders.CreateOrderRequest{
// 			CustomerID: payload.CustomerID,
// 			ProductID: payload.ProductID,
// 			Quantity: payload.Quantity,
// 		})
// 		if err != nil{
// 			log.Fatalf("client errL %v", err)
// 		}
// 		res, err := orderClient.GetOrders(ctx, &orders.GetOrdersRequest{
// 			CustomerID: payload.CustomerID,
// 		})
// 		if err != nil{
// 			log.Fatalf("client err %v", err)
// 		}
// 		t := template.Must(template.New("order").Parse(ordersTemplate))

// 		if err := t.Execute(c.Writer,res.GetOrders()); err != nil{
// 			log.Fatalf("template err %v", err)
// 		}
// 	})
// 	log.Println("Starting server on", s.addr)
// 	return http.ListenAndServe(s.addr, router)
// }

// var ordersTemplate = `
// <!DOCTYPE html>
// <html>
// <head>
//     <title>Kitchen Orders</title>
// </head>
// <body>
//     <h1>Orders List</h1>
//     <table border="1">
//         <tr>
//             <th>Order ID</th>
//             <th>Customer ID</th>
//             <th>Quantity</th>
//         </tr>
//         {{range .}}
//         <tr>
//             <td>{{.OrderID}}</td>
//             <td>{{.CustomerID}}</td>
//             <td>{{.Quantity}}</td>
//         </tr>
//         {{end}}
//     </table>
// </body>
// </html>`