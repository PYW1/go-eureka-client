package eureka

import (
	"strings"
	"fmt"
)

func (c *Client) SendHeartbeat(appId, instanceId string) error {
	values := []string{"apps", appId, instanceId}
	path := strings.Join(values, "/")
	resp, err := c.Put(path, nil)
	if err != nil {
		return err
	}
	if resp == nil {
		return fmt.Errorf("eureka respose is nil")
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("Server response error: %d", resp.StatusCode)
	}
	return err
}
