package abstractfactory

import "fmt"

type Chair interface {
	SitOn()
}

type Table interface {
	Use()
}

type ModernChair struct{}

func (m *ModernChair) SitOn() {
	fmt.Println("sit on modern chair")
}

type ClassicChair struct{}

func (c *ClassicChair) SitOn() {
	fmt.Println("sit on classic chair")
}

type ModernTable struct{}

func (m *ModernTable) Use() {
	fmt.Println("use modern table")
}

type ClassicTable struct{}

func (c *ClassicTable) Use() {
	fmt.Println("use classic table")
}

type FurnitureFactory interface {
	CreateChair() Chair
	CreateTable() Table
}

type ModernFurnitureFactory struct{}

func (m *ModernFurnitureFactory) CreateChair() Chair {
	return &ModernChair{}
}

func (m *ModernFurnitureFactory) CreateTable() Table {
	return &ModernTable{}
}

type ClassicFurnitureFactory struct{}

func (c *ClassicFurnitureFactory) CreateChair() Chair {
	return &ClassicChair{}
}

func (c *ClassicFurnitureFactory) CreateTable() Table {
	return &ClassicTable{}
}
