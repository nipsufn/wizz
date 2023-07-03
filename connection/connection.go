package connection

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/FerdinaKusumah/wizz/models"
	log "github.com/sirupsen/logrus"
)

const Port = "38899"

func SendUdpMessage(host string, message *models.RequestPayload) (*models.ResponsePayload, error) {
	var (
		err             error
		response        = make([]byte, 4096)
		responsePayload = new(models.ResponsePayload)
		remoteAddr      = new(net.UDPAddr)
		conn            = new(net.UDPConn)
		payload         []byte
	)
	// initialize connection to UDP
	if remoteAddr, err = net.ResolveUDPAddr("udp", fmt.Sprintf(`%s:%s`, host, Port)); err != nil {
		log.Errorf(`Unable to resolve to udp: %s`, err)
		return nil, err
	}
	if conn, err = net.DialUDP("udp", nil, remoteAddr); err != nil {
		log.Errorf(`Unable to dial up to udp: %s`, err)
		return nil, err
	}
	if err = conn.SetReadDeadline(time.Now().Add(30 * time.Second)); err != nil {
		log.Errorf(`Unable to set UDP timeout: %s`, err)
		return nil, err
	}
	defer conn.Close()
	if payload, err = json.Marshal(message); err != nil {
		log.Errorf(`Unable to marshal payload: %s`, err)
		return nil, err
	}
	payloadString := string(payload)
	log.Tracef(fmt.Sprintf(`Payload string: %s`, payloadString))
	// send payload to bulb
	if _, err = conn.Write(payload); err != nil {
		log.Errorf(`Unable to send message to UDP: %s`, err)
		return nil, err
	}
	// read response from bulb
	if _, err = bufio.NewReader(conn).Read(response); err != nil {
		log.Errorf(`Unable to read message from UDP: %s`, err)
		return nil, err
	}
	result := []byte(strings.Trim(string(response), "\x00'"))
	if err = json.Unmarshal(result, responsePayload); err != nil {
		log.Errorf(`'Unable to unmarshal response: %s'`, err)
		return nil, err
	}
	return responsePayload, nil
}
