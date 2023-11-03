package abstractfactory

import "testing"

func TestAbstractFactory(t *testing.T) {
	furnitureType := "modern"

	var furnitureFactory FurnitureFactory
	switch furnitureType {
	case "modern":
		furnitureFactory = &ModernFurnitureFactory{}
	case "classic":
		furnitureFactory = &ClassicFurnitureFactory{}
	}

	chair := furnitureFactory.CreateChair()
	table := furnitureFactory.CreateTable()

	chair.SitOn()
	table.Use()
}
