package commands

import (
	"delivery-planner-bot/services"
	"fmt"
	"strings"
)

// PreviewCommand handles the /preview command which shows a preview of a delivery route
type PreviewCommand struct {
	routeService *services.RouteService
}

// Name returns the command name
func (c *PreviewCommand) Name() string {
	return "preview"
}

// Description returns the command description
func (c *PreviewCommand) Description() string {
	return "Preview a delivery route before optimizing"
}

// NewPreviewCommand creates a new preview command handler
func NewPreviewCommand(routeService *services.RouteService) *PreviewCommand {
	return &PreviewCommand{
		routeService: routeService,
	}
}

func (c *PreviewCommand) Execute(args []string) (string, error) {
	if len(args) == 0 || args[0] == "/preview" {
		return c.getFormatInstructions(), nil
	}

	// Parse comma-separated addresses
	// Remove command from text
	text := strings.TrimSpace(strings.TrimPrefix(args[0], "/preview"))
	addresses := strings.Split(text, ",")
	if len(addresses) < 2 {
		return c.getFormatError(), nil
	}

	origin := strings.TrimSpace(addresses[0])
	destination := strings.TrimSpace(addresses[1])
	deliveries := make([]string, 0)

	// Any additional addresses are delivery points
	for _, addr := range addresses[2:] {
		deliveries = append(deliveries, strings.TrimSpace(addr))
	}

	// Call preview service
	resp, err := c.routeService.PreviewRoute(origin, destination, deliveries)
	if err != nil {
		return fmt.Sprintf(`❌ Error
───────────────
%v`, err), nil
	}

	// Format response
	return c.formatRoutePreview(resp), nil
}

func (c *PreviewCommand) getFormatInstructions() string {
	return `🗺️ Route Preview Format
───────────────
Enter addresses separated by commas in this order:
1. Starting point
2. Final destination
3. Delivery stops

Example:
31 Joo Chiat St, 456 Park Ave, 789 River Rd`
}

func (c *PreviewCommand) getFormatError() string {
	return `❌ Format Error
───────────────
Please include at least:
1. Starting point
2. Final destination

Example: 123 Main St, 456 Park Ave

Need help? Type /preview to see the full guide.`
}

func (c *PreviewCommand) formatRoutePreview(resp *services.PreviewRouteResponse) string {
	var stopsText string
	if len(resp.Deliveries) > 0 {
		stopsText = "\n\n📦 Delivery Stops"
		stopsText += "\n───────────────"
		for i, stop := range resp.Deliveries {
			stopsText += fmt.Sprintf("\n%d. 🔸 %s", i+1, stop)
		}
	}

	return fmt.Sprintf(`✨ Route Preview
───────────────
🚚 From: %s
🏁 To: %s%s

⏱️ Estimated Time: %s
📱 Navigation: %s

💡 Use /optimize to create a full delivery plan`,
		resp.Origin,
		resp.Destination,
		stopsText,
		resp.EstimatedTime,
		resp.GoogleMapsURL)
}
