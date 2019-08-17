package xmeye

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"net"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/kihamo/boggart/components/boggart/providers/xmeye/internal"
)

const (
	DefaultTimeout = time.Second
	DefaultPort    = 34567

	defaultPayloadBuffer = 2048
	defaultReadTimeout   = time.Second * 5
	defaultWriteTimeout  = time.Second * 5

	timeLayout = "2006-01-02 15:04:05"

	CmdLoginResponse                uint16 = 1000
	CmdLogoutResponse               uint16 = 1002
	CmdKeepAliveResponse            uint16 = 1006
	CmdSystemInfoRequest            uint16 = 1020
	CmdConfigGetRequest             uint16 = 1042
	CmdDefaultConfigGetRequest      uint16 = 1044
	CmdConfigChannelTitleSetRequest uint16 = 1046
	CmdConfigChannelTitleGetRequest uint16 = 1048
	CmdAbilityGetRequest            uint16 = 1360
	CmdLogSearchRequest             uint16 = 1442
	CmdSysManagerResponse           uint16 = 1451
	CmdTimeRequest                  uint16 = 1452
	CmdGuardRequest                 uint16 = 1500
	CmdUnGuardRequest               uint16 = 1502
	CmdAlarmRequest                 uint16 = 1504
	CmdConfigExportRequest          uint16 = 1542
	CmdLogExportRequest             uint16 = 1544

	CodeOK                                  = 100
	CodeUnknownError                        = 101
	CodeUnsupportedVersion                  = 102
	CodeRequestNotPermitted                 = 103
	CodeUserAlreadyLoggedIn                 = 104
	CodeUserUserIsNotLoggedIn               = 105
	CodeUsernameOrPasswordIsIncorrect       = 106
	CodeUserDoesNotHaveNecessaryPermissions = 107
	CodeRequestWrongFormat                  = 117
	CodePasswordIsIncorrect                 = 203
	CodeUpgradeSuccessful                   = 515
	CodeConfigurationIsNotExists            = 607
)

type Client struct {
	host     string
	username string
	password []byte
	debug    uint32

	connection net.Conn

	mutex        sync.RWMutex
	mutexRequest sync.Mutex
	done         chan struct{}

	sessionID      uint32
	sequenceNumber uint32

	resetConnection uint32
	alarmStarted    uint32
}

func New(host, username, password string) *Client {
	if _, _, err := net.SplitHostPort(host); err != nil {
		host = host + ":" + strconv.Itoa(DefaultPort)
	}

	return &Client{
		host:     host,
		username: username,
		password: []byte(password),

		sessionID:      0,
		sequenceNumber: 0,
	}
}

func (c *Client) WithDebug(flag bool) *Client {
	var debug uint32
	if flag {
		debug = 1
	}

	atomic.StoreUint32(&c.debug, debug)

	return c
}

func (c *Client) sessionIDAsString() string {
	session := make([]byte, 4)
	binary.LittleEndian.PutUint32(session, atomic.LoadUint32(&c.sessionID))

	return "0x" + hex.EncodeToString([]byte{session[3], session[2], session[1], session[0]})
}

func (c *Client) connect() (conn net.Conn, err error) {
	c.mutex.RLock()
	conn = c.connection
	c.mutex.RUnlock()

	if conn == nil {
		conn, err = net.Dial("tcp", c.host)
		if err != nil {
			if e := ConnectionCheck(conn); e != nil {
				c.reset()
			}

			return nil, err
		}

		tcpConn := conn.(*net.TCPConn)
		err = tcpConn.SetKeepAlive(true)

		if err == nil {
			err = tcpConn.SetReadBuffer(defaultPayloadBuffer * 2)
		}

		if err == nil {
			err = tcpConn.SetWriteBuffer(defaultPayloadBuffer * 2)
		}

		if err != nil {
			conn.Close()
			return nil, err
		}

		c.mutex.Lock()
		c.connection = conn
		c.mutex.Unlock()

		atomic.StoreUint32(&c.sequenceNumber, 0)

		// обрыв соединения обрывает так же сессию
		if atomic.LoadUint32(&c.resetConnection) > 0 {
			atomic.StoreUint32(&c.resetConnection, 0)

			if atomic.LoadUint32(&c.sessionID) > 0 {
				err = c.Login()
			}

			if err == nil && atomic.LoadUint32(&c.alarmStarted) > 0 {
				err = c.AlarmStart()
			}
		}
	}

	return conn, err
}

func (c *Client) reset() {
	c.mutex.Lock()
	if c.connection != nil {
		c.connection.Close()
		c.connection = nil
	}
	c.mutex.Unlock()

	atomic.StoreUint32(&c.resetConnection, 1)
}

func (c *Client) request(code uint16, payload interface{}) (*internal.Packet, error) {
	conn, err := c.connect()
	if err != nil {
		return nil, err
	}

	c.mutexRequest.Lock()
	defer c.mutexRequest.Unlock()

	//
	// --- REQUEST ---
	//
	requestPacket := internal.NewPacket()
	requestPacket.SessionID = atomic.LoadUint32(&c.sessionID)
	requestPacket.SequenceNumber = atomic.LoadUint32(&c.sequenceNumber)
	requestPacket.MessageID = code

	if err := requestPacket.LoadPayload(payload); err != nil {
		return nil, err
	}

	requestPacketBytes := requestPacket.Marshal()

	err = conn.SetWriteDeadline(time.Now().Add(defaultWriteTimeout))
	if err == nil {
		_, err = conn.Write(requestPacketBytes)
	}

	if err != nil {
		if e := ConnectionCheck(conn); e != nil {
			c.reset()
		}

		return nil, err
	}

	debug := atomic.LoadUint32(&c.debug)
	if debug > 0 {
		fmt.Println(">>> request")
		fmt.Println(hex.Dump(requestPacketBytes))
	}

	atomic.AddUint32(&c.sequenceNumber, 1)

	//
	// --- RESPONSE ---
	//
	var responsePacketHead []byte

	err = conn.SetReadDeadline(time.Now().Add(defaultReadTimeout))
	if err == nil {
		responsePacketHead = make([]byte, 0x14) // read head
		_, err = conn.Read(responsePacketHead)
	}

	if err != nil {
		if e := ConnectionCheck(conn); e != nil {
			c.reset()
		}

		return nil, err
	}

	// save session id
	responsePacket := internal.PacketUnmarshal(responsePacketHead)
	if responsePacket.PayloadLen == 0 {
		return &responsePacket, nil
	}

	bufSize := defaultPayloadBuffer
	if bufSize > responsePacket.PayloadLen {
		bufSize = responsePacket.PayloadLen
	}
	buf := make([]byte, bufSize)
	var n int

	for {
		n, err = conn.Read(buf)

		if err != nil {
			return &responsePacket, err
		}

		responsePacket.Payload.Write(buf[:n])

		if responsePacket.Payload.Len() >= responsePacket.PayloadLen {
			break
		}
	}

	if debug > 0 {
		fmt.Println("<<< response")
		fmt.Println(hex.Dump(responsePacket.Marshal()))
		fmt.Println(responsePacket.Payload.String())
	}

	return &responsePacket, nil
}

func (c *Client) Call(code uint16, payload interface{}) (*internal.Packet, error) {
	p, err := c.request(code, payload)
	if err != nil {
		return nil, err
	}

	err = c.PayloadError(p.Payload)
	return p, err
}

func (c *Client) CallWithResult(code uint16, payload, result interface{}) error {
	packet, err := c.Call(code, payload)
	if err != nil {
		return err
	}

	return packet.Payload.UnmarshalJSON(result)
}

func (c *Client) Cmd(code uint16, cmd string) (*internal.Packet, error) {
	return c.Call(code, map[string]string{
		"Name":      cmd,
		"SessionID": c.sessionIDAsString(),
	})
}

func (c *Client) CmdWithResult(code uint16, cmd string, result interface{}) error {
	return c.CallWithResult(code, map[string]string{
		"Name":      cmd,
		"SessionID": c.sessionIDAsString(),
	}, &result)
}

func (c *Client) PayloadError(payload *internal.Payload) error {
	var result internal.Response

	if err := payload.UnmarshalJSON(&result); err == nil {
		switch result.Ret {
		case CodeOK, CodeUpgradeSuccessful, 0:
			return nil

		case CodeUnknownError:
			return errors.New("unknown error")

		case CodeUnsupportedVersion:
			return errors.New("unsupported version")

		case CodeRequestNotPermitted:
			return errors.New("request not permitted")

		case CodeUserAlreadyLoggedIn:
			return errors.New("user already logged in")

		case CodeUserUserIsNotLoggedIn:
			return errors.New("user is not logged in")

		case CodeRequestWrongFormat:
			return errors.New("request wrong format")

		case CodeUsernameOrPasswordIsIncorrect:
			return errors.New("username or password is incorrect")

		case CodeUserDoesNotHaveNecessaryPermissions:
			return errors.New("user does not have necessary permissions")

		case CodePasswordIsIncorrect:
			return errors.New("password is incorrect")

		case CodeConfigurationIsNotExists:
			return errors.New("configuration is not exists")

		default:
			return errors.New("unsupported unknown error")
		}
	}

	return nil
}

func (c *Client) Close() (err error) {
	c.mutex.Lock()

	if c.done != nil {
		close(c.done)
		c.done = nil
	}

	if c.connection != nil {
		err = c.connection.Close()
		c.connection = nil
	}

	c.mutex.Unlock()

	return err
}
