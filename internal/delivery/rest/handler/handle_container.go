package handler

type HandlerContainer struct {
	*AuthHandler
	*BankHandler
	*ProductHandler
	*ImageHandler
}

func NewHandlerContainer(
	authHandler *AuthHandler,
	bankHandler *BankHandler,
	productHandler *ProductHandler,
	imageHandler *ImageHandler,
) *HandlerContainer {
	return &HandlerContainer{
		AuthHandler:    authHandler,
		BankHandler:    bankHandler,
		ProductHandler: productHandler,
		ImageHandler:   imageHandler,
	}
}
