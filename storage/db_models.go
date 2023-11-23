package storage

import (
	"gorm.io/gorm"
	"time"
)

type Vendor struct {
	gorm.Model
	Title    string `gorm:"type:varchar(32);not null;default:''"`
	Address  string `gorm:"type:varchar(32);not null;default:''"`
	City     string `gorm:"type:varchar(32);not null;default:''"`
	Rate     int    `gorm:"type:int;not null;default:0"`
	IsActive bool   `gorm:"index;type:bool;default:false"`
}

type Courier struct {
	gorm.Model
	FirstName   string `gorm:"type:varchar(32);not null;default:''"`
	LastName    string `gorm:"type:varchar(32);not null;default:''"`
	Rate        int    `gorm:"type:int;not null;default:0"`
	IsAvailable bool   `gorm:"index;type:bool;default:false"`
}

type Customer struct {
	gorm.Model
	FirstName string `gorm:"type:varchar(32);not null;default:''"`
	LastName  string `gorm:"type:varchar(32);not null;default:''"`
	City      string `gorm:"type:varchar(32);not null;default:''"`
	Address   string `gorm:"type:varchar(256);not null;default:''"`
}

type Product struct {
	gorm.Model
	Title    string `gorm:"type:varchar(32);not null;default:''"`
	Vendor   Vendor `gorm:"foreignKey:VendorId"`
	VendorId uint
}

type Order struct {
	gorm.Model
	Vendor        Vendor    `gorm:"foreignKey:VendorId"`
	Customer      Customer  `gorm:"foreignKey:CustomerId"`
	State         string    `gorm:"type:varchar(32);not null;default:''"`
	RegisterTime  time.Time `gorm:"index;type:timestamptz"`
	DeliveryTime  int       `gorm:"type:int;not null;default:0"`
	HasDelay      bool      `gorm:"index;type:bool;default:false"`
	DeliveredTime time.Time `gorm:"index;type:timestamptz"`
	VendorId      uint
	CustomerId    uint
}

type Agent struct {
	gorm.Model
	FirstName   string `gorm:"type:varchar(32);not null;default:''"`
	LastName    string `gorm:"type:varchar(32);not null;default:''"`
	IsAvailable bool   `gorm:"index;type:bool;default:false"`
}

type Trip struct {
	gorm.Model
	Order    Order  `gorm:"foreignKey:OrderId"`
	State    string `gorm:"type:varchar(32);not null;default:''"`
	HasDelay bool   `gorm:"index;type:bool;default:false"`
	OrderId  uint
}

type DelayReport struct {
	gorm.Model
	Order            Order     `gorm:"foreignKey:OrderId"`
	Agent            Agent     `gorm:"foreignKey:AgentId"`
	NewEstimatedTime time.Time `gorm:"index;type:timestamptz"`
	AgentId          uint
	OrderId          uint
}
