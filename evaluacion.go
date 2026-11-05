package main

import (
	"fmt"
	"time"
)

type Medicine struct {
	name            string
	manufacturer    string
	manufactureDate time.Time
	shelfLife       int
}

// Getters
func (m *Medicine) GetName() string { return m.name }
func (m *Medicine) GetManufacturer() string { return m.manufacturer }
func (m *Medicine) GetManufactureDate() time.Time { return m.manufactureDate }
func (m *Medicine) GetShelfLife() int { return m.shelfLife }

// Setters con validación básica
func (m *Medicine) SetName(name string) {
	if name != "" {
		m.name = name
	}
}

func (m *Medicine) SetManufacturer(manufacturer string) {
	if manufacturer != "" {
		m.manufacturer = manufacturer
	}
}

func (m *Medicine) SetManufactureDate(dateStr string) {
	if f, err := time.Parse("2006-01-02", dateStr); err == nil {
		m.manufactureDate = f
	}
}

func (m *Medicine) SetShelfLife(months int) {
	if months > 0 {
		m.shelfLife = months
	}
}

// Constructor base
func NewMedicine(name, manufacturer, date string, shelfLife int) Medicine {
	m := Medicine{}
	m.SetName(name)
	m.SetManufacturer(manufacturer)
	m.SetManufactureDate(date)
	m.SetShelfLife(shelfLife)
	return m
}

// Método de caducidad
func (m *Medicine) ExpirationDate() time.Time {
	return m.manufactureDate.AddDate(0, m.shelfLife, 0)
}

type Tablet struct {
	Medicine
	dosePerTablet float64
}

func NewTablet(name, manufacturer, date string, shelfLife int, dose float64) *Tablet {
	if dose <= 0 {
		return nil
	}
	base := NewMedicine(name, manufacturer, date, shelfLife)
	return &Tablet{Medicine: base, dosePerTablet: dose}
}

func (t *Tablet) ShowDetails() {
	fmt.Println("--- Tableta ---")
	fmt.Println("Nombre:", t.name)
	fmt.Println("Fabricante:", t.manufacturer)
	fmt.Println("Fabricación:", t.manufactureDate.Format("2006-01-02"))
	fmt.Println("Vida útil:", t.shelfLife, "meses")
	fmt.Println("Dosis:", t.dosePerTablet, "mg")
}

type Syrup struct {
	Medicine
	volume float64
	flavor string
}

func NewSyrup(name, manufacturer, date string, shelfLife int, volume float64, flavor string) *Syrup {
	if volume <= 0 {
		return nil
	}
	base := NewMedicine(name, manufacturer, date, shelfLife)
	return &Syrup{Medicine: base, volume: volume, flavor: flavor}
}

func (s *Syrup) ShowDetails() {
	fmt.Println("--- Jarabe ---")
	fmt.Println("Nombre:", s.name)
	fmt.Println("Fabricante:", s.manufacturer)
	fmt.Println("Fabricación:", s.manufactureDate.Format("2006-01-02"))
	fmt.Println("Vida útil:", s.shelfLife, "meses")
	fmt.Println("Volumen:", s.volume, "ml")
	fmt.Println("Sabor:", s.flavor)
}

func main() {
	// Crear inventario de tabletas
	tablets := []*Tablet{
		NewTablet("Paracetamol", "Genfar", "2025-03-15", 24, 500),
		NewTablet("Amoxicilina", "Sanofi", "2025-06-01", 18, 875),
	}

	// Crear inventario de jarabes
	syrups := []*Syrup{
		NewSyrup("Ambroxol", "Bayer", "2025-09-10", 12, 120, "Fresa"),
		NewSyrup("Ibuprofeno", "Pfizer", "2025-01-20", 36, 100, "Naranja"),
	}

	fmt.Println("=== INVENTARIO ===")
	for _, t := range tablets {
		t.ShowDetails()
		fmt.Println()
	}
	for _, s := range syrups {
		s.ShowDetails()
		fmt.Println()
	}

	fmt.Println("=== FECHAS DE CADUCIDAD ===")
	for _, t := range tablets {
		fmt.Println(t.GetName(), "caduca el:", t.ExpirationDate().Format("2006-01-02"))
	}
	for _, s := range syrups {
		fmt.Println(s.GetName(), "caduca el:", s.ExpirationDate().Format("2006-01-02"))
	}
}