package config_test

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http/httptest"
	"testing"
	"util"

	"github.com/plutov/paypal/v4"
)

func TestConfigPaypal(t *testing.T) {
	var viper *viper.Viper = util.LoadConfig("../", "blanja.yaml", "yaml")
	conn, connErr := paypal.NewClient(viper.GetString("paypal_clientid"), viper.GetString("paypal_secret"), paypal.APIBaseSandBox)
	if connErr != nil {
		log.Println(connErr)
	}

	_, err := conn.GetAccessToken(context.TODO())
	if err != nil {
		t.Log(err)
	}

	fmt.Println(conn.Token)
}

func TestCreateOrder(t *testing.T) {
	request := httptest.NewRequest("POST", "/order", nil)
	recorder := httptest.NewRecorder()
	request.Header.Set("Authorization", "")

	conn, connerr := paypal.NewClient("AWXJVRsgEPwRAndR74q5ohEIlWw7duJHSHNPxT16PBp-ZAzgut6umNUtCC1DwmOF_EnTQmxiUUSSpJWx", "EEJ3POLxQQseX22-E1JyoHWw2Earw6QzNY-2yJcGJfkyuvm90ZOFSp3G9gnjHwSz9KKZz0G0YwL_17Is", paypal.APIBaseSandBox)
	if connerr != nil {
		t.Log(connerr)
	}
	authorizeid, authorizeerr := conn.GetAuthorization(context.TODO(), "laksjdfklasldfjaksdfj")
	if authorizeerr != nil {
		t.Log(authorizeerr)
	}
	t.Log(authorizeid)
	_, err := conn.GetAccessToken(context.TODO())
	if err != nil {
		t.Log(err)
	}

	order, err := conn.CreateOrder(context.Background(), paypal.OrderIntentCapture, []paypal.PurchaseUnitRequest{paypal.PurchaseUnitRequest{
		ReferenceID: "ref-id",
		Items: []paypal.Item{
			{Name: "product1"},
			{Name: "product2"},
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
