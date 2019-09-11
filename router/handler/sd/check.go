package sd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"net/http"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func HealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "server is running")
}

func DiskCheck(c *gin.Context) {
	stat, _ := disk.Usage("/")

	usedMB := int(stat.Used) / MB
	usedGB := int(stat.Used) / GB
	totalMB := int(stat.Total) / MB
	totalGB := int(stat.Total) / GB
	usedPercent := int(stat.UsedPercent)

	status := http.StatusOK
	text := "OK"

	if usedPercent > 95 {
		text = "CRITICAL"
	} else if usedPercent > 90 {
		text = "WARNING"
		status = http.StatusTooManyRequests
	}

	message := fmt.Sprintf("%s - Disk Usage : %dMB(%dGB) / %dMB(%dGB) | Used %d%%", text, usedMB, usedGB, totalMB, totalGB, usedPercent)
	c.String(status, message)
}

func CPUCheck(c *gin.Context) {
	core, _ := cpu.Counts(false)

	stat, _ := load.Avg()
	load1 := stat.Load1
	load5 := stat.Load15
	load15 := stat.Load15

	status := http.StatusOK
	text := "OK"

	if load5 >= float64(core-1) {
		status = http.StatusInternalServerError
		text = "CRITICAL"
	} else if load5 >= float64(core-2) {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}

	message := fmt.Sprintf("%s - Load Average : %.2f %.2f %.2f | Cores : %d ", text, load1, load5, load15, core)
	c.String(status, message)
}

func RAMCheck(c *gin.Context) {
	stat, _ := mem.VirtualMemory()

	usedMB := int(stat.Used) / MB
	usedGB := int(stat.Used) / GB
	totalMB := int(stat.Total) / MB
	totalGB := int(stat.Total) / GB
	usedPercent := int(stat.UsedPercent)

	status := http.StatusOK
	text := "OK"

	if usedPercent > 95 {
		text = "CRITICAL"
	} else if usedPercent > 90 {
		text = "WARNING"
		status = http.StatusTooManyRequests
	}

	message := fmt.Sprintf("%s - Memory Usage : %dMB(%dGB) / %dMB(%dGB) | Used %d%%", text, usedMB, usedGB, totalMB, totalGB, usedPercent)
	c.String(status, message)
}
