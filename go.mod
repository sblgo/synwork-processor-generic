module sbl.system/synwork/generic

go 1.17

replace sbl.systems/go/synwork/plugin-sdk => ../synwork/plugin-sdk

replace sbl.systems/go/synwork => ../synwork

require (
	golang.org/x/net v0.0.0-20211216030914-fe4d6282115f
	sbl.systems/go/synwork/plugin-sdk v0.0.0-00010101000000-000000000000
)

require (
	github.com/aws/aws-sdk-go v1.43.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/hashicorp/go-version v1.4.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	golang.org/x/text v0.3.6 // indirect
	sbl.systems/go/synwork v0.0.0-00010101000000-000000000000 // indirect
)
