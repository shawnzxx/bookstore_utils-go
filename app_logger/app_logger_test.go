package app_logger_test

import (
	"encoding/json"
	"fmt"
	"github.com/shawnzxx/bookstore_utils-go/app_logger"
	"os"
	"testing"
)

var logger = app_logger.GetLogger()

// 多行申明, 标注json.Marshal出来的是小写key，这样符合json输出规则
type ShoppingCart struct {
	Id     string  `json:"id"`
	Total  float32 `json:"total"`
	Orders []Order `json:"orders"`
}

type Order struct {
	ProductName string  `json:"productName"`
	Quantity    int     `json:"quantity"`
	Amount      float32 `json:"amount"`
}

var cart = &ShoppingCart{
	Id:    "001",
	Total: 50.25,
	Orders: []Order{
		{Amount: 10.10, ProductName: "chicken rice", Quantity: 2},
		{Amount: 30.05, ProductName: "pasta", Quantity: 1},
	},
}

// TestMain is main entrypoint for each of package
// M.run mean run all test cases in this test file
func TestMain(m *testing.M) {
	fmt.Println("about to start app_logger test")
	os.Exit(m.Run())
}

func TestLogInfo(t *testing.T) {
	logger.Info("I am info log")
	logger.Info("I am info log, order from %s, total amount is %v", "shawnzxx", 10)
	cartStr, _ := json.Marshal(cart)
	logger.Info("I am info log, order from %s, struct is %v", "shawnzxx", string(cartStr))
}

func TestDebugInfo(t *testing.T) {
	logger.Debug("I am debug log")
	logger.Debug("I am debug log, order from %s, total amount is %v", "shawnzxx", 10)
	cartStr, _ := json.Marshal(cart)
	logger.Debug("I am debug log, order from %s, struct is %v", "shawnzxx", string(cartStr))
}

func TestWarningInfo(t *testing.T) {
	logger.Warning("I am warning log")
	logger.Warning("I am warning log, order from %s, total amount is %v", "shawnzxx", 10)
	cartStr, _ := json.Marshal(cart)
	logger.Warning("I am warning log, order from %s, struct is %v", "shawnzxx", string(cartStr))
}

func TestErrorInfo(t *testing.T) {
	logger.Error("I am error log")
	logger.Error("I am error log, order from %s, total amount is %v", "shawnzxx", 10)
	cartStr, _ := json.Marshal(cart)
	logger.Error("I am error log, order from %s, struct is %v", "shawnzxx", string(cartStr))
}
