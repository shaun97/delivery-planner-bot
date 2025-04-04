package commands

type StartCommand struct {
	// Add dependencies if needed
}

func NewStartCommand() *StartCommand {
	return &StartCommand{}
}

func (c *StartCommand) Execute(args []string) (string, error) {
	return `🚚 Delivery Planner Bot
───────────────
Optimize your delivery routes with these commands:

📍 /preview - Preview a delivery route
📦 /newdelivery - Add a new delivery
🔄 /optimize - Generate optimized routes
👤 /assign - Assign routes to drivers
📊 /status - Check delivery status
🗺️ /myroute - View your assigned route

Type any command for detailed instructions.`, nil
}
