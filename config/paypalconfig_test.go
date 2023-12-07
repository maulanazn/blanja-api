package config_test

import (
	"context"
	"fmt"
	"log"
	"net/http/httptest"
	"testing"
	"util"

	"github.com/spf13/viper"

	"github.com/plutov/paypal/v4"
)

func TestConfigPaypal(t *testing.T) {
	var viper *viper.Viper = util.LoadConfig("../", "blanja.yaml", "yaml")
	conn, connErr := paypal.NewClient(viper.GetString("thirdparty.paypal_clientid"), viper.GetString("thirdparty.paypal_secret"), paypal.APIBaseSandBox)
	if connErr != nil {
		log.Println(connErr)
	}

	resultAuthorize, err := conn.GetAccessToken(context.TODO())

	if err != nil {
		t.Log(err)
	}

	fmt.Println(resultAuthorize.Token)
}

func TestCreateOrder(t *testing.T) {
	var viper *viper.Viper = util.LoadConfig("../", "blanja.yaml", "yaml")
	conn, connerr := paypal.NewClient(viper.GetString("thirdparty.paypal_clientid"), viper.GetString("thirdparty.paypal_secret"), paypal.APIBaseSandBox)

	request := httptest.NewRequest("POST", "/order", nil)
	recorder := httptest.NewRecorder()

	resultAuthorize, _ := conn.GetAccessToken(context.TODO())

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+resultAuthorize.Token)

	if connerr != nil {
		t.Error(connerr)
	}

	order, err := conn.CreateOrder(context.Background(), paypal.OrderIntentCapture, []paypal.PurchaseUnitRequest{
		{
			Amount: &paypal.PurchaseUnitAmount{
				Currency: "100",
			},
		}},
		&paypal.CreateOrderPayer{
			Name: &paypal.CreateOrderPayerName{
				GivenName: "maulanazn",
			},
			EmailAddress: "maulanazn@example.com",
		},
		&paypal.ApplicationContext{
			BrandName: "lkasdlfj",
		},
	)
	recorder.Result()

	if err != nil {
		t.Log(err)
	}
	t.Log(order.ID)
}

func TestAuthentication(t *testing.T) {
	request := httptest.NewRequest("POST", "/order", nil)
	request.Header.Set("Authorization", "")
	conn, connerr := paypal.NewClient("AWXJVRsgEPwRAndR74q5ohEIlWw7duJHSHNPxT16PBp-ZAzgut6umNUtCC1DwmOF_EnTQmxiUUSSpJWx", "EEJ3POLxQQseX22-E1JyoHWw2Earw6QzNY-2yJcGJfkyuvm90ZOFSp3G9gnjHwSz9KKZz0G0YwL_17Is", paypal.APIBaseSandBox)
	if connerr != nil {
		t.Log(connerr)
	}

	accesstoken, accesstokenerr := conn.GetAccessToken(context.Background())
	if accesstokenerr != nil {
		t.Log(accesstokenerr)
	}

	auth, err := conn.GetAuthorization(context.Background(), accesstoken.Token)
	if err != nil {
		t.Log(err)
	}

	fmt.Println(auth)
}
