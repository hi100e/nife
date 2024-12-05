package myip

import (
	"context"
	"fmt"
	"net"

	"github.com/hi100e/nife"
)

func init() {
	cmd := &nife.Cmd{
		Name:       "myip",
		Title:      "My IP Addresses",
		Usage:      "myip",
		Short:      "Get your private and public IP address",
		Long:       "Get your private and public IP address using dns, http and ifconfig.co",
		CmdHandler: Run,
	}
	nife.RegisterCmd(cmd)
}

func Run(ctx context.Context, cmd string, args []string) error {
	select {
	case <-ctx.Done():
		return nil
	default:
		ip, err := FindMyIPWithDNS()
		if err != nil {
			return err
		}
		fmt.Println("Public IP:", ip)
		fmt.Println("Local IP Addresses:")
		localIPs, err := getLocalIPAddresses()
		if err != nil {
			return err
		}
		for _, ip := range localIPs {
			fmt.Println(ip)
		}
	}

	return nil
}

// FindMyIPWithDNS returns the public IP address of the current machine using DNS
func FindMyIPWithDNS() (string, error) {
	//dig +short txt ch whoami.cloudflare @1.0.0.1
	//dig -6 TXT +short o-o.myaddr.l.google.com @ns1.google.com
	//dig +short myip.opendns.com @resolver1.opendns.com
	txts, err := lookupTXTRecord("o-o.myaddr.l.google.com", "ns1.google.com:53")
	if err != nil {
		return "", err
	}
	for _, txt := range txts {
		if ip := net.ParseIP(txt); ip != nil {
			return ip.String(), nil
		}
	}
	return "", nil
}

// function to perform DNS TXT record lookup
func lookupTXTRecord(domain, dnsServer string) ([]string, error) {
	// Specify the resolver to use the custom DNS server
	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			return net.Dial("udp", dnsServer)
		},
	}

	// Perform the TXT lookup
	txtRecords, err := resolver.LookupTXT(context.Background(), domain)
	if err != nil {
		return nil, err
	}
	return txtRecords, nil
}

func getLocalIPAddresses() ([]string, error) {
	var ipAddresses []string

	// Get all network interfaces
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, iface := range interfaces {
		// Skip the loopback interface
		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		// Skip down interfaces
		if iface.Flags&net.FlagUp == 0 {
			continue
		}
		// Skip virtual interfaces (e.g., Docker, VPN)
		if iface.Flags&net.FlagPointToPoint != 0 || iface.Flags&net.FlagMulticast == 0 {
			continue
		}
		// Skip tunnel interfaces
		if iface.Flags&net.FlagBroadcast == 0 {
			continue
		}

		// Get all addresses for each interface
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			localIP := fmt.Sprintf("%v\t%v", iface.Name, ip.String())

			ipAddresses = append(ipAddresses, localIP)
		}
	}

	return ipAddresses, nil
}
