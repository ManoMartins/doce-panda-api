package gorm

import (
	couponModel "doce-panda/infra/gorm/coupon/model"
	orderModel "doce-panda/infra/gorm/order/model"
	paymentModel "doce-panda/infra/gorm/payment/model"
	productModel "doce-panda/infra/gorm/product/model"
	"doce-panda/infra/gorm/user/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/lib/pq"
	"log"
)

type Database struct {
	Db            *gorm.DB
	Dsn           string
	DsnTest       string
	DbType        string
	DbTypeTest    string
	Debug         bool
	AutoMigrateDb bool
	Env           string
}

func NewDb() *gorm.DB {
	dbInstance := &Database{
		Dsn:           "dbname=doce-panda sslmode=disable user=postgres password=mysecretpassword host=localhost",
		DbType:        "postgres",
		Debug:         false,
		AutoMigrateDb: true,
		Env:           "production",
	}

	connection, err := dbInstance.Connect()

	if err != nil {
		log.Fatalf("Test db error: %v", err)
	}

	return connection
}

func NewDbTest() *gorm.DB {
	dbInstance := &Database{}

	dbInstance.Env = "test"
	dbInstance.DbTypeTest = "sqlite3"
	dbInstance.DsnTest = ":memory:"
	dbInstance.AutoMigrateDb = true
	dbInstance.Debug = true

	connection, err := dbInstance.Connect()

	if err != nil {
		log.Fatalf("Test db error: %v", err)
	}

	return connection
}

func (d *Database) Connect() (*gorm.DB, error) {
	var err error

	if d.Env != "test" {
		d.Db, err = gorm.Open(d.DbType, d.Dsn)
	} else {
		d.Db, err = gorm.Open(d.DbTypeTest, d.DsnTest)
	}

	if err != nil {
		return nil, err
	}

	if d.Debug {
		d.Db.LogMode(true)
	}

	if d.AutoMigrateDb {
		d.Db.AutoMigrate(&productModel.Product{}, &model.User{}, &model.Address{}, &orderModel.Order{}, &orderModel.OrderItem{}, &paymentModel.CreditCard{}, &orderModel.OrderPayment{}, &couponModel.Coupon{})
		d.Db.Model(model.Address{}).AddForeignKey("user_id", "users (id)", "CASCADE", "CASCADE")
		d.Db.Model(orderModel.Order{}).AddForeignKey("address_id", "addresses (id)", "CASCADE", "CASCADE")
		d.Db.Model(orderModel.Order{}).AddForeignKey("user_id", "users (id)", "CASCADE", "CASCADE")
		d.Db.Model(orderModel.Order{}).AddForeignKey("coupon_id", "coupons (id)", "CASCADE", "CASCADE")
		d.Db.Model(orderModel.OrderItem{}).AddForeignKey("order_id", "orders (id)", "CASCADE", "CASCADE")
		d.Db.Model(orderModel.OrderItem{}).AddForeignKey("product_id", "products (id)", "CASCADE", "CASCADE")
		d.Db.Model(orderModel.OrderPayment{}).AddForeignKey("order_id", "orders (id)", "CASCADE", "CASCADE")
		d.Db.Model(orderModel.OrderPayment{}).AddForeignKey("credit_card_id", "credit_cards (id)", "CASCADE", "CASCADE")
		d.Db.Model(couponModel.Coupon{}).AddForeignKey("user_id", "users (id)", "CASCADE", "CASCADE")
	}

	return d.Db, nil
}
