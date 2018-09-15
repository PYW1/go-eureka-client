package eureka

import (
	"encoding/json"
	"strings"
	"fmt"
)

func (c *Client) RegisterInstance(appId string, instanceInfo *InstanceInfo) error {
	values := []string{"apps", appId}
	path := strings.Join(values, "/")
	instance := &Instance{
		Instance: instanceInfo,
	}
	body, err := json.Marshal(instance)
	if err != nil {
		return err
	}

	resp, err := c.Post(path, body)
	if err != nil {
		return err
	}
	if resp == nil {
		return fmt.Errorf("eureka respose is nil")
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("Server response error: %d", resp.StatusCode)
	}
	return err
}
