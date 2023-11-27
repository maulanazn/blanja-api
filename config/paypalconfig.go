package config

import (
	"context"
	"log"
	"os"

	"github.com/plutov/paypal/v4"
)

func GetPaypalConfig() *paypal.Client {
  conn, connErr := paypal.NewClient("AWXJVRsgEPwRAndR74q5ohEIlWw7duJHSHNPxT16PBp-ZAzgut6umNUtCC1DwmOF_EnTQmxiUUSSpJWx", "EEJ3POLxQQseX22-E1JyoHWw2Earw6QzNY-2yJcGJfkyuvm90ZOFSp3G9gnjHwSz9KKZz0G0YwL_17Is", paypal.APIBaseSandBox)
  if connErr != nil {
    log.Println(connErr)
  }

  conn.SetLog(os.Stdout)

  _, err := conn.GetAccessToken(context.TODO())
  if err != nil {
    log.Println(err)
  }

  return conn
}
