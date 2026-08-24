package main

import (
	"github.com/zhangkui/smart-park-operations/internal/modules/alerts"
	"github.com/zhangkui/smart-park-operations/internal/modules/audit"
	"github.com/zhangkui/smart-park-operations/internal/modules/energy"
	"github.com/zhangkui/smart-park-operations/internal/modules/inspections"
	"github.com/zhangkui/smart-park-operations/internal/modules/parking"
	"github.com/zhangkui/smart-park-operations/internal/modules/parks"
	"github.com/zhangkui/smart-park-operations/internal/modules/reports"
	"github.com/zhangkui/smart-park-operations/internal/modules/spaces"
	"github.com/zhangkui/smart-park-operations/internal/modules/tenants"
	"github.com/zhangkui/smart-park-operations/internal/modules/users"
	"github.com/zhangkui/smart-park-operations/internal/modules/visitors"
	"github.com/zhangkui/smart-park-operations/internal/modules/workorders"
	"github.com/zhangkui/smart-park-operations/internal/platform"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "development-only-change-me"
	}
	hub := platform.NewHub()
	_ = hub
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		platform.JSON(w, 200, map[string]any{"service": "smart-park-operations", "time": time.Now().UTC()})
	})
	usersRepo := users.NewRepository()
	usersService := users.NewService(usersRepo, func(action, id string) {})
	usersHandler := users.NewHandler(usersService)
	usersHandler.Routes(mux)
	parksRepo := parks.NewRepository()
	parksService := parks.NewService(parksRepo, func(action, id string) {})
	parksHandler := parks.NewHandler(parksService)
	parksHandler.Routes(mux)
	tenantsRepo := tenants.NewRepository()
	tenantsService := tenants.NewService(tenantsRepo, func(action, id string) {})
	tenantsHandler := tenants.NewHandler(tenantsService)
	tenantsHandler.Routes(mux)
	spacesRepo := spaces.NewRepository()
	spacesService := spaces.NewService(spacesRepo, func(action, id string) {})
	spacesHandler := spaces.NewHandler(spacesService)
	spacesHandler.Routes(mux)
	parkingRepo := parking.NewRepository()
	parkingService := parking.NewService(parkingRepo, func(action, id string) {})
	parkingHandler := parking.NewHandler(parkingService)
	parkingHandler.Routes(mux)
	visitorsRepo := visitors.NewRepository()
	visitorsService := visitors.NewService(visitorsRepo, func(action, id string) {})
	visitorsHandler := visitors.NewHandler(visitorsService)
	visitorsHandler.Routes(mux)
	workordersRepo := workorders.NewRepository()
	workordersService := workorders.NewService(workordersRepo, func(action, id string) {})
	workordersHandler := workorders.NewHandler(workordersService)
	workordersHandler.Routes(mux)
	inspectionsRepo := inspections.NewRepository()
	inspectionsService := inspections.NewService(inspectionsRepo, func(action, id string) {})
	inspectionsHandler := inspections.NewHandler(inspectionsService)
	inspectionsHandler.Routes(mux)
	energyRepo := energy.NewRepository()
	energyService := energy.NewService(energyRepo, func(action, id string) {})
	energyHandler := energy.NewHandler(energyService)
	energyHandler.Routes(mux)
	alertsRepo := alerts.NewRepository()
	alertsService := alerts.NewService(alertsRepo, func(action, id string) {})
	alertsHandler := alerts.NewHandler(alertsService)
	alertsHandler.Routes(mux)
	reportsRepo := reports.NewRepository()
	reportsService := reports.NewService(reportsRepo, func(action, id string) {})
	reportsHandler := reports.NewHandler(reportsService)
	reportsHandler.Routes(mux)
	auditRepo := audit.NewRepository()
	auditService := audit.NewService(auditRepo, func(action, id string) {})
	auditHandler := audit.NewHandler(auditService)
	auditHandler.Routes(mux)
	mux.Handle("/ws", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := platform.Upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		hub.Add(c)
		defer hub.Remove(c)
		for {
			if _, _, err := c.ReadMessage(); err != nil {
				return
			}
		}
	}))
	log.Printf("smart park operations listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", platform.RequestID(platform.Auth(secret, mux))))
}
