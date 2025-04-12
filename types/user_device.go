package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeviceData struct {
	DeviceID string `bson:"deviceId"`
	User     User   `bson:"user"`
}

type User struct {
	Auth                       Auth                         `bson:"auth"`
	UserAvailablePaymentMethod []UserAvailablePaymentMethod `bson:"userAvailablePaymentMethod"`
}

type Auth struct {
	Email string `bson:"email"`
}

type UserAvailablePaymentMethod struct {
	PaymentMethod PlatformPaymentMethod `bson:"paymentMethod"`
}

type PlatformPaymentMethod struct {
	ID                          primitive.ObjectID `bson:"_id"`
	ProviderName                string             `bson:"providerName"`
	ProviderType                string             `bson:"providerType"`
	CodeName                    string             `bson:"codeName"`
	Icon                        string             `bson:"icon,omitempty"`
	Status                      string             `bson:"status"`
	ReceivingChanelNumberOrName string             `bson:"receivingChanelNumberOrName"`
	ExchangeRate                float64            `bson:"exchangeRate"`
	AutoConfirmationFromApp     bool               `bson:"autoConfirmationFromApp"`
	AllowWithdrawal             bool               `bson:"allowWithdrawal"`
	CreatedAt                   time.Time          `bson:"createdAt"`
	UpdatedAt                   time.Time          `bson:"updatedAt"`
}
type ActiveAgent struct {
	Uid     string   `json:"uid" bson:"uid"`
	Name    string   `json:"name" bson:"name"`
	Email   string   `json:"email"  bson:"email"`
	Devices []Device `json:"devices" bson:"devices"`
}

type UserInfo struct {
	Uid        string   `json:"uid" bson:"uid"`
	Name       string   `json:"name" bson:"name"`
	Email      string   `json:"email"  bson:"email"`
	DeviceId   string   `json:"deviceId" bson:"deviceId"`
	DeviceName string   `json:"deviceName" bson:"deviceName"`
	SimNumbers []string `json:"simNumbers" bson:"simNumbers"`
}
type Device struct {
	Id           string        `json:"id"`
	Name         string        `json:"name"`
	PhoneNumbers []PhoneNumber `json:"phoneNumbers"`
}

type PhoneNumber struct {
	Number         string          `json:"number" bson:"Number"`
	PaymentMethods []PaymentMethod `json:"paymentMethods" bson:"PaymentMethods"`
}

type PaymentMethod struct {
	Id   primitive.ObjectID `json:"id" bson:"ID"`
	Name string             `json:"name" bson:"Name"`
	Type string             `json:"type" bson:"Type"`
	Icon string             `json:"icon" bson:"Icon"`
}
