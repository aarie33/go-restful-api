//go:build wireinject
// +build wireinject

package simple

import (
	"github.com/google/wire"
	"io"
	"os"
)

func InitializedService(isError bool) (*SimpleService, error) {
	wire.Build(NewSimpleRepository, NewSimpleService)
	return nil, nil
}

func InitializedDatabaseRepository() *DatabaseRepository {
	wire.Build(NewDatabasePostgreSQL, NewDatabaseMongoDB, NewDatabaseRepository)

	return nil
}

var fooSet = wire.NewSet(NewFooRepository, NewFooService)
var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializedFooBarService() *FooBarService {
	wire.Build(fooSet, barSet, NewFooBarService)

	return nil
}

// Binding interface
// contoh kode salah
// func InitializedHelloService() *HelloService {
// 	wire.Build(NewHelloService, NewSayHelloImpl)

// 	return nil
// }

var helloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)),
)

func InitializedHelloService() *HelloService {
	wire.Build(helloSet, NewHelloService)

	return nil
}

// Struct Provider
// jarang digunakan
func InitializedFooBar() *FooBar {
	// wire.Build(NewFoo, NewBar, wire.Struct(new(FooBar), "Foo", "Bar"))
	wire.Build(NewFoo, NewBar, wire.Struct(new(FooBar), "*"))

	return nil
}

var fooValue = &Foo{}
var barValue = &Bar{}

func InitializedFooBarUsingValue() *FooBar {
	wire.Build(wire.Value(fooValue), wire.Value(barValue), wire.Struct(new(FooBar), "*"))

	return nil
}

func InitializedReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))

	return nil
}

func InitializedConfiguration() *Configuration {
	// application := NewApplication()
	// configuration := application.Configuration

	// return configuration
	// shoule be like this

	wire.Build(
		NewApplication,
		wire.FieldsOf(new(*Application), "Configuration"),
	)

	return nil
}

func InitializedConnection(name string) (*Connection, func()) {
	wire.Build(NewConnection, NewFile)
	return nil, nil
}