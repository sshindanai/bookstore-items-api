package models

func (c *Cart) CalculateTotalPrice() {
	if len(c.Items) == 0 {
		c.TotalPrice = 0
		return
	}

	for _, item := range c.Items {
		c.TotalPrice += item.Price * float32(item.Quantity)
	}

	return
}
