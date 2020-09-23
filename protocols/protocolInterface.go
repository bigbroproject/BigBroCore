package protocols

import "github.com/moneye/models"

type ProtocolInterface interface {
	CheckService(Protocol models.Protocol) error // Check if protocol is OK or ERROR
}

func RegisterInterface(registeredInterfaces *map[string]ProtocolInterface, protocolInterfaceName string, protocolInterface ProtocolInterface) {
	(*registeredInterfaces)[protocolInterfaceName] = protocolInterface
}

func DefaultRegisteredInterfaces() (map[string]ProtocolInterface, error) {
	registeredInterfaces := make(map[string]ProtocolInterface, 0)
	registeredInterfaces["https"] = Https{}
	registeredInterfaces["http"] = Http{}
	registeredInterfaces["icmp"] = Icmp{}
	registeredInterfaces["icmp6"] = Icmp6{}
	return registeredInterfaces, nil
}

func IsRegistered(registeredInterfaces *map[string]ProtocolInterface, funcName string) bool {
	interf := (*registeredInterfaces)[funcName]

	if interf != nil {
		return true
	} else {
		return false
	}
}
