// learning from https://github.com/picadoh/gostreamer/blob/master/streamer/config.go

package streamer


/**
Config err struct represents a config err msg
*/

type ConfigError struct {
	error
	message string
}

/**
Config interface provides the means to access configuration.
*/

type Config interface {
	GetString(key string) string
	GetInt(key string) int
	ToString() string
}

/**
Properties config is a key/value pair based configuration structure.
*/

type PropertiesConfig struct {
	
}
