package app

var cfg = &config{
	livenessIsOk: true,
	readnessIsOk: true,
}

type config struct {
	livenessIsOk bool
	readnessIsOk bool
}
