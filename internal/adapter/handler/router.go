package handler

import "github.com/gin-gonic/gin"

type Router struct {
	*gin.Engine
}

func NewRouter(
	customerHandler *CustomerHandler,
	attendantHandler *AttendantHandler,
	productHandler *ProductHandler,
	orderHandler *OrderHandler,
) (*Router, error) {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the API GoFood",
		})
	})

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "UP",
		})
	})

	customers := router.Group("/customers")
	{
		customers.GET("/", customerHandler.GetCustomers)
		customers.GET("/:id", customerHandler.GetCustomer)
		customers.GET("/cpf/:cpf", customerHandler.GetCustomerByCPF)
		customers.POST("/", customerHandler.CreateCustomer)
		customers.PUT("/:id", customerHandler.UpdateCustomer)
		customers.DELETE("/:id", customerHandler.DeleteCustomer)
	}

	attendants := router.Group("/attendants")
	{
		attendants.GET("/", attendantHandler.GetAttendants)
		attendants.GET("/:id", attendantHandler.GetAttendant)
		attendants.POST("/", attendantHandler.CreateAttendant)
		attendants.PUT("/:id", attendantHandler.UpdateAttendant)
		attendants.DELETE("/:id", attendantHandler.DeleteAttendant)
	}

	products := router.Group("/products")
	{
		products.GET("/", productHandler.GetProducts)
		products.GET("/:id", productHandler.GetProductById)
		products.POST("/", productHandler.CreateProduct)
		products.PUT("/:id", productHandler.UpdateProduct)
		products.DELETE("/:id", productHandler.DeleteProduct)
	}

	orders := router.Group("/orders")
	{
		orders.POST("/", orderHandler.StartOrder)
		orders.PUT("/:id/product", orderHandler.AddItemToOrder)
		orders.GET("/:id", orderHandler.GetOrderById)
		orders.PUT("/:id/confirmation", orderHandler.ConfirmationOrder)
		orders.PUT("/:id/payment", orderHandler.PaymentOrder)
	}

	return &Router{
		router,
	}, nil
}
