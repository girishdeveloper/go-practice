package occupancy

const highLimit = 70
const mediumLimit = 50

func Level(occupancyRate float32) (level string) {
	switch {
	case occupancyRate > highLimit:
		{
			level = "High"
		}
	case occupancyRate > mediumLimit:
		{
			level = "Medium"
		}
	default:
		{
			level = "Low"
		}
	}
	return
}

func Rate(occupiedRooms int, totalRooms int) (rate float32) {
	rate = (float32(occupiedRooms) / float32(totalRooms)) * 100
	return
}
