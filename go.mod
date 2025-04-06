module mycrack

go 1.22.4

// github.com/gtlyy/mycoin v0.0.0-20250406155554-290da113752d // indirect
require github.com/gtlyy/mycoin v0.0.2

require (
	github.com/btcsuite/btcd v0.22.2 // indirect
	github.com/btcsuite/btcutil v1.0.3-0.20201208143702-a53e38424cce // indirect
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37 // indirect
)

// 断网离线使用
// replace github.com/gtlyy/mycoin => ../mycoin
