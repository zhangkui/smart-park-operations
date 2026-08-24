package platform

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	HTTPAddr, DatabaseURL, RedisURL, JWTSecret string
	AccessTokenTTL                             time.Duration
}
type Dependencies struct {
	DB     *pgxpool.Pool
	Cache  *redis.Client
	Log    *slog.Logger
	Config Config
}
type Claims struct {
	UserID  string   `json:"user_id"`
	RoleIDs []string `json:"role_ids"`
	jwt.RegisteredClaims
}

func IssueToken(secret, userID string, roles []string, ttl time.Duration) (string, error) {
	if secret == "" {
		return "", errors.New("jwt secret is required")
	}
	now := time.Now()
	c := Claims{UserID: userID, RoleIDs: roles, RegisteredClaims: jwt.RegisteredClaims{Subject: userID, IssuedAt: jwt.NewNumericDate(now), ExpiresAt: jwt.NewNumericDate(now.Add(ttl))}}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString([]byte(secret))
}
func ParseToken(secret, raw string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(raw, &Claims{}, func(t *jwt.Token) (any, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	c, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return c, nil
}
func JSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
func Decode(r *http.Request, v any) error {
	if r.Body == nil {
		return errors.New("empty body")
	}
	return json.NewDecoder(r.Body).Decode(v)
}
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Request-ID", time.Now().UTC().Format("20060102150405.000000000"))
		next.ServeHTTP(w, r)
	})
}
func Auth(secret string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		raw := r.Header.Get("Authorization")
		if len(raw) < 8 || raw[:7] != "Bearer " {
			JSON(w, 401, map[string]string{"error": "missing bearer token"})
			return
		}
		c, err := ParseToken(secret, raw[7:])
		if err != nil {
			JSON(w, 401, map[string]string{"error": err.Error()})
			return
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), claimsKey{}, c)))
	})
}

type claimsKey struct{}

func CurrentClaims(ctx context.Context) *Claims { c, _ := ctx.Value(claimsKey{}).(*Claims); return c }

type Hub struct {
	mu      sync.RWMutex
	clients map[*websocket.Conn]struct{}
}

func NewHub() *Hub                   { return &Hub{clients: map[*websocket.Conn]struct{}{}} }
func (h *Hub) Add(c *websocket.Conn) { h.mu.Lock(); defer h.mu.Unlock(); h.clients[c] = struct{}{} }
func (h *Hub) Remove(c *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.clients, c)
	_ = c.Close()
}
func (h *Hub) Broadcast(v any) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	for c := range h.clients {
		_ = c.WriteJSON(v)
	}
}

var Upgrader = websocket.Upgrader{CheckOrigin: func(*http.Request) bool { return true }}
