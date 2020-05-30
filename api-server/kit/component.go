package kit

import (
	"fmt"
	"sync"

	nats "github.com/nats-io/nats.go"
	"github.com/nats-io/nuid"
)

// Component contains reusable logic related
// to hanndling of the connection to NATS in the system.
type Component struct {
	// cmu is the lock form the component.
	cmu sync.Mutex

	// id is a unique identifier used for this component.
	id string

	// nc is the connection to NATS.
	nc *nats.Conn

	// kind is the type of component
	kind string
}

// NewComponent create a Component
func NewComponent(kind string) *Component {
	id := nuid.Next()
	return &Component{
		id:   id,
		kind: kind,
	}
}

// SetupConnectionToNATS connects to NATS
func (c *Component) SetupConnectionToNATS(servers string, options ...nats.Option) error {
	options = append(options, nats.Name(c.Name()))
	c.cmu.Lock()
	defer c.cmu.Unlock()

	nc, err := nats.Connect(servers, options...)
	if err != nil {
		return err
	}
	c.nc = nc

	return nil
}

// NATS returns the current NATS connection.
func (c *Component) NATS() *nats.Conn {
	c.cmu.Lock()
	defer c.cmu.Unlock()
	return c.nc
}

// ID returns the ID from the component.
func (c *Component) ID() string {
	c.cmu.Lock()
	defer c.cmu.Unlock()
	return c.id
}

// Name is the label used to identify the NATS connection.
func (c *Component) Name() string {
	c.cmu.Lock()
	defer c.cmu.Unlock()
	return fmt.Sprintf("%s:%s", c.kind, c.id)
}

// Shutdown makes the component go away.
func (c *Component) Shutdown() error {
	c.NATS().Close()
	return nil
}
