package config

import (
	"errors"
	"flag"
	"net"
	"strconv"
	"strings"
)

const (
	defaultServerAddress   = "localhost:8080"
	defaultDatabaseDSN     = ""
	defaultAccrualAddress  = ""
	defaultRequestTimeout  = int64(20)
	defaultPasswordHashKey = "hash"
	defaultTokenSignKey    = "token"
	defaultIssuer          = "gophermart"
	defaultTokenDuration   = int64(3)
)

type NetAddress struct {
	Host string
	Port int
}

func ParseConfigFlags() (netAddress string, databaseDSN string, accrualAddress string, tokenSignKey string, hashKey string, issuer string, tokenDuration int64, requestTimeout int64) {
	serverAddress := NetAddress{}
	accrualSystemAddress := NetAddress{}
	_ = flag.Value(&serverAddress)

	flag.Var(&serverAddress, "a", "Net address host:port")
	flag.StringVar(&databaseDSN, "d", defaultDatabaseDSN, "Postgres database connection string")
	flag.StringVar(&accrualAddress, "r", defaultAccrualAddress, "Accrual net address host:port")

	flag.StringVar(&tokenSignKey, "s", defaultTokenSignKey, "Key for signing JWT")
	flag.StringVar(&hashKey, "p", defaultPasswordHashKey, "Key for hashing password")
	flag.StringVar(&issuer, "i", defaultIssuer, "Token Issuer")
	flag.Int64Var(&tokenDuration, "t", defaultTokenDuration, "Token expiration time in hours")

	flag.Int64Var(&requestTimeout, "rt", defaultRequestTimeout, "Timeout for every request in seconds")

	flag.Parse()

	return serverAddress.String(), databaseDSN, accrualSystemAddress.String(), tokenSignKey, hashKey, issuer, tokenDuration, requestTimeout
}

func (a *NetAddress) String() string {
	if a.Host == "" && a.Port == 0 {
		return defaultServerAddress
	}

	return a.Host + ":" + strconv.Itoa(a.Port)
}

func (a *NetAddress) Set(s string) error {
	hostAndPort := strings.Split(s, ":")
	if len(hostAndPort) != 2 {
		return errors.New("need address in a form `host:port`")
	}

	host := hostAndPort[0]
	port, err := strconv.Atoi(hostAndPort[1])
	if err != nil {
		return err
	}

	if port < 1 {
		return errors.New("port number is a positive integer")
	}

	if host != "localhost" {
		ip := net.ParseIP(hostAndPort[0])
		if ip == nil {
			return errors.New("incorrect IP-address provided")
		}
	}

	a.Host = host
	a.Port = port
	return nil
}
